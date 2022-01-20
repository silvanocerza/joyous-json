package add

import (
	"fmt"
	"os"

	"github.com/silvanocerza/joyous-json/pkg/processor"
	"github.com/silvanocerza/joyous-json/pkg/step"
	"github.com/spf13/cobra"
)

var Example string = `  jj add incident_id 6502
  jj add new_team team-w`

func NewCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "add <key> <value>",
		Short: "Adds a new key value pair",
		Long: "Adds a new key value pair to all processed JSON objects in the stream." +
			"If the key already exist it's overwritten.",
		Example: Example,
		Args:    cobra.ExactArgs(2),
		Run:     runAdd,
	}
}

func runAdd(cmd *cobra.Command, args []string) {
	p := processor.New(os.Stdin, os.Stdout)
	key, value := args[0], args[1]
	p.AddStep(step.NewAdd(key, value))
	err := p.ReadAll()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error processing: %v\n", err)
	}
}
