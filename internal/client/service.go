package client

import (
	"fmt"
	"log"
	"time"
	"wgwd/internal/wireguard"
)

// UpdateWireGuardEndpoint 更新 wireguard endpoint
func UpdateWireGuardEndpoint(id string, url string, netInterface string, wgInterface string, wgPeerKey string, username string, password string, skipCertificateVerify bool) (err error) {
	// 从服务端 netinfo 服务获取指定 id 的记录
	record, err := GetRecord(id, url, username, password, skipCertificateVerify)
	if err != nil {
		return fmt.Errorf("unable to fetch record from %s\n%v", url, err)
	}
	// 从记录中获取新 endpoint 的地址
	newEndpointAddr, err := record.GetNewEndpointAddr(netInterface)
	if err != nil {
		return fmt.Errorf("unable to parse endpoint from %v\n%v", record, err)
	}

	// 指定了 peer key 就只更新匹配的 endpoint, 未指定则更新所有的 endpoint
	if wgPeerKey != "" {
		// 查询匹配的 endpoint
		endpoint, err := wireguard.GetEndpointByKey(wgInterface, wgPeerKey)
		if err != nil {
			return fmt.Errorf("unable to get wireguard endpoint from interface %s using key %s\n%v", wgInterface, wgPeerKey, err)
		}
		// 更新 endpoint
		err = wireguard.UpdateWireGuardEndpointByKey(endpoint, newEndpointAddr, wgInterface, wgPeerKey)
		if err != nil {
			return fmt.Errorf("unable to update wireguard endpoint\n%v", err)
		}
	} else {
		// 查询所有的endpoint
		endpoints, err := wireguard.GetEndpoints(wgInterface)
		if err != nil {
			return fmt.Errorf("unable to get wireguard endpoints from interface %s\n%v", wgInterface, err)
		}
		// 循环遍历并更新所有 endpoint
		for _, endpoint := range endpoints {
			err = wireguard.UpdateWireGuardEndpointByKey(endpoint, newEndpointAddr, wgInterface, endpoint.PublicKey)
			if err != nil {
				return fmt.Errorf("unable to update wireguard endpoint\n%v", err)
			}
		}
	}
	return nil
}

// UpdateWireGuardEndpointLoop 定时循环更新 wireguard endpoint
func UpdateWireGuardEndpointLoop(id string, url string, netInterface string, wgInterface string, wgPeerKey string, username string, password string, skipCertificateVerify bool, interval time.Duration) {
	for {
		err := UpdateWireGuardEndpoint(id, url, netInterface, wgInterface, wgPeerKey, username, password, skipCertificateVerify)
		if err != nil {
			log.Println(err)
		} else {
			log.Println("updating wireguard endpoint succeeded")
		}
		time.Sleep(interval)
	}
}
