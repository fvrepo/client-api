package handler

import (
	"github.com/sirupsen/logrus"

	"github.com/client-api/cmd/config"
	"github.com/client-api/internal/server/restapi/operations"
	portApi "github.com/port-domain/pkg/grpcapi/port"
)

var l = logrus.New()

type Handler struct {
	config config.Config
	port   portApi.PortServiceClient
}

func New(config config.Config, port portApi.PortServiceClient) *Handler {
	return &Handler{config: config, port: port}
}

// todo add internal error codes
func (h *Handler) ConfigureHandlers(api *operations.ClientAPI) {
	api.Logger = l.Printf
	api.GetAllPortsHandler = operations.GetAllPortsHandlerFunc(h.GetAllPorts)
	api.PostPortsHandler = operations.PostPortsHandlerFunc(h.PostPorts)
}
