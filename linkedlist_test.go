package linkedlist

import (
	"fmt"
	"strings"
	"testing"
)

func TestAppendOnly(t *testing.T) {
	expected := 42
	t.Run(fmt.Sprintf("%#v to a linkedlist", expected), func(t *testing.T) {
		// GIVEN
		list := LinkedList{}
		assertEmpty(t, list)
		// WHEN
		list.Append(&Node{Data: expected})
		// THEN
		assertNode(t, "Head", list.Head, expected)
		if list.Head.next != nil {
			t.Error("Next pointer for the List Head should be nil")
		}
		expected2 := expected + 1
		// WHEN
		list.Append(&Node{Data: expected2})
		// THEN
		assertNode(t, "Head", list.Head, expected)
		assertNode(t, "Head.next", list.Head.next, expected2)
	})
}

func TestAppendValue(t *testing.T) {
	expected := -1
	t.Run(fmt.Sprintf("%#v to a linkedlist", expected), func(t *testing.T) {
		list := LinkedList{}
		assertEmpty(t, list)
		list.AppendValue(expected)
		assertNode(t, "Head", list.Head, expected)
		if list.Head.next != nil {
			t.Error("Next pointer for the List Head should be nil")
		}
		expected2 := expected - 1
		list.AppendValue(expected2)
		assertNode(t, "Head", list.Head, expected)
		assertNode(t, "Head.next", list.Head.next, expected2)
	})
}

func TestPrependOnly(t *testing.T) {
	expected := 42
	t.Run(fmt.Sprintf("%#v to a linkedlist", expected), func(t *testing.T) {
		list := LinkedList{}
		assertEmpty(t, list)
		list.Prepend(&Node{Data: expected})
		assertNode(t, "Head", list.Head, expected)
		if list.Head.next != nil {
			t.Error("Next pointer for the List Head should be nil")
		}
		expected2 := expected + 1
		list.Prepend(&Node{Data: expected2})
		assertNode(t, "Head", list.Head, expected2)
		assertNode(t, "Head.next", list.Head.next, expected)
	})
}

func TestPrependValue(t *testing.T) {
	expected := 0
	list := LinkedList{}
	assertEmpty(t, list)
	list.PrependValue(expected)
	assertNode(t, "Head", list.Head, expected)
	if list.Head.next != nil {
		t.Error("Next pointer for the List Head should be nil")
	}
	for i := expected + 1; i < 1002; i = i + 100 {
		t.Run(fmt.Sprintf("%#v to a linkedlist", i), func(t *testing.T) {
			previous := list.Head.Data
			list.PrependValue(i)
			assertNode(t, "Head", list.Head, i)
			assertNode(t, "Head.next", list.Head.next, previous)
		})
	}
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

func TestRemoveSuccess(t *testing.T) {
	var testCases = []struct {
		nodeValues        []int
		count             int
		expectedLength    int
		expectedLastValue int
	}{
		{nodeValues: []int{-1, 42}, count: 1, expectedLength: 1, expectedLastValue: -1},
		{nodeValues: []int{100, -1, 42}, count: 1, expectedLength: 2, expectedLastValue: -1},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Reducing %v from linkedlist %#v", tc.count, tc.nodeValues), func(t *testing.T) {
			list := createList(tc.nodeValues)
			list.Reduce()
			if tc.expectedLength != list.Length() {
				t.Error("\nExpected list length:", tc.expectedLength, "\nReceived list length: ", list.Length())
			}
			lastNode := list.Get(tc.expectedLength - 1)
			if tc.expectedLastValue != lastNode.Data {
				t.Error("\nExpected node with value:", tc.expectedLastValue, "\nReceived Node: ", lastNode)
			}
		})
	}
}

func TestRemoveEdgeCases(t *testing.T) {
	t.Run(fmt.Sprintf("Reducing an empty linkedlist"), func(t *testing.T) {
		list := LinkedList{}
		assertEmpty(t, list)
		list.Reduce()
		assertEmpty(t, list)

	})
	t.Run(fmt.Sprintf("Reducing a linkedlist beyond empty"), func(t *testing.T) {
		list := createList([]int{100, 99})
		list.Reduce()
		list.Reduce()
		list.Reduce()
		list.Reduce()
		assertEmpty(t, list)
	})
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

func assertEmpty(t *testing.T, list LinkedList) {
	t.Helper()
	if 0 != list.Length() {
		t.Error("\nExpected list length:", 0, "\nReceived list length: ", list.Length())
	}
	if list.Head != nil {
		t.Error("Head pointer for the List should still be nil")
	}
}

func assertNode(t *testing.T, hint string, n *Node, expectedData int) {
	t.Helper()
	if n == nil {
		t.Error("node pointer for", hint, "in the List is unexpectedly nil")
	}
	if expectedData != n.Data {
		t.Error("\n", hint, "data expected:", expectedData, "\nReceived: ", n.Data)
	}
}
