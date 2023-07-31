package wireguard

import (
    "fmt"
    "github.com/gek64/gek/gExec"
    "net/netip"
    "os/exec"
    "strings"
)

// https://manpages.debian.org/unstable/wireguard-tools/wg.8.en.html#set

type Endpoint struct {
    PublicKey string
    AddrPort  netip.AddrPort
}

// GetEndpoints 获取 endpoints
func GetEndpoints(interfaceName string) (endpoints []Endpoint, err error) {
    endpointsString, err := gExec.Output(exec.Command("wg", "show", interfaceName, "endpoints"))
    if err != nil {
        return nil, err
    }
    endpointArray := strings.Split(endpointsString, "\n")

    for _, endpoint := range endpointArray {
        e := strings.Split(endpoint, "\u0009")
        if len(e) >= 2 {
            var ep Endpoint
            ep.PublicKey = e[0]
            ep.AddrPort, err = netip.ParseAddrPort(strings.ReplaceAll(e[1], "\r", ""))
            if err == nil {
                endpoints = append(endpoints, ep)
            }
        }
    }

    if len(endpoints) <= 0 {
        return nil, fmt.Errorf("no endpoints")
    }
    return endpoints, nil
}

// GetEndpointAddrPort 获取指定 key 的 endpoint 地址与端口
func GetEndpointAddrPort(interfaceName string, peerKey string) (addrPort netip.AddrPort, err error) {
    endpoints, err := GetEndpoints(interfaceName)
    if err != nil {
        return netip.AddrPort{}, err
    }

    for _, endpoint := range endpoints {
        if endpoint.PublicKey == peerKey {
            return endpoint.AddrPort, nil
        }
    }
    return netip.AddrPort{}, fmt.Errorf("no ip address and port")
}

// SetEndpoint 设置 endpoint
func SetEndpoint(interfaceName string, peerKey string, endpoint string) (err error) {
    return gExec.Run(exec.Command("wg", "set", interfaceName, "peer", peerKey, "endpoint", endpoint))
}
