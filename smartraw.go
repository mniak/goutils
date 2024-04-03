package goutils

func SmartRaw(input string) string {
	lines := SmartSplitLines(input)
	emptyFirstLine := lines[0] == ""
	if !emptyFirstLine {
		return input
	}
	return input
}
