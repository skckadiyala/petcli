/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"context"
	"fmt"

	"github.com/skckadiyala/kubecrt-vms/utils"
	"github.com/skckadiyala/petstore"
	"github.com/spf13/cobra"
)

// petsCmd represents the pets command
var (
	petAddCmd = &cobra.Command{
		Use:   "pet",
		Short: "Add a pet",
		Long: `A longer description that spans multiple lines and likely contains examples
		and usage of using your command. For example:

		Cobra is a CLI library for Go that empowers applications.
		This application is a tool to generate the needed files
		to quickly create a Cobra application.`,
		Run: addPet,
	}
	petListCmd = &cobra.Command{
		Use:   "pets",
		Short: "list a pet",
		Long: `A longer description that spans multiple lines and likely contains examples
		and usage of using your command. For example:

		Cobra is a CLI library for Go that empowers applications.
		This application is a tool to generate the needed files
		to quickly create a Cobra application.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("pets called")
		},
	}
	petDelCmd = &cobra.Command{
		Use:   "pet",
		Short: "Add a pet",
		Long: `A longer description that spans multiple lines and likely contains examples
		and usage of using your command. For example:

		Cobra is a CLI library for Go that empowers applications.
		This application is a tool to generate the needed files
		to quickly create a Cobra application.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("pets called")
		},
	}
	petEditCmd = &cobra.Command{
		Use:   "pet",
		Short: "Add a pet",
		Long: `A longer description that spans multiple lines and likely contains examples
		and usage of using your command. For example:

		Cobra is a CLI library for Go that empowers applications.
		This application is a tool to generate the needed files
		to quickly create a Cobra application.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("pets called")
		},
	}
)

func init() {
	addCmd.AddCommand(petAddCmd)
	listCmd.AddCommand(petListCmd)
	deleteCmd.AddCommand(petDelCmd)
	editCmd.AddCommand(petEditCmd)

	addCmd.Flags().StringVarP(&file, "swagger", "f", "", "The filename of the swagger api to be stored")
	addCmd.MarkFlagRequired("swagger")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// petsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// petsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func addPet(cmd *cobra.Command, args []string) {
	// utils.PrettyPrintInfo("Creating backend API")

	cfg := getConfig()
	fmt.Println("Host in CFG", cfg.Host)
	// orgID := getOrganizationByName(args)

	client := petstore.NewAPIClient(cfg)

	// file, err := os.Open(file)
	// if err != nil {
	// 	utils.PrettyPrintErr("Error Opening file: %v", err)
	// }

	body := *petstore.NewPet("doggie", []string{"PhotoUrls_example"})
	// body := *petstore.NewPetWithDefaults()

	resp, err := client.PetApi.AddPet(context.Background()).Body(body).Execute()
	// beAPI, _, err := client.APIRepositoryApi.ApirepoImportPost(context.Background(), orgID, apiName, "swagger", file)
	if err != nil {
		utils.PrettyPrintErr("Error Adding a pet: %v", err)
		return
	}
	utils.PrettyPrintInfo("Pet Added to the request %v ", resp.Body)
	return
}
