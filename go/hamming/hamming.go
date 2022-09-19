package hamming

import "errors"

func Distance(a, b string) (int, error) {
	dist := 0
	if len(a) != len(b) {
		return 0, errors.New("Input Strings must be same length")
	}
	for i, _ := range a {
		if a[i] != b[i] {
			dist += 1
		}
	}
	return dist, nil
}
