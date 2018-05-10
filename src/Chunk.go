package src

func getChunk(a []string, b []string) []bool {

	if len(b) > len(a) {
		c := b
		a = c
		b = a
	}

	var match []bool

	for i := range a {
		if i < len(b) {
			match = append(match, a[i] == b[i])
		} else {
			match = append(match, false)
		}

	}

	return match
}
