package emptyint

// MultitypeSum adds a list of different types of numbers together, producing a float64.
// It can handle the following types:
//			float64, int64, uint64
// If the input list contains any other type, that element of the list is ignored.
// If the list is empty (or consists of only non-supported types) returns 0.0
func MultitypeSum(input []interface{}) float64 {
	return 0.0 // TODO implement
}

// HasStackTrace is an interface matching any error that has a stack trace associated with it
// Do not modify this interface!
type HasStackTrace interface {

	// GetStackTrace gets the error's stack trace
	GetStackTrace() string

	// SetStackTrace set's the error's stack trace
	SetStackTrace(st string)

	// Error gets the message of the root error (excluding stack trace)
	Error() string
}

// A simple implementation of HasStackTrace
// Do not modify this struct!
type stackTraceError struct {
	err        error
	stackTrace string
}

func (e *stackTraceError) Error() string {
	return e.err.Error()
}

func (e *stackTraceError) GetStackTrace() string {
	return e.stackTrace
}

func (e *stackTraceError) SetStackTrace(st string) {
	e.stackTrace = st
}

// AddStackTrace either adds a stack trace to an error which does not have one (by returning a stackTraceError),
// or adds another line to the stack trace of an error that already has a stack trace.
// If the input err is nil, returns nil
// The stacktrace format is:
//			"${stack line} :: ${current stack trace of err}"
// Or if the input error does not yet have a stack trace:
// 			"${stack line} :: ${message of err}"
//
// Note: This method should work with any type that matches HasStackTrace, not only stackTraceError.
func AddStackTrace(err error, stackLine string) error {
	return nil // TODO implement
}
