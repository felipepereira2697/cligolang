package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{Use: "aryago"}
	var name, pass string
	var cmd = &cobra.Command{
		Use:   "create",
		Short: "Create an user",
		Run: func(cmd *cobra.Command, args []string) {
			if name == "" {
				fmt.Println("Sorry, name cannot be blank")
				return
			}

			if len(pass) < 10 {
				fmt.Println("Sorry, pass to weak, please provide a new pass")
				return
			}

			fmt.Printf("Name: %s\nPass: %s\n", name, pass)
		},
	}

	cmd.Flags().StringVarP(&name, "name", "n", "", "User name")
	cmd.Flags().StringVarP(&pass, "pass", "p", "", "Password")

	//Command sample: ./cligolang create -n Stark -p Stark1234
	rootCmd.AddCommand(cmd)
	rootCmd.Execute()

}
