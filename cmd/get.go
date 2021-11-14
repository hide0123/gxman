package cmd

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

func install(url string) {
	path := "/opt/devkitpro/tools/bin"
	if _, err := os.Stat(path); os.IsNotExist(err) {
		fmt.Println("error:", path, "is not found. Please install devkitPro.")
		os.Exit(1)
	}

	fmt.Println("Installing...")

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("error", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	file, err := os.Create(path + "/3gxtool")
	if err != nil {
		fmt.Println("error", err)
		os.Exit(1)
	}
	defer file.Close()

	io.Copy(file, resp.Body)
	os.Chmod(path+"/3gxtool", 0755)

	fmt.Println("3gxtool was successfully installed in", path)
}

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Install the latest 3gxtool",
	Run: func(cmd *cobra.Command, args []string) {
		normal := "https://cdn.discordapp.com/attachments/463776354899460101/718621100665208842/3gxtool"
		patched := "https://cdn.discordapp.com/attachments/479233444249862174/908696881297760296/3gxtool"

		unlimited, err := cmd.Flags().GetBool("unlimited")
		if err != nil {
			fmt.Println("error", err)
			os.Exit(1)
		}

		if unlimited {
			install(patched)
		} else {
			install(normal)
		}
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
	getCmd.Flags().BoolP("unlimited", "u", false, "1 MiB limit removed")
}
