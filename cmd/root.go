package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gxman",
	Short: "A tool for 3gxtool",
	Long: `gxman is a CLI tool for 3gxtool.
This tool is used to install the latest one
or to check the type of the installed one.`,
	Run: func(cmd *cobra.Command, args []string) {
		version, err := cmd.Flags().GetBool("version")
		if err != nil {
			fmt.Println("error", err)
			os.Exit(1)
		}

		if version {
			fmt.Println("gxman version 1.0.1")
		} else {
			cmd.Help()
		}
	},
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.Flags().BoolP("version", "v", false, "print the version")
}
