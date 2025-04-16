package endpoint_test

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ianhecker/eth-gas-watcher/internal/endpoint"
	"github.com/ianhecker/eth-gas-watcher/internal/endpoint/payload"
)

type handlerFunc func(w http.ResponseWriter, r *http.Request)

type Response struct {
	Message string `json:"message"`
}

func MakeServerHandler(message string) handlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(Response{Message: message})
	}
}

func MakeServerHandlerWithStatus(status int) handlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
	}
}

func MakeServerHandlerWithNoBody() handlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
	}
}

func NewHTTPRequest(t *testing.T, method, url, message string) *http.Request {
	payloadBytes, _ := json.Marshal(payload.MakePayload())
	buffer := bytes.NewBuffer(payloadBytes)
	request, err := http.NewRequest(method, url, buffer)
	require.NoError(t, err)
	request.Header.Set("Content-Type", "application/json")
	return request
}

func NewHTTPResponse(t *testing.T) *http.Response {
	return &http.Response{
		Status:        "200 OK",
		StatusCode:    http.StatusOK,
		Body:          io.NopCloser(strings.NewReader("")),
		Header:        make(http.Header),
		ContentLength: int64(len("")),
	}
}

func TestClient_Get(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		expected := "hello world"

		handler := MakeServerHandler(expected)
		ts := httptest.NewServer(http.HandlerFunc(handler))
		defer ts.Close()

		client := endpoint.NewClient(ts.URL)

		request := NewHTTPRequest(t, "GET", ts.URL, "request")
		response, err := client.GetWithRequest(request)
		require.Nil(t, err)

		out := bytes.NewBuffer(response)

		var got Response
		err = json.NewDecoder(out).Decode(&got)

		assert.Nil(t, err)
		assert.Equal(t, string(expected), string(got.Message))
	})

	t.Run("get url error", func(t *testing.T) {
		handler := MakeServerHandler("'")
		ts := httptest.NewServer(http.HandlerFunc(handler))
		ts.Close()

		client := endpoint.NewClient(ts.URL)
		request := NewHTTPRequest(t, "GET", ts.URL, "request")

		_, err := client.GetWithRequest(request)
		assert.ErrorContains(t, err, "could not get")
	})

	t.Run("response status not OK", func(t *testing.T) {
		expected := http.StatusNotFound

		handler := MakeServerHandlerWithStatus(expected)
		ts := httptest.NewServer(http.HandlerFunc(handler))
		defer ts.Close()

		client := endpoint.NewClient(ts.URL)
		request := NewHTTPRequest(t, "GET", ts.URL, "request")

		_, err := client.GetWithRequest(request)
		assert.ErrorContains(t, err, "status code not OK. Got: '404'")
	})

	t.Run("bad response body", func(t *testing.T) {
		handler := MakeServerHandlerWithNoBody()
		ts := httptest.NewServer(http.HandlerFunc(handler))
		defer ts.Close()

		client := endpoint.NewClient(ts.URL)
		request := NewHTTPRequest(t, "GET", ts.URL, "request")

		_, err := client.GetWithRequest(request)
		assert.ErrorContains(t, err, "body is empty")
	})
}

func TestClient_MakeRequestWithPayload(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		url := "url"
		client := endpoint.NewClient(url)

		method := http.MethodPost
		payload := payload.MakePayload()

		expectedPayload, err := json.Marshal(payload)
		require.NoError(t, err)

		request, err := client.MakeRequestWithPayload(method, payload)
		assert.Nil(t, err)
		assert.Equal(t, method, request.Method)
		assert.Equal(t, url, request.URL.String())
		assert.Equal(t, "application/json", request.Header.Get("Content-Type"))

		requestBody, err := ioutil.ReadAll(request.Body)
		require.NoError(t, err)
		assert.Equal(t, expectedPayload, requestBody)
	})

	t.Run("bad method", func(t *testing.T) {
		client := endpoint.NewClient("")

		method := "bad method"
		payload := payload.MakePayload()

		_, err := client.MakeRequestWithPayload(method, payload)
		assert.ErrorContains(t, err, "could not form request: 'net/http: invalid method \"bad method\"'")
	})
}
