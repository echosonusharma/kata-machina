package example

func LinearSearchList(haystack []int, needle int) bool {
	for _, v := range haystack {
		if v == needle {
			return true
		}
	}

	return false
}
