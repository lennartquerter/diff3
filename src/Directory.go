package src

import (
	"os"
	"strings"
	"io/ioutil"
	"CopyModels/src/Helpers"
	exp "CopyModels/src/exported"
)

func MergeDir(path string, sourceFiles []exp.Dir, historyPath string, debug bool) error {

	for _, v := range sourceFiles {
		src, err := Helpers.FilePathWalkDir(path + v.Source)

		if err != nil {
			return err
		}

		dest, err := Helpers.FilePathWalkDir(path + v.Destination)

		if err != nil {
			return err
		}

		var oldFiles []string
		var newFiles []string

		// get all common files

		for _, f := range src {
			l := strings.Split(f, v.Source)

			newFiles = append(newFiles, l[len(l)-1])
		}

		for _, f := range dest {
			l := strings.Split(f, v.Destination)

			oldFiles = append(oldFiles, l[len(l)-1])
		}

		// get classes in common
		merge := Helpers.Common(newFiles, oldFiles)

		for _, m := range merge {
			file, err := MergeFile(path, historyPath+v.Source+m, v.Source+m, v.Destination+m, debug)
			if err != nil {
				return err
			}

			p := strings.Split(m, "/")

			pa := path + v.Destination + strings.Join(p[:len(p)-1], "/") + "/"

			ha := historyPath + v.Destination + strings.Join(p[:len(p)-1], "/") + "/"

			os.MkdirAll(pa, os.ModePerm)
			os.MkdirAll(ha, os.ModePerm)

			f, err := os.Create(path + v.Destination + m)
			h, err := os.Create(historyPath + v.Destination + m)

			f.Write([]byte(strings.Join(file, "\n")))
			h.Write([]byte(strings.Join(file, "\n")))
			if err != nil {
				return err
			}

			f.Close()
			h.Close()

			if err != nil {
				return err
			}

		}

		add := Helpers.Except(newFiles, oldFiles)
		remove := Helpers.Except(oldFiles, newFiles)

		// remove classes that are not in new

		for _, r := range remove {
			os.Remove(path + v.Destination + r)
			os.Remove(historyPath + v.Destination + r)

			ok, err := Helpers.IsEmpty(path + v.Destination)

			if err != nil {
				return err
			}

			if ok {
				os.Remove(path + v.Destination)
			}

			ok, err = Helpers.IsEmpty(historyPath + v.Destination)

			if err != nil {
				return err
			}

			if ok {
				os.Remove(path + v.Destination)
			}
		}

		// add classes that are not in old

		for _, a := range add {

			dat, err := ioutil.ReadFile(path + v.Source + a)

			if err != nil {
				return err
			}

			p := strings.Split(a, "/")

			pa := path + v.Destination + strings.Join(p[:len(p)-1], "/") + "/"

			ha := historyPath + v.Destination + strings.Join(p[:len(p)-1], "/") + "/"

			os.MkdirAll(pa, os.ModePerm)
			os.MkdirAll(ha, os.ModePerm)

			f, err := os.Create(path + v.Destination + a)
			h, err := os.Create(historyPath + v.Destination + a)

			f.Write(dat)
			h.Write(dat)

			if err != nil {
				return err
			}

			f.Close()
			h.Close()

			if err != nil {
				return err
			}
		}
	}

	return nil
}
