package main

import "strings"

// StringEmpty returns true if the string is empty.
func StringEmpty(s string) bool { return len(s) < 1 }

// StringNotEmpty returns true if the string is not empty.
func StringNotEmpty(s string) bool { return len(s) > 0 }

// StringContains returns a function that returns true if the
// string contains the substring.
func StringContains(substr string) func(string) bool {
	return func(s string) bool { return strings.Contains(s, substr) }
}

// SliceEmpty returns true if the slice is empty.
func SliceEmpty[T any](arr []T) bool { return len(arr) < 1 }

// SliceNotEmpty returns true if the slice is not empty.
func SliceNotEmpty[T any](arr []T) bool { return len(arr) > 0 }

// Nil returns true if the value is nil.
func Nil(v any) bool { return v == nil }

// NotNil returns true if the value is not nil.
func NotNil(v any) bool { return v != nil }
