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

// addcolorCmd represents the addcolor command
var addcolorCmd = &cobra.Command{
	Use:   "addcolor",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		addColor(args)
	},
}

func addColor(args []string) {
	hex := args[0]
	colorName := args[1]

	content, err := ioutil.ReadFile("colornames.min.json")

	if err != nil {
		fmt.Printf("Error while reading file: %v", err)
	}

	var hexMap map[string]string
	_ = json.Unmarshal(content, &hexMap)

	colorName, ok := hexMap[hex]

	if ok {
		fmt.Printf("Hex already exists. Color name is: %s\n", colorName)
	} else {
		hexMap[hex] = colorName

		hexJSON, _ := json.Marshal(hexMap)

		err = ioutil.WriteFile("colornames.min.json", hexJSON, 0777)

		if err != nil {
			fmt.Printf("Error while writing file: %v", err)
		}

		fmt.Println("Hex code added sucessfully!")
	}
}

func init() {
	rootCmd.AddCommand(addcolorCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addcolorCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addcolorCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
