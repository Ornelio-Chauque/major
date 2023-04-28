package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	majorCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("You are using version 0.3.0")
	},

	Use:   "version",
	Short: "",
	Long:  "",
}
