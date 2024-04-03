package goutils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSmartRaw(t *testing.T) {
	testCases := []struct {
		name   string
		input  string
		output string
	}{
		{
			name:   "Simple string",
			input:  "hello world",
			output: "hello world",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			output := SmartRaw(tc.input)
			assert.Equal(t, tc.output, output)
		})
	}
}
