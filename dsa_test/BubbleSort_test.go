package dsa_test

import (
	"slices"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	testCases := []struct {
		name string
		arr  []int
	}{
		{"expected * ", []int{23, 34, 45, 1}},
		{"expected * ", []int{23, 334, 45, 345, 55}},
		{"expected * ", []int{23, 34, 45, 88, 990}},
		{"expected * ", []int{-87, -1, 0, 23, 34, 45}},
		{"expected * ", []int{45, 567, 23, 34, 45}},
		{"expected * ", []int{}},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			arr := append([]int(nil), tt.arr...)
			DsaStore.bubbleSortFunc(arr)

			expected := append([]int(nil), tt.arr...)
			slices.Sort(expected)

			if !slices.Equal(arr, expected) {
				t.Errorf("got %v, want %v", arr, expected)
			}
		})
	}
}
