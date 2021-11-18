package structs

// AdderSubber maintains an internal sum which can be modified using the interface's methods
// Do not modify this interface!
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

type Calculator struct {
	res int
}

func (c *Calculator) Add(amount int) {
	c.res += amount
}
func (c *Calculator) Subtract(amount int) {
	c.res -= amount
}
func (c *Calculator) GetCurrentValue() int {
	return c.res
}

func NewAdderSubber(initialSum int) AdderSubber {
	cal := Calculator{res: initialSum}
	cal.Add(initialSum)
	cal.Subtract(initialSum)
	cal.GetCurrentValue()
	return &cal
}

// Dog is an interface representing an individual of the species Canis Familiaris
// Do not modify this interface!
type Dog interface {

	// MakeNoise returns a dog noise.
	MakeNoise() string

	// RollOver returns a boolean indicating whether the dog rolled over as instructed.
	// If the dog is a good dog it will roll over. If it is not a good dog, it will not.
	RollOver() bool

	// SetIsGoodDog sets a flag indicating whether this Dog is a good dog.
	SetIsGoodDog(isGoodDog bool)
}

// Use this struct when writing your Dog implementation.
// Do not modify this struct! Hint: remember that we talked about embedded structs.
type Barker struct{}

func (n *Barker) MakeNoise() string {
	return "BARK BARK!!!!"
}

type GoodDog struct {
	barker    Barker
	isGoodDog bool
}

func (gd *GoodDog) MakeNoise() string {
	return gd.barker.MakeNoise()
}
func (gd *GoodDog) SetIsGoodDog(isGoodDog bool) {
	gd.isGoodDog = isGoodDog
}
func (gd *GoodDog) RollOver() bool {
	return gd.isGoodDog
}

// NewDog returns a concrete struct which meets the Dog interface (see above).
// To solve this problem, you need use the Barker struct in your Dog implementation.
// It doesn't matter whether your impl is or is not a good dog by default.
func NewDog() Dog {
	goodDog := new(GoodDog)
	goodDog.SetIsGoodDog(true)
	goodDog.RollOver()
	goodDog.MakeNoise()
	return goodDog
}
