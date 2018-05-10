package main

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

func MergeDir(path string, list []Model, debug bool) error {

	for _, v := range list {
		src, err := FilePathWalkDir(path + v.src)

		if err != nil {
			log.Fatal(err)
		}

		dest, err := FilePathWalkDir(path + v.dest)

		var oldFiles []string
		var newFiles []string

		// get all common files

		for _, f := range src {
			l := strings.Split(f, v.src)

			oldFiles = append(oldFiles, l[len(l)-1])

		}

		for _, f := range dest {
			l := strings.Split(f, v.dest)

			newFiles = append(newFiles, l[len(l)-1])
		}

		println("OLD")
		for _, f := range oldFiles {
			println(f)
		}

		println("NEW")
		for _, f := range newFiles {
			println(f)
		}

		// get classes in common

		// remove classes that are not in new

		// add classes that are not in old

	}

	return nil
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
