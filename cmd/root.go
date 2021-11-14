package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gxman",
	Short: "A tool for 3gxtool",
	Long: `gxman is a CLI tool for 3gxtool.
This tool is used to install the latest one
or to check the type of the installed one.`,
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
