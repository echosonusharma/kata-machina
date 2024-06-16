package dsa_test

import (
	"testing"

	dsa "github.com/echosonusharma/kata-machina/dsa/example"
)

func TestBS_List(t *testing.T) {
	testCases := []struct {
		name     string
		haystack []int
		needle   int
		expected bool
	}{
		{"expected true", []int{23, 34, 45}, 23, true},
		{"expected false", []int{23, 45, 55, 334, 345}, 323, false},
		{"expected true", []int{23, 34, 45, 88, 990}, 88, true},
		{"expected false", []int{-87, -1, 23, 34, 45}, 2300, false},
		{"expected false", []int{23, 34, 45}, 0, false},
		{"expected false", []int{}, 33, false},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			res := dsa.BS_List(tt.haystack, tt.needle)

			// t.Log("input - ", tt.haystack, tt.needle)

			if res != tt.expected {
				t.Errorf("got %t, want %t", res, tt.expected)
			}
		})
	}

}
