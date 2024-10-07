package dsa_test

import (
	"slices"
	"testing"

	dsa "github.com/echosonusharma/kata-machina/dsa/example"
)

func TestPreOrderSearch(t *testing.T) {
	bt1 := &dsa.BinaryNode[int]{
		Value: 20,
		Right: &dsa.BinaryNode[int]{
			Value: 50,
			Right: &dsa.BinaryNode[int]{
				Value: 100,
				Right: nil,
				Left:  nil,
			},
			Left: &dsa.BinaryNode[int]{
				Value: 30,
				Right: &dsa.BinaryNode[int]{
					Value: 45,
					Right: nil,
					Left:  nil,
				},
				Left: &dsa.BinaryNode[int]{
					Value: 29,
					Right: nil,
					Left:  nil,
				},
			},
		},
		Left: &dsa.BinaryNode[int]{
			Value: 10,
			Right: &dsa.BinaryNode[int]{
				Value: 15,
				Right: nil,
				Left:  nil,
			},
			Left: &dsa.BinaryNode[int]{
				Value: 5,
				Right: &dsa.BinaryNode[int]{
					Value: 7,
					Right: nil,
					Left:  nil,
				},
				Left: nil,
			},
		},
	}

	res := bt1.PreOrderSearch()
	exp := []int{
		20,
		10,
		5,
		7,
		15,
		50,
		30,
		29,
		45,
		100,
	}

	ck := slices.Compare(res, exp)

	if ck != 0 {
		t.Errorf("invalid traversal - got result > %v expected > %v", res, exp)
	}
}
