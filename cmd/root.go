package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	rootCmd = &cobra.Command{
		Use:              "password-manager",
		Short:            "Password manager app",
		Long:             "Server to manage user's passwords accross multiple apps",
		TraverseChildren: true,
	}
)

func init() {
	cobra.OnInitialize(initEnv)

	rootCmd.AddCommand(serverCmd)
}

func initEnv() {
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		log.Fatalln("Using config file:", viper.ConfigFileUsed())
	}
}

// Execute root command
func Execute() error {
	return rootCmd.Execute()
}
