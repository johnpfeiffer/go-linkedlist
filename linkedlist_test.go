package linkedlist

import (
	"fmt"
	"strings"
	"testing"
)

func TestAppendOnly(t *testing.T) {
	expected := 42
	t.Run(fmt.Sprintf("Append %#v to a linkedlist", expected), func(t *testing.T) {
		list := LinkedList{}
		if list.Head != nil {
			t.Error("Head pointer for the List should still be nil")
		}
		list.Append(&Node{Data: expected})
		if list.Head == nil {
			t.Error("Head pointer for the List is nil")
		}
		if expected != list.Head.Data {
			t.Error("\nExpected:", expected, "\nReceived: ", list.Head.Data)
		}
		if list.Head.next != nil {
			t.Error("Next pointer for the List should be nil")
		}

		expected2 := expected + 1
		list.Append(&Node{Data: expected2})
		if list.Head == nil {
			t.Error("Head pointer for the List is nil")
		}
		if list.Head.next == nil {
			t.Error("Next pointer for the List should not be nil")
		}
		if expected2 != list.Head.next.Data {
			t.Error("\nExpected:", expected2, "\nReceived: ", list.Head.Data)
		}
	})
}

func TestAppendValue(t *testing.T) {
	expected := -1
	t.Run(fmt.Sprintf("AppendValue %#v to a linkedlist", expected), func(t *testing.T) {
		list := LinkedList{}
		if list.Head != nil {
			t.Error("Head pointer for the List should still be nil")
		}
		list.AppendValue(expected)
		if list.Head == nil {
			t.Error("Head pointer for the List is nil")
		}
		if expected != list.Head.Data {
			t.Error("\nExpected:", expected, "\nReceived: ", list.Head.Data)
		}
		if list.Head.next != nil {
			t.Error("Next pointer for the List should be nil")
		}

		expected2 := expected - 1
		list.AppendValue(expected2)
		if list.Head == nil {
			t.Error("Head pointer for the List is nil")
		}
		if list.Head.next == nil {
			t.Error("Next pointer for the List should not be nil")
		}
		if expected2 != list.Head.next.Data {
			t.Error("\nExpected:", expected2, "\nReceived: ", list.Head.next.Data)
		}
	})
}

func TestLength(t *testing.T) {
	var testCases = []struct {
		nodeValues []int
		expected   int
	}{
		{nodeValues: []int{}, expected: 0},
		{nodeValues: []int{1}, expected: 1},
		{nodeValues: []int{-1, 42}, expected: 2},
		{nodeValues: []int{-1, 100, 9}, expected: 3},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Appending %#v to a linkedlist", tc.nodeValues), func(t *testing.T) {
			list := createList(tc.nodeValues)
			assertLengthEqual(t, tc.expected, list.Length())
		})
	}
}

func TestDisplay(t *testing.T) {
	t.Run(fmt.Sprintf("Displaying a linkedlist"), func(t *testing.T) {
		list := LinkedList{}
		list.AppendValue(42)
		result := list.Display()
		expected := "&{<nil> 42}"
		if expected != strings.TrimSpace(result) {
			fmt.Println(len(expected))
			fmt.Println(len(result))
			t.Error("\nExpected:", expected, "\nReceived: ", result)
		}
	})
}

func TestFindFound(t *testing.T) {
	var testCases = []struct {
		nodeValues []int
		location   int
		expected   int
	}{
		{nodeValues: []int{1}, expected: 1},
		{nodeValues: []int{-1, 42}, expected: -1},
		{nodeValues: []int{-1, 42}, expected: 42},
		{nodeValues: []int{-1, 100, 9}, expected: 100},
		// TODO: a case where we show only retrieving the first result?
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Finding %v from linkedlist %#v", tc.expected, tc.nodeValues), func(t *testing.T) {
			list := createList(tc.nodeValues)
			result := list.Find(tc.expected)
			if tc.expected != result.Data {
				t.Error("\nExpected node with value:", tc.expected, "\nReceived Node: ", result)
			}
		})
	}
}

func TestFindNotFound(t *testing.T) {
	var testCases = []struct {
		nodeValues []int
		target     int
	}{
		{nodeValues: []int{}, target: 1},
		{nodeValues: []int{1}, target: 2},
		{nodeValues: []int{-1, 42}, target: -2},
		{nodeValues: []int{-1, 100, 9}, target: 101},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Should not find %v in linkedlist %#v", tc.target, tc.nodeValues), func(t *testing.T) {
			list := createList(tc.nodeValues)
			result := list.Find(tc.target)
			if result != nil {
				t.Error("\nExpected to not find the value:", tc.target, "\nReceived Node: ", result)
			}
		})
	}
}

func TestGetSuccess(t *testing.T) {
	var testCases = []struct {
		nodeValues []int
		location   int
		expected   int
	}{
		{nodeValues: []int{1}, location: 0, expected: 1},
		{nodeValues: []int{-1, 42}, location: 0, expected: -1},
		{nodeValues: []int{-1, 42}, location: 1, expected: 42},
		{nodeValues: []int{-1, 100, 9}, location: 0, expected: -1},
		{nodeValues: []int{-1, 100, 9}, location: 1, expected: 100},
		{nodeValues: []int{-1, 100, 9}, location: 2, expected: 9},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Getting index %v from linkedlist %#v", tc.location, tc.nodeValues), func(t *testing.T) {
			list := createList(tc.nodeValues)
			result := list.Get(tc.location)
			if tc.expected != result.Data {
				t.Error("\nExpected node with value:", tc.expected, "\nReceived Node: ", result)
			}
		})
	}
}

func TestGetErrors(t *testing.T) {
	var testCases = []struct {
		nodeValues []int
		location   int
		expected   int
	}{
		{nodeValues: []int{}, location: 0},
		{nodeValues: []int{-1, 42}, location: -1},
		{nodeValues: []int{-1, 42}, location: 2},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Getting index %v from linkedlist %#v", tc.location, tc.nodeValues), func(t *testing.T) {
			list := createList(tc.nodeValues)
			result := list.Get(tc.location)
			if result != nil {
				t.Error("\nExpected nil but received Node: ", result)
			}
		})
	}
}

// HELPER FUNCTIONS
func createList(a []int) LinkedList {
	list := LinkedList{}
	for _, v := range a {
		list.AppendValue(v)
	}
	return list
}

func assertLengthEqual(t *testing.T, expected int, result int) {
	t.Helper()
	if expected != result {
		t.Error("\nExpected Length:", expected, "\nReceived Length: ", result)
	}
}
