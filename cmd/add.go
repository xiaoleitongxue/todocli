/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"log"

	"todocli/todo"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new todo",
	Long:  `Add will create a new todo item to the list`,
	Run:   addRun,
}
var priority int

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	addCmd.Flags().IntVarP(&priority, "priority", "p", 2, "Priority: 1, 2, 3")
}

// Function is a type in go.
func addRun(cmd *cobra.Command, args []string) {
	// An expression that creates a new value each time it is evaluated
	// var items = []todo.Item{}
	items, err := todo.ReadItems(dataFile)
	if err != nil {
		log.Printf("%v", err)
	}
	for _, x := range args {
		item := todo.Item{Text: x}
		item.SetPriority(priority)
		items = append(items, item)
	}
	if err := todo.SaveItems(dataFile, items); err != nil {
		fmt.Println(fmt.Errorf("%v", err))
	}
}
