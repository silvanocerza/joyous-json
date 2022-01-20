package cmd

import (
	"github.com/silvanocerza/joyous-json/cmd/add"
	"github.com/spf13/cobra"
)

func Root() *cobra.Command {
	root := &cobra.Command{
		Use:     "jj",
		Short:   "TODO",
		Long:    "TODO",
		Example: "TODO",
	}

	root.AddCommand(add.NewCommand())

	return root
}
