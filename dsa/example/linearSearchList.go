package example

func LinearSearchList(heystack []int, needle int) bool {
	for _, v := range heystack {
		if v == needle {
			return true
		}
	}

	return false
}
