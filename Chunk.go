package main

func getChunk(a []string, b []string) []bool {

	if len(b) > len(a) {
		c := b
		a = c
		b = a
	}

	var match []bool

	for i := range a {
		match = append(match, a[i] == b[i])
	}

	return match
}
