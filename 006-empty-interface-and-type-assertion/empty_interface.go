package emptyint

// MultitypeSum adds a list of different types of numbers together, producing a float64.
// It can handle the following types:
//			float64, int64, uint64
// If the input list contains any other type, that element of the list is ignored.
// If the list is empty (or consists of only non-supported types) returns 0.0
func MultitypeSum(input []interface{}) float64 {
	multiSum := float64(0)
	for _, value := range input {
		switch val := value.(type) {
		case float64:
			multiSum += val
		case int64:
			multiSum += float64(val)
		case uint64:
			multiSum += float64(val)
		default:
		}
	}
	return multiSum
}

// Stringer is an interface for something which can be represented as a string
// Do not modify this interface!
type Stringer interface {
	// String gets the string representation of this Stringer
	String() string
}

type defaultStringer struct {
	str string
}

func (ds defaultStringer) String() string {
	return ds.str
}
func (ds defaultStringer) AddString(str string) {
	ds.str += str
}

// AppendIfStringer checks if "input" is an implementation of Stringer and, if yes, returns a new Stringer
// with the value of "str" appended to the output of input.String():
//		inputIfStringer.String() + str
// If "input" is not a Stringer, returns a new Stringer containing only the value of "str".
func AppendIfStringer(input interface{}, str string) Stringer {
	if input == nil {
		return &defaultStringer{str: str}
	}
	if input, ok := input.(Stringer); ok {
		return defaultStringer{str: input.String() + str}
	} else {
		return defaultStringer{str: str}
	}
}
