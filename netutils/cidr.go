package netutils

import "net"

// SpanningCIDR computes network covers given IP addresses
func SpanningCIDR(first, last net.IP) *net.IPNet {
	_, bits := last.DefaultMask().Size()

	var network net.IPNet
	for ones := bits; !network.Contains(first); ones-- {
		network.Mask = net.CIDRMask(ones, bits)
		network.IP = last.Mask(network.Mask)
	}
	return &network
}

// CIDRPartition partitions a network on an exclude IP
func CIDRPartition(network *net.IPNet, exclude net.IP) {
}

// iPRangeToCIDRs convert an IP range to minimal number of CIDRs
// See: https://pythonhosted.org/netaddr/_modules/netaddr/ip.html#cidr_merge
func iPRangeToCIDRs(first, last net.IP) []*net.IPNet {
	// spanNetwork := SpanningCIDR(first, last)
	// spanFirst, spanLast := NetworkRange(spanNetwork)

	return nil
}

// CIDRMerge merges multiple IP networks
func CIDRMerge(ipnets []net.IPNet) []*net.IPNet {
	// ranges := make([]iPRange, len(ipnets))
	// for i, ipnet := range ipnets {
	// 	ranges[i].ipnet = &ipnet
	// 	ranges[i].init()
	// }

	// sort.Sort(byLast(ranges))

	// // Merge overlapped ranges
	// for i := len(ranges) - 1; i > 0; i-- {
	// 	if ranges[i].first <= ranges[i-1].last+1 {
	// 		if ranges[i].first < ranges[i-1].first {
	// 			ranges[i-1].first = ranges[i].first
	// 		}
	// 		ranges[i-1].last = ranges[i].last
	// 		ranges[i-1].ipnet = nil

	// 		ranges = append(ranges[:i], ranges[i+1:]...)
	// 	}
	// }

	// var merged []*net.IPNet
	// for _, ir := range ranges {
	// 	if ir.ipnet != nil {
	// 		merged = append(merged, ir.ipnet)
	// 	} else {
	// 	}
	// }

	// return merged
	return nil
}
