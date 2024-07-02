package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// This is a test function to verify the functionality of the callApi method.
// It creates a mock server, sends a request to the server, and compares the results.
func TestCallApi(t *testing.T) {
	tests := []struct {
		relay      Relay
		serverType string
	}{
		{Relay{Type: "wireguard"}, "wireguard"},
		{Relay{Type: "openvpn"}, "openvpn"},
		{Relay{Type: "bridge"}, "bridge"},
	}

	for _, test := range tests {
		// Create a mock server.
		server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			_, err := rw.Write([]byte(test.relay.Type))
			if err != nil {
				t.Errorf("Error writing response: %v", err)
			}
		}))

		// Close the server when test finishes.
		defer server.Close()

		// Call the callApi method, passing the mock server's URL as the API endpoint.
		data, err := callApi(server.URL, *serverType)
		if err != nil {
			t.Errorf("Error calling API: %v", err)
		}

		// Compare the returned data with the have data.
		if string(data) != test.serverType {
			t.Errorf("want %v, got %v", test.serverType, string(data))
		}
	}
}
