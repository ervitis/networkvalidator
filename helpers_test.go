package networkvalidator

import (
	"net"
	"testing"
)

func TestIpRange (t *testing.T) {
	ipnet1 := &net.IPNet{
		IP:   net.ParseIP("192.168.2.20").To4(),
		Mask: net.IPMask(net.ParseIP("255.0.0.0").To4()),
	}

	if getiprange(ipnet1, &ipnet1.Mask) == nil {
		t.Error("error correct data")
	}
}

func TestIpRangeSameIp (t *testing.T) {
	ipnet2 := &net.IPNet{
		IP: net.ParseIP("192.168.2.20").To4(),
		Mask: net.IPMask(net.ParseIP("255.255.255.255").To4()),
	}

	if v := getiprange(ipnet2, &ipnet2.Mask); v == nil {
		t.Error("error correct data")
	} else if v.start.String() != v.end.String() {
		t.Error("not equals ips in range")
	}
}

func TestIpRangeNilData (t *testing.T) {
	if getiprange(nil, nil) != nil {
		t.Error("error not nil")
	}
}

func TestIptoint (t *testing.T) {
	_, i, _ := iptoint(net.ParseIP("10.0.0.2").To4())
	if i != 32 {
		t.Error("not ipv4")
	}

	_, i, _ = iptoint(net.ParseIP("2001:0db8:85a3:0000:0000:8a2e:0370:7334").To16())
	if i != 128 {
		t.Error("not ipv6")
	}

	_, _, e := iptoint(net.ParseIP("hola caracola").To4())
	if e == nil {
		t.Error("not a valid ip to convert to int")
	}
}