package nazuna

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

// StringContainsAny returns a function that returns true if the
// string contains any of the substrings.
func StringContainsAny(substrs []string) func(string) bool {
	return func(s string) bool {
		for _, substr := range substrs {
			if strings.Contains(s, substr) {
				return true
			}
		}
		return false
	}
}

// StringContainsAll returns a function that returns true if the
// string contains all of the substrings.
func StringContainsAll(substrs []string) func(string) bool {
	return func(s string) bool {
		for _, substr := range substrs {
			if !strings.Contains(s, substr) {
				return false
			}
		}
		return true
	}
}

// StringContainsRune returns a function that returns true if the
// string contains the rune.
func StringContainsRune(r rune) func(string) bool {
	return func(s string) bool { return strings.ContainsRune(s, r) }
}

// StringContainsAnyRune returns a function that returns true if the
// string contains any of the runes.
func StringContainsAnyRune(runes []rune) func(string) bool {
	return func(s string) bool {
		for _, r := range runes {
			if strings.ContainsRune(s, r) {
				return true
			}
		}
		return false
	}
}

// StringContainsAllRune returns a function that returns true if the
// string contains all of the runes.
func StringContainsAllRune(runes []rune) func(string) bool {
	return func(s string) bool {
		for _, r := range runes {
			if !strings.ContainsRune(s, r) {
				return false
			}
		}
		return true
	}
}

// StringHasPrefix returns a function that returns true if the
// string has the prefix.
func StringHasPrefix(prefix string) func(string) bool {
	return func(s string) bool { return strings.HasPrefix(s, prefix) }
}

// StringHasSuffix returns a function that returns true if the
// string has the suffix.
func StringHasSuffix(suffix string) func(string) bool {
	return func(s string) bool { return strings.HasSuffix(s, suffix) }
}

// StringEqual returns a function that returns true if the
// string is equal to the other string.
func StringEqual(other string) func(string) bool {
	return func(s string) bool { return s == other }
}

// StringNotEqual returns a function that returns true if the
// string is not equal to the other string.
func StringNotEqual(other string) func(string) bool {
	return func(s string) bool { return s != other }
}

// StringEqualFold returns a function that returns true if the
// string is equal to the other string, ignoring case.
func StringEqualFold(other string) func(string) bool {
	return func(s string) bool { return strings.EqualFold(s, other) }
}

// StringNotEqualFold returns a function that returns true if the
// string is not equal to the other string, ignoring case.
func StringNotEqualFold(other string) func(string) bool {
	return func(s string) bool { return !strings.EqualFold(s, other) }
}
