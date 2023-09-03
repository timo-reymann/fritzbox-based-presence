package config

import (
	"net"
)

type IpRangeDecoder net.IPNet

func (iprd *IpRangeDecoder) Decode(value string) error {
	_, cidr, err := net.ParseCIDR(value)
	if err != nil {
		return err
	}

	decoded := IpRangeDecoder(*cidr)
	*iprd = decoded
	return nil
}

// AllowsIP checks if the given IP is allowed by auth CIDR
func AllowsIP(ip net.IP) bool {
	ipnet := net.IPNet(config.AuthIpRange)
	return ipnet.Contains(ip)
}
