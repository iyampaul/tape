package actions

import "os/exec"

func cmd(cmd string, args string) ([]byte, error) {
	if len(args) > 0 {
		return (exec.Command(cmd, args).CombinedOutput())
	} else {
		return (exec.Command(cmd).CombinedOutput())
	}
}
