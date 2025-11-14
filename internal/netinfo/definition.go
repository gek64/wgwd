package netinfo

import (
	"net/netip"
	"time"
)

type NetInfo struct {
	UpdatedAt     time.Time      `json:"updatedAt,omitempty" xml:"updatedAt,omitempty" form:"updatedAt,omitempty"`
	NetInterfaces []NetInterface `json:"netInterfaces" xml:"netInterfaces" form:"netInterfaces" binding:"required"`
}

type NetInterface struct {
	Name string       `json:"name"`
	IPs  []netip.Addr `json:"ips"`
	Mac  string       `json:"mac,omitempty"`
}
