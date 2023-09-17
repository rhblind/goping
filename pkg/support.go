package pkg

import "net/netip"

// Hosts returns a slice of all the usable IP addresses in a given CIDR
func Hosts(cidr string) []netip.Addr {
	prefix, err := netip.ParsePrefix(cidr)
	if err != nil {
		panic(err)
	}

	var ips []netip.Addr
	for addr := prefix.Addr(); prefix.Contains(addr); addr = addr.Next() {
		ips = append(ips, addr)
	}

	if len(ips) < 2 {
		return ips
	}

	return ips[1 : len(ips)-1]
}

// FilterBy takes a slice of any type and a test function and returns a slice of the same type
func FilterBy[T any](arg []T, test func(T) bool) (ret []T) {
	for _, x := range arg {
		if test(x) {
			ret = append(ret, x)
		}
	}
	return
}
