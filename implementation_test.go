package prefixcalc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEvaluatePrefixExpression_Simple(t *testing.T) {
	result, err := EvaluatePrefixExpression("+ 2 3")
	assert.NoError(t, err)
	assert.Equal(t, float64(5), result)
}

func TestEvaluatePrefixExpression_Complex(t *testing.T) {
	expr := "+ 5 * - 4 2 ^ 3 2" // 5 + ((4 - 2) * (3 ^ 2)) = 5 + (2 * 9) = 23
	result, err := EvaluatePrefixExpression(expr)
	assert.NoError(t, err)
	assert.Equal(t, float64(23), result)
}

func TestEvaluatePrefixExpression_InvalidToken(t *testing.T) {
	_, err := EvaluatePrefixExpression("+ 2 a")
	assert.Error(t, err)
}

func TestEvaluatePrefixExpression_EmptyInput(t *testing.T) {
	_, err := EvaluatePrefixExpression("")
	assert.Error(t, err)
}

func TestEvaluatePrefixExpression_DivisionByZero(t *testing.T) {
	_, err := EvaluatePrefixExpression("/ 4 0")
	assert.Error(t, err)
}
func ExampleEvaluatePrefixExpression() {
	result, err := EvaluatePrefixExpression("+ 1 * 2 3")
	if err != nil {
		panic(err)
	}
	println(result)
}
