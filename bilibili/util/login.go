package util

import (
	"io/ioutil"
	"os"
	"path"
	"strings"
)

var (
	pwd, _     = os.Getwd()
	cookieDir  = path.Join(pwd, "/cookie")
	cookieFile = ".bilibili_cookie.txt"
)

func ReadCookieFromFile() ([]string, error) {
	file, err := os.Open(path.Join(cookieDir, cookieFile))
	if err != nil {
		return nil, err
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	return strings.Split(string(data), "\n"), nil
}

func CreateLoginDir() error {
	_, err := os.Stat(cookieDir)
	switch err {
	case nil:
	default:
		if os.IsNotExist(err) {
			if err := os.MkdirAll(cookieDir, 0o755); err != nil {
				return err
			}
		} else {
			return err
		}
	}
	return nil
}

func SaveCookieFile(cookie []string) error {
	err := CreateLoginDir()
	if err != nil {
		return err
	}

	file, err := os.OpenFile(path.Join(cookieDir, cookieFile), os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.Write([]byte(strings.Join(cookie, "\n")))
	return err
}
