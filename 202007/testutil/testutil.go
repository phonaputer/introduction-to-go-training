package testutil

import "testing"

func AssertEqualMsgf(t *testing.T, expected, actual interface{}, format string, args ...interface{}) {
	if actual != expected {
		sprintArgs := []interface{}{expected, actual}
		sprintArgs = append(sprintArgs, args...)
		t.Fatalf(format, sprintArgs...)
	}
}

func AssertEqual(t *testing.T, expected, actual interface{}) {
	AssertEqualMsgf(t, expected, actual, "Expected '%v' but got '%v'")
}

func AssertEqualMsg(t *testing.T, expected, actual interface{}, msg string) {
	if actual != expected {
		t.Fatal(msg)
	}
}

func AssertErrNil(t *testing.T, theErr error, shouldBeNil bool) {
	if shouldBeNil && theErr != nil {
		t.Fatal("The result error should be nil!")
	}

	if !shouldBeNil && theErr == nil {
		t.Fatal("The result error should NOT be nil!")
	}
}
