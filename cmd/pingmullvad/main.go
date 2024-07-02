package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
)

var (
	serverType  = flag.String("type", "all", "Server type: wireguard, openvpn, or bridge.")
	countryCode = flag.String("country", "all", "Country code: us, hk, jp, etc.")
	version     = flag.Bool("version", false, "Print the current version.")
	ver, build  string
)

// This function parses command line flags, retrieves server information from an API,
// and performs a ping operation on the obtained server relays.
func main() {
	flag.Parse()

	// Print the version and exit.
	if *version {
		fmt.Printf("Version: %v.%v\n", ver, build)
		os.Exit(0)
	}

	// Call the API to get the server relays.
	res, err := callApi("https://api.mullvad.net/www/relays", *serverType)
	if err != nil {
		fmt.Printf("Error calling API: %v", err)
		os.Exit(1)
	}

	// Unmarshal the response into a slice of Relay structs.
	var relays []Relay
	err = json.Unmarshal(res, &relays)
	if err != nil {
		fmt.Printf("Error unmarshalling response: %v\n", err)
		os.Exit(1)
	}

	// Ping the server relays.
	p := &Ping{Relays: &relays}
	ping(p)
}
