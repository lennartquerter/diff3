package main

import (
	"log"
	"os"
	"strings"
	"CopyModels/Util"
	"io/ioutil"
)

func MergeDir(path string, list []Model, debug bool) error {

	for _, v := range list {
		src, err := Util.FilePathWalkDir(path + v.src)

		if err != nil {
			log.Fatal(err)
		}

		dest, err := Util.FilePathWalkDir(path + v.dest)

		var oldFiles []string
		var newFiles []string

		// get all common files

		for _, f := range src {
			l := strings.Split(f, v.src)

			newFiles = append(newFiles, l[len(l)-1])
		}

		for _, f := range dest {
			l := strings.Split(f, v.dest)

			oldFiles = append(oldFiles, l[len(l)-1])
		}

		// get classes in common
		merge := Util.Common(newFiles, oldFiles)

		for _, m := range merge {
			file, err := MergeFile(path, v.past, v.src+m, v.dest+m, debug)
			if err != nil {
				return err
			}

			f, err := os.Create(path + v.dest + m)

			f.Write([]byte(strings.Join(file, "\n")))
			if err != nil {
				return err
			}

			f.Close()

			if err != nil {
				return err
			}

		}

		add := Util.Except(newFiles, oldFiles)
		remove := Util.Except(oldFiles, newFiles)

		// remove classes that are not in new

		for _, r := range remove {
			os.Remove(path + v.dest + r)
			ok, err := Util.IsEmpty(path + v.dest)

			if err != nil {
				return err
			}

			if ok {
				os.Remove(path + v.dest)
			}
		}

		// add classes that are not in old

		for _, a := range add {

			dat, err := ioutil.ReadFile(path + v.src + a)

			if err != nil {
				return err
			}

			p := strings.Split(a, "/")

			pa := path + v.dest + strings.Join(p[:len(p)-1], "/") + "/"

			os.MkdirAll(pa, os.ModePerm)

			f, err := os.Create(path + v.dest + a)

			f.Write(dat)

			if err != nil {
				return err
			}

			f.Close()

			if err != nil {
				return err
			}
		}
	}

	return nil
}
