package prefix

import (
	"fmt"
	"os"

	"github.com/silvanocerza/joyous-json/pkg/processor"
	"github.com/silvanocerza/joyous-json/pkg/step"
	"github.com/spf13/cobra"
)

var Example string = `  jj prefix team old_
  jj prefix id incident_`

func NewCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "prefix <key> <prefix>",
		Short: "Prefixes a key with a custom prefix",
		Long: "Prefixes a key with a custom prefix if the key is found in a JSON object.\n" +
			"If the prefixed key already exists it's overwritten.",
		Example: Example,
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
