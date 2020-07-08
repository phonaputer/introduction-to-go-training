package closures

import (
	"../testutil"
	"testing"
)

// Tests for GetAdderSubber

func TestGetAdderSubber_InitialStateIsZero(t *testing.T) {
	_, _, curval := GetAdderSubber()

	testutil.AssertEqual(t, 0, curval())
}

func TestGetAdderSubber_InitialStateCanBeModifiedWithAdderAndSubber(t *testing.T) {
	addr, subbr, curval := GetAdderSubber()

	addr(10)
	testutil.AssertEqual(t, 10, curval())
	addr(1)
	testutil.AssertEqual(t, 11, curval())
	addr(0)
	testutil.AssertEqual(t, 11, curval())
	addr(-1)
	testutil.AssertEqual(t, 10, curval())

	subbr(0)
	testutil.AssertEqual(t, 10, curval())
	subbr(5)
	testutil.AssertEqual(t, 5, curval())
	subbr(1)
	testutil.AssertEqual(t, 4, curval())
	subbr(-22)
	testutil.AssertEqual(t, 26, curval())
}

// Tests for GetAggregator

func TestGetAggregator_AnyInput_InternalStateStartsAtZero(t *testing.T) {
	aggr := GetAggregator(func(a, b int) int { return a })

	testutil.AssertEqual(t, 0, aggr(0))
}

func TestGetAggregator_AdditionFunc_AggrWorksForAddition(t *testing.T) {
	aggr := GetAggregator(func(a, b int) int { return a + b + 5 })

	testutil.AssertEqual(t, 5, aggr(0))
	testutil.AssertEqual(t, 0, aggr(-10))
	testutil.AssertEqual(t, 7, aggr(2))
}

func TestGetAggregator_MultiplicationFunc_AggrWorksForMultiplication(t *testing.T) {
	aggr := GetAggregator(func(a, b int) int { return a*b + 1 })

	testutil.AssertEqual(t, 1, aggr(0))
	testutil.AssertEqual(t, -9, aggr(-10))
	testutil.AssertEqual(t, -17, aggr(2))
}

func TestGetAggregator_SubtractionFunc_AggrSubtractsInputFromState(t *testing.T) {
	aggr := GetAggregator(func(a, b int) int { return a - b })

	testutil.AssertEqual(t, 0, aggr(0))
	testutil.AssertEqual(t, 10, aggr(-10))
	testutil.AssertEqual(t, 8, aggr(2))
}
