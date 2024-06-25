/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"consul-companion/api_consul"
	"consul-companion/internal/cfg"

	"github.com/spf13/cobra"
)

// apiCmd represents the api command
var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "A brief description of your command",
	Long:  `.`,
	Run: func(cmd *cobra.Command, args []string) {
		t, _ := cmd.Flags().GetString("target")

		if args[0] == "deregister" {
			config := cfg.GetConfig()
			svcList := api_consul.GetNodeServices(config, t)

			for _, r := range svcList.Services {
				api_consul.DeregisterService(config, svcList.Node.Node, r.ID, svcList.Node.Address)
			}
		}

	},
}

func init() {
	rootCmd.AddCommand(apiCmd)
	apiCmd.Flags().StringP("target", "t", "127.0.0.1", "host address")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// apiCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// apiCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	sdCmd.Flags().SortFlags = true
}
