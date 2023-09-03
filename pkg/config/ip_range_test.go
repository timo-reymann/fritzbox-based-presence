package config

import (
	"net"
	"testing"
)

func TestIpRangeDecoder_Decode(t *testing.T) {
	decoder := IpRangeDecoder{}
	err := decoder.Decode("1.1.1.1/24")
	if err != nil {
		t.Errorf("Expected decoder to parse IP range accordingly")
	}

	if decoder.Mask.String() != net.IPv4Mask(255, 255, 255, 0).String() {
		t.Errorf("deocder is not parsing mask properly")
	}

	if err = decoder.Decode("adsf"); err == nil {
		t.Errorf("Expected invalid CIDR to lead to an error")
	}
}

func TestAllowsIP(t *testing.T) {
	config = AppConfig{
		AuthIpRange: IpRangeDecoder(net.IPNet{
			IP:   net.ParseIP("1.1.1.1"),
			Mask: net.CIDRMask(24, 32),
		}),
	}
	if !AllowsIP(net.ParseIP("1.1.1.1")) {
		t.Errorf("Should allow IP in range")
	}

	if AllowsIP(net.ParseIP("8.8.8.8")) {
		t.Errorf("Should not allow IP out of range")
	}
}
