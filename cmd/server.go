package cmd

import (
	"log"
	"password-manager/app"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	serverCmd = &cobra.Command{
		Use:   "server",
		Short: "run app server",
		Run:   serverRun,
	}

	port int

	flagPort = "port"
)

func init() {
	serverCmd.Flags().IntVarP(&port, flagPort, "p", 8080, "server port used")
	viper.BindPFlag(flagPort, serverCmd.Flags().Lookup(flagPort))
}

func serverRun(cmd *cobra.Command, args []string) {
	jwtSecret := viper.GetString("JWT_SECRET")
	if jwtSecret == "" {
		log.Fatalln("missing jwt secret")
	}

	a := app.App{
		Port:      viper.GetInt(flagPort),
		JWTSecret: jwtSecret,
	}

	a.Start()
}
