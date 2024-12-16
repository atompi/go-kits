package os

import (
	"os"
	"os/exec"
	"os/signal"
	"syscall"

	"go.uber.org/zap"
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

func GracefulExit() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	s := <-c
	zap.S().Warnf("a %v signal is received, exiting...", s)
	os.Exit(0)
}
