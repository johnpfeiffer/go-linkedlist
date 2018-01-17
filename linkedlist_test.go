package linkedlist

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"testing"
)

// TODO: probably could be improved by leveraging https://github.com/stretchr/testify

var testCases = []struct {
	dataValues []int
	length     int
}{
	{dataValues: []int{}, length: 0},
	{dataValues: []int{1}, length: 1},
	{dataValues: []int{-1, 42}, length: 2},
	{dataValues: []int{-1, 100, 9}, length: 3},
	{dataValues: []int{1, 999, 99, 10001}, length: 4},
}

func TestAppendOnly(t *testing.T) {
	expected := 42
	t.Run(fmt.Sprintf("%#v to a linkedlist", expected), func(t *testing.T) {
		// GIVEN
		list := LinkedList{}
		assertEmpty(t, list)
		// WHEN
		list.Append(&Node{Data: expected})
		// THEN
		assertHead(t, list.Head, expected)

		expected2 := expected + 1
		// WHEN
		list.Append(&Node{Data: expected2})
		// THEN
		assertNode(t, "Head", list.Head, expected)
		assertNode(t, "Head.next", list.Head.next, expected2)
		assertList(t, &list, []int{expected, expected2})
	})
}

func TestAppendValue(t *testing.T) {
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%#v to a linkedlist", tc.dataValues), func(t *testing.T) {
			list := LinkedList{}
			for i := 0; i < len(tc.dataValues); i++ {
				list.AppendValue(tc.dataValues[i])
				if i == 0 {
					assertHead(t, list.Head, tc.dataValues[0])
				}
			}
			assertLengthEqual(t, tc.length, list.Length())
			assertList(t, &list, tc.dataValues)
		})
	}
}

func TestPrependOnly(t *testing.T) {
	expected := 42
	t.Run(fmt.Sprintf("%#v to a linkedlist", expected), func(t *testing.T) {
		list := LinkedList{}
		assertEmpty(t, list)
		list.Prepend(&Node{Data: expected})
		assertHead(t, list.Head, expected)
		assertLengthEqual(t, 1, list.Length())
		expected2 := expected + 1
		list.Prepend(&Node{Data: expected2})
		assertNode(t, "Head", list.Head, expected2)
		assertNode(t, "Head.next", list.Head.next, expected)
		assertLengthEqual(t, 2, list.Length())
		assertList(t, &list, []int{expected2, expected})
	})
}

func TestPrependValue(t *testing.T) {
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%#v to a linkedlist", tc.dataValues), func(t *testing.T) {
			list := LinkedList{}
			if len(tc.dataValues) > 0 {
				list.PrependValue(tc.dataValues[0])
				assertHead(t, list.Head, tc.dataValues[0])
				assertLengthEqual(t, 1, list.Length())
			}
			for i := 1; i < len(tc.dataValues); i++ {
				previous := list.Head.Data

				list.PrependValue(tc.dataValues[i])

				assertNode(t, "Head", list.Head, tc.dataValues[i])
				assertNode(t, "Head.next", list.Head.next, previous)
				assertLengthEqual(t, i+1, list.Length())
			}
			assertLengthEqual(t, tc.length, list.Length())
			assertList(t, &list, reverseData(tc.dataValues))
		})
	}
}

func TestInsertValue(t *testing.T) {
	expected := 0
	list := LinkedList{}
	assertEmpty(t, list)
	list.InsertValue(expected, 0)
	assertHead(t, list.Head, expected)
	expectedLength := 1
	assertLengthEqual(t, expectedLength, list.Length())
	for i := expected + 1; i < 1002; i = i + 100 {
		t.Run(fmt.Sprintf("%#v to a linkedlist", i), func(t *testing.T) {
			list.InsertValue(i, 1)
			assertNode(t, "Head", list.Head, expected)
			assertNode(t, "Head.next", list.Head.next, i)
			expectedLength++
			assertLengthEqual(t, expectedLength, list.Length())
		})
	}
	last := list.Find(1)
	assertNode(t, "Last", last, 1)
	if last.next != nil {
		t.Error("Next pointer for the last node should be nil but it has value:", last.next.Data)
	}
}

func TestInsertSimple(t *testing.T) {
	list := LinkedList{}
	list.Insert(&Node{Data: 2}, 0)
	assertHead(t, list.Head, 2)
	assertLengthEqual(t, 1, list.Length())

	list.Insert(&Node{Data: 1}, 0)
	assertNode(t, "Head", list.Head, 1)
	assertNode(t, "Last", list.Head.next, 2)
	assertLengthEqual(t, 2, list.Length())

	list.Insert(&Node{Data: 5}, 3)
	assertLengthEqual(t, 3, list.Length())

	list.Insert(&Node{Data: 3}, 2)
	assertLengthEqual(t, 4, list.Length())

	list.Insert(&Node{Data: 4}, 1)
	assertNode(t, "Head", list.Head, 1)
	assertLengthEqual(t, 5, list.Length())

	last := list.Find(5)
	assertNode(t, "Last", last, 5)
	if last.next != nil {
		t.Error("Next pointer for the last node should be nil but it has value:", last.next.Data)
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

func TestValuesOnly(t *testing.T) {
	expected := 42
	t.Run(fmt.Sprintf("%#v to a linkedlist", expected), func(t *testing.T) {
		list := LinkedList{}
		if "" != list.Values() {
			t.Error("Expected empty linked list to have values as empty string")
		}
		list.Head = &Node{Data: expected}
		s := strconv.Itoa(expected)
		if s != list.Values() {
			t.Errorf("Expected linked list to have values %d but received %s", expected, list.Values())
		}
		list.Head.next = &Node{Data: expected + 1}
		s = s + " " + strconv.Itoa(expected+1)
		if s != list.Values() {
			t.Errorf("Expected linked list to have values %s but received %s", s, list.Values())
		}
	})
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

func TestReduceSuccess(t *testing.T) {
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Reducing from linkedlist %#v", tc.dataValues), func(t *testing.T) {
			list := createList(tc.dataValues)
			previousLength := list.Length()

			list.Reduce()

			expectedLength := previousLength - 1
			if previousLength < 2 {
				expectedLength = 0
			} else {
				lastNode := list.Get(expectedLength - 1)
				assertNode(t, "Last Node", lastNode, tc.dataValues[expectedLength-1])
			}
			assertLengthEqual(t, expectedLength, list.Length())
			assertList(t, &list, tc.dataValues[:expectedLength])
		})
	}
}

func TestReduceEdgeCases(t *testing.T) {
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

func TestDeleteHeadSuccess(t *testing.T) {
	for _, tc := range testCases {
		if len(tc.dataValues) > 0 {
			t.Run(fmt.Sprintf("Deleting index 0 from linkedlist %#v", tc.dataValues), func(t *testing.T) {
				list := createList(tc.dataValues)
				previousLength := list.Length()
				expectedNode := list.Head.next

				err := list.Delete(0)
				if err != nil {
					t.Errorf("No error was expected but received: %v", err)
				}

				expectedLength := previousLength - 1
				if previousLength == 1 {
					expectedLength = 0
					assertEmpty(t, list)
				} else {
					assertNode(t, "New Head", list.Head, expectedNode.Data)
				}
				assertLengthEqual(t, expectedLength, list.Length())
				assertList(t, &list, tc.dataValues[1:])
			})
		}
	}
}

func TestDeleteTailSuccess(t *testing.T) {
	for _, tc := range testCases {
		if len(tc.dataValues) > 0 {
			lastIndex := len(tc.dataValues) - 1
			t.Run(fmt.Sprintf("Deleting the last item (index %v) from linkedlist %#v", lastIndex, tc.dataValues), func(t *testing.T) {
				list := createList(tc.dataValues)
				err := list.Delete(lastIndex)
				if err != nil {
					t.Errorf("No error was expected but received: %v", err)
				}
				assertLengthEqual(t, len(tc.dataValues)-1, list.Length())
				assertList(t, &list, tc.dataValues[:len(tc.dataValues)-1])
			})
		}
	}
}

func TestDeleteSuccess(t *testing.T) {
	for _, tc := range testCases {
		if len(tc.dataValues) > 1 {
			t.Run(fmt.Sprintf("Deleting index 1 from linkedlist %#v", tc.dataValues), func(t *testing.T) {
				list := createList(tc.dataValues)
				previousLength := list.Length()

				err := list.Delete(1)
				if err != nil {
					t.Errorf("No error was expected but received: %v", err)
				}

				assertLengthEqual(t, previousLength-1, list.Length())
				expectedValues := []int{}
				expectedValues = append(expectedValues, tc.dataValues[0])
				expectedValues = append(expectedValues, tc.dataValues[2:]...)
				assertList(t, &list, expectedValues)
			})
		}
	}
}

func TestDeleteEdgeCases(t *testing.T) {
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Negative Index on list %v", tc.dataValues), func(t *testing.T) {
			expected := "Cannot remove an index that is less than zero"
			list := createList(tc.dataValues)
			err := list.Delete(-1)
			assertError(t, expected, err)
		})
		t.Run(fmt.Sprintf("Out of range index on list %v", tc.dataValues), func(t *testing.T) {
			outOfRange := len(tc.dataValues) + 1
			expected := fmt.Sprintf("Index %d is out of range of the length of the list: %d", outOfRange, len(tc.dataValues))
			list := createList(tc.dataValues)
			err := list.Delete(outOfRange)
			assertError(t, expected, err)
		})
	}
}

func TestReverseEdgeCases(t *testing.T) {
	t.Run(fmt.Sprintf("empty or one item"), func(t *testing.T) {
		list := LinkedList{}
		list.Reverse()
		assertEmpty(t, list)
		list.AppendValue(1)
		assertHead(t, list.Head, 1)
		list.Reverse()
		assertHead(t, list.Head, 1)
	})
}

func TestReverseSuccess(t *testing.T) {
	var testCasesOrdered = []struct {
		dataValues []int
		length     int
	}{
		{dataValues: []int{}, length: 0},
		{dataValues: []int{1}, length: 1},
		{dataValues: []int{2, 1}, length: 2},
		{dataValues: []int{3, 2, 1}, length: 3},
		{dataValues: []int{4, 3, 2, 1}, length: 4},
	}
	for _, tc := range testCasesOrdered {
		t.Run(fmt.Sprintf("%v", tc.dataValues), func(t *testing.T) {
			list := createList(tc.dataValues)
			list.Reverse()
			assertLengthEqual(t, tc.length, list.Length())
			sorted := make([]int, len(tc.dataValues))
			copy(sorted, tc.dataValues)
			sort.Ints(sorted)
			assertList(t, &list, sorted)
		})
	}
	for _, tc := range testCasesOrdered {
		t.Run(fmt.Sprintf("ArrayMethod %v", tc.dataValues), func(t *testing.T) {
			list := createList(tc.dataValues)
			list.ReverseEasy()
			assertLengthEqual(t, tc.length, list.Length())
			sorted := make([]int, len(tc.dataValues))
			copy(sorted, tc.dataValues)
			sort.Ints(sorted)
			assertList(t, &list, sorted)
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

func reverseData(original []int) []int {
	a := make([]int, len(original))
	copy(a, original)
	last := len(a) - 1
	for i := 0; i < len(a)/2; i++ {
		a[i], a[last-i] = a[last-i], a[i]
	}
	return a
}

func assertLengthEqual(t *testing.T, expected, result int) {
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

func assertHead(t *testing.T, n *Node, expectedData int) {
	t.Helper()
	assertNode(t, "Head", n, expectedData)
	if n.next != nil {
		t.Error("Next pointer for the List Head should be nil")
	}
}

func assertList(t *testing.T, target *LinkedList, expectedData []int) {
	t.Helper()
	result := target.Values()
	if len(expectedData) == 0 {
		if "" != result {
			t.Errorf("List should have been empty but received: %v", result)
		}
		return
	}
	parts := strings.Fields(result)
	for i := 0; i < len(parts); i++ {
		current, err := strconv.Atoi(parts[i])
		if err != nil {
			t.Errorf("List index %d was expected: %d but error: %v", i, expectedData[i], err)
		}
		if current != expectedData[i] {
			t.Errorf("List index %d was expected: %d but received: %d", i, expectedData[i], current)
		}
	}
}

func assertError(t *testing.T, expected string, err error) {
	t.Helper()
	if (err == nil) || (err.Error() != expected) {
		t.Errorf("Should have received an error: %v", expected)
	}
}
