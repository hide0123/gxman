package cmd

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"runtime"

	"github.com/spf13/cobra"
)

func install(name string, url string) {
	dkp := os.Getenv("DEVKITPRO")
	if dkp == "" {
		fmt.Println("error: Please set the DEVKITPRO environment.")
		os.Exit(1)
	}
	path := dkp + "/tools/bin"

	fmt.Println("Installing...")

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("error", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	file, err := os.Create(path + "/" + name)
	if err != nil {
		fmt.Println("error", err)
		os.Exit(1)
	}
	defer file.Close()

	io.Copy(file, resp.Body)

	if runtime.GOOS == "linux" {
		os.Chmod(path+"/"+name, 0755)
	}

	fmt.Println("3gxtool was successfully installed in", path)
}

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Install the latest 3gxtool",
	Run: func(cmd *cobra.Command, args []string) {
		var bin, normal, patched string
		if runtime.GOOS == "windows" {
			bin = "3gxtool.exe"
			normal = "https://cdn.discordapp.com/attachments/479233979271086090/707634663765573753/3gxtool.exe"
			patched = "https://cdn.discordapp.com/attachments/512385640852357150/825415395036758057/3gxtool.exe"
		} else if runtime.GOOS == "linux" {
			bin = "3gxtool"
			normal = "https://cdn.discordapp.com/attachments/463776354899460101/718621100665208842/3gxtool"
			patched = "https://cdn.discordapp.com/attachments/479233444249862174/908696881297760296/3gxtool"
		} else {
			fmt.Println("error: Unsupported OS.")
			os.Exit(1)
		}

		unlimited, err := cmd.Flags().GetBool("unlimited")
		if err != nil {
			fmt.Println("error", err)
			os.Exit(1)
		}

		if unlimited {
			install(bin, patched)
		} else {
			install(bin, normal)
		}
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
	getCmd.Flags().BoolP("unlimited", "u", false, "1 MiB limit removed")
}
