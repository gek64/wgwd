package wireguard

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
