package goutils

import "strings"

func SmartRaw(input string) string {
	lines := SplitLines(input)
	firstLineIsEmpty := lines[0] == ""
	lines = lines[1:]
	if !firstLineIsEmpty {
		return input
	}
	DedentLines(lines)
	return strings.Join(lines, "\n")
}

func CountLineIndentOfKind2(line string, kind IndentKind) (int, IndentKind) {
	// for charIndex, char := range line {
	// 	switch kind {
	// 	case KIND_SPACES:
	// 		if char != ' ' {
	// 			return charIndex, KIND_SPACES
	// 		}
	// 	case KIND_TABS:
	// 		if char != '\t' {
	// 			return charIndex, KIND_TABS
	// 		}
	// 	case KIND_UNDEFINED:
	// 		kind = DiscoverIndentKind(line)
	// 		if kind == KIND_UNDEFINED {
	// 			return 0, KIND_UNDEFINED
	// 		}
	// 	default:
	return 0, KIND_UNDEFINED
	// 	}
	// }
}

func CountLineIndent(line string, kind IndentKind) int {
	var expectedChar rune
	switch kind {
	case KIND_SPACES:
		expectedChar = ' '
	case KIND_TABS:
		expectedChar = '\t'
	default:
		return 0
	}
	for charIndex, char := range line {
		if char != expectedChar {
			return charIndex
		}
	}
	return len(line)
}

func DiscoverIndentKind(line string) IndentKind {
	if len(line) == 0 {
		return KIND_UNDEFINED
	}

	switch line[0] {
	case ' ':
		return KIND_SPACES
	case '\t':
		return KIND_TABS
	default:
		return KIND_UNDEFINED
	}
}
