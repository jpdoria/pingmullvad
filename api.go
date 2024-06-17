package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/schollz/progressbar/v3"
)

// callApi makes a GET request to the specified API endpoint and returns the response body as a byte slice.
// If the server type is "wireguard", "openvpn", or "bridge", the server type is appended to the API endpoint.
// The progress of the request is displayed using a progress bar.
func callApi(apiEndpoint, serverType string) ([]byte, error) {
	e := fmt.Sprintf("%v/all", apiEndpoint)
	if serverType == "wireguard" || serverType == "openvpn" || serverType == "bridge" {
		e = fmt.Sprintf("%v/%v", apiEndpoint, serverType)
	}

	req, err := http.NewRequest("GET", e, nil)
	if err != nil {
		return []byte{}, err
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return []byte{}, err
	}
	defer res.Body.Close()

	bar := progressbar.DefaultBytes(
		res.ContentLength,
		"fetching server list",
	)

	r, err := io.ReadAll(io.TeeReader(res.Body, bar))
	if err != nil {
		return []byte{}, err
	}

	return r, nil
}
