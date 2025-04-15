package cmd

import (
	"cobra_learning/student"
	"fmt"
	"github.com/spf13/cobra"
)

var name string
var math float32
var english float32
var age int

func init() {

	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(showCmd)
	//本地标志
	addCmd.Flags().StringVarP(&name, "name", "n", "", "student name set")
	addCmd.Flags().Float32VarP(&math, "math", "m", 0, "student math set")
	addCmd.Flags().IntVarP(&age, "age", "a", 0, "student age set")
	addCmd.Flags().Float32VarP(&english, "english", "e", 0, "student english set")
}

var addCmd = &cobra.Command{
	Use:   "add <student name>",
	Short: "add a student",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Add called")
		//可以在此进行flag判断来进行分支处理
		if cmd.Flags().Changed("name") {
			//如果指定了--name
			fmt.Println("You add a student named", name)
		}

	},
}

var showCmd = &cobra.Command{
	Use:   "show",
	Short: "show all students",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Show called")
		student.AddStudent("zhangsan", 15, 33, 23)
	},
}
