package Util

import (
	"io/ioutil"
)

func PrintStringToFile(content string, path string) error {
	d1 := []byte(content)
	err := ioutil.WriteFile(path, d1, 0644)

	return err
}

func PrintByteToFile(d1 []byte, path string) error {
	err := ioutil.WriteFile(path, d1, 0644)

	return err
}
