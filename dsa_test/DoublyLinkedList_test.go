package dsa_test

import (
	"testing"

	dsa "github.com/echosonusharma/kata-machina/dsa/example"
)

func TestDoublyLinkedList(t *testing.T) {
	numList := &dsa.DoublyLinkedList[string]{}
	initialValues := []string{
		"Alfa",
		"Bravo",
		"Charlie",
		"Delta",
		"Echo",
		"Foxtrot",
		"Golf",
		"Hotel",
		"India",
	}

	for _, v := range initialValues {
		numList.Append(v)
	}

	testCases := []struct {
		name          string
		operation     func(*dsa.DoublyLinkedList[string])
		index         uint
		expectedValue string
		expectErr     bool
	}{
		{name: "Initial Append", operation: func(l *dsa.DoublyLinkedList[string]) {}, index: 1, expectedValue: initialValues[1]},
		{name: "Initial Append", operation: func(l *dsa.DoublyLinkedList[string]) {}, index: 5, expectedValue: initialValues[5]},
		{name: "Prepend", operation: func(l *dsa.DoublyLinkedList[string]) { l.Prepend("Juliett") }, index: 0, expectedValue: "Juliett"},
		{name: "Append After Prepend", operation: func(l *dsa.DoublyLinkedList[string]) { l.Append("Kilo") }, index: 10, expectedValue: "Kilo"},
		{name: "Append Another Value", operation: func(l *dsa.DoublyLinkedList[string]) { l.Append("Lima") }, index: 11, expectedValue: "Lima"},
		{name: "Remove Element", operation: func(l *dsa.DoublyLinkedList[string]) { l.Remove("Mike") }, index: 12, expectedValue: "", expectErr: true},
		{name: "Insert At Specific Index", operation: func(l *dsa.DoublyLinkedList[string]) { l.InsertAt("November", 7) }, index: 7, expectedValue: "November"},
		{name: "Remove At Specific Index", operation: func(l *dsa.DoublyLinkedList[string]) { l.RemoveAt(7) }, index: 7, expectedValue: "Golf"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.operation(numList)
			value, err := numList.Get(tc.index)
			if err != nil && !tc.expectErr {
				t.Errorf("Error accessing index %d: %v", tc.index, err)
			}
			if value != tc.expectedValue {
				t.Errorf("Expected %s at index %d but got %s", tc.expectedValue, tc.index, value)
			}
		})
	}
}
