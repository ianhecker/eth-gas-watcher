package endpoint

import (
	"io/ioutil"
	"net/http"

	"github.com/ianhecker/eth-gas-watcher/internal/desist"
)

type Endpoint interface {
	Get() ([]byte, error)
}

type Client struct {
	URL string
}

func NewClient(url string) *Client {
	return &Client{
		URL: url,
	}
}

func (c Client) Get() ([]byte, error) {
	resp, err := http.Get(c.URL)
	if err != nil {
		return nil, desist.Error("could not fetch endpoint", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, desist.Error("status code not OK. Got", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, desist.Error("could not read body", err)
	}

	return body, nil
}
