package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"runtime"

	"github.com/spf13/cobra"
)

var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Check the type of your 3gxtool",
	Run: func(cmd *cobra.Command, args []string) {
		dkp := os.Getenv("DEVKITPRO")
		if dkp == "" {
			fmt.Println("error: Please set the DEVKITPRO environment.")
			os.Exit(1)
		}

		var offset int
		path := dkp

		if runtime.GOOS == "windows" {
			path += "/tools/bin/3gxtool.exe"
			offset = 0x11B1
		} else if runtime.GOOS == "linux" {
			path += "/tools/bin/3gxtool"
			offset = 0x2AF7
		} else {
			fmt.Println("error: Unsupported OS.")
			os.Exit(1)
		}

		file, err := ioutil.ReadFile(path)
		if err != nil {
			fmt.Println("error:", err)
			os.Exit(1)
		}

		flag := file[offset]

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
