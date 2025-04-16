package feehistory

type BaseTenResults struct {
	BaseFeePerGas     []uint64
	GasUsedRatio      []float64
	BaseFeePerBlobGas []uint64
	BlobGasUsedRatio  []float64
	OldestBlock       uint64
	Reward            [][]uint64
}
