/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"os"

	"github.com/YunosukeY/kind-backend/internal/migration"
	"github.com/YunosukeY/kind-backend/internal/util"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

func main() {
	execute()
}

const (
	UP     = "up"
	DOWN   = "down"
	SCHEMA = "schema"
	RECORD = "record"
)

var rootCmd = &cobra.Command{
	Run: func(cmd *cobra.Command, args []string) {
		op, _ := cmd.Flags().GetString("operation")
		t, _ := cmd.Flags().GetString("type")

		if (op != UP && op != DOWN) && (t != SCHEMA && t != RECORD) {
			log.Info().Msg("invalid params")
			return
		}

		if err := util.LoadEnv(); err != nil {
			log.Panic().Err(err).Msg("")
			panic(err)
		}
		if op == UP {
			if t == SCHEMA {
				migration.UpSchema()
			} else if t == RECORD {
				migration.UpTestData()
			}
		} else if op == DOWN {
			if t == SCHEMA {
				migration.DownSchema()
			} else if t == RECORD {
				migration.DownTestData()
			}
		}
	},
}

func execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringP("operation", "o", "", "up or down")
	rootCmd.Flags().StringP("type", "t", "", "schema or record")
}
