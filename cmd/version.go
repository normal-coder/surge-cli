package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "查看当前版本号",
	Long:  `All software has versions. This is Surge-CLI's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Surge-CLI version is v1.0")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
