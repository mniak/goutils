package goutils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSplitLines(t *testing.T) {
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
			output := SplitLines(tc.input)
			assert.Equal(t, tc.expected, output)
		})
	}
}

func TestCountLineIndent(t *testing.T) {
	testCases := []struct {
		name          string
		line          string
		kind          IndentKind
		expectedCount int
	}{
		{
			name:          "Empty UNDEFINED",
			line:          "",
			kind:          KIND_UNDEFINED,
			expectedCount: 0,
		},
		{
			name:          "Empty SPACES",
			line:          "",
			kind:          KIND_SPACES,
			expectedCount: 0,
		},
		{
			name:          "Empty TABS",
			line:          "",
			kind:          KIND_TABS,
			expectedCount: 0,
		},
		{
			name:          "One Space",
			line:          " text",
			kind:          KIND_SPACES,
			expectedCount: 1,
		},
		{
			name:          "Three Spaces",
			line:          "   text",
			kind:          KIND_SPACES,
			expectedCount: 3,
		},
		{
			name:          "Seven Spaces",
			line:          "       text",
			kind:          KIND_SPACES,
			expectedCount: 7,
		},
		{
			name:          "One Tab",
			line:          "\ttext",
			kind:          KIND_TABS,
			expectedCount: 1,
		},
		{
			name:          "Three Tabs",
			line:          "\t\t\ttext",
			kind:          KIND_TABS,
			expectedCount: 3,
		},
		{
			name:          "Seven Tabs",
			line:          "\t\t\t\t\t\t\ttext",
			kind:          KIND_TABS,
			expectedCount: 7,
		},
		{
			name:          "Tabs followed by spaces",
			line:          "\t\t    text",
			kind:          KIND_TABS,
			expectedCount: 2,
		},
		{
			name:          "Spaces followed by tabs",
			line:          "    \t\ttext",
			kind:          KIND_SPACES,
			expectedCount: 4,
		},
		{
			name:          "Expecting spaces gets tabs",
			line:          "\t\ttext",
			kind:          KIND_SPACES,
			expectedCount: 0,
		},
		{
			name:          "Expecting tabs gets spaces",
			line:          "   text",
			kind:          KIND_TABS,
			expectedCount: 0,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			count := CountLineIndent(tc.line, tc.kind)
			assert.Equal(t, tc.expectedCount, count)
		})
	}
}

func TestCountIndent(t *testing.T) {
	testCases := []struct {
		name          string
		text          string
		expectedCount int
		expectedKind  IndentKind
	}{
		// Empty
		{
			name:          "Empty",
			text:          "",
			expectedCount: 0,
			expectedKind:  KIND_UNDEFINED,
		},
		// Single line
		{
			name:          "Single line no indentation",
			text:          "one",
			expectedCount: 0,
			expectedKind:  KIND_UNDEFINED,
		},
		{
			name:          "Single line with spaces indentation",
			text:          "    one",
			expectedCount: 4,
			expectedKind:  KIND_SPACES,
		},
		{
			name:          "Single line with tabs indentation",
			text:          "\t\t\tone",
			expectedCount: 3,
			expectedKind:  KIND_TABS,
		},
		// Two lines
		{
			name:          "Two lines no indentation on first",
			text:          "one\n  two",
			expectedCount: 0,
			expectedKind:  KIND_UNDEFINED,
		},
		{
			name:          "Two lines no indentation on second",
			text:          "  one\ntwo",
			expectedCount: 0,
			expectedKind:  KIND_UNDEFINED,
		},
		{
			name:          "Two lines with spaces indentation",
			text:          "    one\n   two",
			expectedCount: 3,
			expectedKind:  KIND_SPACES,
		},
		{
			name:          "Two lines with tabs indentation",
			text:          "\t\t\tone\n\ttwo",
			expectedCount: 1,
			expectedKind:  KIND_TABS,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			lines := SplitLines(tc.text)

			count, kind := CountIndent(lines)

			assert.Equal(t, tc.expectedCount, count)
			assert.Equal(t, tc.expectedKind, kind)
		})
	}
}
