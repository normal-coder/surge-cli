package cmd

import (
	"fmt"
	"surge/api/feature/system_proxy"

	"github.com/spf13/cobra"
)

var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "停止 Surge（取消设置为系统代理）", //Example: cfgFile,
	Run: func(cmd *cobra.Command, args []string) {
		_, err := system_proxy.Reviser(false)
		if err != nil {
			fmt.Print("failed", err)
		} else {
			fmt.Println("success")
		}

	},
}

func init() {
	rootCmd.AddCommand(stopCmd)
}
