/*
Copyright © 2026
*/
package cmd

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

// logoutCmd represents the logout command
var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "Remove local authentication config",
	RunE: func(cmd *cobra.Command, args []string) error {
		home, err := os.UserHomeDir()
		if err != nil {
			return fmt.Errorf("cannot resolve home directory: %w", err)
		}

		path := filepath.Join(home, ".xplr-mq", "config.yaml")
		if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
			fmt.Println("Already logged out")
			return nil
		}

		if err := os.Remove(path); err != nil {
			return fmt.Errorf("failed to remove auth config: %w", err)
		}

		fmt.Println("✅ Logged out")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(logoutCmd)
}
