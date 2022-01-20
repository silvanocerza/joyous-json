// package step contains some utilities function to create the
// most used processor.StepFunc
package step

import (
	"fmt"

	"github.com/silvanocerza/joyous-json/pkg/processor"
)

// NewFilterInIfEqual returns a step that returns true if key exists in m and has identical value
func NewFilterInIfEqual(key string, value string) processor.StepFunc {
	return func(m *map[string]interface{}) (bool, error) {
		v, ok := (*m)[key]
		if !ok {
			return false, nil
		}

		return fmt.Sprint(v) == value, nil
	}
}

// NewFilterInIfDifferent returns a step that returns true if key exists in m and has different value
func NewFilterInIfDifferent(key string, value string) processor.StepFunc {
	return func(m *map[string]interface{}) (bool, error) {
		v, ok := (*m)[key]
		if !ok {
			return false, nil
		}
		return fmt.Sprint(v) != value, nil
	}
}

// NewFilterOutIfEqual returns a step that returns true if key exists in m and has different value
func NewFilterOutIfEqual(key string, value string) processor.StepFunc {
	return NewFilterInIfDifferent(key, value)
}

// NewFilterOutIfDifferent returns a step that returns true if key exists in m and has identical value
func NewFilterOutIfDifferent(key string, value string) processor.StepFunc {
	return NewFilterInIfEqual(key, value)
}

func filterNumber(key string, number float64, compare func(a, b float64) bool) processor.StepFunc {
	return func(m *map[string]interface{}) (bool, error) {
		v, ok := (*m)[key]
		if !ok {
			return false, nil
		}
		n, ok := v.(float64)
		if !ok {
			return false, fmt.Errorf("value of %s is not a number", key)
		}
		return compare(n, number), nil
	}
}

// NewFilterInGreaterThan returns a step that return true if key value exists in m and is greater than number.
func NewFilterInGreaterThan(key string, number float64) processor.StepFunc {
	return filterNumber(
		key,
		number,
		func(a, b float64) bool { return a > b },
	)
}

// NewFilterInGreaterOrEqualThan returns a step that return true if key value exists in m and is greater or equal than number.
func NewFilterInGreaterOrEqualThan(key string, number float64) processor.StepFunc {
	return filterNumber(
		key,
		number,
		func(a, b float64) bool { return a >= b },
	)
}

// NewFilterInLessThan returns a step that return true if key value exists in m and is less than number.
func NewFilterInLessThan(key string, number float64) processor.StepFunc {
	return filterNumber(
		key,
		number,
		func(a, b float64) bool { return a < b },
	)
}

// NewFilterInLessOrEqualThan returns a step that return true if key value exists in m and is less or equal than number.
func NewFilterInLessOrEqualThan(key string, number float64) processor.StepFunc {
	return filterNumber(
		key,
		number,
		func(a, b float64) bool { return a <= b },
	)
}

// NewFilterInGreaterThan returns a step that return false if key value exists in m and is greater than number.
func NewFilterOutGreaterThan(key string, number float64) processor.StepFunc {
	return NewFilterInLessOrEqualThan(key, number)
}

// NewFilterOutGreaterOrEqualThan returns a step that return false if key value exists in m and is greater or equal than number.
func NewFilterOutGreaterOrEqualThan(key string, number float64) processor.StepFunc {
	return NewFilterInLessThan(key, number)
}

// NewFilterOutLessThan returns a step that return false if key value exists in m and is less than number.
func NewFilterOutLessThan(key string, number float64) processor.StepFunc {
	return NewFilterInGreaterOrEqualThan(key, number)
}

// NewFilterOutLessOrEqualThan returns a step that return false if key value exists in m and is less or equal than number.
func NewFilterOutLessOrEqualThan(key string, number float64) processor.StepFunc {
	return NewFilterInGreaterThan(key, number)
}

// NewFilterInStep returns a step that returns true if key is found in m
func NewFilterIn(key string) processor.StepFunc {
	return func(m *map[string]interface{}) (bool, error) {
		_, ok := (*m)[key]
		return ok, nil
	}
}

// NewFilterOutStep returns a step that returns false if key is found in m
func NewFilterOut(key string) processor.StepFunc {
	return func(m *map[string]interface{}) (bool, error) {
		_, ok := (*m)[key]
		return !ok, nil
	}
}

// NewPrefixStep returns a step that adds a prefix to key if found in m
func NewPrefix(key, prefix string) processor.StepFunc {
	return func(m *map[string]interface{}) (bool, error) {
		if value, ok := (*m)[key]; ok {
			delete(*m, key)
			(*m)[prefix+key] = value
		}
		return true, nil
	}
}

// NewAddStep returns a step that adds a new key, value pair in m
func NewAdd(key string, value interface{}) processor.StepFunc {
	return func(m *map[string]interface{}) (bool, error) {
		(*m)[key] = value
		return true, nil
	}
}

// NewAddStep returns a step that adds a new key, value pair in m or an error
// if key is already set.
func NewSafeAdd(key string, value interface{}) processor.StepFunc {
	return func(m *map[string]interface{}) (bool, error) {
		if _, ok := (*m)[key]; ok {
			return false, fmt.Errorf("key %s already exist", key)
		}
		(*m)[key] = value
		return true, nil
	}
}
