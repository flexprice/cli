package cmd

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"

	"github.com/spf13/cobra"

	"github.com/flexprice/cli/internal/ui"
	"github.com/flexprice/cli/internal/update"
)

func TestWantsUpdateNotice(t *testing.T) {
	cases := []struct {
		name    string
		command string
		version string
		quiet   bool
		env     string
		tty     bool
		want    bool
	}{
		{name: "resource command on a tty", command: "customers", version: "1.0.1", tty: true, want: true},
		{name: "bare root", command: "flexprice", version: "1.0.1", tty: true, want: true},
		{name: "version", command: "version", version: "1.0.1", tty: true, want: false},
		{name: "update", command: "update", version: "1.0.1", tty: true, want: false},
		{name: "help", command: "help", version: "1.0.1", tty: true, want: false},
		{name: "completion", command: "completion", version: "1.0.1", tty: true, want: false},
		{name: "completion internals", command: cobra.ShellCompRequestCmd, version: "1.0.1", tty: true, want: false},
		{name: "dev build", command: "customers", version: "dev", tty: true, want: false},
		{name: "quiet", command: "customers", version: "1.0.1", quiet: true, tty: true, want: false},
		{name: "not a terminal", command: "customers", version: "1.0.1", tty: false, want: false},
		{name: "opted out", command: "customers", version: "1.0.1", tty: true, env: "1", want: false},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			t.Setenv("FLEXPRICE_NO_UPDATE", c.env)
			got := wantsUpdateNotice(&cobra.Command{Use: c.command}, &Globals{Quiet: c.quiet}, c.version, c.tty)
			if got != c.want {
				t.Errorf("wantsUpdateNotice = %v, want %v", got, c.want)
			}
		})
	}
}

func TestWantsUpdateNotice_NestedCompletion(t *testing.T) {
	parent := &cobra.Command{Use: "completion"}
	child := &cobra.Command{Use: "zsh"}
	parent.AddCommand(child)
	if wantsUpdateNotice(child, &Globals{}, "1.0.1", true) {
		t.Error("completion subcommands must not trigger a notice")
	}
}

// Serves the /releases/latest redirect and, at "/", an install script that
// records the environment it was given.
func fakeReleases(t *testing.T, tag string) *httptest.Server {
	t.Helper()
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/releases/latest":
			http.Redirect(w, r, "/releases/tag/"+tag, http.StatusFound)
		case "/":
			_, _ = w.Write([]byte("#!/bin/sh\nprintf '%s %s\\n' \"$FLEXPRICE_VERSION\" \"$FLEXPRICE_INSTALL_DIR\" > \"$INSTALL_LOG\"\n"))
		default:
			http.NotFound(w, r)
		}
	}))
	t.Cleanup(srv.Close)
	return srv
}

func testChecker(t *testing.T, srv *httptest.Server) *update.Checker {
	t.Helper()
	return &update.Checker{
		ReleasesURL: srv.URL + "/releases",
		CachePath:   filepath.Join(t.TempDir(), "update-check.json"),
	}
}

func testUI(out, errOut *bytes.Buffer) *ui.UI {
	return ui.New(ui.Options{Out: out, Err: errOut, StderrTTY: true, Term: "dumb"})
}

// The two lines are a contract: exactly this text, stderr only.
func TestPrintUpdateNotice_ExactText(t *testing.T) {
	srv := fakeReleases(t, "v2.0.1")
	var out, errOut bytes.Buffer
	g := &Globals{UI: testUI(&out, &errOut)}

	printUpdateNotice(context.Background(), g, "1.0.1", testChecker(t, srv))

	want := "A new version of flexprice is available: v1.0.1 → v2.0.1\n" +
		"Run `flexprice update` to upgrade.\n"
	if errOut.String() != want {
		t.Errorf("stderr:\n%q\nwant:\n%q", errOut.String(), want)
	}
	if out.Len() != 0 {
		t.Errorf("the notice must never touch stdout, got:\n%s", out.String())
	}
}

func TestPrintUpdateNotice_SilentWhenCurrent(t *testing.T) {
	srv := fakeReleases(t, "v1.0.1")
	var out, errOut bytes.Buffer
	g := &Globals{UI: testUI(&out, &errOut)}

	printUpdateNotice(context.Background(), g, "1.0.1", testChecker(t, srv))
	if errOut.Len() != 0 {
		t.Errorf("no notice expected when current, got:\n%s", errOut.String())
	}
}

// The hook runs before every command and must never get in its way.
func TestPrintUpdateNotice_SilentOnFailure(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "down", http.StatusServiceUnavailable)
	}))
	t.Cleanup(srv.Close)
	var out, errOut bytes.Buffer
	g := &Globals{UI: testUI(&out, &errOut)}

	printUpdateNotice(context.Background(), g, "1.0.1", testChecker(t, srv))
	if errOut.Len() != 0 {
		t.Errorf("a failed check must be silent outside --debug, got:\n%s", errOut.String())
	}
}

func runUpdate(t *testing.T, srv *httptest.Server, version string, args ...string) (string, string, error) {
	t.Helper()
	var out, errOut bytes.Buffer
	cmd := newUpdateCommand(&Globals{UI: testUI(&out, &errOut)}, version, testChecker(t, srv), srv.URL+"/")
	cmd.SetOut(&out)
	cmd.SetErr(&errOut)
	cmd.SetArgs(args)
	err := cmd.Execute()
	return out.String(), errOut.String(), err
}

func TestUpdateCommand_UpToDate(t *testing.T) {
	srv := fakeReleases(t, "v1.0.1")
	out, _, err := runUpdate(t, srv, "1.0.1")
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(out, "up to date") {
		t.Errorf("expected an up-to-date message, got:\n%s", out)
	}
}

func TestUpdateCommand_CheckOnly(t *testing.T) {
	srv := fakeReleases(t, "v2.0.1")
	t.Setenv("INSTALL_LOG", filepath.Join(t.TempDir(), "never-written"))
	out, _, err := runUpdate(t, srv, "1.0.1", "--check")
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(out, "v2.0.1") {
		t.Errorf("stdout should name the newer version, got:\n%s", out)
	}
	if _, err := os.Stat(os.Getenv("INSTALL_LOG")); err == nil {
		t.Error("--check must not run the install script")
	}
}

func TestUpdateCommand_RunsInstallScriptPinnedToBinaryDir(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("the install script is not run on Windows")
	}
	srv := fakeReleases(t, "v2.0.1")
	log := filepath.Join(t.TempDir(), "install.log")
	t.Setenv("INSTALL_LOG", log)

	if _, _, err := runUpdate(t, srv, "1.0.1"); err != nil {
		t.Fatalf("update: %v", err)
	}

	exe, err := os.Executable()
	if err != nil {
		t.Fatal(err)
	}
	if resolved, err := filepath.EvalSymlinks(exe); err == nil {
		exe = resolved
	}
	got, err := os.ReadFile(log)
	if err != nil {
		t.Fatalf("install script did not run: %v", err)
	}
	want := "v2.0.1 " + filepath.Dir(exe) + "\n"
	if string(got) != want {
		t.Errorf("script env = %q, want %q", got, want)
	}
}

func TestUpdateCommand_FailsWhenUnreachable(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "down", http.StatusServiceUnavailable)
	}))
	t.Cleanup(srv.Close)
	if _, _, err := runUpdate(t, srv, "1.0.1"); err == nil {
		t.Fatal("an explicit update that cannot reach the release server must fail loudly")
	}
}
