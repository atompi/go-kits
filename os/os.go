package os

import (
	"os"
	"os/exec"
)

func ExecCmd(command string) (output string, err error) {
	cmd := exec.Command("/bin/sh", "-c", command)
	outputBuf, err := cmd.Output()
	if err != nil {
		output = ""
		return
	}
	output = string(outputBuf)
	return
}

func PathExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}
