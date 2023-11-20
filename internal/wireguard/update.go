package wireguard

import (
	"fmt"
)

// UpdateEndpoint 更新 endpoint
func UpdateEndpoint(wgInterface string, wgPeerKey string, newEndpointAddr string, newEndpointPort int) (err error) {
	// 指定了 peer key 就只更新匹配的 EndpointConfig, 未指定则更新所有的 EndpointConfig
	if wgPeerKey != "" {
		// 获取单个 EndpointConfig
		endpointConfig, err := GetEndpointConfigByKey(wgInterface, wgPeerKey)
		if err != nil {
			return fmt.Errorf("unable to get wireguard endpoint from interface %s using key %s\n%v", wgInterface, wgPeerKey, err)
		}
		err = endpointConfig.ApplyNewEndpointConfig(wgInterface, wgPeerKey, newEndpointAddr, newEndpointPort)
		if err != nil {
			return err
		}
	} else {
		// 获取所有 EndpointConfig
		endpointConfigs, err := GetEndpointConfigs(wgInterface)
		if err != nil {
			return fmt.Errorf("unable to get wireguard endpoints from interface %s\n%v", wgInterface, err)
		}
		for _, endpointConfig := range endpointConfigs {
			err = endpointConfig.ApplyNewEndpointConfig(wgInterface, wgPeerKey, newEndpointAddr, newEndpointPort)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
