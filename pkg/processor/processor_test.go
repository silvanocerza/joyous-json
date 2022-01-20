package processor

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func BenchmarkProcessor(b *testing.B) {
	for i := 0; i < b.N; i++ {
		file, _ := os.Open(filepath.Join(".", "testdata", "test.json"))
		p := New(file, io.Discard)
		p.AddStep(func(m *map[string]interface{}) (bool, error) {
			(*m)["test"] = "value"
			return true, nil
		})
		p.ReadAll()
		file.Close()
	}
}

func TestNextWithoutSteps(t *testing.T) {
	reader := strings.NewReader(`{"team": "team-c"}`)
	buf := new(bytes.Buffer)
	p := New(reader, buf)
	require.NotNil(t, p)

	// Reads everything
	require.NoError(t, p.Next())
	require.ErrorIs(t, p.Next(), io.EOF)

	require.Equal(t, "{\"team\":\"team-c\"}\n", buf.String())
}

func TestNextWithSingleTrueStep(t *testing.T) {
	reader := strings.NewReader(`{"team": "team-c"}`)
	buf := new(bytes.Buffer)
	p := New(reader, buf)
	require.NotNil(t, p)

	// Add step that accepts the value as is
	step := func(value *map[string]interface{}) (bool, error) {
		return true, nil
	}
	p.AddStep(step)

	// Reads everything
	require.NoError(t, p.Next())
	require.ErrorIs(t, p.Next(), io.EOF)

	// Verifies value is in output
	require.Equal(t, "{\"team\":\"team-c\"}\n", buf.String())

}

func TestNextWithSingleFalseStep(t *testing.T) {
	reader := strings.NewReader(`{"team": "team-c"}`)
	buf := new(bytes.Buffer)
	p := New(reader, buf)
	require.NotNil(t, p)

	// Add step that doesn't accept value
	step := func(value *map[string]interface{}) (bool, error) {
		return false, nil
	}
	p.AddStep(step)

	// Reads everything
	require.NoError(t, p.Next())
	require.ErrorIs(t, p.Next(), io.EOF)

	// Verify value from reader is not in the output
	require.Equal(t, "", buf.String())
}

func TestNextWithFailingStep(t *testing.T) {
	reader := strings.NewReader(`{"team": "team-c"}`)
	buf := new(bytes.Buffer)
	p := New(reader, buf)
	require.NotNil(t, p)

	// Add step that only returns an error
	step := func(value *map[string]interface{}) (bool, error) {
		return true, fmt.Errorf("Step failed")
	}
	p.AddStep(step)

	require.Error(t, p.Next(), "Step failed")
	require.ErrorIs(t, p.Next(), io.EOF)
}

func TestNextWithSideEffectStep(t *testing.T) {
	reader := strings.NewReader(`{"team": "team-c"}`)
	buf := new(bytes.Buffer)
	p := New(reader, buf)
	require.NotNil(t, p)

	// Add step that adds a new field
	step := func(value *map[string]interface{}) (bool, error) {
		(*value)["severity"] = 5
		return true, nil
	}
	p.AddStep(step)

	// Reads everything
	require.NoError(t, p.Next())
	require.ErrorIs(t, p.Next(), io.EOF)

	// Verifies new field has been added
	require.Equal(t, "{\"severity\":5,\"team\":\"team-c\"}\n", buf.String())
}

func TestNextWithMultipleSteps(t *testing.T) {
	reader := strings.NewReader(`
{"team": "team-x", "severity": 5}
{"team": "team-c", "severity": 2}
`)
	buf := new(bytes.Buffer)
	p := New(reader, buf)
	require.NotNil(t, p)

	// Add some steps
	p.AddStep(func(value *map[string]interface{}) (bool, error) {
		sev, ok := (*value)["severity"].(float64)
		if !ok {
			return false, nil
		}
		if sev > 2.0 {
			return true, nil
		}
		return false, nil
	})
	p.AddStep(func(value *map[string]interface{}) (bool, error) {
		(*value)["incident"] = 1234
		return true, nil
	})

	// Reads everything
	require.NoError(t, p.Next())
	require.NoError(t, p.Next())
	require.ErrorIs(t, p.Next(), io.EOF)

	// Verifies new field has been added
	require.Equal(t, "{\"incident\":1234,\"severity\":5,\"team\":\"team-x\"}\n", buf.String())
}

func TestReadAll(t *testing.T) {
	reader := strings.NewReader(`
{"team": "team-x", "severity": 5}
{"team": "team-c", "severity": 2}
`)
	buf := new(bytes.Buffer)
	p := New(reader, buf)
	require.NotNil(t, p)

	require.NoError(t, p.ReadAll())
	expectedOutput := `{"severity":5,"team":"team-x"}
{"severity":2,"team":"team-c"}
`
	require.Equal(t, expectedOutput, buf.String())

}
