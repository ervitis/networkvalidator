package networkvalidator

import (
	"bytes"
	"fmt"
	"net"
)

func NewIpV4Data(ipaddress, netmask, ipgateway, ipsubnet, dnsprimary, dnssecondary string) *ipV4Data {
	d := &ipV4Data{
		IpAddress: net.ParseIP(ipaddress).To4(),
		Netmask:   net.IPMask(net.ParseIP(netmask).To4()),
		IpGateway: net.ParseIP(ipgateway).To4(),
		IpSubnet:  net.ParseIP(ipsubnet).To4(),
		Dnss:      []net.IP{net.ParseIP(dnsprimary).To4(), net.ParseIP(dnssecondary).To4()},
	}

	o, _ := d.Netmask.Size()
	_, in, _ := net.ParseCIDR(fmt.Sprintf("%s/%d", d.IpAddress, o))

	d.Cidr = in
	d.ipRange = getiprange(in, &d.Netmask)

	return d
}

type ValidatorForIpv4 struct {
	data       *ipV4Data
}

func (v *ValidatorForIpv4) Isinrange() bool {
	if bytes.Compare([]byte(v.data.IpAddress), []byte(v.data.ipRange.start)) < 0 {
		return false
	}

	return bytes.Compare([]byte(v.data.IpAddress), []byte((v.data.ipRange.end))) <= 0
}

func (v *ValidatorForIpv4) OverlapsWith(ipnet *net.IPNet) bool {
	return v.data.Cidr.Contains(ipnet.IP) || ipnet.Contains(v.data.Cidr.IP)
}

func (v *ValidatorForIpv4) SubnetIsCorrect() bool {
	return v.data.Cidr.IP.Equal(net.ParseIP(v.data.IpSubnet.String()).To4())
}

func (v *ValidatorForIpv4) RangeNetworkIsCorrect() bool {
	return v.data.ipRange != nil
}

func (v *ValidatorForIpv4) AddressIpIsLocalhost() bool {
	return v.data.IpAddress.IsLoopback()
}

func (v *ValidatorForIpv4) AddressIpIsMulticast() bool {
	return v.data.IpAddress.IsMulticast()
}

func (v *ValidatorForIpv4) AllIpAreCorrect() bool {
	return v.data.IpAddress != nil &&
		v.data.IpSubnet != nil &&
		v.data.Netmask != nil &&
		v.data.IpGateway != nil
}

func (v *ValidatorForIpv4) CIDRIsCorrect() bool {
	return v.data.Cidr != nil
}

func (v *ValidatorForIpv4) IpGatewayIsinrange() bool {
	return v.data.Cidr.Contains(v.data.IpGateway)
}

func NewValidatorForIpv4(ipv4data *ipV4Data) Validators {
	v := &ValidatorForIpv4{
		data:       ipv4data,
	}

	return v
}
