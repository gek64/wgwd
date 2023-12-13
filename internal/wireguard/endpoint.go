package wireguard

import (
	"encoding/hex"
	"fmt"
	"github.com/gek64/gek/gExec"
	"net/netip"
	"os/exec"
	"strings"
)

type EndpointConfig struct {
	PublicKey string
	AddrPort  netip.AddrPort
}

// NewEndpointConfig 新建 EndpointConfig
func NewEndpointConfig(publicKey string, addrPort string) (*EndpointConfig, error) {
	parseAddrPort, err := netip.ParseAddrPort(addrPort)
	if err != nil {
		return nil, err
	}

	return &EndpointConfig{
		PublicKey: publicKey,
		AddrPort:  parseAddrPort,
	}, nil
}

// SetPublicKey 为当前 EndpointConfig 设置新的公钥
func (e *EndpointConfig) SetPublicKey(publicKey string) {
	e.PublicKey = publicKey
}

// SetAddr 为当前 EndpointConfig 设置新的地址
func (e *EndpointConfig) SetAddr(addr string) error {
	parseAddr, err := netip.ParseAddr(addr)
	if err != nil {
		return err
	}
	parseAddrPort := netip.AddrPortFrom(parseAddr, e.AddrPort.Port())
	return e.SetAddrPort(parseAddrPort)
}

// SetPort 为当前 EndpointConfig 设置新的端口
func (e *EndpointConfig) SetPort(port int) error {
	parseAddrPort := netip.AddrPortFrom(e.AddrPort.Addr(), uint16(port))
	return e.SetAddrPort(parseAddrPort)
}

// SetAddrPort 为当前 EndpointConfig 设置 AddrPort
func (e *EndpointConfig) SetAddrPort(addrPort netip.AddrPort) error {
	e.AddrPort = addrPort
	return nil
}

// ApplyEndpointConfig 应用 endpoint
func (e *EndpointConfig) ApplyEndpointConfig(wgInterface string) (err error) {
	return gExec.Run(exec.Command("wg", "set", wgInterface, "peer", e.PublicKey, "endpoint", e.AddrPort.String()))
}

// ApplyNewEndpointConfig 对当前 endpoint 进行修改后再应用
func (e *EndpointConfig) ApplyNewEndpointConfig(wgInterface string, newEndpointKey string, newEndpointAddr string, newEndpointPort int) (err error) {
	if newEndpointKey != "" {
		e.SetPublicKey(newEndpointKey)
	}

	if newEndpointAddr != "" {
		err = e.SetAddr(newEndpointAddr)
		if err != nil {
			return err
		}
	}

	if newEndpointPort >= 0 && newEndpointPort <= 65535 {
		err = e.SetPort(newEndpointPort)
		if err != nil {
			return err
		}
	}

	// 应用 endpointConfig
	err = e.ApplyEndpointConfig(wgInterface)
	if err != nil {
		return fmt.Errorf("unable to update wireguard endpoint\n%v", err)
	}
	return nil
}

//----------------------------------------------------------------------------------------------------------------------

// GetEndpointConfigByKey 获取指定 key 的 endpointConfig
func GetEndpointConfigByKey(wgInterface string, peerKey string) (endpointConfig *EndpointConfig, err error) {
	endpoints, err := GetEndpointConfigs(wgInterface)
	if err != nil {
		return nil, err
	}

	for _, e := range endpoints {
		if e.PublicKey == peerKey {
			return e, nil
		}
	}
	return nil, fmt.Errorf("no endpoint with peer key %s found", peerKey)
}

// GetEndpointConfigs 获取 endpointConfigs
func GetEndpointConfigs(wgInterface string) (endpointConfigs []*EndpointConfig, err error) {
	endpointsString, err := gExec.Output(exec.Command("wg", "show", wgInterface, "endpoints"))
	if err != nil {
		return nil, err
	}

	endpointArray := strings.Split(hex.EncodeToString(endpointsString), "\n")

	for _, endpoint := range endpointArray {
		e := strings.Split(endpoint, "\u0009")
		if len(e) >= 2 {
			var ep EndpointConfig
			ep.PublicKey = e[0]
			ep.AddrPort, err = netip.ParseAddrPort(strings.ReplaceAll(e[1], "\r", ""))
			if err == nil {
				endpointConfigs = append(endpointConfigs, &ep)
			}
		}
	}

	if len(endpointConfigs) <= 0 {
		return nil, fmt.Errorf("no endpoints with interface %s found", wgInterface)
	}
	return endpointConfigs, nil
}
