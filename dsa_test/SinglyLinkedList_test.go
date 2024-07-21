package dsa_test

import (
	"testing"

	dsa "github.com/echosonusharma/kata-machina/dsa/example"
)

func TestSinglyLinkedList(t *testing.T) {
	numList := &dsa.LinkedList[int]{}

	listToAppend := []int{324, 2345, 876, 6784, 467473, 234}

	for _, v := range listToAppend {
		numList.Append(v)
	}

	if x, _ := numList.Get(0); x != listToAppend[0] {
		t.Errorf("Get failed")
	}

	if x, _ := numList.Get(4); x != listToAppend[4] {
		t.Errorf("Get failed")
	}

	y := 99

	numList.Prepend(y)
	if x, _ := numList.Get(0); x != y {
		t.Errorf("Get failed")
	}

	if x, _ := numList.Get(4); x != listToAppend[3] {
		t.Errorf("Get failed")
	}

	numList.Remove()

	if x, _ := numList.Get(5); x != listToAppend[4] {
		t.Errorf("Get failed")
	}

	numList.RemoveAt(2)

	numList.InsertAt(879, 2)

	if x, _ := numList.Get(2); x != 879 {
		t.Errorf("Get failed")
	}
}
