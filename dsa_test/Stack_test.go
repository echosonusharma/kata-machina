package dsa_test

import (
	"testing"

	dsa "github.com/echosonusharma/kata-machina/dsa/example"
)

func TestStack(t *testing.T) {
	strS := &dsa.Stack[string]{}

	strS.Push("a")
	strS.Push("a")
	strS.Push("a")
	strS.Push("a")
	strS.Push("a")

	if x := strS.Size(); x != uint(5) {
		t.Errorf("string size failed")
	}

	strS.Push("x")
	strS.Push("y")
	strS.Push("z")

	if x := strS.Pop(); x != "z" {
		t.Errorf("string Pop failed")
	}
	if x := strS.Pop(); x != "y" {
		t.Errorf("string Pop failed")
	}

	if x := strS.Peek(); x != "x" {
		t.Errorf("string Pop failed")
	}

	if x := strS.Pop(); x != "x" {
		t.Errorf("string Pop failed")
	}
}
