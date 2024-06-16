package example

import "math"

func TwoCrystalBalls(breaks []bool) int {
	l := len(breaks)
	var jmpAmount int = int(math.Sqrt(float64(l)))

	i := jmpAmount

	for ; i < l; i += jmpAmount {
		if breaks[i] {
			break
		}
	}

	i -= jmpAmount

	for j := 0; j <= jmpAmount && i < l; j += 1 {
		if breaks[i] {
			return i
		}

		i += 1
	}

	return -1
}
