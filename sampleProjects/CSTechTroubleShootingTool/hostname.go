package hostname

import (
	"errors"
	"os/exec"
	"runtime"
	"strings"
)

func GetHostname() (string, error) {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("hostname")
	} else if runtime.GOOS == "darwin" {
		cmd = exec.Command("scutil", "--get", "LocalHostName")
	} else {
		return "", errors.New("unsupported operating system")
	}

	output, err := cmd.Output()
	if err != nil {
		return "", err
	}

	hostname := string(output)
	hostname = strings.TrimSpace(hostname)

	return hostname, nil
}
