/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"consul-companion/internal/cfg"
	"consul-companion/internal/core"

	"github.com/spf13/cobra"
)

// sdCmd represents the sd command
var sdCmd = &cobra.Command{
	Use:   "sd",
	Short: "Service Discovery on host",
	Long:  `Service Discovery on host`,
	Run: func(cmd *cobra.Command, args []string) {
		search, _ := cmd.Flags().GetString("search")
		confdir, _ := cmd.Flags().GetString("conf-dir")

		if args[0] == "watch" {
			cfg.SDConfig(search, confdir)
			core.RunWatch()
		}
	},
}

func init() {
	rootCmd.AddCommand(sdCmd)
	sdCmd.Flags().StringP("search", "s", "/opt", "path to search project")
	sdCmd.Flags().StringP("conf-dir", "c", "/etc/consul/consul.d", "path to config")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// sdCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// sdCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	sdCmd.Flags().SortFlags = true
}
