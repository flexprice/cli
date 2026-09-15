package cmd

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"github.com/spf13/cobra"
	"github.com/flexprice/cli/internal/update"
)

const installScriptURL = "https://cli.flexprice.io"

var noUpdateNotice = map[string]bool{
	"version":                       true,
	"update":                        true,
	"help":                          true,
	"completion":                    true,
	cobra.ShellCompRequestCmd:       true,
	cobra.ShellCompNoDescRequestCmd: true,
}

func wantsUpdateNotice(cmd *cobra.Command, g *Globals, version string, stderrTTY bool) bool {
	if !stderrTTY || g.Quiet || version == "dev" {
		return false
	}
	if os.Getenv("FLEXPRICE_NO_UPDATE") != "" {
		return false
	}
	for c := cmd; c != nil; c = c.Parent() {
		if noUpdateNotice[c.Name()] {
			return false
		}
	}
	return true
}

func printUpdateNotice(ctx context.Context, g *Globals, version string, c *update.Checker) {
	if ctx == nil {
		ctx = context.Background()
	}
	latest, err := c.Latest(ctx)
	if err != nil && g.Debug {
		g.UI.Info("update check skipped: %v", err)
	}
	if !update.IsNewer(version, latest) {
		return
	}
	g.UI.Info("A new version of flexprice is available: %s → %s", withV(version), latest)
	g.UI.Info("Run `flexprice update` to upgrade.")
}

// Builds stamp "1.0.1"; release tags read "v1.0.1". Shown the tag way.
func withV(v string) string {
	if v == "" || v[0] == 'v' {
		return v
	}
	return "v" + v
}

func newUpdateCommand(g *Globals, version string, c *update.Checker, scriptURL string) *cobra.Command {
	var check bool
	cmd := &cobra.Command{
		Use: "update",
		Short: "Update the CLI to the latest release",
		Long: "Check GitHub for a newer release and, if there is one, run the install " +
			"script (" + installScriptURL + ") to replace this binary in place.\n\n" +
			"Every other command checks once a day and prints a notice when a newer " +
			"release exists; set FLEXPRICE_NO_UPDATE=1 to turn that off.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			ctx := cmd.Context()
			if ctx == nil {
				ctx = context.Background()
			}
			latest, err := c.Fetch(ctx)
			if err != nil {
				return err
			}
			if !update.IsNewer(version, latest) {
				g.UI.Data("flexprice %s is up to date (latest release: %s)", withV(version), latest)
				return nil
			}
			if check {
				g.UI.Data("flexprice %s is available (installed: %s)", latest, withV(version))
				g.UI.Info("Run `flexprice update` to upgrade.")
				return nil
			}
			if runtime.GOOS == "windows" {
				return fmt.Errorf("the install script does not support Windows; download %s from %s/latest", latest, update.ReleasesURL)
			}
			g.UI.Info("Updating flexprice %s to %s...", withV(version), latest)
			return runInstallScript(ctx, scriptURL, latest, cmd.OutOrStdout(), cmd.ErrOrStderr())
		},
	}
	cmd.Flags().BoolVar(&check, "check", false, "report whether a newer release exists without installing it")
	return cmd
}

func runInstallScript(ctx context.Context, scriptURL, version string, stdout, stderr io.Writer) error {
	exe, err := os.Executable()
	if err != nil {
		return err
	}
	if resolved, err := filepath.EvalSymlinks(exe); err == nil {
		exe = resolved
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, scriptURL, nil)
	if err != nil {
		return err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("download install script: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("download install script: unexpected %s from %s", resp.Status, scriptURL)
	}
	script, err := io.ReadAll(io.LimitReader(resp.Body, 1<<20))
	if err != nil {
		return fmt.Errorf("download install script: %w", err)
	}

	sh := exec.CommandContext(ctx, "sh", "-s")
	sh.Stdin = bytes.NewReader(script)
	sh.Stdout = stdout
	sh.Stderr = stderr
	sh.Env = append(os.Environ(),
		"FLEXPRICE_VERSION="+version,
		"FLEXPRICE_INSTALL_DIR="+filepath.Dir(exe),
	)
	return sh.Run()
}
