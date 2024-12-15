package cmd

import (
	"fmt"
	"surge/api/modules"

	"github.com/spf13/cobra"
)

var disableCmd = &cobra.Command{
	Use:   "disable",
	Short: "禁用模块",
	Run: func(cmd *cobra.Command, args []string) {
		_, err := modules.Disable(args)
		if err != nil {
			fmt.Print("failed", err)
		} else {
			fmt.Println("success")
		}
	},
}

func init() {
	rootCmd.AddCommand(disableCmd)
}
