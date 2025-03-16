package dsa_test

import (
	"testing"

	ty "github.com/echosonusharma/kata-machina/types"
)

type testCase struct {
	name          string
	operation     func(*ty.LinkedList)
	index         uint
	expectedValue string
	expectErr     bool
}

func TestSinglyLinkedList(t *testing.T) {
	strList := &ty.LinkedList{}
	initialValues := []string{"sam", "james", "micky", "donald", "buddha", "jack"}
	for _, v := range initialValues {
		DsaStore.linkedListFunc.Append(strList, v)
	}

	testCases := []testCase{
		{name: "Initial Append", operation: func(l *ty.LinkedList) {}, index: 1, expectedValue: initialValues[1]},
		{name: "Initial Append", operation: func(l *ty.LinkedList) {}, index: 5, expectedValue: initialValues[5]},
		{name: "Prepend", operation: func(l *ty.LinkedList) { DsaStore.linkedListFunc.Prepend(l, "99") }, index: 0, expectedValue: "99"},
		{name: "Append After Prepend", operation: func(l *ty.LinkedList) { DsaStore.linkedListFunc.Append(l, "879") }, index: 7, expectedValue: "879"},
		{name: "Append Another Value", operation: func(l *ty.LinkedList) { DsaStore.linkedListFunc.Append(l, "991") }, index: 8, expectedValue: "991"},
		{name: "Remove Element", operation: func(l *ty.LinkedList) { DsaStore.linkedListFunc.Remove(l, "991") }, index: 8, expectedValue: "", expectErr: true},
		{name: "Insert At Specific Index", operation: func(l *ty.LinkedList) { DsaStore.linkedListFunc.InsertAt(l, "991", 7) }, index: 7, expectedValue: "991"},
		{name: "Remove At Specific Index", operation: func(l *ty.LinkedList) { DsaStore.linkedListFunc.RemoveAt(l, 7) }, index: 7, expectedValue: "879"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.operation(strList)
			value, err := DsaStore.linkedListFunc.Get(strList, tc.index)
			if err != nil && !tc.expectErr {
				t.Errorf("Error accessing index %d: %v", tc.index, err)
			}
			if value != tc.expectedValue {
				t.Errorf("Expected %s at index %d but got %s", tc.expectedValue, tc.index, value)
			}
		})
	}
}
