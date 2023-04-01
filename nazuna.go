package nazuna

import (
	"golang.org/x/exp/constraints"
)

// Nil returns true if the value is nil.
func Nil(v any) bool { return v == nil }

// NotNil returns true if the value is not nil.
func NotNil(v any) bool { return v != nil }

// IsGreater returns a function that returns true if the
// value is greater than the anchor value.
func IsGreater[T constraints.Ordered](anchor T) func(T) bool {
	return func(i T) bool { return i > anchor }
}

// IsGreaterEqual returns a function that returns true if the
// value is greater than or equal to the anchor value.
func IsGreaterEqual[T constraints.Ordered](anchor T) func(T) bool {
	return func(i T) bool { return i >= anchor }
}

// IsEqual returns a function that returns true if the
// value is equal to the anchor value.
func IsEqual[T comparable](anchor T) func(T) bool {
	return func(i T) bool { return i == anchor }
}

// IsLess returns a function that returns true if the
// value is less than the anchor value.
func IsLess[T constraints.Ordered](anchor T) func(T) bool {
	return func(i T) bool { return i < anchor }
}

// IsLessEqual returns a function that returns true if the
// value is less than or equal to the anchor value.
func IsLessEqual[T constraints.Ordered](anchor T) func(T) bool {
	return func(i T) bool { return i <= anchor }
}
