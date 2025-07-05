package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	cobra.OnInitialize(init1, init2)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cobra.yaml)")
	rootCmd.PersistentFlags().StringP("author", "a", "YOUR NAME", "author name for copyright attribution")
	rootCmd.Flags().StringVarP(&userLicense, "license", "l", "", "name of license for the project")
	rootCmd.Flags().Bool("viper", true, "use Viper for configuration")
	viper.BindPFlag("author", rootCmd.PersistentFlags().Lookup("author"))
	viper.BindPFlag("useViper", rootCmd.Flags().Lookup("viper"))
	viper.SetDefault("author", "NAME HERE <EMAIL ADDRESS>")
	viper.SetDefault("license", "apache")

	rootCmd.AddCommand(subCmd)
}

var (
	// Global vars used for flags
	cfgFile, userLicense string

	// Root command
	rootCmd = &cobra.Command{
		Use:   "root",
		Short: "Root short description",
		Long:  "Root long description",
	}
)

func Execute() error {
	return rootCmd.Execute()
}

func init1() {
	return
}

func init2() {
	return
}
