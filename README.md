# About

Ping Mullvad servers. This tool is best run when not connected to the VPN.

# Usage

```bash
❯ chmod +x pingmullvad-arm64-0.3.0
❯ ./pingmullvad-arm64-0.3.0 -help
  -country string
    	Country code: us, hk, jp, etc. (default "all")
  -type string
    	Server type: wireguard, openvpn, or bridge. (default "all")
  -version
    	Print the current version.
❯ ./pingmullvad-arm64-0.3.0 -type=wireguard
pinging servers 100% |█████████████████████████████████████████████████████████████████████████████████████████████████████████| (468/468, 58 it/min)
hostname          ip                  bandwidth     ownership            provider          city                   country            latency
al-tia-wg-001     31.171.153.66       10Gbps        Rented               iRegister         Tirana                 Albania            Timeout
al-tia-wg-002     31.171.154.50       10Gbps        Rented               iRegister         Tirana                 Albania            272.593ms
at-vie-wg-001     146.70.116.98       10Gbps        Rented               M247              Vienna                 Austria            536.601ms
# truncated
us-uyk-wg-103     173.205.85.34       10Gbps        Rented               Quadranet         Secaucus, NJ           USA                249.108ms
za-jnb-wg-001     154.47.30.130       10Gbps        Rented               DataPacket        Johannesburg           South Africa       442.094ms
za-jnb-wg-002     154.47.30.143       10Gbps        Rented               DataPacket        Johannesburg           South Africa       513.294ms

top 5 fastest servers:
hostname          ip                 bandwidth     ownership     provider       city          country       latency
hk-hkg-wg-201     103.125.233.18     10Gbps        Rented        xtom           Hong Kong     Hong Kong     25.062ms
hk-hkg-wg-202     103.125.233.3      10Gbps        Rented        xtom           Hong Kong     Hong Kong     27.091ms
hk-hkg-wg-302     146.70.224.66      10Gbps        Rented        M247           Hong Kong     Hong Kong     30.737ms
hk-hkg-wg-301     146.70.224.2       10Gbps        Rented        M247           Hong Kong     Hong Kong     31.124ms
sg-sin-wg-002     138.199.60.15      10Gbps        Rented        DataPacket     Singapore     Singapore     35.619ms
❯