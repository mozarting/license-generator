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
			if license == "" || name == "" || year == "" {
				fmt.Println("ERROR: 'license', 'name', and 'year' flags are required.")
				os.Exit(1)
			}

			// get license files  from directory
			licenses, err := os.ReadDir("./licenses")
			if err != nil {
				fmt.Println("ERROR: Failed to read 'licenses' directory: ", err)
				os.Exit(1)
			}
			licenseFound := false
			// create license file
			for _, l := range licenses {
				if strings.EqualFold(l.Name(), license+".txt") {
					licenseFound = true
					file_content, err := os.ReadFile("licenses/" + l.Name())
					if err != nil {
						fmt.Printf("ERROR: Failed to read file '%s': %v\n", l.Name(), err)
						return
					}
					r := strings.NewReplacer("<year>", year, "<fullname>", name)
					fc := r.Replace(string(file_content))
					err = os.WriteFile("LICENSE", []byte(fc), 0644)
					if err != nil {
						fmt.Println("ERROR: Failed to create 'LICENSE' file: ", err)
						os.Exit(1)
					}

					fmt.Println("LICENSE file created successfully!")
					break
				}
			}

			if !licenseFound {
				fmt.Printf("ERROR: License '%s' not found in './licenses'.\n", license)
				fmt.Println("Available licenses", licenses)
				os.Exit(1)
			}

			fmt.Printf("License generated!\nLicense Type: %s\nAuthor Name: %s\nYear: %s\n", license, name, year)
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
