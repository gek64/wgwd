package netinfo

import (
	"net/netip"
	"time"
)

type Data struct {
	ID            string         `json:"id" xml:"id" form:"id" binding:"required"`
	UpdatedAt     time.Time      `json:"updatedAt,omitempty" xml:"updatedAt,omitempty" form:"updatedAt,omitempty"`
	RequestIP     netip.Addr     `json:"requestIP,omitempty" xml:"requestIP,omitempty" form:"requestIP,omitempty"`
	NetInterfaces []NetInterface `json:"netInterfaces" xml:"netInterfaces" form:"netInterfaces" binding:"required"`
}

type NetInterface struct {
	Name string       `json:"name"`
	IPs  []netip.Addr `json:"ips"`
	Mac  string       `json:"mac,omitempty"`
}
