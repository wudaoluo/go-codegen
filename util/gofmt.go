package util

import (
	"github.com/wudaoluo/goutil"
	"os/exec"
)

//TODO 这些命令整理下
func Gofmt(path string) bool {
	if goutil.FileIsExist(path) {
		if !ExecCommand("goimports", "-l", "-w", path) {
			if !ExecCommand("gofmt", "-l", "-w", path) {
				return ExecCommand("go", "fmt", path)
			}
		}
		return true
	}
	return false
}


func ExecCommand(name string, args ...string) bool {
	cmd := exec.Command(name, args...)
	_, err := cmd.Output()
	if err != nil {
		return false
	}
	return true
}
