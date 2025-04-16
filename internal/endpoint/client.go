package endpoint

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/ianhecker/eth-gas-watcher/internal/desist"
	"github.com/ianhecker/eth-gas-watcher/internal/endpoint/payload"
)

type Client interface {
	GetWithRequest(http.Request) ([]byte, error)
	MakeRequestWithPayload(payload.Payload) ([]byte, error)
}

type EndpointClient struct {
	HTTPClient
	URL string
}

func NewEndpointClient(url string) *EndpointClient {
	return &EndpointClient{
		&http.Client{},
		url,
	}
}

func NewEndpointClientFromRaw(
	client HTTPClient,
	url string,
) *EndpointClient {
	return &EndpointClient{
		client,
		url,
	}
}

func (client EndpointClient) GetWithRequest(
	request *http.Request,
) ([]byte, error) {
	resp, err := client.Do(request)
	if err != nil {
		return nil, desist.Error("could not get", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, desist.Error("status code not OK. Got", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, desist.Error("could not read body", err)
	}

	if len(body) == 0 {
		return nil, errors.New("body is empty")
	}

	return body, nil
}

func (client EndpointClient) MakeRequestWithPayload(
	method string,
	payload payload.Payload,
) (*http.Request, error) {

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		desist.Error("could not marsha payload", err)
	}

	payloadBuffer := bytes.NewBuffer(payloadBytes)

	request, err := http.NewRequest(method, client.URL, payloadBuffer)
	if err != nil {
		return nil, desist.Error("could not form request", err)
	}

	request.Header.Set("Content-Type", "application/json")

	return request, nil
}
