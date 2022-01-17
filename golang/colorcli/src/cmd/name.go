/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"encoding/json"
	"io/ioutil"
	"github.com/spf13/cobra"
)

// nameCmd represents the name command
var nameCmd = &cobra.Command{
	Use:   "name",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println("name called")
		hexToName(args)
	},
}

func hexToName(args []string) {
	var hexMap map[string]string

	//read the input file
	content, err := ioutil.ReadFile("colornames.min.json")

	// display error if unable to read file
	if err != nil {
		fmt.Printf("Error while reading file: %v", err)
	}

	// parse json structure
	_ = json.Unmarshal(content, &hexMap)

	name, ok := hexMap[args[0]]  
	if ok {
		fmt.Printf("Name: %s, Hex %s\n", name, args[0])
	} else {
		fmt.Println("Color name not found")
	}

}

func init() {
	rootCmd.AddCommand(nameCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// nameCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// nameCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
