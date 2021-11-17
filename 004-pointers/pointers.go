package pointers

// AddToPointer adds amountToAdd to the integer referenced by ptr.
// If ptr is nil, does nothing.
// E.g.
//      intVar := 3
//      AddToPointer(&intVar, 3)
//		fmt.Println(intVar) // 6
func AddToPointer(ptr *int, amountToAdd int) {
	if ptr == nil {
		return
	}
	*ptr = *ptr + amountToAdd
}

// SwapStrings switches the values of two string pointers.
// If either of the two pointers is nil, does nothing.
// E.g.
//		strA := "abc"
//		strB := "xyz"
// 		SwapStrings(strA, strB)
// 		fmt.Println(strA + ", " + strB) // "xyz, abc"
func SwapStrings(a, b *string) {
	if a == nil || b == nil {
		return
	}
	temp := *a
	*a = *b
	*b = temp
}

// SumOptionalList sums all integers in the input int pointer list.
// It skips any pointers which are nil.
// If the list is empty, or contains only nil pointers, returns 0.
// E.g. for the input list { 4, nil, 5, 2 } it will return 11
func SumOptionalList(maybeInts []*int) int {
	sum := 0
	if len(maybeInts) == 0 {
		return sum
	}
	for _, iPtr := range maybeInts {
		if iPtr != nil {
			sum += *iPtr
		}
	}
	return sum
}
