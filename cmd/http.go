/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
	"rmock/service"
)

// httpCmd represents the http command
var httpCmd = &cobra.Command{
	Use:   "http",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		port, err := cmd.Flags().GetInt("port")
		if err != nil {
			panic(err)
		}
		json, err := cmd.Flags().GetString("json")
		if err != nil {
			panic(err)
		}
		service.ServeHttp(port, json)
	},
}

func init() {
	rootCmd.AddCommand(httpCmd)
	httpCmd.Flags().IntP("port", "p", 8083, "Port to run the server on")
	httpCmd.Flags().StringP("json", "j", "", "Location of the JSON file to serve")
}
