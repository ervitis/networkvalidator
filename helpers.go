package networkvalidator

import (
	"errors"
	"math/big"
	"net"
)

func getiprange(ipnet *net.IPNet) *rangeIp {
	fip := inc(ipnet.IP)

	l, bits := ipnet.Mask.Size()
	if l == bits {
		lip := make([]byte, len(fip))
		copy(lip, fip)
		return &rangeIp{fip, lip}
	}

	fipint, bits, err := iptoint(fip)
	if err != nil {
		return nil
	}

	hlen := uint(bits) - uint(l)
	lipint := big.NewInt(1)
	lipint.Lsh(lipint, hlen)
	lipint.Sub(lipint, big.NewInt(1))
	lipint.Or(lipint, fipint)

	e := dec(inttoip(lipint, bits))

	return &rangeIp{fip, e}
}

func iptoint(ip net.IP) (*big.Int, int, error) {
	v := &big.Int{}
	v.SetBytes([]byte(ip))

	if len(ip) == net.IPv4len {
		return v, 32, nil
	} else if len(ip) == net.IPv6len {
		return v, 128, nil
	} else {
		return nil, -1, errors.New("unsupported length")
	}
}

func inttoip(ip *big.Int, bits int) net.IP {
	b := ip.Bytes()
	ret := make([]byte, bits/8)

	for i := 1; i <= len(b); i++ {
		ret[len(ret)-i] = b[len(b)-i]
	}

	return net.IP(ret)
}

func inc(ip net.IP) net.IP {
	incip := make([]byte, len(ip))
	copy(incip, ip)

	for i := len(incip) - 1; i >= 0; i-- {
		incip[i]++
		if incip[i] > 0 {
			break
		}
	}

	return incip
}

func dec(ip net.IP) net.IP {
	decip := make([]byte, len(ip))
	copy(decip, ip)

	for i := len(decip) - 1; i >= 0; i -- {
		decip[i] --
		if decip[i] < 255 {
			break
		}
	}

	return decip
}
