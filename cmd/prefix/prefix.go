package prefix

import (
	"fmt"
	"os"

	"github.com/silvanocerza/joyous-json/pkg/processor"
	"github.com/silvanocerza/joyous-json/pkg/step"
	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	return &cobra.Command{
		Use:     "prefix",
		Short:   "",
		Long:    "",
		Example: "",
		Args:    cobra.ExactArgs(2),
		Run:     runPrefix,
	}
}

func runPrefix(cmd *cobra.Command, args []string) {
	p := processor.New(os.Stdin, os.Stdout)
	key, prefix := args[0], args[1]
	p.AddStep(step.NewPrefix(key, prefix))
	err := p.ReadAll()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error processing: %v\n", err)
	}
}
