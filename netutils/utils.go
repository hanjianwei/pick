package netutils

import (
	"encoding/binary"
	"net"
)

// From https://github.com/docker/libnetwork

// IPToUint32 converts an ipv4 to a uint32
func IPToUint32(ip net.IP) uint32 {
	return binary.BigEndian.Uint32(ip.To4())
}

// Uint32ToIP converts a uint32 to an ipv4
func Uint32ToIP(ip uint32) net.IP {
	addr := net.IPv4(0, 0, 0, 0)
	binary.BigEndian.PutUint32(addr, ip)
	return addr
}

// GetIPCopy returns a copy of the passed IP address
func GetIPCopy(from net.IP) net.IP {
	to := make(net.IP, len(from))
	copy(to, from)
	return to
}

// GetIPNetCopy returns a copy of the passed IP Network
func GetIPNetCopy(from *net.IPNet) *net.IPNet {
	if from == nil {
		return nil
	}
	bm := make(net.IPMask, len(from.Mask))
	copy(bm, from.Mask)
	return &net.IPNet{IP: GetIPCopy(from.IP), Mask: bm}
}

// NetworkRange calculates the first and last IP addresses in an IPNet
func NetworkRange(network *net.IPNet) (net.IP, net.IP) {
	if network == nil {
		return nil, nil
	}

	firstIP := network.IP.Mask(network.Mask)
	lastIP := GetIPCopy(firstIP)
	for i := 0; i < len(firstIP); i++ {
		lastIP[i] = firstIP[i] | ^network.Mask[i]
	}

	if network.IP.To4() != nil {
		firstIP = firstIP.To4()
		lastIP = lastIP.To4()
	}

	return firstIP, lastIP
}

// NetworkOverlaps detects overlap between one IPNet and another
func NetworkOverlaps(netX *net.IPNet, netY *net.IPNet) bool {
	return netX.Contains(netY.IP) || netY.Contains(netX.IP)
}

// IPDistance computes distance of IP first and last
func IPDistance(first net.IP, last net.IP) int {
	f := IPToUint32(first)
	l := IPToUint32(last)
	if f > l {
		return -int(f - l)
	}

	return int(l - f)
}
