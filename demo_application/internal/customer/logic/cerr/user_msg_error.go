package cerr

import "errors"

type userMsgError struct {
	msg     string
	wrapped error
}

func (u *userMsgError) Error() string {
	if u.wrapped != nil {
		return u.wrapped.Error()
	}

	return ""
}

func (u *userMsgError) Unwrap() error {
	return u.wrapped
}

func NewUserMsg(msg string) error {
	return SetUserMsg(errors.New(msg), msg)
}

func SetUserMsg(err error, msg string) error {
	return &userMsgError{
		msg:     msg,
		wrapped: err,
	}
}

func GetUserMsg(err error) (string, bool) {
	var userMsgErr *userMsgError
	if errors.As(err, &userMsgErr) {
		return userMsgErr.msg, true
	}

	return "", false
}
