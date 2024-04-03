package strings

import (
	"strings"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
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

func TestSmartRaw_SimpleScenarios_Empties(t *testing.T) {
	testCases := []struct {
		name   string
		input  string
		output string
	}{
		{
			name:   "Empty",
			input:  "",
			output: "",
		},
		{
			name:   "Spaces",
			input:  "              ",
			output: "              ",
		},
		{
			name:   "Spaces after new line",
			input:  "\n              ",
			output: "",
		},
		{
			name:   "Empty after new line LF",
			input:  "\n",
			output: "",
		},
		{
			name:   "New line only CRLF",
			input:  "\r\n",
			output: "",
		},
		{
			name:   "New line only LFCR",
			input:  "\n\r",
			output: "",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			output := SmartRaw(tc.input)
			assert.Equal(t, tc.output, output)
		})
	}
}

func TestSmartRaw_ComplexScenarios(t *testing.T) {
	lines := make([]string, 10)
	gofakeit.Slice(&lines)

	// Add different indentations to the lines
	for i := 0; i < len(lines); i++ {
		lines[i] = strings.Repeat(" ", 3-(i%4)) + lines[i]
	}
	baseText := strings.Join(lines, "\n")

	// Add common indentation to all lines
	commonIndentation := strings.Repeat(" ", gofakeit.IntRange(4, 10))
	for i := 0; i < len(lines); i++ {
		lines[i] = commonIndentation + lines[i]
	}
	inputText := strings.Join(lines, "\n")

	t.Run("When first line is not empty, return the lines as they are", func(t *testing.T) {
		output := SmartRaw(inputText)
		assert.Equal(t, inputText, output)
	})
	t.Run("When first line is empty, should remove indentation", func(t *testing.T) {
		output := SmartRaw("\n" + inputText)
		assert.Equal(t, baseText, output)
	})
}

func TestSmartRaw_Examples(t *testing.T) {
	testCases := []struct {
		name   string
		input  string
		output string
	}{
		{
			name: "#1",
			input: `
				SELECT *
				FROM TABLE1
				  WHERE ID = 2`,
			output: "SELECT *\nFROM TABLE1\n  WHERE ID = 2",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			output := SmartRaw(tc.input)
			assert.Equal(t, tc.output, output)
		})
	}
}
