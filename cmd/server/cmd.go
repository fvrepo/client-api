package server

import (
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime/middleware"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/client-api/cmd/config"
	"github.com/client-api/internal/server/handler"
	serverMW "github.com/client-api/internal/server/middleware"
	"github.com/client-api/internal/server/restapi"
	"github.com/client-api/internal/server/restapi/operations"
	"github.com/client-api/internal/utils"
)

var cfg config.Config

var l = logrus.New()

func init() {
	Cmd.Flags().AddFlagSet(cfg.Flags())
}

var Cmd = &cobra.Command{
	Use:   "client-api",
	Short: "run server",
	RunE: func(cmd *cobra.Command, args []string) error {
		utils.BindEnv(cmd)

		l.Info("start ClientApi server")
		defer l.Info("stop ClientApi server")

		// load embedded swagger file
		swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
		if err != nil {
			return errors.WithStack(err)
		}

		// create new service API
		api := operations.NewClientAPI(swaggerSpec)
		server := restapi.NewServer(api)
		defer func() {
			if err := server.Shutdown(); err != nil {
				l.WithError(err).Error()
			}
		}()

		// set the port this service will be run on
		server.Port = cfg.Port
		server.Host = cfg.Host
		server.ReadTimeout = cfg.ReadTimeout
		server.WriteTimeout = cfg.WriteTimeout

		handler.New(cfg).ConfigureHandlers(api)

		server.SetHandler(serverMW.PanicRecovery(serverMW.Logger(api.Serve(middleware.PassthroughBuilder))))

		// serve API
		if err := server.Serve(); err != nil {
			return errors.WithStack(err)
		}

		return nil
	},
}
