package dsa_test

import (
	"testing"

	dsa "github.com/echosonusharma/kata-machina/dsa/example"
)

type testCase struct {
	name          string
	operation     func(*dsa.LinkedList[int64])
	index         uint
	expectedValue int64
	expectErr     bool
}

func TestSinglyLinkedList(t *testing.T) {
	numList := &dsa.LinkedList[int64]{}
	initialValues := []int64{324, 2345, 876, 6784, 467473, 234}
	for _, v := range initialValues {
		numList.Append(v)
	}

	testCases := []testCase{
		{name: "Initial Append", operation: func(l *dsa.LinkedList[int64]) {}, index: 1, expectedValue: initialValues[1]},
		{name: "Initial Append", operation: func(l *dsa.LinkedList[int64]) {}, index: 5, expectedValue: initialValues[5]},
		{name: "Prepend", operation: func(l *dsa.LinkedList[int64]) { l.Prepend(99) }, index: 0, expectedValue: 99},
		{name: "Append After Prepend", operation: func(l *dsa.LinkedList[int64]) { l.Append(879) }, index: 7, expectedValue: 879},
		{name: "Append Another Value", operation: func(l *dsa.LinkedList[int64]) { l.Append(991) }, index: 8, expectedValue: 991},
		{name: "Remove Element", operation: func(l *dsa.LinkedList[int64]) { l.Remove(991) }, index: 8, expectedValue: 0, expectErr: true},
		{name: "Insert At Specific Index", operation: func(l *dsa.LinkedList[int64]) { l.InsertAt(991, 7) }, index: 7, expectedValue: 991},
		{name: "Remove At Specific Index", operation: func(l *dsa.LinkedList[int64]) { l.RemoveAt(7) }, index: 7, expectedValue: 879},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.operation(numList)
			value, err := numList.Get(tc.index)
			if err != nil && !tc.expectErr {
				t.Errorf("Error accessing index %d: %v", tc.index, err)
			}
			if value != tc.expectedValue {
				t.Errorf("Expected %d at index %d but got %d", tc.expectedValue, tc.index, value)
			}
		})
	}
}
