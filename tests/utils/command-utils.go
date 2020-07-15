package utils

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path"
	"runtime"
)

func RunScriptFromBaseDir(command string, args []string, envVars []string) (string, string, error) {
	_, filename, _, _ := runtime.Caller(0)
	dir := path.Join(path.Dir(filename), "..", "..")
	return RunCommand(fmt.Sprintf("%s/%s", dir, command), args, envVars)
}

func RunCommand(command string, args []string, envVars []string) (string, string, error) {
	cmd := exec.Command(command, args...)
	cmd.Env = append(os.Environ(), envVars...)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	return string(stdout.Bytes()), string(stderr.Bytes()), err
}
