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

type ClientInterface interface {
	GetWithRequest(request *http.Request) ([]byte, error)
	MakeRequestWithPayload(method string, payload payload.Payload) (*http.Request, error)
}

type Client struct {
	HTTPClient
	URL string
}

func NewClient(url string) ClientInterface {
	return &Client{
		&http.Client{},
		url,
	}
}

func NewClientFromRaw(
	client HTTPClient,
	url string,
) ClientInterface {
	return &Client{
		client,
		url,
	}
}

func (client Client) GetWithRequest(
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

func (client Client) MakeRequestWithPayload(
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
