package feehistory_test

import (
	"testing"

	"github.com/ianhecker/eth-gas-watcher/internal/endpoint/feehistory"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestConvertHexToBaseTen(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		expected := uint64(1)
		got := uint64(0)
		err := feehistory.ConvertHexToBaseTen("0x1", &got)
		assert.Nil(t, err)
		assert.Equal(t, expected, got)
	})

	t.Run("bad hex", func(t *testing.T) {
		got := uint64(0)
		err := feehistory.ConvertHexToBaseTen("bad hex", &got)
		assert.ErrorContains(t, err, "invalid hex")
	})

	t.Run("pointer is nil", func(t *testing.T) {
		err := feehistory.ConvertHexToBaseTen("0x1", nil)
		assert.ErrorContains(t, err, "given nil pointer")
	})
}

func ConvertHexToBaseTenMatrix(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		in := []string{"0x1", "0x2", "0x3"}
		expected := []uint64{1, 2, 3}
		got := []uint64{0, 0, 0}

		err := feehistory.ConvertHexToBaseTenArray(in, &got)
		assert.Nil(t, err)
		assert.Equal(t, expected, got)
	})

	t.Run("bad hex", func(t *testing.T) {
		in := []string{"0x1", "0x2", "bad hex"}
		got := []uint64{0, 0, 0}
		err := feehistory.ConvertHexToBaseTenArray(in, &got)
		assert.ErrorContains(t, err, "could not convert hex")
	})

	t.Run("pointer is nil", func(t *testing.T) {
		in := []string{"0x1", "0x2", "bad hex"}
		err := feehistory.ConvertHexToBaseTenArray(in, nil)
		assert.ErrorContains(t, err, "given nil pointer")
	})
}

func ConvertHexToBaseTenArray(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		in := [][]string{
			{"0x1", "0x2", "0x3"},
			{"0x4", "0x5", "0x6"},
		}
		expected := [][]uint64{
			{1, 2, 3},
			{4, 5, 6},
		}
		got := [][]uint64{
			{0, 0, 0},
			{0, 0, 0},
		}
		err := feehistory.ConvertHexToBaseTenMatrix(in, &got)
		assert.Nil(t, err)
		assert.Equal(t, expected, got)
	})

	t.Run("bad hex", func(t *testing.T) {
		in := [][]string{
			{"0x1", "0x2", "0x3"},
			{"0x4", "0x5", "0x6"},
		}
		got := [][]uint64{
			{0, 0, 0},
			{0, 0, 0},
		}
		err := feehistory.ConvertHexToBaseTenMatrix(in, &got)
		assert.ErrorContains(t, err, "could not convert hex matrix row")
	})

	t.Run("pointer is nil", func(t *testing.T) {
		in := [][]string{
			{"0x1", "0x2", "0x3"},
			{"0x4", "0x5", "0x6"},
		}
		err := feehistory.ConvertHexToBaseTenMatrix(in, nil)
		assert.ErrorContains(t, err, "given nil pointer")
	})
}

func TestResult_ToBaseTen(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		result := feehistory.Results{
			BaseFeePerGas:     []string{"0x1", "0x2", "0x3"},
			GasUsedRatio:      []float64{1.1, 2.2, 3.3},
			BaseFeePerBlobGas: []string{"0x4", "0x5", "0x6"},
			BlobGasUsedRatio:  []float64{4.4, 5.5, 6.6},
			OldestBlock:       "0x123",
			Reward: [][]string{
				{"0x11", "0x22"},
				{"0x33", "0x44"},
			},
		}
		expected := feehistory.BaseTenResults{
			BaseFeePerGas:     []uint64{1, 2, 3},
			GasUsedRatio:      result.GasUsedRatio,
			BaseFeePerBlobGas: []uint64{4, 5, 6},
			BlobGasUsedRatio:  result.BlobGasUsedRatio,
			OldestBlock:       uint64(291),
			Reward: [][]uint64{
				{17, 34},
				{51, 68},
			},
		}
		got, err := result.ToBaseTen()
		assert.Nil(t, err)
		require.NotNil(t, got)
		assert.Equal(t, expected, *got)
	})

	t.Run("bad hex", func(t *testing.T) {
		result := feehistory.Results{
			BaseFeePerGas: []string{"bad hex"},
		}
		_, err := result.ToBaseTen()
		assert.ErrorContains(t, err, "bad hex")
	})
}
