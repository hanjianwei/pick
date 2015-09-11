package netutils

import (
	"bytes"
	"net"
)

// NetworkBundle represents a range of IP.
type NetworkBundle struct {
	first    net.IP
	last     net.IP
	networks []*net.IPNet
}

// ByLast sort network bundles by their last ip
type ByLast []NetworkBundle

func (a ByLast) Len() int           { return len(a) }
func (a ByLast) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByLast) Less(i, j int) bool { return bytes.Compare(a[i].last, a[j].last) < 0 }

func (nb *NetworkBundle) append(network *net.IPNet) {
	if network == nil {
		return
	}

	first, last := NetworkRange(network)

	if len(nb.networks) == 0 {
		nb.first, nb.last = first, last
	} else {
		if bytes.Compare(first, nb.first) < 0 {
			nb.first = first
		}
		if bytes.Compare(last, nb.last) > 0 {
			nb.last = last
		}
	}
	nb.networks = append(nb.networks, network)
}

func (nb *NetworkBundle) extend(other *NetworkBundle) {
	if other != nil {
		for _, n := range other.networks {
			nb.append(n)
		}
	}
}
