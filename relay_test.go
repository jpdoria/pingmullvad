package main

import (
	"reflect"
	"testing"
	"time"
)

// TestFilterRelaysByCountryCode tests the filterRelaysByCountryCode method.
func TestFilterRelaysByCountryCode(t *testing.T) {
	// Create a Ping instance with some relays.
	p := &Ping{
		Relays: &[]Relay{
			{CountryCode: "us"},
			{CountryCode: "sg"},
			{CountryCode: "hk"},
			{CountryCode: "jp"},
		},
	}

	// Call the function with the country code "us".
	result := p.filterRelaysByCountryCode("us")

	// Check that the result contains the correct relays.
	if len(result) != 2 {
		t.Errorf("Expected 2 relays, got %d", len(result))
	}
	for _, relay := range result {
		if relay.CountryCode != "US" {
			t.Errorf("Expected relay with country code 'us', got '%s'", relay.CountryCode)
		}
	}

	// Call the function with the country code "sg".
	result = p.filterRelaysByCountryCode("sg")

	// Check that the result contains the correct relay.
	if len(result) != 1 {
		t.Errorf("Expected 1 relay, got %d", len(result))
	}
	if result[0].CountryCode != "hk" {
		t.Errorf("Expected relay with country code 'hk', got '%s'", result[0].CountryCode)
	}
}

// TestFormatter tests the formatter method.
func TestFormatter(t *testing.T) {
	p := &Ping{}
	relay := &Relay{
		Hostname:         "testHostname",
		Ipv4AddrIn:       "127.0.0.1",
		NetworkPortSpeed: 1,
		Owned:            true,
		Provider:         "testProvider",
		CityName:         "testCity",
		CountryName:      "testCountry",
		Latency:          100 * time.Millisecond,
	}

	expected := "testHostname\t127.0.0.1\t1Gbps\tOwned by Mullvad\ttestProvider\ttestCity\ttestCountry\t100ms"
	result := p.formatter(relay)

	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

// TestPinger tests the pinger method.
func TestPinger(t *testing.T) {
	p := &Ping{}
	relay := &Relay{Ipv4AddrIn: "127.0.0.1"}

	latency, err := p.pinger(relay)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if latency <= 0 {
		t.Errorf("Expected latency to be greater than 0, got %v", latency)
	}
}

// TestShowRelays tests the showRelays method.
func TestShowRelays(t *testing.T) {
	p := &Ping{
		Relays: &[]Relay{
			{Ipv4AddrIn: "127.0.0.1"},
			{Ipv4AddrIn: "192.168.1.1"},
		},
	}

	relays := p.showRelays()

	if !reflect.DeepEqual(relays, p.Relays) {
		t.Errorf("Expected %v, got %v", p.Relays, relays)
	}
}

// TestPing tests the ping method.
func TestPing(t *testing.T) {
	pm := &Ping{
		Relays: &[]Relay{
			{Hostname: "testHostname1", Ipv4AddrIn: "127.0.0.1", NetworkPortSpeed: 1, Owned: true, Provider: "testProvider1", CityName: "testCity1", CountryName: "testCountry1", Latency: 0},
			{Hostname: "testHostname2", Ipv4AddrIn: "192.168.1.1", NetworkPortSpeed: 10, Owned: false, Provider: "testProvider2", CityName: "testCity2", CountryName: "testCountry2", Latency: 150},
		},
	}

	ping(pm)

	// Check if the latency of the relays has been updated.
	for _, relay := range *pm.Relays {
		if relay.Latency <= 0 {
			t.Errorf("Expected latency to be greater than 0, got %v", relay.Latency)
		}
	}
}
