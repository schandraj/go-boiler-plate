/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"base-plate/app"
	"base-plate/app/api"
	"base-plate/config"
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"log"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Run the service",
	Long:  `RUN RUN RUN`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := config.Load(); err != nil {
			log.Fatal(err)
		}
		fmt.Println("serve called")

		ctx := context.Background()

		c := app.NewContainer(ctx)

		//go listener.Listen(ctx, c)

		api.Serve(ctx, c)
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
