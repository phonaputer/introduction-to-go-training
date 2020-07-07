package testutil

import "testing"

func AssertEqualMsg(t *testing.T, expected, actual interface{}, format string, args ...interface{}) {
	if actual != expected {
		sprintArgs := []interface{}{expected, actual}
		sprintArgs = append(sprintArgs, args...)
		t.Fatalf(format, sprintArgs...)
	}
}

func AssertEqual(t *testing.T, expected, actual interface{}) {
	AssertEqualMsg(t, expected, actual, "Expected '%v' but got '%v'")
}
