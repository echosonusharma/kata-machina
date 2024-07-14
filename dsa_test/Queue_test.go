package dsa_test

import (
	"testing"

	dsa "github.com/echosonusharma/kata-machina/dsa/example"
)

type person struct {
	name string
	age  string
}

func TestQueue(t *testing.T) {
	numQ := &dsa.Queue[int64]{}

	numQ.Enqueue(87987)
	numQ.Enqueue(343)
	numQ.Enqueue(87878767)

	if x, _ := numQ.Dequeue(); x != int64(87987) {
		t.Errorf("int64 dequeue failed")
	}

	if x := numQ.Size(); x != uint(2) {
		t.Errorf("int64 size failed")
	}

	personQ := &dsa.Queue[person]{}

	personQ.Enqueue(person{
		name: "zzz",
		age:  "zzz",
	})

	personQ.Enqueue(person{
		name: "yyy",
		age:  "yyy",
	})

	if x, _ := personQ.Dequeue(); x.name != "zzz" {
		t.Errorf("person dequeue failed")
	}
}
