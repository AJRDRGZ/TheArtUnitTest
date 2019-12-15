package calculator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMemCalculator_Sum_ByDefault_ReturnsZero(t *testing.T) {
	calc := makeCalc()
	lastSum := calc.Sum()

	assert.Equal(t, 0, lastSum, "it should be zero")
}

func TestMemCalculator_Sum_Add_WhenCalled_ChangesSum(t *testing.T) {
	calc := makeCalc()
	calc.Add(1)
	sum := calc.Sum()

	assert.Equal(t, 1, sum, "it should be one")
}

// makeCalc is a factory method to initialize MemCalculator
// This is a good idea, because it saves time writing the tests, makes the code
// inside each test smaller and a little more readable, and makes sure
// MemCalculator is always initialized the same way. It’s also better for test
// maintainability, because if the constructor for MemCalculator changes, you
// only need to change the initialization in one place instead of going through
// each test and changing the new call.
func makeCalc() *MemCalculator {
	return &MemCalculator{}
}
