package main

import (
	"reflect"
	"testing"
	"time"
)

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
