# About

Ping Mullvad servers. This tool is best run when not connected to the VPN.

# Usage

```bash
❯ chmod +x pingmullvad-arm64-0.1.0
❯ ./ping-mullvad-arm64-0.1.0 -help
  -type string
    	Server type: wireguard, openvpn, or bridge. (default "all")
  -version
    	Print the current version.
❯ ./ping-mullvad-arm64-0.1.0 -type=bridge
pinging servers 100% |████████████████████████████████████████████████████████████████████████████████████████████████████████████| (27/27, 1 it/s)
hostname          ip                  bandwidth     ownership            provider     city                   country            latency
au-syd-br-001     146.70.141.154      10Gbps        Rented               M247         Sydney                 Australia          299.491ms
ca-mtr-br-001     217.138.213.18      10Gbps        Rented               M247         Montreal               Canada             219.36ms
ca-tor-br-101     198.44.140.226      10Gbps        Rented               Tzulo        Toronto                Canada             252.742ms
ch-zrh-br-001     193.32.127.117      10Gbps        Owned by Mullvad     31173        Zurich                 Switzerland        311.357ms
cz-prg-br-101     217.138.199.106     1Gbps         Rented               M247         Prague                 Czech Republic     319.682ms
de-fra-br-001     185.213.155.117     20Gbps        Owned by Mullvad     31173        Frankfurt              Germany            296.225ms
fi-hel-br-101     193.138.7.132       10Gbps        Owned by Mullvad     Blix         Helsinki               Finland            307.056ms
fr-par-br-001     193.32.126.117      20Gbps        Owned by Mullvad     31173        Paris                  France             287.635ms
gb-mnc-br-001     89.238.134.58       10Gbps        Rented               M247         Manchester             UK                 372.115ms
hk-hkg-br-201     103.125.233.210     10Gbps        Rented               xtom         Hong Kong              Hong Kong          296.114ms
jp-tyo-br-201     185.242.4.34        10Gbps        Rented               M247         Tokyo                  Japan              260.66ms
nl-ams-br-001     185.65.134.116      10Gbps        Owned by Mullvad     31173        Amsterdam              Netherlands        Timeout
no-osl-br-001     91.90.44.10         1Gbps         Owned by Mullvad     Blix         Oslo                   Norway             Timeout
no-svg-br-001     194.127.199.245     10Gbps        Owned by Mullvad     Blix         Stavanger              Norway             299.788ms
se-got-br-001     185.213.154.117     1Gbps         Owned by Mullvad     31173        Gothenburg             Sweden             295.054ms
se-mma-br-001     193.138.218.71      10Gbps        Owned by Mullvad     31173        Malmö                  Sweden             288.806ms
se-sto-br-001     185.65.135.115      1Gbps         Owned by Mullvad     31173        Stockholm              Sweden             295.196ms
sg-sin-br-101     146.70.192.38       10Gbps        Rented               M247         Singapore              Singapore          Timeout
us-atl-br-101     66.115.180.241      1Gbps         Rented               100TB        Atlanta, GA            USA                200.838ms
us-chi-br-001     68.235.44.130       10Gbps        Rented               Tzulo        Chicago, IL            USA                205.869ms
us-dal-br-101     174.127.113.18      1Gbps         Rented               100TB        Dallas, TX             USA                185.168ms
us-lax-br-401     62.133.44.202       10Gbps        Rented               M247         Los Angeles, CA        USA                152.298ms
us-mia-br-101     146.70.183.34       10Gbps        Rented               M247         Miami, FL              USA                211.515ms
us-nyc-br-501     212.103.48.226      10Gbps        Rented               M247         New York, NY           USA                212.441ms
us-nyc-br-601     38.132.121.146      10Gbps        Rented               M247         New York, NY           USA                213.193ms
us-rag-br-101     198.54.130.178      1Gbps         Rented               Tzulo        Raleigh, NC            USA                219.468ms
us-slc-br-101     69.4.234.146        1Gbps         Rented               100TB        Salt Lake City, UT     USA                186.717ms
```
