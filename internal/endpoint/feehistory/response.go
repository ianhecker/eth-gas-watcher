package feehistory

type Response struct {
	JSONRPC string  `json:"jsonrpc"`
	ID      int     `json:"id"`
	Result  Results `json:"result"`
}
