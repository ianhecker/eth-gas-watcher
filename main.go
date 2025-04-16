package main

import (
	"encoding/json"
	"fmt"

	"github.com/ianhecker/eth-gas-watcher/internal/desist"
	"github.com/ianhecker/eth-gas-watcher/internal/endpoint/blast"
	"github.com/ianhecker/eth-gas-watcher/internal/endpoint/payload"
)

func main() {
	d := desist.NewDesistor()

	blast := blast.NewBlastEndpoint()

	payload := payload.MakePayloadForFeeHistory(5, "latest", []int{10, 50, 90})

	result, err := blast.GetFeeHistory(payload)
	d.WithError(err).FatalOnError("getting fee history")

	baseTen, err := result.ToBaseTen()
	d.WithError(err).FatalOnError("could not convert hexidecimal to base ten")

	bytes, err := json.MarshalIndent(baseTen, "", "  ")
	d.WithError(err).FatalOnError("could not marshal base ten")

	fmt.Println(string(bytes))
}
