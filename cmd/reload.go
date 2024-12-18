package cmd

import (
	"fmt"
	"surge/api/profile"

	"github.com/spf13/cobra"
)

var reloadCmd = &cobra.Command{
	Use:   "reload",
	Short: "重载当前 Surge 配置",
	Run: func(cmd *cobra.Command, args []string) {
		_, err := profile.Reload()
		if err != nil {
			fmt.Print("failed", err)
		} else {
			fmt.Println("success")
		}
	},
}

func init() {
	rootCmd.AddCommand(reloadCmd)
}
