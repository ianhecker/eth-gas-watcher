package endpoint

import "github.com/ianhecker/eth-gas-watcher/internal/endpoint/feehistory"

type EndpointInterface interface {
	GetFeeHistory() (feehistory.Result, error)
}

type Endpoint struct {
}
