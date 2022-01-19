package cmd

import "github.com/spf13/cobra"

func Root() *cobra.Command {
	root := &cobra.Command{
		Use:     "jj",
		Short:   "TODO",
		Long:    "TODO",
		Example: "TODO",
	}

	return root
}
