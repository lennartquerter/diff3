package src

import (
	"io/ioutil"
	"strings"
	"fmt"
	"CopyModels/src/Helpers"
)

func MergeFile(bd string, past string, src string, dest string, debug bool) ([]string, error) {

	pf, pastErr := ioutil.ReadFile(bd + past)

	if pastErr != nil {
		var readErr error
		pf, readErr = ioutil.ReadFile(bd + dest)
		if readErr != nil {
			return nil, pastErr
		}
	}

	srcf, err := ioutil.ReadFile(bd + src)

	if err != nil {
		return nil, err
	}

	destf, err := ioutil.ReadFile(bd + dest)

	if err != nil {
		return nil, err
	}

	pastLines := strings.Split(string(pf), "\n")
	srcLines := strings.Split(string(srcf), "\n")
	destLines := strings.Split(string(destf), "\n")

	srcChunk := getChunk(pastLines, srcLines)
	destChunk := getChunk(pastLines, destLines)

	length := Helpers.Max(srcChunk, destChunk)

	var ml []string

	// Implementation

	var conflicts []string

	for i := 0; i <= length; i++ {

		// if Line is the same as old version
		if (i < len(pastLines) && i < len(srcLines) && i < len(destLines)) &&
			(pastLines[i] == destLines[i] && destLines[i] == srcLines[i]) {

			if debug {
				pastLines[i] = "0" + pastLines[i]
			}

			ml = append(ml, pastLines[i])

		} else {
			// unstable
			if i < len(srcChunk) && srcChunk[i] {
				// change in A

				if i < len(destChunk) && !destChunk[i] {
					// change in both / Conflict

					conflicts = append(conflicts, fmt.Sprintf("Conflict in line %d \n File: %s%s", i, bd, src))
					ml = append(ml, []string{">>>>>>>>>> Version 2 (New)",
						destLines[i],
						"==========",
						srcLines[i],
						"<<<<<<<<<< Version 1 (old)"}...)
				} else {
					// only change in src

					if debug {
						srcLines[i] = "1" + srcLines[i]
					}
					ml = append(ml, srcLines[i])
				}

			} else if i < len(destChunk) && destChunk[i] {
				// change in B

				if i < len(srcChunk) && !srcChunk[i] {
					// change in both / Conflict
					conflicts = append(conflicts, fmt.Sprintf("Conflict in line %d \n File: %s%s", i, bd, src))
					ml = append(ml, []string{">>>>>>>>>> Version 2 (New)",
						srcLines[i],
						"==========",
						destLines[i],
						"<<<<<<<<<< Version 1 (old)"}...)
				} else {
					// only change in dest
					if debug {
						destLines[i] = "2" + destLines[i]
					}
					ml = append(ml, destLines[i])
				}
			}
		}
	}

	return ml, nil
}

