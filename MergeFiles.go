package main

import (
	"io/ioutil"
	"strings"
	"fmt"
)

func MergeFile(bd string, past string, v1 string, v2 string, debug bool) ([]string, error) {

	pf, err := ioutil.ReadFile(bd + past)

	if err != nil {
		pf, err = ioutil.ReadFile(bd + v1)
		if err != nil {
			return nil, err
		}
	}

	v1f, err := ioutil.ReadFile(bd + v1)

	if err != nil {
		return nil, err
	}

	v2f, err := ioutil.ReadFile(bd + v2)

	if err != nil {
		return nil, err
	}

	pastLines := strings.Split(string(pf), "\n")
	version1Lines := strings.Split(string(v1f), "\n")
	version2Lines := strings.Split(string(v2f), "\n")

	aChunk := getChunk(pastLines, version1Lines)
	bChunk := getChunk(pastLines, version2Lines)

	length := Max(aChunk, bChunk)

	var ml []string

	// Implementation

	var conflicts []string

	for i := 0; i <= length; i++ {

		// if Line is the same as old version
		if (i < len(pastLines) && i < len(version1Lines) && i < len(version2Lines)) &&
			(pastLines[i] == version2Lines[i] && version2Lines[i] == version1Lines[i]) {

			if debug {
				pastLines[i] = "0" + pastLines[i]
			}

			ml = append(ml, pastLines[i])

		} else {
			// unstable
			if i < len(aChunk) && aChunk[i] {
				// change in A

				if i < len(bChunk) && !bChunk[i] {
					// change in both / Conflict

					conflicts = append(conflicts, fmt.Sprintf("Conflict in line %d \n File: %s%s", i, bd, v1))
					ml = append(ml, []string{">>>>>>>>>> Version 2 (New)",
						version2Lines[i],
						"==========",
						version1Lines[i],
						"<<<<<<<<<< Version 1 (old)"}...)
				} else {
					// only change in VER1

					if debug {
						version1Lines[i] = "1" + version1Lines[i]
					}
					ml = append(ml, version1Lines[i])
				}

			} else if i < len(bChunk) && bChunk[i] {
				// change in B
				if i < len(aChunk) && !aChunk[i] {
					// change in both / Conflict
					conflicts = append(conflicts, fmt.Sprintf("Conflict in line %d \n File: %s%s", i, bd, v2))
					ml = append(ml, []string{">>>>>>>>>> Version 2 (New)",
						version2Lines[i],
						"==========",
						version1Lines[i],
						"<<<<<<<<<< Version 1 (old)"}...)
				} else {
					// only change in VER2
					if debug {
						version2Lines[i] = "2" + version2Lines[i]
					}
					ml = append(ml, version2Lines[i])
				}
			}
		}
	}

	return ml, nil
}

func Max(list1 []bool, list2 []bool) (int) {
	l := 0
	for _, e := range []int{len(list1), len(list2)} {
		if e > l {
			l = e
		}
	}

	return l
}
