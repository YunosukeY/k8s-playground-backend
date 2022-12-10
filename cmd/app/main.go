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
	appCmd.Flags().BoolP("dummy", "d", false, "run with dummy middleware")
	rootCmd.AddCommand(authCmd)
	authCmd.Flags().BoolP("dummy", "d", false, "run with dummy middleware")
	rootCmd.AddCommand(mailCmd)
	mailCmd.Flags().BoolP("dummy", "d", false, "run with dummy middleware")
}

var appCmd = &cobra.Command{
	Use:   "app",
	Short: "Run an app server",
	Run: func(cmd *cobra.Command, args []string) {
		dummy, _ := cmd.Flags().GetBool("dummy")
		app.Run(dummy)
	},
}

var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "Run an auth server",
	Run: func(cmd *cobra.Command, args []string) {
		dummy, _ := cmd.Flags().GetBool("dummy")
		auth.Run(dummy)
	},
}

var mailCmd = &cobra.Command{
	Use:   "mail",
	Short: "Run an mail server",
	Run: func(cmd *cobra.Command, args []string) {
		dummy, _ := cmd.Flags().GetBool("dummy")
		mail.Run(dummy)
	},
}
