package emptyint

// MultitypeSum adds a list of different types of numbers together, producing a float64.
// It can handle the following types:
//
//	float64, int64, uint64
//
// If the input list contains any other type, that element of the list is ignored.
// If the list is empty (or consists of only non-supported types) returns 0.0
func MultitypeSum(input []interface{}) float64 {
	var result float64 = 0.0
	for _, element := range input {
		switch concreteValue := element.(type) {
		case float64:
			result += concreteValue
		case int64:
			result += float64(concreteValue)
		case uint64:
			result += float64(concreteValue)

		}
	}
	return result
}

// Stringer is an interface for something which can be represented as a string
// Do not modify this interface!
type Stringer interface {

	// String gets the string representation of this Stringer
	String() string
}

type ConcreteStringer struct {
	stringRepresentation string
}

func (concreteStringer *ConcreteStringer) String() string {
	return concreteStringer.stringRepresentation
}

// AppendIfStringer checks if "input" is an implementation of Stringer and, if yes, returns a new Stringer
// with the value of "str" appended to the output of input.String():
//
//	inputIfStringer.String() + str
//
// If "input" is not a Stringer, returns a new Stringer containing only the value of "str".
func AppendIfStringer(input interface{}, str string) Stringer {
	inputStringer, ok := input.(Stringer)
	if ok {
		return &ConcreteStringer{inputStringer.String() + str}
	}
	return &ConcreteStringer{str}
}
