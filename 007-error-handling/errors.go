package errhandle

import "errors"

// Divide divides numerator by denominator (numerator / denominator) and returns the result.
// If denominator is 0, it returns an error (it doesn't matter what number you return in the error case).
func Divide(numerator, denominator float64) (float64, error) {
	return 0.0, nil // TODO implement
}

// A UsernameGetter is a function which selects a user from storage using their user ID.
// If the user is not found, returns NotFoundError. If there is an unexpected error, returns that error.
// Do not modify this type!
type UsernameGetter func(userId int) (username string, err error)

// Do not modify this variable!
var NotFoundError = errors.New("user ID not found!")

// GetUsernameFromDb uses the provided selectUsername function to get a user from storage
//(see UsernameGetter comments above).
// Returns true and the retrieved username if there is no error.
// If selectUsername returns NotFoundError. Returns false and no error or username.
// If selectUsername returns any other error, returns that error
// (doesn't matter what string or bool is returned in this case).
func GetUsernameFromDb(userId int, selectUsername UsernameGetter) (username string, doesUserExist bool, err error) {
	return "", false, nil // TODO implement
}
