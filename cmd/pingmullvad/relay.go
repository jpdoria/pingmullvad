package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
	"text/tabwriter"
	"time"

	"github.com/fatih/color"
	probing "github.com/prometheus-community/pro-bing"
	"github.com/schollz/progressbar/v3"
)

// Relay struct contains the relay information like the hostname, location, IP address, and the provider.
type Relay struct {
	Hostname         string `json:"hostname"`
	CountryCode      string `json:"country_code"`
	CountryName      string `json:"country_name"`
	CityCode         string `json:"city_code"`
	CityName         string `json:"city_name"`
	Active           bool   `json:"active"`
	Owned            bool   `json:"owned"`
	Provider         string `json:"provider"`
	Ipv4AddrIn       string `json:"ipv4_addr_in"`
	NetworkPortSpeed int    `json:"network_port_speed"`
	Type             string `json:"type"`
	Latency          time.Duration
	ServerOwnership  string
}

// PingModel interface defines the methods that a Ping model must implement.
type PingModel interface {
	pinger(relay *Relay) (time.Duration, error)
	formatter(relay *Relay) string
	showRelays() *[]Relay
	filterRelaysByCountryCode(countryCode string) []Relay
}

// Ping struct contains the relay information and implements the PingModel interface.
type Ping struct {
	Relays *[]Relay
}

// formatter func formats the output message with the relay information.
func (p *Ping) formatter(relay *Relay) string {
	// Format the output message.
	var msg string
	relay.ServerOwnership = "Owned by Mullvad"
	if !relay.Owned {
		relay.ServerOwnership = "Rented"
	}

	msg = fmt.Sprintf("%v\t%v\t%vGbps\t%v\t%v\t%v\t%v\t%v",
		relay.Hostname,
		relay.Ipv4AddrIn,
		relay.NetworkPortSpeed,
		relay.ServerOwnership,
		relay.Provider,
		relay.CityName,
		relay.CountryName,
		relay.Latency,
	)
	return msg
}

// pinger func that actually pings the IP address of the relay.
func (p *Ping) pinger(relay *Relay) (time.Duration, error) {
	// Create a new pinger with the relay's IP address.
	pr, err := probing.NewPinger(relay.Ipv4AddrIn)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Set the number of pings and the timeout duration.
	pr.Count = 3
	pr.Timeout = 1 * time.Second
	err = pr.Run()
	if err != nil {
		return 0, err
	}

	// Set the latency of the relay.
	return pr.Statistics().AvgRtt, nil
}

// filterRelaysByCountryCode func filters the relays by the country code.
func (p *Ping) filterRelaysByCountryCode(countryCode string) []Relay {
	var filtered []Relay
	for _, relay := range *p.Relays {
		if relay.CountryCode == countryCode {
			filtered = append(filtered, relay)
		}
	}

	return filtered
}

// showRelays func returns the relays.
func (p *Ping) showRelays() *[]Relay {
	return p.Relays
}

// ping func pings the IP address of the relay.
func ping(pm PingModel) {
	relays := pm.showRelays()

	// Filter the relays by country code.
	if strings.ToLower(*countryCode) != "" && strings.ToLower(*countryCode) != "all" {
		*relays = pm.filterRelaysByCountryCode(strings.ToLower(*countryCode))
	}

	// Create a progress bar to display the progress of the ping operation.
	bar := progressbar.Default(int64(len(*relays)), "pinging servers")

	// Create a tabwriter to format the output.
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 5, ' ', 0)
	defer w.Flush()
	color.New(color.FgWhite).Fprintln(w, "hostname\tip\tbandwidth\townership\tprovider\tcity\tcountry\tlatency")

	// Ping each relay and display the results.
	var newList []Relay
	for _, relay := range *relays {
		bar.Add(1)

		// Call the pinger and formatter methods on the relay.
		latency, _ := pm.pinger(&relay)
		relay.Latency = latency
		msg := pm.formatter(&relay)

		// Color the output based on the latency.
		// Red: Timeout
		// Yellow: Latency > 150ms
		// Green: Latency < 150ms
		if latency <= 0*time.Millisecond {
			msg = fmt.Sprintf("%v\t%v\t%vGbps\t%v\t%v\t%v\t%v\t%v",
				relay.Hostname,
				relay.Ipv4AddrIn,
				relay.NetworkPortSpeed,
				relay.ServerOwnership,
				relay.Provider,
				relay.CityName,
				relay.CountryName,
				"Timeout",
			)
			color.New(color.FgRed).Fprintln(w, msg)
		} else if latency > 150.00*time.Millisecond {
			color.New(color.FgYellow).Fprintln(w, msg)
			newList = append(newList, relay)
		} else if latency < 150.00*time.Millisecond && latency != 0 {
			color.New(color.FgGreen).Fprintln(w, msg)
			newList = append(newList, relay)
		}
	}

	// Show the top 5 relays.
	sort.Slice(newList, func(i, j int) bool {
		return newList[i].Latency < newList[j].Latency
	})
	newList = newList[:5]
	color.New(color.FgWhite).Fprintln(w, "\ntop 5 fastest servers:")
	color.New(color.FgWhite).Fprintln(w, "hostname\tip\tbandwidth\townership\tprovider\tcity\tcountry\tlatency")
	for _, relay := range newList {
		msg := pm.formatter(&relay)
		color.New(color.FgBlue).Fprintln(w, msg)
	}
}
