package dsa_test

import (
	"os"
	"testing"

	dsa "github.com/echosonusharma/kata-machina/example"
	t "github.com/echosonusharma/kata-machina/types"
)

// type signature for all the dsa main functions
type linearSearch func([]int, int) bool
type binarySearch func([]int, int) bool
type twoCrystalBalls func([]bool) int
type bubbleSort func([]int)
type LinkedListF struct {
	Get       func(*t.LinkedList, uint) (string, error)
	Append    func(*t.LinkedList, string)
	Prepend   func(*t.LinkedList, string)
	Remove    func(*t.LinkedList, string) (string, error)
	InsertAt  func(*t.LinkedList, string, uint) error
	RemoveAt  func(*t.LinkedList, uint) error
	PrintList func(*t.LinkedList)
}

type dsaList struct {
	linearSearchFunc    linearSearch
	binarySearchFunc    binarySearch
	twoCrystalBallsFunc twoCrystalBalls
	bubbleSortFunc      bubbleSort
	linkedListFunc      LinkedListF
}

var DsaStore *dsaList

func TestMain(m *testing.M) {
	DsaStore = &dsaList{
		linearSearchFunc:    dsa.LinearSearchList,
		binarySearchFunc:    dsa.BS_List,
		twoCrystalBallsFunc: dsa.TwoCrystalBalls,
		bubbleSortFunc:      dsa.BubbleSort,
		linkedListFunc: LinkedListF{
			Get:       dsa.LlGet,
			Append:    dsa.LlAppend,
			Prepend:   dsa.LlPrepend,
			Remove:    dsa.LlRemove,
			InsertAt:  dsa.LlInsertAt,
			RemoveAt:  dsa.LlRemoveAt,
			PrintList: dsa.LlPrintList,
		},
	}

	code := m.Run()
	os.Exit(code)
}
