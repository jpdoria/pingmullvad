package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"text/tabwriter"
	"time"

	"github.com/fatih/color"
	probing "github.com/prometheus-community/pro-bing"
	"github.com/schollz/progressbar/v3"
)

const ApiEndpoint = "https://api.mullvad.net/www/relays"

var (
	serverType = flag.String("type", "all", "Server type: wireguard, openvpn, or bridge.")
	version    = flag.Bool("version", false, "Print the current version.")
	ver, build string
)

// Relays struct contains the relay information like the hostname, location, IP address, and the provider.
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

// Caller interface defines the callApi method.
type ApiCaller interface {
	callApi(serverType string) []byte
}

// RealApiCaller struct implements the Caller interface.
type RealApiCaller struct{}

// callApi calls the Mullvad API and returns the response body.
func (RealApiCaller) callApi(serverType string) []byte {
	req, err := http.NewRequest("GET", ApiEndpoint+"/"+serverType, nil)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	defer res.Body.Close()

	bar := progressbar.DefaultBytes(
		res.ContentLength,
		"fetching server list",
	)

	r, err := io.ReadAll(io.TeeReader(res.Body, bar))
	if err != nil {
		fmt.Printf("Error reading response body: %v", err)
		os.Exit(1)
	}

	return r
}

// formatter func formats the output of the ping command.
func formatter(relay *Relay) string {
	var msg string
	relay.ServerOwnership = "Owned by Mullvad"
	if !relay.Owned {
		relay.ServerOwnership = "Rented"
	}

	msg = fmt.Sprintf("%v\t%v\t%vGbps\t%v\t%v\t%v\t%v\t%v", relay.Hostname, relay.Ipv4AddrIn, relay.NetworkPortSpeed, relay.ServerOwnership, relay.Provider, relay.CityName, relay.CountryName, relay.Latency)
	return msg
}

// pinger func that actually pings the IP address of the relay.
func pinger(relay *Relay) {
	p, err := probing.NewPinger(relay.Ipv4AddrIn)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	p.Count = 3
	p.Timeout = 1 * time.Second
	err = p.Run()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	relay.Latency = p.Statistics().AvgRtt
}

// ping func pings the IP address of the relay.
func ping(relays *[]Relay) {
	bar := progressbar.Default(int64(len(*relays)), "pinging servers")
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 5, ' ', 0)
	defer w.Flush()
	color.New(color.FgWhite).Fprintln(w, "hostname\tip\tbandwidth\townership\tprovider\tcity\tcountry\tlatency")

	for _, relay := range *relays {
		bar.Add(1)
		pinger(&relay)
		msg := formatter(&relay)

		if relay.Latency <= 0*time.Millisecond {
			msg = fmt.Sprintf("%v\t%v\t%vGbps\t%v\t%v\t%v\t%v\t%v", relay.Hostname, relay.Ipv4AddrIn, relay.NetworkPortSpeed, relay.ServerOwnership, relay.Provider, relay.CityName, relay.CountryName, "Timeout")
			color.New(color.FgRed).Fprintln(w, msg)
		} else if relay.Latency > 150.00*time.Millisecond {
			color.New(color.FgYellow).Fprintln(w, msg)
		} else if relay.Latency < 150.00*time.Millisecond && relay.Latency != 0 {
			color.New(color.FgGreen).Fprintln(w, msg)
		}
	}
}

// get func gets the locations of the relays.
func get(caller ApiCaller, serverType string) []Relay {
	res := caller.callApi(serverType)
	var relays []Relay
	json.Unmarshal(res, &relays)
	return relays
}

func main() {
	flag.Parse()

	if *version {
		fmt.Printf("Version: %v.%v\n", ver, build)
		os.Exit(0)
	}

	realCaller := RealApiCaller{}
	relays := get(realCaller, *serverType)
	ping(&relays)
}
