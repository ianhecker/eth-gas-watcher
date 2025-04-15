package endpoint_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ianhecker/eth-gas-watcher/internal/endpoint"
)

type Response struct {
	Message string `json:"message"`
}

func MakeServerHandler(message string) func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		resp := Response{Message: message}

		err := json.NewEncoder(w).Encode(resp)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func TestClient(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		expected := "hello world"

		handler := MakeServerHandler(expected)
		ts := httptest.NewServer(http.HandlerFunc(handler))
		defer ts.Close()

		client := endpoint.NewClient(ts.URL)

		response, err := client.Get()
		require.Nil(t, err)

		buffer := bytes.NewBuffer(response)

		var got Response
		err = json.NewDecoder(buffer).Decode(&got)

		assert.Nil(t, err)
		assert.Equal(t, string(expected), string(got.Message))
	})
}
