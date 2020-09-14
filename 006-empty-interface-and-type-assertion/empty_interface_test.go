package emptyint

import (
	"errors"
	"introduction-to-go-training/testutil"
	"testing"
)

// Tests for MultitypeSum

func TestMultitypeSum_EmptyListInput_ReturnsZero(t *testing.T) {
	res := MultitypeSum([]interface{}{})

	testutil.AssertEqual(t, 0.0, res)
}

func TestMultitypeSum_AllSupportedTypes_ReturnsSumOfAllInput(t *testing.T) {
	res := MultitypeSum([]interface{}{0.11, uint64(1), int64(2)})

	exptd := 0.11 + float64(uint64(1)) + float64(int64(2))
	testutil.AssertEqual(t, exptd, res)
}

func TestMultitypeSum_IncludingUnsupportedTypes_ReturnsSumExcludingUnsupported(t *testing.T) {
	res := MultitypeSum([]interface{}{0.11, []int{1}, float32(0.22), uint64(1), int64(2), "dfhsdf"})

	exptd := 0.11 + float64(uint64(1)) + float64(int64(2))
	testutil.AssertEqual(t, exptd, res)
}

func TestMultitypeSum_AllUnsupportedTypes_ReturnsZero(t *testing.T) {
	res := MultitypeSum([]interface{}{[]int{1}, float32(0.22), "dfhsdf"})

	testutil.AssertEqual(t, 0.0, res)
}

// tests for AddStackTrace

func TestAddStackTrace_Nil_ReturnsNil(t *testing.T) {
	res := AddStackTrace(nil, "stacked")

	testutil.AssertErrNil(t, res, true)
}

func TestAddStackTrace_RegularError_AddsStackTrace(t *testing.T) {
	input := errors.New("bigolerror")

	res := AddStackTrace(input, "stacked")

	hasStack, ok := res.(HasStackTrace)
	if ok {
		testutil.AssertEqual(t, input.Error(), hasStack.Error())
		testutil.AssertEqual(t, "stacked :: bigolerror", hasStack.GetStackTrace())
		return
	}
	t.Fatal("Did not return an implementation of HasStackTrace")
}

func TestAddStackTrace_stackTraceError_AddsNewStackLine(t *testing.T) {
	theErr := errors.New("bigolerror")
	input := &stackTraceError{err: theErr, stackTrace: "st1 :: st0"}

	res := AddStackTrace(input, "st2")

	hasStack, ok := res.(HasStackTrace)
	if ok {
		testutil.AssertEqual(t, theErr.Error(), hasStack.Error())
		testutil.AssertEqual(t, "st2 :: st1 :: st0", hasStack.GetStackTrace())
		return
	}
	t.Fatal("Did not return an implementation of HasStackTrace")
}

func TestAddStackTrace_startWithStackTraceErrorAndCallTwice_AddsTwoNewStackLines(t *testing.T) {
	theErr := errors.New("bigolerror")
	input := &stackTraceError{err: theErr, stackTrace: "st1 :: st0"}

	res := AddStackTrace(input, "st2")
	res = AddStackTrace(res, "st3")

	hasStack, ok := res.(HasStackTrace)
	if ok {
		testutil.AssertEqual(t, theErr.Error(), hasStack.Error())
		testutil.AssertEqual(t, "st3 :: st2 :: st1 :: st0", hasStack.GetStackTrace())
		return
	}
	t.Fatal("Did not return an implementation of HasStackTrace")
}

type otherImpl struct {
	err        error
	stackTrace string
}

func (e *otherImpl) Error() string {
	return e.err.Error()
}

func (e *otherImpl) GetStackTrace() string {
	return e.stackTrace
}

func (e *otherImpl) SetStackTrace(st string) {
	e.stackTrace = st
}

func TestAddStackTrace_OtherImplOfHasStackTrace_AddsNewStackLine(t *testing.T) {
	theErr := errors.New("bigolerror")
	input := &otherImpl{err: theErr, stackTrace: "st1 :: st0"}

	res := AddStackTrace(input, "st2")

	hasStack, ok := res.(HasStackTrace)
	if ok {
		testutil.AssertEqual(t, input.Error(), hasStack.Error())
		testutil.AssertEqual(t, "st2 :: st1 :: st0", hasStack.GetStackTrace())
		return
	}
	t.Fatal("Did not return an implementation of HasStackTrace")
}
