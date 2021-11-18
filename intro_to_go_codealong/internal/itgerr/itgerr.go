package itgerr

import "errors"

type itgErr struct {
	wrapped error
	kind    Kind
	msg     string
}

func (i *itgErr) Error() string {
	return i.wrapped.Error()
}

func (i *itgErr) Unwrap() error {
	return i.wrapped
}

func WithKind(err error, k Kind, msg string) error {
	wrapped := err
	if err == nil {
		wrapped = errors.New(msg)
	}

	return &itgErr{
		wrapped: wrapped,
		kind:    k,
		msg:     msg,
	}
}

func GetKind(err error) (Kind, string) {
	var itgerror *itgErr
	if errors.As(err, &itgerror) {
		return itgerror.kind, itgerror.msg
	}

	return KindInternalServer, "internal server error"
}
