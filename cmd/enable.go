package cmd

import (
	"fmt"
	"surge/api/modules"

	"github.com/spf13/cobra"
)

var enableCmd = &cobra.Command{
	Use:   "enable",
	Short: "启用模块",
	Run: func(cmd *cobra.Command, args []string) {
		_, err := modules.Enable(args)
		if err != nil {
			fmt.Print("failed", err)
		} else {
			fmt.Println("success")
		}
	},
}

func init() {
	rootCmd.AddCommand(enableCmd)
}
