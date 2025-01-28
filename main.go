package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

func main() {
	var license string
	var name string
	var year string

	var rootCmd = &cobra.Command{
		Use:   "license-generator",
		Short: "CLI for open-source License generator",
	}

	var addCmd = &cobra.Command{
		Use:   "add",
		Short: "Add license in current directory",
		Run: func(cmd *cobra.Command, args []string) {
			// get license files  from directory
			licenses, err := os.ReadDir("./licenses")
			if err != nil {
				fmt.Println("ERROR: reading directory", err)
			}
			// create license file
			for _, l := range licenses {
				fmt.Println(l.Name())
				if strings.EqualFold(l.Name(), license) {
					file_content, err := os.ReadFile("licenses/" + l.Name())
					if err != nil {
						fmt.Println("ERROR: Failed to read file, ", err)
						return
					}
					r := strings.NewReplacer("<year>", year, "<copyright holder>", name)
					fmt.Println(r.Replace(string(file_content)))
					break
				}
			}
			fmt.Printf("License generated!\n license_type : %s\n author_name : %s\n year: %s\n", license, name, year)
		},
	}

	addCmd.Flags().StringVarP(&license, "license", "l", "", "License name")
	addCmd.Flags().StringVarP(&name, "name", "n", "", "Full name of author")
	addCmd.Flags().StringVarP(&year, "year", "y", "", "The year in license")

	rootCmd.AddCommand(addCmd)
	// Execute the root command
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
