package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/client-api/cmd/server"
)

var l = logrus.New()

var RootCmd = &cobra.Command{
	Use:   "client-api",
	Short: "ClientAPI swagger API",
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		l.WithError(err).Fatal("something goes wrong")
		return
	}
}

func init() {
	RootCmd.AddCommand(server.Cmd)
}
