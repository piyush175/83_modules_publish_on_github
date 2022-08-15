package geometry

import "errors"

func CubeVol(n int) (int, error) {
	if n != 0 {
		return n * n * n, nil
	}
	return 0, errors.New("zero edge not allowed")
}
