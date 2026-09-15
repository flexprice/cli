package update

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

const (
	ReleasesURL = "https://github.com/flexprice/cli/releases"
	CheckInterval = 24 * time.Hour
	fetchTimeout = 3 * time.Second
)

type Checker struct {
	ReleasesURL string
	CachePath   string
	Client      *http.Client
	Now         func() time.Time
}

func NewChecker() *Checker {
	c := &Checker{}
	if home, err := os.UserHomeDir(); err == nil {
		c.CachePath = filepath.Join(home, ".flexprice", "update-check.json")
	}
	return c
}

type cache struct {
	CheckedAt time.Time `json:"checked_at"`
	Latest    string    `json:"latest"`
}

func (c *Checker) Latest(ctx context.Context) (string, error) {
	cached := c.read()
	if c.now().Sub(cached.CheckedAt) < CheckInterval {
		return cached.Latest, nil
	}
	tag, err := c.Fetch(ctx)
	if err != nil {
		c.write(cache{CheckedAt: c.now(), Latest: cached.Latest})
		return cached.Latest, err
	}
	return tag, nil
}

func (c *Checker) Fetch(ctx context.Context) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, fetchTimeout)
	defer cancel()

	base := c.ReleasesURL
	if base == "" {
		base = ReleasesURL
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, base+"/latest", nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("User-Agent", "flexprice-cli")

	client := http.Client{}
	if c.Client != nil {
		client = *c.Client
	}
	client.CheckRedirect = func(*http.Request, []*http.Request) error { return http.ErrUseLastResponse }

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	loc := resp.Header.Get("Location")
	i := strings.LastIndex(loc, "/tag/")
	if resp.StatusCode < 300 || resp.StatusCode >= 400 || i < 0 {
		return "", fmt.Errorf("resolve latest release: unexpected %s from %s/latest", resp.Status, base)
	}
	tag := strings.TrimSuffix(loc[i+len("/tag/"):], "/")
	if _, ok := parse(tag); !ok {
		return "", fmt.Errorf("resolve latest release: %q is not a release tag", tag)
	}
	c.write(cache{CheckedAt: c.now(), Latest: tag})
	return tag, nil
}

func IsNewer(current, latest string) bool {
	cur, ok := parse(current)
	if !ok {
		return false
	}
	lat, ok := parse(latest)
	if !ok {
		return false
	}
	for i := range cur {
		if lat[i] != cur[i] {
			return lat[i] > cur[i]
		}
	}
	return false
}

func parse(v string) ([3]int, bool) {
	var out [3]int
	parts := strings.Split(strings.TrimPrefix(strings.TrimSpace(v), "v"), ".")
	if len(parts) != 3 {
		return out, false
	}
	for i, p := range parts {
		n, err := strconv.Atoi(p)
		if err != nil || n < 0 {
			return out, false
		}
		out[i] = n
	}
	return out, true
}

func (c *Checker) now() time.Time {
	if c.Now != nil {
		return c.Now()
	}
	return time.Now()
}

func (c *Checker) read() cache {
	var s cache
	if c.CachePath == "" {
		return s
	}
	b, err := os.ReadFile(c.CachePath)
	if err != nil || json.Unmarshal(b, &s) != nil {
		return cache{}
	}
	return s
}

func (c *Checker) write(s cache) {
	if c.CachePath == "" {
		return
	}
	if err := os.MkdirAll(filepath.Dir(c.CachePath), 0o700); err != nil {
		return
	}
	b, err := json.Marshal(s)
	if err != nil {
		return
	}
	_ = os.WriteFile(c.CachePath, b, 0o600)
}
