package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestComputeHandler_Valid(t *testing.T) {
	input := strings.NewReader("+ 1 2")
	output := &strings.Builder{}

	handler := &ComputeHandler{Input: input, Output: output}
	err := handler.Compute()

	assert.NoError(t, err)
	assert.Equal(t, "3.00\n", output.String())
}

func TestComputeHandler_Invalid(t *testing.T) {
	input := strings.NewReader("+ 1 x")
	output := &strings.Builder{}

	handler := &ComputeHandler{Input: input, Output: output}
	err := handler.Compute()

	assert.Error(t, err)
}
