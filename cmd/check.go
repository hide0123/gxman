package cmd

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/spf13/cobra"
)

var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Check the type of your 3gxtool",
	Run: func(cmd *cobra.Command, args []string) {
		file, err := ioutil.ReadFile("/opt/devkitpro/tools/bin/3gxtool")
		if err != nil {
			fmt.Println("error:", err)
			os.Exit(1)
		}

		flag := file[0x2AF7]

		if flag == 0x0F {
			fmt.Println("Your 3gxtool is official.")
		} else if flag == 0xFF {
			fmt.Println("Your 3gxtool is patched. (No 1 Mib limit)")
		} else {
			fmt.Println("Couldn't determine your 3gxtool type.")
		}
	},
}

func init() {
	rootCmd.AddCommand(checkCmd)
}
