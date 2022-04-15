/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"

	"golang.org/x/crypto/ssh/terminal"

	"github.com/c-bata/go-prompt"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

// addPetCmd represents the addPet command
var (
	loginCmd = &cobra.Command{
		Use:   "login",
		Short: "Stores the login info of petsotr",
		Long: `Stores the login info of petstore. Stores
		Hostname, 
		Port, 
	`,
		Run: login,
	}

	addCmd = &cobra.Command{
		Use:   "add",
		Short: "add a pet or user",
		Long: `add a pet or user . For example:


	# Add a aoo using the data in user.json
	petcli add pet -f ./org.json`,
		// Run: func(cmd *cobra.Command, args []string) {
		// 	fmt.Println("addPet called")
		// },
	}

	listCmd = &cobra.Command{
		Use:   "list",
		Short: "list a pets or users",
		Long: `list a pets or users . For example:
	
	
		# Add a aoo using the data in user.json
		petcli list pet `,
		// Run: func(cmd *cobra.Command, args []string) {
		// 	fmt.Println("addPet called")
		// },
	}

	deleteCmd = &cobra.Command{
		Use:   "delete",
		Short: "delete a pet or user",
		Long: `delete a pet or user . For example:
	
	
		# Add a aoo using the data in user.json
		petcli delete pet id`,
		// Run: func(cmd *cobra.Command, args []string) {
		// 	fmt.Println("addPet called")
		// },
	}

	editCmd = &cobra.Command{
		Use:   "edit",
		Short: "edit a pet or user",
		Long: `edit a pet or user . For example:
	
	
		# Add a aoo using the data in user.json
		petcli edit pet -f ./org.json`,
		// Run: func(cmd *cobra.Command, args []string) {
		// 	fmt.Println("addPet called")
		// },
	}
)

func init() {
	rootCmd.AddCommand(loginCmd)
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(deleteCmd)
	rootCmd.AddCommand(editCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addPetCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addPetCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func completer(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: "host", Description: "Store PetStore Host"},
		{Text: "port", Description: "Store PetStore Port"},
		{Text: "username", Description: "Store the PetStore username"},
		{Text: "password", Description: "Store the PetStore password"},
	}
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}

func login(cmd *cobra.Command, args []string) {

	fmt.Print("\nPetStore Hostname")
	host := prompt.Input(": ", completer)
	fmt.Print("PetStore Port")
	port := prompt.Input(": ", completer)
	fmt.Print("Username")
	username := prompt.Input(": ", completer)
	fmt.Println("Password")
	password, err := terminal.ReadPassword(0)

	// password := prompt.Input(": ", completer)

	conf := configAPI{}

	fmt.Println("Password", string(password))
	data := []byte(username + ":" + string(password))
	basicAuth := base64.StdEncoding.EncodeToString(data)

	conf.PetHost = host
	conf.PetPort = port
	conf.Authorization = basicAuth

	out, err := yaml.Marshal(conf)
	if err != nil {
		fmt.Println("Error to marshal config yaml", err)
		return
	}
	home := os.Getenv("HOME")
	err = ioutil.WriteFile(home+"/.petcli.yaml", out, 0644)
	if err != nil {
		fmt.Println("Error to write config yaml file", err)

	}
}
