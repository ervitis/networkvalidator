# Network validator for Golang

Utilities for validation network

Example:

```go
package main

import (
	"fmt"
	nv "github.com/ervitis/networkvalidator"
)

type ipdatarequest struct {
	ipaddress, netmask, ipgateway, ipsubnet, dnsprimary, dnssecondary string
}

func main() {
	network := &ipdatarequest{
		ipaddress: "192.168.2.10", 
		netmask: "255.255.255.0", 
		ipgateway: "192.168.2.1", 
		ipsubnet: "192.168.2.0", 
		dnsprimary: "8.8.8.8", 
		dnssecondary: "1.1.1.1",
	}

	validator := nv.NewNetworkValidator(nv.IpV4Name, network.ipaddress, network.netmask, network.ipgateway, network.ipsubnet, network.dnsprimary, network.dnssecondary)

	fmt.Println(validator.AddressIpIsLocalhost())
}
```