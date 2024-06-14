package main

import (
	"reflect"
	"testing"
)

// MockApiCaller is a mock implementation of the ApiCaller interface.
type MockApiCaller struct{}

// callApi returns a mock response.
func (MockApiCaller) callApi(serverType string) []byte {
	mockResponse := `[{
		"hostname": "ab-cde-fg-123",
		"country_code": "ab",
		"country_name": "Alpha",
		"city_code": "bra",
		"city_name": "Bravo",
		"active": true,
		"owned": false,
		"provider": "Foobar",
		"ipv4_addr_in": "1.2.3.4",
		"network_port_speed": 10,
		"type": "wireguard"
	}]`
	return []byte(mockResponse)
}

// TestGet tests the get function.
func TestGet(t *testing.T) {
	mockCaller := MockApiCaller{}
	expectedRelays := []Relay{
		{
			Hostname:         "ab-cde-fg-123",
			CountryCode:      "ab",
			CountryName:      "Alpha",
			CityCode:         "bra",
			CityName:         "Bravo",
			Active:           true,
			Owned:            false,
			Provider:         "Foobar",
			Ipv4AddrIn:       "1.2.3.4",
			NetworkPortSpeed: 10,
			Type:             "wireguard",
		},
	}

	relays := get(mockCaller, "wireguard")
	if !reflect.DeepEqual(relays, expectedRelays) {
		t.Errorf("got %v, want %v", relays, expectedRelays)
	}
}
