package handler

import (
	"github.com/client-api/cmd/config"
	"github.com/client-api/internal/server/restapi/operations"
	"github.com/sirupsen/logrus"
)

var l = logrus.New()

type Handler struct {
	config config.Config
}

func New(config config.Config) *Handler {
	return &Handler{config: config}
}

// todo add internal error codes
func (h *Handler) ConfigureHandlers(api *operations.ClientAPI) {
	api.Logger = l.Printf
	api.GetAllPortsHandler = operations.GetAllPortsHandlerFunc(h.GetAllPorts)
	api.PostPortsHandler = operations.PostPortsHandlerFunc(h.PostPorts)
}
