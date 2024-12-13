package cmd

import (
	"fmt"
	"surge/api/feature/system_proxy"

	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "启动 Surge（设置为系统代理）",
	Run: func(cmd *cobra.Command, args []string) {
		_, err := system_proxy.Reviser(true)
		if err != nil {
			fmt.Print("failed", err)
		} else {
			fmt.Println("success")
		}
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
