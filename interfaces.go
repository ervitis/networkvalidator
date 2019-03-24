package networkvalidator

import (
	"net"
)

type Validators interface {
	Isinrange() bool
	IpGatewayIsinrange() bool
	OverlapsWith(ipnet *net.IPNet) bool
	SubnetIsCorrect() bool
	RangeNetworkIsCorrect() bool
	AddressIpIsLocalhost() bool
	AddressIpIsMulticast() bool
	AllIpAreCorrect() bool
	CIDRIsCorrect() bool
}

func NewNetworkValidator(ipversion, ipaddress, netmask, ipgateway, ipsubnet, dnsprimary, dnssecondary string) Validators {
	if ipversion == IpV4Name {
		return NewValidatorForIpv4(NewIpV4Data(ipaddress, netmask, ipgateway, ipsubnet, dnsprimary, dnssecondary))
	} else if ipversion == IpV6Name {
		return nil
	} else {
		return nil
	}
}
