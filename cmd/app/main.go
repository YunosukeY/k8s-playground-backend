package main

import (
	"os"

	"github.com/YunosukeY/kind-backend/internal/app"
	"github.com/YunosukeY/kind-backend/internal/auth"
	"github.com/YunosukeY/kind-backend/internal/mail"
	"github.com/spf13/cobra"
)

func main() {
	execute()
}

var rootCmd = &cobra.Command{}

func execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
func init() {
	rootCmd.AddCommand(appCmd)
	rootCmd.AddCommand(authCmd)
	rootCmd.AddCommand(mailCmd)
}

var appCmd = &cobra.Command{
	Use:   "app",
	Short: "Run an app server",
	Run: func(cmd *cobra.Command, args []string) {
		app.Run()
	},
}

var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "Run an auth server",
	Run: func(cmd *cobra.Command, args []string) {
		auth.Run()
	},
}

var mailCmd = &cobra.Command{
	Use:   "mail",
	Short: "Run an mail server",
	Run: func(cmd *cobra.Command, args []string) {
		mail.Run()
	},
}
