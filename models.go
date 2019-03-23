package networkvalidator

import "net"

type rangeIp struct {
	start net.IP
	end   net.IP
}

type ipV4Data struct {
	IpAddress net.IP
	Netmask   net.IPMask
	IpGateway net.IP
	IpSubnet  net.IP
	Dnss      []net.IP
	ipRange   *rangeIp
	Cidr      *net.IPNet
}

const (
	IpV4Name = "ipversion4"
	IpV6Name = "ipversion6"
)
