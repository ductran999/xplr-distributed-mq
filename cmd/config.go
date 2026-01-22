/*
Copyright © 2026 Duc Ngo
*/
package cmd

import (
	"fmt"
	"os"
	"xplr-distributed-mq/internal/config"

	"github.com/spf13/cobra"
)

var configPathCmd = &cobra.Command{
	Use:   "path",
	Short: "Show config file path",
	RunE: func(cmd *cobra.Command, args []string) error {
		path, err := config.ConfigPath()
		if err != nil {
			return err
		}
		if _, err := os.Stat(path); err == nil {
			fmt.Println(path)
		} else {
			fmt.Println(path, "(not created yet)")
		}

		return nil
	},
}

var configViewCmd = &cobra.Command{
	Use:   "view",
	Short: "View current config",
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := config.LoadAuthConfig()
		if err != nil {
			return err
		}

		fmt.Println("URL:", cfg.URL)
		fmt.Println("Token:", "******")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(configPathCmd)
	rootCmd.AddCommand(configViewCmd)

}
