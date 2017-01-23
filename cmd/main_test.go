package main

import (
	"bytes"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

func TestMain(t *testing.T) {
	executable, _ := filepath.Abs("../build/minivm")
	filepath.Walk("../test", func(path string, info os.FileInfo, err error) error {
		if strings.HasSuffix(path, ".in") {
			cmd := exec.Command("bash", "-c", executable+" < "+filepath.Base(path))
			cmd.Dir = filepath.Dir(path)
			stderr := new(bytes.Buffer)
			cmd.Stderr = stderr
			output, err := cmd.Output()
			if err != nil {
				t.Errorf("FAIL: execution failed: " + path + ": " + err.Error() + ", " + string(output) + " " + stderr.String())
			} else {
				outfile := strings.TrimSuffix(path, filepath.Ext(path)) + ".out"
				expected, err := ioutil.ReadFile(outfile)
				if err != nil {
					t.Errorf("FAIL: error on reading output file: " + outfile)
				} else if string(output) == string(expected) {
					t.Logf("PASS: " + path + "\n")
				} else {
					t.Errorf("FAIL: output differs: " + path + "\nOutput:" + string(output) + "\nExpected:" + string(expected))
				}
			}
		}
		return nil
	})
}
