package endpoint

import "github.com/ianhecker/eth-gas-watcher/internal/endpoint/feehistory"

type Endpoint interface {
	GetFeeHistory() (feehistory.Result, error)
}
