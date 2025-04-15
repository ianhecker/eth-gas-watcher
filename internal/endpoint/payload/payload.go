package payload

type Payload map[string]interface{}

func MakePayload() Payload {
	return map[string]interface{}{
		"jsonrpc": "2.0",
		"id":      1,
		"method":  "eth_feeHistory",
		"params": []interface{}{
			"0x5",             // Block count (hex encoded)
			"latest",          // Newest block specifier
			[]int{10, 50, 90}, // Reward percentiles
		},
	}
}
