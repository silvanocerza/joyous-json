package filter

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/silvanocerza/joyous-json/pkg/processor"
	"github.com/silvanocerza/joyous-json/pkg/step"
	"github.com/spf13/cobra"
)

var Example = `  jj filter out team team-x
  jj filter out timestamp \>=1642328776
  jj filter in severity \!0`

func NewCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "filter",
		Short: "Filter in or out JSON objects based on their key values",
		Long: "Filter in or out JSON objects. It can filter wether a key exists\n" +
			"or not. It can also filter based on a key value, wther it's greater, less,\n" +
			"equal or unequal.",
		Example: Example,
		Args:    cobra.RangeArgs(2, 3),
		Run:     runFilter,
	}
}

// Possible operations executed by the filter command
const (
	not int32 = iota
	equal
	greaterOrEqual
	lessOrEqual
	greater
	less
	exist
	notExist
)

func runFilter(cmd *cobra.Command, args []string) {
	action, key, value := "", "", ""
	if len(args) == 3 {
		action, key, value = args[0], args[1], args[2]
	} else {
		action, key = args[0], args[1]
	}

	switch action {
	case "in":
	case "out":
	default:
		fmt.Fprintf(os.Stderr, "Unknown action \"%s\"\n\n", action)
		cmd.Help()
		os.Exit(1)
	}

	operation, value := getOperation(action, value)
	s := getStep(operation, key, value)
	if s == nil {
		fmt.Fprintf(os.Stderr, "Error finding step for %d, %s, %s\n", operation, key, value)
		os.Exit(1)
	}

	p := processor.New(os.Stdin, os.Stdout)
	p.AddStep(s)
	err := p.ReadAll()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error processing: %v\n", err)
	}
}

// getOperation returns the appropriate operation to run given the
// action and value set by the user.
// It also returns the value without operations prefixes:
// * !
// * =
// * >=
// * <=
// * >
// * <
func getOperation(action, value string) (int32, string) {
	var operations map[string]int32
	if value == "" {
		operations = map[string]int32{
			"in":  exist,
			"out": notExist,
		}
	} else if strings.HasPrefix(value, "!") {
		value = strings.TrimPrefix(value, "!")
		operations = map[string]int32{
			"in":  not,
			"out": equal,
		}
	} else if strings.HasPrefix(value, "=") {
		value = strings.TrimPrefix(value, "=")
		operations = map[string]int32{
			"in":  equal,
			"out": not,
		}
	} else if strings.HasPrefix(value, ">=") {
		value = strings.TrimPrefix(value, ">=")
		operations = map[string]int32{
			"in":  greaterOrEqual,
			"out": less,
		}
	} else if strings.HasPrefix(value, "<=") {
		value = strings.TrimPrefix(value, "<=")
		operations = map[string]int32{
			"in":  lessOrEqual,
			"out": greater,
		}
	} else if strings.HasPrefix(value, ">") {
		value = strings.TrimPrefix(value, ">")
		operations = map[string]int32{
			"in":  greater,
			"out": lessOrEqual,
		}
	} else if strings.HasPrefix(value, "<") {
		value = strings.TrimPrefix(value, "<")
		operations = map[string]int32{
			"in":  greater,
			"out": lessOrEqual,
		}
	} else {
		operations = map[string]int32{
			"in":  equal,
			"out": not,
		}
	}
	return operations[action], value
}

// getStep returns the appropriate step function to run given the operation, key and value
// Returns nil if operation is not known.
func getStep(operation int32, key, value string) processor.StepFunc {
	switch operation {
	case not:
		return step.NewFilterInIfDifferent(key, value)
	case equal:
		return step.NewFilterInIfEqual(key, value)
	case greaterOrEqual:
		v, _ := strconv.ParseFloat(value, 64)
		return step.NewFilterInGreaterOrEqualThan(key, v)
	case lessOrEqual:
		v, _ := strconv.ParseFloat(value, 64)
		return step.NewFilterInLessOrEqualThan(key, v)
	case greater:
		v, _ := strconv.ParseFloat(value, 64)
		return step.NewFilterInGreaterThan(key, v)
	case less:
		v, _ := strconv.ParseFloat(value, 64)
		return step.NewFilterInLessThan(key, v)
	case exist:
		return step.NewFilterIn(key)
	case notExist:
		return step.NewFilterOut(key)
	}
	return nil
}
