package payload

import "fmt"

type Payload map[string]interface{}

var DefaultPayload = map[string]interface{}{
	"jsonrpc": "2.0",
	"id":      1,
}

func MakePayload(
	method string,
	params []interface{},
) Payload {
	payload := DefaultPayload
	payload["method"] = method
	payload["params"] = params
	return payload
}

func MakePayloadForFeeHistory(
	blockCount int,
	newestBlock string,
	rewardPercentiles []int,
) Payload {
	params := make([]interface{}, 3)
	params[0] = fmt.Sprintf("%#x", blockCount)
	params[1] = newestBlock
	params[2] = rewardPercentiles

	return MakePayload("eth_feeHistory", params)
}
