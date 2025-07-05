package cmd

import "github.com/spf13/cobra"

func init() {
	subCmd.PersistentFlags().StringP("subCommandVarPers", "p", "", "subComand persisted")
	subCmd.Flags().StringVarP(&cmd1var, "var1", "v", "", "subcommandVar")
}

var (
	// vars used for flags
	cmd1var string

	// Root command
	subCmd = &cobra.Command{
		Use:   "subCmd",
		Short: "subCmd short description",
		Long:  "subCmd long description",
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
		},
	}
)
