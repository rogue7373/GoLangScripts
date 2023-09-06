package username

import (
	"errors"
	"os/user"
	"runtime"
)

func GetUsername() (string, error) {
	var u *user.User
	var err error
	if runtime.GOOS == "windows" {
		u, err = user.Current()
	} else if runtime.GOOS == "darwin" {
		u, err = user.Current()
	} else {
		return "", errors.New("unsupported operating system")
	}

	if err != nil {
		return "", err
	}

	return u.Username, nil
}
