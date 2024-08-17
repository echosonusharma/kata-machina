package dsa_test

import (
	"slices"
	"testing"

	dsa "github.com/echosonusharma/kata-machina/dsa/example"
)

func TestQuickSort(t *testing.T) {
	testCases := []struct {
		name string
		arr  []int
	}{
		{"expected * ", []int{10, 7, 8, 9, 1, 5}},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			arr := append([]int(nil), tt.arr...)
			dsa.QuickSort(arr)

			expected := append([]int(nil), tt.arr...)
			slices.Sort(expected)

			if !slices.Equal(arr, expected) {
				t.Errorf("got %v, want %v", arr, expected)
			}
		})
	}

}
