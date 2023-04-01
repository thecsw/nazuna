package nazuna

// SliceEmpty returns true if the slice is empty.
func SliceEmpty[T any](arr []T) bool { return len(arr) < 1 }

// SliceNotEmpty returns true if the slice is not empty.
func SliceNotEmpty[T any](arr []T) bool { return len(arr) > 0 }

// SliceContains returns a function that returns true if the
// slice contains the value.
func SliceContains[T comparable](v T) func([]T) bool {
	return func(arr []T) bool {
		for _, i := range arr {
			if i == v {
				return true
			}
		}
		return false
	}
}

// SliceContainsAny returns a function that returns true if the
// slice contains any of the values.
func SliceContainsAny[T comparable](vs []T) func([]T) bool {
	return func(arr []T) bool {
		for _, v := range vs {
			for _, i := range arr {
				if i == v {
					return true
				}
			}
		}
		return false
	}
}

// SliceContainsAll returns a function that returns true if the
// slice contains all of the values.
func SliceContainsAll[T comparable](vs []T) func([]T) bool {
	return func(arr []T) bool {
		for _, v := range vs {
			found := false
			for _, i := range arr {
				if i == v {
					found = true
					break
				}
			}
			if !found {
				return false
			}
		}
		return true
	}
}

// SliceContainsRune returns a function that returns true if the
// slice contains the rune.
func SliceContainsRune(r rune) func([]rune) bool {
	return func(arr []rune) bool {
		for _, i := range arr {
			if i == r {
				return true
			}
		}
		return false
	}
}

// SliceContainsAnyRune returns a function that returns true if the
// slice contains any of the runes.
func SliceContainsAnyRune(runes []rune) func([]rune) bool {
	return func(arr []rune) bool {
		for _, r := range runes {
			for _, i := range arr {
				if i == r {
					return true
				}
			}
		}
		return false
	}
}

// SliceContainsAllRune returns a function that returns true if the
// slice contains all of the runes.
func SliceContainsAllRune(runes []rune) func([]rune) bool {
	return func(arr []rune) bool {
		for _, r := range runes {
			found := false
			for _, i := range arr {
				if i == r {
					found = true
					break
				}
			}
			if !found {
				return false
			}
		}
		return true
	}
}
