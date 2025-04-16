package endpoint

import (
	"encoding/json"
	"net/http"

	"github.com/ianhecker/eth-gas-watcher/internal/desist"
	"github.com/ianhecker/eth-gas-watcher/internal/endpoint/feehistory"
	"github.com/ianhecker/eth-gas-watcher/internal/endpoint/payload"
)

type EndpointInterface interface {
	GetFeeHistory(payload payload.Payload) (*feehistory.Results, error)
}

type Endpoint struct {
	ClientInterface
}

func NewEndpoint(url string) EndpointInterface {
	client := NewClient(url)
	return NewEndpointFromRaw(client)
}

func NewEndpointFromRaw(client ClientInterface) EndpointInterface {
	return &Endpoint{client}
}

func (e Endpoint) GetFeeHistory(payload payload.Payload) (*feehistory.Results, error) {

	request, err := e.MakeRequestWithPayload(http.MethodPost, payload)
	if err != nil {
		return nil, desist.Error("could not make request", err)
	}

	bytes, err := e.GetWithRequest(request)
	if err != nil {
		return nil, desist.Error("could not get bytes", err)
	}

	var response feehistory.Response
	err = json.Unmarshal(bytes, &response)
	if err != nil {
		return nil, desist.Error("could not unmarshal response", err)
	}

	return &response.Result, nil
}
