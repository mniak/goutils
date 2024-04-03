package goutils

import "regexp"

var smartSplitRegexp = regexp.MustCompile("\r\n|\n\r|\n")

func SmartSplitLines(input string) []string {
	result := smartSplitRegexp.Split(input, -1)
	return result
}
