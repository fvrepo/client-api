package handler

import (
	"encoding/json"
	"io"

	"github.com/go-openapi/runtime/middleware"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"

	"github.com/client-api/internal/server/models"
	"github.com/client-api/internal/server/restapi/operations"
	portApi "github.com/fvrepo/port-domain/pkg/grpcapi/port"
)

func (h *Handler) GetAllPorts(params operations.GetAllPortsParams) middleware.Responder {
	ctx := params.HTTPRequest.Context()

	res, err := h.port.GetAllPorts(ctx, &portApi.GetAllPortsRequest{Limit: int32(params.Limit), Skip: int32(params.Skip)})
	if err != nil {
		return operations.NewPostPortsInternalServerError().WithPayload(&models.Error{Code: 500, Message: "failed to get port data"})
	}

	return operations.NewGetAllPortsOK().WithPayload(grpcToRestModel(res.Ports))
}

func (h *Handler) PostPorts(params operations.PostPortsParams) middleware.Responder {
	defer params.File.Close()

	lr := &io.LimitedReader{N: int64(h.config.MaxFileSize), R: params.File}
	pr := make(chan *portApi.SavePortRequest, h.config.Workers)
	g, gctx := errgroup.WithContext(params.HTTPRequest.Context())

	g.Go(func() error {
		defer close(pr)
		if err := h.parsePortsJson(lr, pr); err != nil {
			return err
		}
		return nil
	})

	for i := 0; i < h.config.Workers; i++ {
		g.Go(func() error {
			for p := range pr {
				if _, err := h.port.SavePort(gctx, p); err != nil {
					l.WithError(err).Errorf("failed to save port by id:%s", p.Id)
				}
			}
			return nil
		})
	}
	if err := g.Wait(); err != nil {
		return operations.NewPostPortsInternalServerError().WithPayload(&models.Error{Code: 500, Message: "failed to load port data"})
	}

	return operations.NewPostPortsOK()
}

func (h *Handler) parsePortsJson(lr *io.LimitedReader, pr chan *portApi.SavePortRequest) error {
	dec := json.NewDecoder(lr)

	// read open bracket
	_, err := dec.Token()
	if err != nil {
		l.WithError(err).Error("failed to read JSON start token")
		return errors.WithStack(err)
	}

	// while the array contains values
	for dec.More() {
		// read port id
		t, err := dec.Token()
		if err != nil {
			l.WithError(err).Error("failed to read key")
			continue
		}
		var key string
		var ok bool
		if key, ok = t.(string); !ok {
			l.WithError(err).Error("key is not a string")
			continue
		}

		var p models.Port
		// decode port
		err = dec.Decode(&p)
		if err != nil {
			l.WithError(err).Error("failed to decode payload to port model")
		}
		pd := &portApi.SavePortRequest{
			Id: key,
			Details: &portApi.PortDetails{
				Code:        p.Code,
				Unlocs:      p.Unlocs,
				Timezone:    p.Timezone,
				Province:    p.Province,
				Coordinates: p.Coordinates,
				Regions:     p.Regions,
				Alias:       p.Alias,
				City:        p.City,
				Country:     p.Country,
				Name:        p.Name,
			},
		}
		pr <- pd
	}
	// read closing bracket
	_, err = dec.Token()
	if err != nil {
		l.WithError(err).Error("failed to read JSON end token")
	}
	return nil
}

func grpcToRestModel(ports *portApi.PortMap) map[string]models.Port {
	pm := make(map[string]models.Port)
	for key, port := range ports.Port {
		pm[key] = models.Port{
			Code:        port.Code,
			Name:        port.Name,
			Country:     port.Country,
			City:        port.City,
			Alias:       port.Alias,
			Regions:     port.Regions,
			Coordinates: port.Coordinates,
			Province:    port.Province,
			Timezone:    port.Timezone,
			Unlocs:      port.Unlocs,
		}
	}
	return pm
}
