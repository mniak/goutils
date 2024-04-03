package goutils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSmartRaw_SimpleScenarios_ShouldReturnVerbatim(t *testing.T) {
	testCases := []struct {
		name   string
		input  string
		output string
	}{
		{
			name:   "Single line",
			input:  "hello world",
			output: "hello world",
		},
		{
			name:   "Two lines LF",
			input:  "hello\nworld",
			output: "hello\nworld",
		},
		{
			name:   "Two lines CRLF",
			input:  "hello\r\nworld",
			output: "hello\r\nworld",
		},
		{
			name:   "Two lines LFCR",
			input:  "hello\n\rworld",
			output: "hello\n\rworld",
		},
		{
			name:   "Two lines, second is idented with four spaces",
			input:  "hello\n  world",
			output: "hello\n  world",
		},
		{
			name:   "Two lines, second is idented with four spaces",
			input:  "hello\n    world",
			output: "hello\n    world",
		},
		{
			name:   "Two lines, second is idented with one tab",
			input:  "hello\n\tworld",
			output: "hello\n\tworld",
		},
		{
			name:   "Two lines, second is idented with two tabs",
			input:  "hello\n\t\tworld",
			output: "hello\n\t\tworld",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			output := SmartRaw(tc.input)
			assert.Equal(t, tc.output, output)
		})
	}
}
