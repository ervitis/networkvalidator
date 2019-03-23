package networkvalidator

import "testing"

type ipdatarequest struct {
	ipaddress, netmask, ipgateway, ipsubnet, dnsprimary, dnssecondary string
}

var (
	networkGood = &ipdatarequest{
		ipaddress: "192.168.2.10",
		netmask: "255.255.255.0",
		ipgateway: "192.168.2.1",
		ipsubnet: "192.168.2.0",
		dnsprimary: "8.8.8.8",
		dnssecondary: "1.1.1.1",
	}
	
	validatorGood = NewNetworkValidator(IpV4Name, networkGood.ipaddress, networkGood.netmask, networkGood.ipgateway, networkGood.ipsubnet, networkGood.dnsprimary, networkGood.dnssecondary)
)

func TestValidatorForIpv4_CIDRIsCorrect(t *testing.T) {
	if !validatorGood.CIDRIsCorrect() {
		t.Error("error in cidr is correct")
	}
}

func TestValidatorForIpv4_AddressIpIsLocalhost(t *testing.T) {
	if validatorGood.AddressIpIsLocalhost() {
		t.Error("address ip is localhost")
	}
}

func TestValidatorForIpv4_AddressIpIsMulticast(t *testing.T) {
	if validatorGood.AddressIpIsMulticast() {
		t.Error("address ip is multicast")
	}
}

func TestValidatorForIpv4_AllIpAreCorrect(t *testing.T) {
	if !validatorGood.AllIpAreCorrect() {
		t.Error("ips are not correct")
	}
}

func TestValidatorForIpv4_RangeNetworkIsCorrect(t *testing.T) {
	if !validatorGood.RangeNetworkIsCorrect() {
		t.Error("range network not correct")
	}
}

func TestValidatorForIpv4_Isinrange(t *testing.T) {
	if !validatorGood.Isinrange() {
		t.Error("ip address is not in range of cidr")
	}
}

func TestValidatorForIpv4_SubnetIsCorrect(t *testing.T) {
	if !validatorGood.SubnetIsCorrect() {
		t.Error("subnet is not correct")
	}
}

func TestValidatorForIpv4_OverlapsWith(t *testing.T) {
	another := NewIpV4Data("192.168.2.20", "255.255.255.0", "192.168.2.1", "192.168.1.0", "1.1.1.1", "2.2.2.2")

	if !validatorGood.OverlapsWith(another.Cidr) {
		t.Error("overlapping ip when it's not overlapped")
	}
}