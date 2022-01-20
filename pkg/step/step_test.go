package step

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewFilterInIfEqual(t *testing.T) {
	s := NewFilterInIfEqual("key", 420)
	input := map[string]interface{}{
		"key": 420,
	}
	expected := input
	ok, err := s(&input)
	require.True(t, ok)
	require.NoError(t, err)
	require.Equal(t, expected, input)

	input = map[string]interface{}{
		"key": 5,
	}
	expected = input
	ok, err = s(&input)
	require.False(t, ok)
	require.NoError(t, err)
	require.Equal(t, expected, input)

	input = map[string]interface{}{
		"otherKey": 5,
	}
	expected = input
	ok, err = s(&input)
	require.False(t, ok)
	require.NoError(t, err)
	require.Equal(t, expected, input)
}

func TestNewFilterInIfDifferent(t *testing.T) {
	s := NewFilterInIfDifferent("key", 5)
	input := map[string]interface{}{
		"key": 420,
	}
	expected := input
	ok, err := s(&input)
	require.True(t, ok)
	require.NoError(t, err)
	require.Equal(t, expected, input)

	input = map[string]interface{}{
		"key": 5,
	}
	expected = input
	ok, err = s(&input)
	require.False(t, ok)
	require.NoError(t, err)
	require.Equal(t, expected, input)

	input = map[string]interface{}{
		"otherKey": 420,
	}
	expected = input
	ok, err = s(&input)
	require.False(t, ok)
	require.NoError(t, err)
	require.Equal(t, expected, input)
}

func TestNewFilterOutIfEqual(t *testing.T) {
	s := NewFilterOutIfEqual("key", 5)
	input := map[string]interface{}{
		"key": 420,
	}
	expected := input
	ok, err := s(&input)
	require.True(t, ok)
	require.NoError(t, err)
	require.Equal(t, expected, input)

	input = map[string]interface{}{
		"key": 5,
	}
	expected = input
	ok, err = s(&input)
	require.False(t, ok)
	require.NoError(t, err)
	require.Equal(t, expected, input)

	input = map[string]interface{}{
		"otherKey": 420,
	}
	expected = input
	ok, err = s(&input)
	require.False(t, ok)
	require.NoError(t, err)
	require.Equal(t, expected, input)
}

func TestNewFilterOutIfDifferent(t *testing.T) {
	s := NewFilterOutIfDifferent("key", 420)
	input := map[string]interface{}{
		"key": 420,
	}
	expected := input
	ok, err := s(&input)
	require.True(t, ok)
	require.NoError(t, err)
	require.Equal(t, expected, input)

	input = map[string]interface{}{
		"key": 5,
	}
	expected = input
	ok, err = s(&input)
	require.False(t, ok)
	require.NoError(t, err)
	require.Equal(t, expected, input)

	input = map[string]interface{}{
		"otherKey": 5,
	}
	expected = input
	ok, err = s(&input)
	require.False(t, ok)
	require.NoError(t, err)
	require.Equal(t, expected, input)
}

func TestNewFilterInGreaterThan(t *testing.T) {
	s := NewFilterInGreaterThan("key", 42.0)
	input := map[string]interface{}{
		"key": 43.0,
	}
	expected := input
	ok, err := s(&input)
	require.True(t, ok)
	require.NoError(t, err)
	require.Equal(t, expected, input)

	input = map[string]interface{}{
		"key": 5.0,
	}
	expected = input
	ok, err = s(&input)
	require.False(t, ok)
	require.NoError(t, err)
	require.Equal(t, expected, input)

	input = map[string]interface{}{
		"otherKey": 5.0,
	}
	expected = input
	ok, err = s(&input)
	require.False(t, ok)
	require.NoError(t, err)
	require.Equal(t, expected, input)
}

func TestNewFilterInGreaterOrEqualThan(t *testing.T) {
	s := NewFilterInGreaterOrEqualThan("key", 42.0)
	input := map[string]interface{}{
		"key": 42.0,
	}
	expected := input
	ok, err := s(&input)
	require.True(t, ok)
	require.NoError(t, err)
	require.Equal(t, expected, input)

	input = map[string]interface{}{
		"key": 5.0,
	}
	expected = input
	ok, err = s(&input)
	require.False(t, ok)
	require.NoError(t, err)
	require.Equal(t, expected, input)

	input = map[string]interface{}{
		"otherKey": 5.0,
	}
	expected = input
	ok, err = s(&input)
	require.False(t, ok)
	require.NoError(t, err)
	require.Equal(t, expected, input)
}

func TestNewFilterInLessThan(t *testing.T) {
	s := NewFilterInLessThan("key", 42.0)
	input := map[string]interface{}{
		"key": 43.0,
	}
	expected := input
	ok, err := s(&input)
	require.False(t, ok)
	require.NoError(t, err)
	require.Equal(t, expected, input)

	input = map[string]interface{}{
		"key": 5.0,
	}
	expected = input
	ok, err = s(&input)
	require.True(t, ok)
	require.NoError(t, err)
	require.Equal(t, expected, input)

	input = map[string]interface{}{
		"otherKey": 5.0,
	}
	expected = input
	ok, err = s(&input)
	require.False(t, ok)
	require.NoError(t, err)
	require.Equal(t, expected, input)
}

func TestNewFilterInLessOrEqualThan(t *testing.T) {
	s := NewFilterInLessOrEqualThan("key", 42.0)
	input := map[string]interface{}{
		"key": 42.0,
	}
	expected := input
	ok, err := s(&input)
	require.True(t, ok)
	require.NoError(t, err)
	require.Equal(t, expected, input)

	input = map[string]interface{}{
		"key": 43.0,
	}
	expected = input
	ok, err = s(&input)
	require.False(t, ok)
	require.NoError(t, err)
	require.Equal(t, expected, input)

	input = map[string]interface{}{
		"otherKey": 5.0,
	}
	expected = input
	ok, err = s(&input)
	require.False(t, ok)
	require.NoError(t, err)
	require.Equal(t, expected, input)
}

func TestNewFilterOutGreaterThan(t *testing.T) {
	s := NewFilterOutGreaterThan("key", 42.0)
	input := map[string]interface{}{
		"key": 43.0,
	}
	expected := input
	ok, err := s(&input)
	require.False(t, ok)
	require.NoError(t, err)
	require.Equal(t, expected, input)

	input = map[string]interface{}{
		"key": 42.0,
	}
	expected = input
	ok, err = s(&input)
	require.True(t, ok)
	require.NoError(t, err)
	require.Equal(t, expected, input)

	input = map[string]interface{}{
		"otherKey": 5.0,
	}
	expected = input
	ok, err = s(&input)
	require.False(t, ok)
	require.NoError(t, err)
	require.Equal(t, expected, input)
}

func TestNewFilterOutGreaterOrEqualThan(t *testing.T) {
	s := NewFilterOutGreaterOrEqualThan("key", 42.0)
	input := map[string]interface{}{
		"key": 42.0,
	}
	expected := input
	ok, err := s(&input)
	require.False(t, ok)
	require.NoError(t, err)
	require.Equal(t, expected, input)

	input = map[string]interface{}{
		"key": 5.0,
	}
	expected = input
	ok, err = s(&input)
	require.True(t, ok)
	require.NoError(t, err)
	require.Equal(t, expected, input)

	input = map[string]interface{}{
		"otherKey": 5.0,
	}
	expected = input
	ok, err = s(&input)
	require.False(t, ok)
	require.NoError(t, err)
	require.Equal(t, expected, input)
}

func TestNewFilterOutLessThan(t *testing.T) {
	s := NewFilterOutLessThan("key", 42.0)
	input := map[string]interface{}{
		"key": 41.0,
	}
	expected := input
	ok, err := s(&input)
	require.False(t, ok)
	require.NoError(t, err)
	require.Equal(t, expected, input)

	input = map[string]interface{}{
		"key": 44.0,
	}
	expected = input
	ok, err = s(&input)
	require.True(t, ok)
	require.NoError(t, err)
	require.Equal(t, expected, input)

	input = map[string]interface{}{
		"otherKey": 5.0,
	}
	expected = input
	ok, err = s(&input)
	require.False(t, ok)
	require.NoError(t, err)
	require.Equal(t, expected, input)
}

func TestNewFilterOutLessOrEqualThan(t *testing.T) {
	s := NewFilterOutLessOrEqualThan("key", 42.0)
	input := map[string]interface{}{
		"key": 42.0,
	}
	expected := input
	ok, err := s(&input)
	require.False(t, ok)
	require.NoError(t, err)
	require.Equal(t, expected, input)

	input = map[string]interface{}{
		"key": 44.0,
	}
	expected = input
	ok, err = s(&input)
	require.True(t, ok)
	require.NoError(t, err)
	require.Equal(t, expected, input)

	input = map[string]interface{}{
		"otherKey": 5.0,
	}
	expected = input
	ok, err = s(&input)
	require.False(t, ok)
	require.NoError(t, err)
	require.Equal(t, expected, input)
}

func TestNewFilterIn(t *testing.T) {
	s := NewFilterIn("myKey")
	input := map[string]interface{}{
		"myKey": "myValue",
	}
	expected := input
	ok, err := s(&input)
	require.True(t, ok)
	require.NoError(t, err)
	require.Equal(t, expected, input)

	input = map[string]interface{}{
		"someoneElseKey": "myValue",
	}
	expected = input
	ok, err = s(&input)
	require.False(t, ok)
	require.NoError(t, err)
	require.Equal(t, expected, input)
}

func TestNewFilterOut(t *testing.T) {
	s := NewFilterOut("myKey")
	input := map[string]interface{}{
		"myKey": "myValue",
	}
	expected := input
	ok, err := s(&input)
	require.False(t, ok)
	require.NoError(t, err)
	require.Equal(t, expected, input)

	input = map[string]interface{}{
		"someoneElseKey": "myValue",
	}
	expected = input
	ok, err = s(&input)
	require.True(t, ok)
	require.NoError(t, err)
	require.Equal(t, expected, input)
}

func TestNewPrefix(t *testing.T) {
	s := NewPrefix("key", "prefix_")
	input := map[string]interface{}{
		"key":      "myValue",
		"otherKey": "foo",
	}
	expected := map[string]interface{}{
		"prefix_key": "myValue",
		"otherKey":   "foo",
	}
	ok, err := s(&input)
	require.True(t, ok)
	require.NoError(t, err)
	require.Equal(t, expected, input)
}

func TestNewAdd(t *testing.T) {
	s := NewAdd("myKey", 420)
	input := map[string]interface{}{}
	expected := map[string]interface{}{
		"myKey": 420,
	}
	ok, err := s(&input)
	require.True(t, ok)
	require.NoError(t, err)
	require.Equal(t, expected, input)
}

func TestNewSafeAdd(t *testing.T) {
	s := NewSafeAdd("myKey", 420)
	input := map[string]interface{}{}
	expected := map[string]interface{}{
		"myKey": 420,
	}
	ok, err := s(&input)
	require.True(t, ok)
	require.NoError(t, err)
	require.Equal(t, expected, input)

	input = map[string]interface{}{
		"myKey": 420,
	}
	ok, err = s(&input)
	require.False(t, ok)
	require.Error(t, err, "key myKey already exist")
}
