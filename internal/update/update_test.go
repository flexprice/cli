package update

import (
	"context"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"sync/atomic"
	"testing"
	"time"
)

func TestIsNewer(t *testing.T) {
	cases := []struct {
		current, latest string
		want            bool
	}{
		{"1.0.1", "v2.0.1", true},
		{"v1.0.1", "v2.0.1", true},
		{"1.2.3", "v1.10.0", true},
		{"2.0.1", "v2.0.1", false},
		{"2.0.1", "v1.9.9", false},
		{"dev", "v9.9.9", false},
		{"1.0.0", "v1.0.1-rc1", false},
		{"1.0.0", "", false},
		{"1.0.0", "nightly", false},
	}
	for _, c := range cases {
		if got := IsNewer(c.current, c.latest); got != c.want {
			t.Errorf("IsNewer(%q, %q) = %v, want %v", c.current, c.latest, got, c.want)
		}
	}
}

// A release server that only answers the /releases/latest redirect. The tag
// can be swapped and the server can be told to fail, both mid-test.
type fakeReleases struct {
	*httptest.Server
	hits atomic.Int32
	tag  atomic.Value
	fail atomic.Bool
}

func newFakeReleases(t *testing.T, tag string) *fakeReleases {
	t.Helper()
	f := &fakeReleases{}
	f.tag.Store(tag)
	f.Server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/releases/latest" {
			http.NotFound(w, r)
			return
		}
		f.hits.Add(1)
		if f.fail.Load() {
			http.Error(w, "down", http.StatusServiceUnavailable)
			return
		}
		http.Redirect(w, r, "/releases/tag/"+f.tag.Load().(string), http.StatusFound)
	}))
	t.Cleanup(f.Close)
	return f
}

func newChecker(t *testing.T, f *fakeReleases, now *time.Time) *Checker {
	t.Helper()
	return &Checker{
		ReleasesURL: f.URL + "/releases",
		CachePath:   filepath.Join(t.TempDir(), "update-check.json"),
		Now:         func() time.Time { return *now },
	}
}

func TestLatest_FetchesOnceThenCaches(t *testing.T) {
	f := newFakeReleases(t, "v2.0.1")
	now := time.Date(2026, 9, 12, 9, 0, 0, 0, time.UTC)
	c := newChecker(t, f, &now)

	for i := 0; i < 3; i++ {
		tag, err := c.Latest(context.Background())
		if err != nil || tag != "v2.0.1" {
			t.Fatalf("Latest = %q, %v; want v2.0.1", tag, err)
		}
	}
	if got := f.hits.Load(); got != 1 {
		t.Errorf("server hit %d times, want 1", got)
	}
}

func TestLatest_RefetchesAfterInterval(t *testing.T) {
	f := newFakeReleases(t, "v2.0.1")
	now := time.Date(2026, 9, 12, 9, 0, 0, 0, time.UTC)
	c := newChecker(t, f, &now)

	if _, err := c.Latest(context.Background()); err != nil {
		t.Fatal(err)
	}
	f.tag.Store("v2.1.0")
	now = now.Add(CheckInterval + time.Minute)

	tag, err := c.Latest(context.Background())
	if err != nil || tag != "v2.1.0" {
		t.Fatalf("Latest after interval = %q, %v; want v2.1.0", tag, err)
	}
	if got := f.hits.Load(); got != 2 {
		t.Errorf("server hit %d times, want 2", got)
	}
}

// Offline users must not pay the timeout on every command, and must not lose
// the tag they already knew about.
func TestLatest_FailureIsCachedAndKeepsLastTag(t *testing.T) {
	f := newFakeReleases(t, "v2.0.1")
	now := time.Date(2026, 9, 12, 9, 0, 0, 0, time.UTC)
	c := newChecker(t, f, &now)

	if _, err := c.Latest(context.Background()); err != nil {
		t.Fatal(err)
	}
	f.fail.Store(true)
	now = now.Add(CheckInterval + time.Minute)

	tag, err := c.Latest(context.Background())
	if err == nil {
		t.Fatal("expected an error from a failing server")
	}
	if tag != "v2.0.1" {
		t.Errorf("failed fetch should return the last known tag, got %q", tag)
	}

	tag, err = c.Latest(context.Background())
	if err != nil || tag != "v2.0.1" {
		t.Errorf("within the interval the failure must be cached: got %q, %v", tag, err)
	}
	if got := f.hits.Load(); got != 2 {
		t.Errorf("server hit %d times, want 2", got)
	}
}

func TestFetch_RejectsNonReleaseRedirect(t *testing.T) {
	f := newFakeReleases(t, "nightly")
	now := time.Now()
	c := newChecker(t, f, &now)
	if _, err := c.Fetch(context.Background()); err == nil {
		t.Error("a redirect to a non-version tag must be an error, not a version")
	}
}

func TestLatest_CorruptCacheIsIgnored(t *testing.T) {
	f := newFakeReleases(t, "v2.0.1")
	now := time.Now()
	c := newChecker(t, f, &now)
	if err := os.WriteFile(c.CachePath, []byte("{not json"), 0o600); err != nil {
		t.Fatal(err)
	}
	tag, err := c.Latest(context.Background())
	if err != nil || tag != "v2.0.1" {
		t.Fatalf("Latest with corrupt cache = %q, %v; want v2.0.1", tag, err)
	}
}
