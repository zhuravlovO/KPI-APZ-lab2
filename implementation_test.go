package lab2

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPostfixToLisp(t *testing.T) {
	testCases := []struct {
		name      string
		input     string
		expected  string
		expectErr bool
	}{
		{
			name:      "Simple addition",
			input:     "4 2 +",
			expected:  "(+ 4 2)",
			expectErr: false,
		},
		{
			name:      "Medium expression from task",
			input:     "4 2 - 3 2 ^ * 5 +",
			expected:  "(+ (* (- 4 2) (pow 3 2)) 5)",
			expectErr: false,
		},
		{
			name:      "Complex expression with 8 operands",
			input:     "1 2 + 3 * 9 3 / - 5 + 6 2 / *",
			expected:  "(* (+ (- (* (+ 1 2) 3) (/ 9 3)) 5) (/ 6 2))",
			expectErr: false,
		},
		{
			name:      "Empty input string",
			input:     "",
			expectErr: true,
		},
		{
			name:      "Invalid expression - not enough operands",
			input:     "4 +",
			expectErr: true,
		},
		{
			name:      "Invalid expression - too many operands",
			input:     "4 2 3 +",
			expectErr: true,
		},
		{
			name:  "Invalid expression - unknown operator",
			input: "5 5 &",

			expectErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := PostfixToLisp(tc.input)
			if tc.expectErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expected, result)
			}
		})
	}
}

func ExamplePostfixToLisp() {
	res, _ := PostfixToLisp("4 2 - 3 2 ^ * 5 +")
	fmt.Println(res)
	// Output: (+ (* (- 4 2) (pow 3 2)) 5)
}
