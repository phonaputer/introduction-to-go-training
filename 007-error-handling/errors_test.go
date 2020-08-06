package errhandle

import (
	"errors"
	"introduction-to-go-training/testutil"
	"testing"
)

// Tests for Divide

func TestDivide_DenominatorNotZero_ReturnsDivisionResult(t *testing.T) {
	var numer, denomin = 1.7, 8.8

	res, err := Divide(numer, denomin)

	testutil.AssertErrNil(t, err, true)
	testutil.AssertEqual(t, numer/denomin, res)
}

func TestDivide_DenominatorIsZero_ReturnsError(t *testing.T) {
	_, err := Divide(1.0, 0.0)

	testutil.AssertErrNil(t, err, false)
}

// Tests for GetUsernameFromDb

func TestGetUsernameFromDb_UsernameGetterReturnsNotFound_ReturnsFalseWithNoError(t *testing.T) {
	inputUserId := 123
	wasCalled := false
	getter := func(userId int) (username string, err error) {
		testutil.AssertEqual(t, inputUserId, userId)
		wasCalled = true
		return "", NotFoundError
	}

	_, resBool, resErr := GetUsernameFromDb(inputUserId, getter)

	testutil.AssertEqualMsg(t, true, wasCalled, "UsernameGetter was not called!")
	testutil.AssertEqual(t, false, resBool)
	testutil.AssertErrNil(t, resErr, true)
}

func TestGetUsernameFromDb_UsernameGetterReturnsUnknownErr_ReturnsError(t *testing.T) {
	inputUserId := 123
	wasCalled := false
	getter := func(userId int) (username string, err error) {
		testutil.AssertEqual(t, inputUserId, userId)
		wasCalled = true
		return "", errors.New("random err")
	}

	_, _, resErr := GetUsernameFromDb(inputUserId, getter)

	testutil.AssertEqualMsg(t, true, wasCalled, "UsernameGetter was not called!")
	testutil.AssertErrNil(t, resErr, false)
}

func TestGetUsernameFromDb_UsernameGetterReturnsSuccessfully_ReturnsResultAndTrue(t *testing.T) {
	inputUserId := 123
	wasCalled := false
	dbUsername := "usernm"
	getter := func(userId int) (username string, err error) {
		testutil.AssertEqual(t, inputUserId, userId)
		wasCalled = true
		return dbUsername, nil
	}

	resUn, resBool, resErr := GetUsernameFromDb(inputUserId, getter)

	testutil.AssertEqualMsg(t, true, wasCalled, "UsernameGetter was not called!")
	testutil.AssertErrNil(t, resErr, true)
	testutil.AssertEqual(t, true, resBool)
	testutil.AssertEqual(t, dbUsername, resUn)
}
