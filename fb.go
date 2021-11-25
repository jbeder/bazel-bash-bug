package fb

import (
	"io"
	"os"
	"os/exec"
	"path"
	"time"

	"github.com/bazelbuild/rules_go/go/tools/bazel"
)

var firebase = "external/npm/firebase-tools/bin/firebase.sh"

// StartFirebaseEmulator starts an emulator that's an interface into the
// Firebase API.
func Run() error {
	rfpath, err := bazel.RunfilesPath()
	if err != nil {
		return err
	}

	binary, err := bazel.Runfile(firebase)
	if err != nil {
		return err
	}
	dir, err := os.MkdirTemp("", "firebase")
	if err != nil {
		return err
	}
	defer os.RemoveAll(dir)
	configFile := path.Join(dir, "firebase.json")
	data := []byte(`{
  "emulators": {
    "auth": {
      "port": 9099
    },
    "ui": {
      "enabled": true
    }
  }
}`)
	if err := os.WriteFile(configFile, data, 0700); err != nil {
		return err
	}

	args := []string{
		"emulators:start",
		"--config", configFile,
		"--project", "demo-foo",
	}
	cmd := exec.Command(binary, args...)
	cmd.Dir = rfpath
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		return err
	}
	go func() {
		io.Copy(os.Stdout, stdout)
	}()
	go func() {
		io.Copy(os.Stderr, stderr)
	}()
	if err := cmd.Start(); err != nil {
		return err
	}
	done := make(chan error, 1)
	go func() {
		done <- cmd.Wait()
	}()

	select {
	case <-time.After(4 * time.Second):
		return cmd.Process.Kill()
	case err := <-done:
		return err
	}
}
