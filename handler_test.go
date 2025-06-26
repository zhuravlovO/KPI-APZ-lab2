package lab2

import (
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type errorReader struct{}

func (r *errorReader) Read(p []byte) (n int, err error) {
	return 0, fmt.Errorf("mock reader error")
}

func TestComputeHandler(t *testing.T) {
	t.Run("Valid expression", func(t *testing.T) {
		input := strings.NewReader("3 4 +")
		output := &bytes.Buffer{}
		handler := &ComputeHandler{Input: input, Output: output}

		err := handler.Compute()

		assert.NoError(t, err)
		assert.Equal(t, "(+ 3 4)", output.String())
	})

	t.Run("Invalid expression", func(t *testing.T) {
		input := strings.NewReader("3 +")
		output := &bytes.Buffer{}
		handler := &ComputeHandler{Input: input, Output: output}

		err := handler.Compute()

		assert.Error(t, err)
	})

	t.Run("Input returns error", func(t *testing.T) {
		input := &errorReader{}
		output := &bytes.Buffer{}
		handler := &ComputeHandler{Input: input, Output: output}

		err := handler.Compute()

		assert.Error(t, err)
		assert.Equal(t, "mock reader error", err.Error())
	})
}
