package goutils

import (
	"regexp"
)

type IndentKind string

const (
	KIND_UNDEFINED IndentKind = "Undefined"
	KIND_SPACES    IndentKind = "Spaces"
	KIND_TABS      IndentKind = "Tabs"
)

func (k IndentKind) String() string {
	switch k {
	case KIND_SPACES:
		return string(k)
	case KIND_TABS:
		return string(k)
	case KIND_UNDEFINED:
		return string(k)
	default:
		return "Invalid Kind"
	}
}

var smartSplitRegexp = regexp.MustCompile("\r\n|\n\r|\n")

func SplitLines(input string) []string {
	result := smartSplitRegexp.Split(input, -1)
	return result
}

func DedentLines(lines []string) {
	indentCount, indentKind := CountIndent(lines)
	if indentCount == 0 || indentKind == KIND_UNDEFINED {
		return
	}
	for i := 0; i < len(lines); i++ {
		lines[i] = lines[i][indentCount:]
	}
}

func CountIndent(lines []string) (int, IndentKind) {
	if len(lines) == 0 {
		return 0, KIND_UNDEFINED
	}

	indentKind := DiscoverIndentKind(lines[0])
	if indentKind == KIND_UNDEFINED {
		return 0, KIND_UNDEFINED
	}
	minIndent := CountLineIndent(lines[0], indentKind)
	for _, line := range lines[1:] {
		if len(line) <= minIndent {
			minIndent = CountLineIndent(line, indentKind)
		} else {
			minIndent = CountLineIndent(line[:minIndent], indentKind)
		}
		if minIndent == 0 {
			return 0, KIND_UNDEFINED
		}
	}
	return minIndent, indentKind
}
