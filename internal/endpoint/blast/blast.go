package blast

import (
	"github.com/ianhecker/eth-gas-watcher/internal/endpoint"
)

const DefaultBlastEndpointURL = "https://eth-mainnet.public.blastapi.io"

type BlastEndpoint struct {
	endpoint.EndpointInterface
}

func NewBlastEndpoint() *BlastEndpoint {
	endpoint := endpoint.NewEndpoint(DefaultBlastEndpointURL)
	return &BlastEndpoint{endpoint}
}
