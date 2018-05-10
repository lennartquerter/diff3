package Util

import (
	"io/ioutil"
	"path/filepath"
	"os"
	"io"
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

func Except(list []string, expectList []string) ([]string) {
	var result []string

	for _, v := range list {

		if !Contains(expectList, v) {
			result = append(result, v)
		}
	}

	return result
}

func Common(list1 []string, list2 []string) ([]string) {
	var result []string

	for _, v := range list1 {

		if Contains(list2, v) {
			result = append(result, v)
		}

	}

	return result
}

func Contains(list []string, item string) bool {
	for _, v := range list {
		if v == item {
			return true
		}
	}
	return false
}


func FilePathWalkDir(root string) ([]string, error) {
	var files []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	return files, err
}

func IsEmpty(name string) (bool, error) {
	f, err := os.Open(name)
	if err != nil {
		return false, err
	}
	defer f.Close()

	_, err = f.Readdirnames(1) // Or f.Readdir(1)
	if err == io.EOF {
		return true, nil
	}
	return false, err // Either not empty or error, suits both cases
}

