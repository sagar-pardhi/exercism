package hamming

import "errors"

func Distance(a, b string) (int, error) {
	hD := 0

	if len(a) != len(b) {
		return 0, errors.New("length not same")
	}

	for i := range a {
		if a[i] != b[i] {
			hD += 1
		}
	}

	return hD, nil
}
