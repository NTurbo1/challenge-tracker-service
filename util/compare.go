package util

import (
	"cmp"
	"fmt"
)

func Max[T cmp.Ordered](v ...T) (T, error) {
	var maxVal T
	vLen := len(v)
	if vLen == 0 {
		return maxVal, fmt.Errorf(
			"Can't find the max value from an empty slice, motherfucker! Think, bro, THINK!!!",
		)
	}

	maxVal = v[0]
	for i := 1; i < vLen; i++ {
		if v[i] > maxVal {
			maxVal = v[i]
		}
	}

	return maxVal, nil
}
