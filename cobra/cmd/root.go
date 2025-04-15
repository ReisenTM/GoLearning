package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var str string
var rootCmd = &cobra.Command{
	Use:   "stu",
	Short: "stu manager",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		fmt.Println("Root PersistentPreRun called")
	},

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("PersistentFlag(info)", str)
	},

	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		fmt.Println("Root PersistentPostRun called")
	},
}

// bind flags
func init() {
	//全局标志
	rootCmd.PersistentFlags().StringVarP(&str, "info", "i", "", "A info here")

}
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
