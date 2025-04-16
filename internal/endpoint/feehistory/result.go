package feehistory

import (
	"fmt"
	"strconv"

	"github.com/ianhecker/eth-gas-watcher/internal/desist"
)

type Results struct {
	BaseFeePerGas     []string   `json:"baseFeePerGas"`
	GasUsedRatio      []float64  `json:"gasUsedRatio"`
	BaseFeePerBlobGas []string   `json:"baseFeePerBlobGas"`
	BlobGasUsedRatio  []float64  `json:"blobGasUsedRatio"`
	OldestBlock       string     `json:"oldestBlock"`
	Reward            [][]string `json:"reward"`
}

type BaseTenResults struct {
	BaseFeePerGas     []uint64
	GasUsedRatio      []float64
	BaseFeePerBlobGas []uint64
	BlobGasUsedRatio  []float64
	OldestBlock       uint64
	Reward            [][]uint64
}

func (in Results) ToBaseTen() (*BaseTenResults, error) {
	var out BaseTenResults

	err := ConvertHexToBaseTenArray(in.BaseFeePerGas, &out.BaseFeePerGas)
	if err != nil {
		return nil, desist.Error("could not convert base fee per gas", err)
	}

	out.GasUsedRatio = in.GasUsedRatio

	err = ConvertHexToBaseTenArray(in.BaseFeePerBlobGas, &out.BaseFeePerBlobGas)
	if err != nil {
		return nil, desist.Error("could not convert base fee per blob gas", err)
	}

	out.BlobGasUsedRatio = in.BlobGasUsedRatio

	err = ConvertHexToBaseTen(in.OldestBlock, &out.OldestBlock)
	if err != nil {
		return nil, desist.Error("could not convert oldest block", err)
	}

	err = ConvertHexToBaseTenMatrix(in.Reward, &out.Reward)
	if err != nil {
		return nil, desist.Error("could not convert reward matrix", err)
	}

	return &out, nil
}

func ConvertHexToBaseTenMatrix(in [][]string, out *[][]uint64) error {
	tmp := make([][]uint64, len(in))
	for i, row := range in {
		var converted []uint64
		err := ConvertHexToBaseTenArray(row, &converted)
		if err != nil {
			return desist.Error("could not convert hex matrix row", err)
		}
		tmp[i] = converted
	}
	*out = tmp
	return nil
}

func ConvertHexToBaseTenArray(in []string, out *[]uint64) error {
	tmp := make([]uint64, len(in))
	for i, s := range in {
		var v uint64
		err := ConvertHexToBaseTen(s, &v)
		if err != nil {
			return desist.Error("could not covert hex", err)
		}
		tmp[i] = v
	}
	*out = tmp
	return nil
}

func ConvertHexToBaseTen(in string, out *uint64) error {
	if out == nil {
		return fmt.Errorf("given nil pointer")
	}

	v, err := strconv.ParseUint(in, 0, 64)
	if err != nil {
		return fmt.Errorf("invalid hex: '%q': %w", in, err)
	}
	*out = v
	return nil
}
