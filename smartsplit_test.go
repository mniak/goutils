package goutils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSmartSplitLines(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected []string
	}{
		{
			name:     "Empty",
			input:    "",
			expected: []string{""},
		},
		{
			name:     "Empty first line",
			input:    "\ntwo\nthree",
			expected: []string{"", "two", "three"},
		},
		{
			name:     "Single line",
			input:    "single line",
			expected: []string{"single line"},
		},
		{
			name:     "Three lines LF",
			input:    "one\ntwo\nthree",
			expected: []string{"one", "two", "three"},
		},
		{
			name:     "With new line at the end",
			input:    "one\ntwo\nthree\n",
			expected: []string{"one", "two", "three", ""},
		},
		{
			name:     "Three lines CRLF",
			input:    "one\r\ntwo\r\nthree",
			expected: []string{"one", "two", "three"},
		},
		{
			name:     "Three lines LFCR",
			input:    "one\n\rtwo\n\rthree",
			expected: []string{"one", "two", "three"},
		},
		{
			name:     "Four lines mixed LF CRLF LFCR",
			input:    "one\ntwo\r\nthree\n\rfour",
			expected: []string{"one", "two", "three", "four"},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			output := SmartSplitLines(tc.input)
			assert.Equal(t, tc.expected, output)
		})
	}
}
