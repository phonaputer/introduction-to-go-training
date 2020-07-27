package structs

// AdderSubber maintains an internal sum which can be modified using the interface's methods
type AdderSubber interface {

	// Add an integer to the internal sum
	Add(amount int)

	// Subtract an integer from the internal sum
	Subtract(amount int)

	// GetCurrentValue returns the current internal sum
	GetCurrentValue() int
}

// NewAdderSubber returns a concrete implementation of the AdderSubber interface.
// The returned AdderSubber will have its sum set to initialSum
// Note: feel free to define any new types you need to achieve this.
//		 And remember that receiver arguments (like all Go arguments) are pass-by-value.
func NewAdderSubber(initialSum int) AdderSubber {
	return nil // TODO implement
}

// Dog is an interface representing an individual of the species Canis Familiaris
type Dog interface {

	// MakeNoise returns a dog noise.
	MakeNoise() string

	// RollOver returns a boolean depending on whether the dog rolled over as instructed.
	// If the dog is a good boy he will roll over. If he is not, he will not.
	RollOver() bool

	// SetIsGoodBoy sets a flag indicating whether this Dog is a good boy.
	SetIsGoodBoy(isGoodBoy bool)
}

// Use this struct when writing your Dog implementation (But don't modify it. Hint: you can embed structs)
type Canid struct{}

func (n Canid) MakeNoise() string {
	return "BARK BARK!!!!"
}

// NewDog returns a concrete struct which meets the Dog interface (see above).
// To solve this problem, you need use the Canid struct in your Dog implementation.
// It doesn't matter whether your impl is or is not a good boy by default.
func NewDog() Dog {
	return nil // TODO implement
}
