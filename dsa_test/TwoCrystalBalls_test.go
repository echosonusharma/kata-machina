package dsa_test

import (
	"fmt"
	"math/rand/v2"
	"testing"

	dsa "github.com/echosonusharma/kata-machina/dsa/example"
)

type TwoCrystalBallsTestTable struct {
	name     string
	breaks   []bool
	expected int
}

func newTwoCrystalBallsTestCase() TwoCrystalBallsTestTable {
	maxArr := rand.IntN(10_000-1) + 1
	breaksArr := make([]bool, maxArr)

	foundIdx := rand.IntN(maxArr)
	for i := foundIdx; i < len(breaksArr); i++ {
		breaksArr[i] = true
	}

	newCase := TwoCrystalBallsTestTable{
		name:     fmt.Sprintf("expected %d", foundIdx),
		breaks:   breaksArr,
		expected: foundIdx,
	}

	return newCase
}

func TestTwoCrystalBalls(t *testing.T) {
	testCases := make([]TwoCrystalBallsTestTable, 100)

	for i := 0; i < 100; i++ {
		testCases[i] = newTwoCrystalBallsTestCase()
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			res := dsa.TwoCrystalBalls(tt.breaks)

			if res != tt.expected {
				t.Log("input breaks arr - ", tt.breaks)
				t.Errorf("got %d, want %d", res, tt.expected)
			}
		})
	}

}
