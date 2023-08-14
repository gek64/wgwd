package wireguard

import "strconv"

type Config struct {
	Name      string
	Interface Interface
	Peer      Peer
}

type Interface struct {
	PrivateKey string
	Address    []string
	DNS        []string
	MTU        uint
}

type Peer struct {
	PublicKey  string
	AllowedIPs []string
	Endpoint   string
}

func NewEndpointConfig(name string, peerKey string, endpoint string) (c Config) {
	c.Name = name
	c.Peer.PublicKey = peerKey
	c.Peer.Endpoint = endpoint
	return c
}

func (c Config) ChangeEndpoint() (err error) {
	return SetEndpoint(c.Name, c.Peer.PublicKey, c.Peer.Endpoint)
}

// UpdateWireGuardEndpointByKey 按 peer key 更新指定的 endpoint
func UpdateWireGuardEndpointByKey(endpoint Endpoint, newEndpointAddr string, wgInterface string, wgPeerKey string) (err error) {
	// 当前的 endpoint 与新的 endpoint 不一样时才进行更新
	if endpoint.AddrPort.Addr().String() != newEndpointAddr {
		c := NewEndpointConfig(wgInterface, wgPeerKey, newEndpointAddr+":"+strconv.Itoa(int(endpoint.AddrPort.Port())))
		return c.ChangeEndpoint()
	}
	return nil
}
