/*
	This chapter will give you an introduction to generics, dispel reservations you may have about them and,
	give you an idea how to simplify some of your code in the future.
	After reading this you'll know how to write:
	- A function that takes generic arguments
	- A generic data-structure
*/

package fundamentalstest

import "testing"

func TestAssertFunctions(t *testing.T) {
	t.Run("asserting on integers", func(t *testing.T) {
		AssertEqual(t, 1, 1)
		AssertNotEqual(t, 1, 2)
	})

	t.Run("asserting on strings", func(t *testing.T) {
		AssertEqual(t, "hello", "hello")
		AssertNotEqual(t, "hello", "world")
	})
}

// "comparable" mean can compare, == or !=
// "any" can not compare
func AssertEqual[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got != want {
		t.Error("got", got, "want", want)
	}
}

func AssertNotEqual[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got == want {
		t.Error("didn't want", got)
	}
}

/*
	We're going to create a stack data type.
	Stacks should be fairly straightforward to understand from a requirements point of view.
	They're a collection of items where you can Push items to the "top" and
	to get items back again you Pop items from the top (LIFO - last in, first out).
	For the sake of brevity I've omitted the TDD process that arrived me at the following code
	for a stack of ints, and a stack of strings.
*/

type StacksOfInts struct {
	values []int
}

func (s *StacksOfInts) Push(value int) {
	s.values = append(s.values, value)
}

func (s *StacksOfInts) IsEmpty() bool {
	return len(s.values) == 0
}

func (s *StacksOfInts) Pop() (int, bool) {
	if s.IsEmpty() {
		return 0, false
	}

	index := len(s.values) - 1
	el := s.values[index]
	s.values = s.values[:index] // values = potong atas 1 element, dan sisanya

	return el, true
}

type StacksOfStrings struct {
	values []string
}

func (s *StacksOfStrings) Push(value string) {
	s.values = append(s.values, value)
}

func (s *StacksOfStrings) IsEmpty() bool {
	return len(s.values) == 0
}

func (s *StacksOfStrings) Pop() (string, bool) {
	if s.IsEmpty() {
		return "", false
	}

	index := len(s.values) - 1
	el := s.values[index]
	s.values = s.values[:index]

	return el, true
}

func AssertTrue(t *testing.T, got bool) {
	t.Helper()

	if !got {
		t.Error("got", got, "want true")
	}
}

func AssertFalse(t *testing.T, got bool) {
	t.Helper()

	if got {
		t.Error("got", got, "want false")
	}
}

func TestStack(t *testing.T) {
	t.Run("integer stack", func(t *testing.T) {
		myStackInt := StacksOfInts{}

		// check stack if empty
		AssertTrue(t, myStackInt.IsEmpty())

		// add a thing, then check it's not empty
		myStackInt.Push(10)
		AssertFalse(t, myStackInt.IsEmpty())

		// add another thing, pop it back again
		myStackInt.Push(12)
		value, _ := myStackInt.Pop()
		AssertEqual(t, value, 12)
		value, _ = myStackInt.Pop()
		AssertEqual(t, value, 10)
		AssertTrue(t, myStackInt.IsEmpty())
	})

	t.Run("string stack", func(t *testing.T) {
		myStackString := StacksOfStrings{}

		// check stack if empty
		AssertTrue(t, myStackString.IsEmpty())

		// add a thing, then check it's not empty
		myStackString.Push("10")
		AssertFalse(t, myStackString.IsEmpty())

		// add another thing, pop it back again
		myStackString.Push("12")
		value, _ := myStackString.Pop()
		AssertEqual(t, value, "12")
		value, _ = myStackString.Pop()
		AssertEqual(t, value, "10")
		AssertTrue(t, myStackString.IsEmpty())
	})
}

/*
	Problems
	- The code for both StackOfStrings and StackOfInts is almost identical.
		Whilst duplication isn't always the end of the world, it's more code to read, write and maintain.
	-	As we're duplicating the logic across two types, we've had to duplicate the tests too.

	Without generics, this is what we could do:
*/

type StackOfInts2 = Stack
type StackOfStrings2 = Stack

type Stack struct {
	values []interface{}
}

func (s *Stack) Push(value interface{}) {
	s.values = append(s.values, value)
}

func (s *Stack) IsEmpty() bool {
	return len(s.values) == 0
}

func (s *Stack) Pop() (interface{}, bool) {
	if s.IsEmpty() {
		var zero interface{}
		return zero, false
	}

	index := len(s.values) - 1
	el := s.values[index]
	s.values = s.values[:index]

	return el, true
}

func AssertEqual2(t *testing.T, got, want interface{}) {
	if got != want {
		t.Error("got", got, "want", want)
	}
}

func TestStack2(t *testing.T) {
	t.Run("interface stack dx is horrid", func(t *testing.T) {
		myStackOfInts2 := StackOfInts2{}

		myStackOfInts2.Push(1)
		myStackOfInts2.Push(2)

		value1, _ := myStackOfInts2.Pop()
		value2, _ := myStackOfInts2.Pop()

		// need to check we definitely got an int out of the interface{}, siapa yang tau itu int atau bukan ?
		intValue1, ok := value1.(int)
		AssertTrue(t, ok)

		intValue2, ok := value2.(int)
		AssertTrue(t, ok)

		AssertEqual2(t, intValue1+intValue2, 3)
	})
}

// Generic data structures to the rescue

type StackGeneric[T any] struct {
	values []T
}

func (s *StackGeneric[T]) Push(value T) {
	s.values = append(s.values, value)
}

func (s *StackGeneric[T]) IsEmpty() bool {
	return len(s.values) == 0
}

func (s *StackGeneric[T]) Pop() (T, bool) {
	if s.IsEmpty() {
		var zero T
		return zero, false
	}

	index := len(s.values) - 1
	el := s.values[index]
	s.values = s.values[:index]

	return el, true
}

func TestStackGeneric(t *testing.T) {
	t.Run("integer stacks", func(t *testing.T) {
		myStackOfInts := StackGeneric[int]{}

		// check stack is empty
		AssertTrue(t, myStackOfInts.IsEmpty())

		// add a thing, then check it's not empty
		myStackOfInts.Push(1)
		AssertFalse(t, myStackOfInts.IsEmpty())

		// add another thing, pop it back again
		myStackOfInts.Push(2)
		value1, _ := myStackOfInts.Pop()
		AssertEqual(t, value1, 2)
		value2, _ := myStackOfInts.Pop()
		AssertEqual(t, value2, 1)
		AssertTrue(t, myStackOfInts.IsEmpty())

		// we can get the numbers we put is an numbers, not untype as interface{}

		myStackOfInts.Push(1)
		myStackOfInts.Push(2)

		firstNum, _ := myStackOfInts.Pop()  // 2
		secondNum, _ := myStackOfInts.Pop() // 1
		AssertEqual(t, firstNum+secondNum, 3)

	})

	t.Run("string stacks", func(t *testing.T) {
		myStackOfStrings := StackGeneric[string]{}

		// check stack is empty
		AssertTrue(t, myStackOfStrings.IsEmpty())

		// add a thing, then check it's not empty
		myStackOfStrings.Push("a")
		AssertFalse(t, myStackOfStrings.IsEmpty())

		// add another thing, pop it back again
		myStackOfStrings.Push("b")
		value1, _ := myStackOfStrings.Pop()
		AssertEqual(t, value1, "b")

		value2, _ := myStackOfStrings.Pop()
		AssertEqual(t, value2, "a")
	})
}
