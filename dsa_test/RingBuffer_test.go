package dsa_test

import (
	"fmt"
	"testing"

	dsa "github.com/echosonusharma/kata-machina/dsa/example"
)

func TestRingBuffer(t *testing.T) {
	rb := dsa.NewRingBuffer[string](5)

	rb.Push("entry-1")
	rb.Push("entry-2")
	rb.Push("entry-3")
	rb.Push("entry-4")
	rb.Push("entry-5")
	rb.Push("entry-6")
	rb.Push("entry-7")
	rb.Push("entry-8")
	rb.Push("entry-9")

	x, _ := rb.Pop()
	fmt.Println(x)

	x1, _ := rb.Pop()
	fmt.Println(x1)

	rb.Pop()

	x2, _ := rb.Pop()
	fmt.Println(x2)

	rb.Print()
}
