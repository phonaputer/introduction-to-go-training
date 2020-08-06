package closures

// GetAdderSubber returns three functions which share the same state: a sum (initialized to zero).
// The first two returned functions should take integer input:
// 		The adder function should add that input to the internal state.
//		The subber function should subtract that input from the internal state.
// The third function just returns the current value of the internal sum.
// E.g.
//      adder, subber, curVal := GetAdderSubber()
// 		adder(5) // internal sum is 5
// 		adder(11) // internal sum is 16
// 		subber(6) // internal sum  is 10
// 		curVal() // returns 10
func GetAdderSubber() (adder func(int), subber func(int), curVal func() int) {
	return nil, nil, nil // TODO implement
}

// GetAggregator returns a function which maintains an internal state: a counter initialized to zero.
// It also contains a function which can be used to combine two integers into one integer.
// The returned function should take string input, combine it with the counter using the above-mentioned function
// 		(the internal counter should be the first argument to the function),
// and return the new value of the counter.
// E.g.
//      aggr := GetAggregator( func(x, y int) int { return x - y } )
// 		aggr(5) // returns -5
// 		aggr(6) // returns -11
func GetAggregator(aggregate func(int, int) int) func(int) int {
	return nil // TODO implement
}
