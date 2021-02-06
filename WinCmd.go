package main

import(
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func execute_cmd_command(command string) error {
	cmd := exec.Command("cmd", "/c",command)

	cmd.Stdout = os.Stdout
	cmd.Stdin  = os.Stdin
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	return err
}

func clear_str(str string) string {
	if runtime.GOOS == "windows" {
		return strings.TrimRight(str,"\r\n")
	} else {
		return strings.TrimRight(str,"\n")
	}
}