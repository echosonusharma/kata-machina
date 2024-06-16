package example

func BS_List(haystack []int, needle int) bool {
	// lo is inclusive and hi is exclusive
	lo := 0
	hi := len(haystack)

	for lo < hi {
		var m int = lo + (hi-lo)/2
		v := haystack[m]

		if v == needle {
			return true
		} else if v > needle {
			hi = m
		} else {
			lo = m + 1
		}
	}

	return false
}
