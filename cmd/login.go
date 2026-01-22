/*
Copyright © 2026
*/
package cmd

import (
	"fmt"
	"xplr-distributed-mq/internal/config"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	loginURL   string
	loginToken string
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Configure endpoint and access token",
	RunE: func(cmd *cobra.Command, args []string) error {
		if loginURL == "" || loginToken == "" {
			return fmt.Errorf("--url and --token are required")
		}

		cfg := config.AuthConfig{
			URL:   loginURL,
			Token: loginToken,
		}

		if err := config.SaveAuthConfig(cfg); err != nil {
			return err
		}

		fmt.Println("✅ Login successful")
		fmt.Println("Endpoint:", loginURL)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
	loginCmd.Flags().StringVar(
		&loginURL,
		"url",
		"",
		"control plane URL (or env XPLR_URL)",
	)

	loginCmd.Flags().StringVar(
		&loginToken,
		"token",
		"",
		"access token (or env XPLR_TOKEN)",
	)

	_ = viper.BindPFlag("url", loginCmd.Flags().Lookup("url"))
	_ = viper.BindPFlag("token", loginCmd.Flags().Lookup("token"))

	viper.SetEnvPrefix("xplr_mq")
	viper.AutomaticEnv()
}
