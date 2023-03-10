package slices

import "testing"

func TestSlicesAreEqualAreEqual(t *testing.T) {
	first := []int{1, 2, 3, 4, 5}
	second := []int{1, 2, 3, 4, 5}

	if !SlicesAreEqual(first, second) {
		t.Fatal("Slices are not equal")
	}
}

func TestSlicesAreEqualAreNotEqual(t *testing.T) {
	first := []int{1, 2, 3, 4, 5}
	second := []int{1, 2, 3, 4, 6}

	if SlicesAreEqual(first, second) {
		t.Fatal("Slices are equal when they shouldn't be")
	}
}

func TestIntToSlicePadded(t *testing.T) {
	num := 123456
	slice := IntToSlicePadded(num, 6)
	if !SlicesAreEqual(slice, []int{1, 2, 3, 4, 5, 6}) {
		t.Fatal("Slices are not equal")
	}
}

func TestIntToSlicePaddedNotEqual(t *testing.T) {
	num := 123456
	slice := IntToSlicePadded(num, 6)
	if SlicesAreEqual(slice, []int{1, 2, 3, 4, 5, 5}) {
		t.Fatal("Slices are equal when they should not be")
	}
}

func TestIntToSlicePaddedWithPadding(t *testing.T) {
	num := 1234
	slice := IntToSlicePadded(num, 6)
	if !SlicesAreEqual(slice, []int{0, 0, 1, 2, 3, 4}) {
		t.Fatal("Slices are not equal")
	}
}

func TestSlicesAreEqualAreNotEqualLength(t *testing.T) {
	first := []int{}
	second := []int{1, 2, 3, 4, 6}

	if SlicesAreEqual(first, second) {
		t.Fatal("Slices are equal when they shouldn't be")
	}
}
