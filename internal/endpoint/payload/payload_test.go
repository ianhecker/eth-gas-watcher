package payload_test

import (
	"testing"

	"github.com/ianhecker/eth-gas-watcher/internal/endpoint/payload"
	"github.com/stretchr/testify/assert"
)

func TestMakePayload(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		method := "method"
		params := []interface{}{
			0: "one",
			1: "two",
		}
		payload := payload.MakePayload(method, params)

		assert.Equal(t, method, payload["method"])
		assert.Equal(t, params, payload["params"])
	})
}

func TestMakePayloadForFeeHistory(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		blockCount := 1
		newestBlock := ""
		rewardPercentiles := []int{1, 2, 3}

		payload := payload.MakePayloadForFeeHistory(
			blockCount,
			newestBlock,
			rewardPercentiles,
		)

		expected := []interface{}{
			0: "0x1",
			1: newestBlock,
			2: rewardPercentiles,
		}

		assert.Equal(t, "eth_feeHistory", payload["method"])
		assert.Equal(t, expected, payload["params"])
	})
}
