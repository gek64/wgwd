package netinfo

import (
	"encoding/json"
	"fmt"
	"github.com/gek64/gek/gNet"
	"wgwd/internal/receive/preload"
)

// GetPublicIP 从网络信息中获取公共 IP
func (r *Data) GetPublicIP(interfaceName string) (ip string, err error) {
	for _, netInterface := range r.NetInterfaces {
		if netInterface.Name == interfaceName {
			for _, ip := range netInterface.IPs {
				isPublic, _ := gNet.IsPublic(ip.String())
				if isPublic {
					return ip.String(), nil
				}
			}
		}
	}
	return "", fmt.Errorf("no valid public IP found in network infomation data")
}

// GetFromJsonBytes 从加密的比特切片中获取 *Data
func GetFromJsonBytes(jsonBytes []byte, encryptionKey []byte) (data *Data, err error) {
	// 解密, encryptionKey 长度为 0 的情况, 会直接返回输入的密文
	jsonBytes, err = preload.GetDecryptedPreload(jsonBytes, encryptionKey, preload.AssociatedDataSize)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(jsonBytes, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
