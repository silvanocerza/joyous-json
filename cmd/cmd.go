package cmd

import (
	"github.com/silvanocerza/joyous-json/cmd/add"
	"github.com/silvanocerza/joyous-json/cmd/filter"
	"github.com/silvanocerza/joyous-json/cmd/prefix"
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
	root.AddCommand(filter.NewCommand())
	root.AddCommand(prefix.NewCommand())

	return root
}
