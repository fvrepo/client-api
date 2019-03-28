package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/onetwoclimb/cmd/server"
)

var l = logrus.New()

var RootCmd = &cobra.Command{
	Use:   "OneTwoClimbAPI",
	Short: "OneTwoClimb swagger API",
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
