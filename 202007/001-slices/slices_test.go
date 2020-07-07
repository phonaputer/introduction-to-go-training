package slices

import (
	"math/rand"
	"testing"
)

// Tests checking the answer to SumFloats

func TestSumFloats_Nil_ReturnsZero(t *testing.T) {
	res := SumFloats(nil)

	assertEqual(t, 0.0, res)
}

func TestSumFloats_EmptySlice_ReturnsZero(t *testing.T) {
	res := SumFloats([]float64{})

	assertEqual(t, 0.0, res)
}

func TestSumFloats_LengthOne_ReturnsTheOne(t *testing.T) {
	expected := rand.Float64()

	res := SumFloats([]float64{expected})

	assertEqual(t, expected, res)
}

func TestSumFloats_MultipleValues_ReturnsSumOfValues(t *testing.T) {
	res := SumFloats([]float64{1.0, 2.0, 3.3333, 4.4444})

	assertEqual(t, 10.7777, res)
}

// Tests checking the answer to GetIntSlice

func TestGetIntSlice_FirstEqualsLast_ReturnsOne(t *testing.T) {
	res := GetIntSlice(7, 7)

	assertEqualMsg(t, 1, len(res), "Expected length %v but got length %v")
	assertEqual(t, 7, res[0])
}

func TestGetIntSlice_FirstIsGreaterthanLast_ReturnsEmptySlice(t *testing.T) {
	res := GetIntSlice(2, 1)

	assertEqualMsg(t, 0, len(res), "Expected length %v but got length %v")
}

func TestGetIntSlice_FirstIsLessThanLast_ReturnsAllIntsBetweenFirstAndLastInclusive(t *testing.T) {
	res := GetIntSlice(6, 11)

	expected := []int {6, 7, 8, 9, 10, 11}
	assertEqualMsg(t, len(expected), len(res), "Expected length %v but got length %v")
	for itr, val := range res {
		assertEqualMsg(t, expected[itr], val,"Expected %v but got %v (slice index %v)", itr)
	}
}

// Tests checking the answer to ConcatenateStringSlices

func TestConcatenateStringSlices_BothNil_ReturnsEmptySlice(t *testing.T) {
	res := ConcatenateStringSlices(nil, nil)

	assertEqualMsg(t, 0, len(res), "Expected length %v but got length %v")
}

func TestConcatenateStringSlices_BothEmpty_ReturnsEmptySlice(t *testing.T) {
	res := ConcatenateStringSlices([]string{}, []string{})

	assertEqualMsg(t, 0, len(res), "Expected length %v but got length %v")
}

func TestConcatenateStringSlices_Bnil_ReturnsA(t *testing.T) {
	res := ConcatenateStringSlices([]string{"1", "b"}, nil)

	expected := []string {"1", "b"}
	assertEqualMsg(t, len(expected), len(res), "Expected length %v but got length %v")
	for itr, val := range res {
		assertEqualMsg(t, expected[itr], val,"Expected %v but got %v (slice index %v)", itr)
	}
}

func TestConcatenateStringSlices_Anil_ReturnsB(t *testing.T) {
	res := ConcatenateStringSlices(nil, []string{"c", "100"})

	expected := []string {"c", "100"}
	assertEqualMsg(t, len(expected), len(res), "Expected length %v but got length %v")
	for itr, val := range res {
		assertEqualMsg(t, expected[itr], val,"Expected %v but got %v (slice index %v)", itr)
	}
}

func TestConcatenateStringSlices_BothHaveValues_ReturnsConcatenation(t *testing.T) {
	slicea := []string{"a", "b", "c"}
	sliceb := []string{"1", "2", "3"}

	res := ConcatenateStringSlices(slicea, sliceb)

	expected := []string {"a", "b", "c", "1", "2", "3"}
	assertEqualMsg(t, len(expected), len(res), "Expected length %v but got length %v")
	for itr, val := range res {
		assertEqualMsg(t, expected[itr], val,"Expected %v but got %v (slice index %v)", itr)
	}
}

// helpers

func assertEqualMsg(t *testing.T, expected, actual interface{}, format string, args... interface{}) {
	if actual != expected {
		sprintArgs := []interface{}{expected, actual}
		sprintArgs = append(sprintArgs, args...)
		t.Fatalf(format, sprintArgs...)
	}
}

func assertEqual(t *testing.T, expected, actual interface{}) {
	assertEqualMsg(t, expected, actual, "Expected '%v' but got '%v'")
}