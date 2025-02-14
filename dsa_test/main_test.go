package dsa_test

import (
	"os"
	"testing"

	dsa "github.com/echosonusharma/kata-machina/example"
)

// type signature for all the dsa main functions
type linearSearch func([]int, int) bool
type binarySearch func([]int, int) bool
type twoCrystalBalls func([]bool) int
type bubbleSort func([]int)

type dsaList struct {
	linearSearchFunc    linearSearch
	binarySearchFunc    binarySearch
	twoCrystalBallsFunc twoCrystalBalls
	bubbleSortFunc      bubbleSort
}

var DsaStore *dsaList

func TestMain(m *testing.M) {
	DsaStore = &dsaList{
		linearSearchFunc:    dsa.LinearSearchList,
		binarySearchFunc:    dsa.BS_List,
		twoCrystalBallsFunc: dsa.TwoCrystalBalls,
		bubbleSortFunc:      dsa.BubbleSort,
	}

	code := m.Run()
	os.Exit(code)
}
