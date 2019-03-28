package handler

import (
	"github.com/go-openapi/runtime/middleware"

	"github.com/client-api/internal/server/restapi/operations"
)

func (h *Handler) GetAllPorts(params operations.GetAllPortsParams) middleware.Responder {
	return nil
}

func (h *Handler) PostPorts(params operations.PostPortsParams) middleware.Responder {
	return nil
}
