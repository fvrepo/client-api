package handler

import (
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/stretchr/testify/mock"

	"github.com/client-api/cmd/config"
	"github.com/client-api/internal/server/handler/mocks"
	"github.com/client-api/internal/server/restapi/operations"
	portApi "github.com/fvrepo/port-domain/pkg/grpcapi/port"
)

//go:generate mockery -case=snake -dir=./../../../vendor/github.com/fvrepo/port-domain/pkg -outpkg=mocks -output=./mocks -name=.*Client -recursive

// todo fix test later
func TestHandler_PostPorts(t *testing.T) {
	t.Skip()
	portDomain := &mocks.PortServiceClient{}

	h := New(config.Config{MaxFileSize: 500000}, portDomain)

	jsonReq := `
		{
  			"AEAJM": {
    			"name": "Ajman",
    			"city": "Ajman",
    			"country": "United Arab Emirates",
				"alias": [],
    			"regions": [],
    			"coordinates": [
      				55.5136433,
      				25.4052165
    			],
    			"province": "Ajman",
    			"timezone": "Asia/Dubai",
    			"unlocs": [
      				"AEAJM"
    			],
    			"code": "52000"
  			},
  			"AEAUH": {
    			"name": "Abu Dhabi",
    			"coordinates": [
      				54.37,
      				24.47
    			],
    			"city": "Abu Dhabi",
    			"province": "Abu Z¸aby [Abu Dhabi]",
    			"country": "United Arab Emirates",
    			"alias": [],
    			"regions": [],
				"timezone": "Asia/Dubai",
    			"unlocs": [
      				"AEAUH"
    			],
    			"code": "52001"
			}
		}
	`

	stringReader := strings.NewReader(jsonReq)
	stringReadCloser := ioutil.NopCloser(stringReader)

	savePortReq := operations.PostPortsParams{
		HTTPRequest: &http.Request{},
		File:        stringReadCloser,
	}

	portDomain.On("SavePort", mock.Anything, mock.Anything).Return(&portApi.SavePortResponse{}, nil).Twice()

	res := h.PostPorts(savePortReq)
	result := res.(*operations.PostPortsOK)
	require.NotNil(t, result)
}
