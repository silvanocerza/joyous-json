package cmd

import (
	"strings"

	"github.com/silvanocerza/joyous-json/cmd/add"
	"github.com/silvanocerza/joyous-json/cmd/filter"
	"github.com/silvanocerza/joyous-json/cmd/prefix"
	"github.com/spf13/cobra"
)

func Root() *cobra.Command {
	root := &cobra.Command{
		Use:   "jj",
		Short: "jj is used to process JSON objects streams",
		Long: "jj is a command line tool to manipulate JSON objects stream.\n" +
			"It supports adding new key value pairs, prefixing existing keys and\n" +
			"filtering keys based on their values.",
		Example: strings.Join([]string{add.Example, filter.Example, prefix.Example}, "\n"),
	}

	root.AddCommand(add.NewCommand())
	root.AddCommand(filter.NewCommand())
	root.AddCommand(prefix.NewCommand())

	return root
}
