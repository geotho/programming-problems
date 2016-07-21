package main

import (
	"strings"
	"unicode"
)

func MatchPatterns(pattern string, classNames []string) []string {
	var matched []string
	for _, c := range classNames {
		if match(pattern, c) {
			matched = append(matched, c)
		}
	}
	return matched
}

// match returns true if the classname matches the pattern
func match(pattern, classname string) bool {
	patterns := split(pattern)
	classnames := split(classname)

	if len(patterns) > len(classnames) {
		return false
	}

	for i, p := range patterns {
		if !strings.HasPrefix(classnames[i], p) {
			return false
		}
	}

	return true
}

// split splits a string on its uppercase letters. "FooBar" -> ["Foo", "Bar"]
func split(s string) []string {
	start := 0
	var splits []string
	for i, r := range s {
		if unicode.IsUpper(r) && i > 0 {
			splits = append(splits, s[start:i])
			start = i
		}
	}
	splits = append(splits, s[start:])
	return splits
}
