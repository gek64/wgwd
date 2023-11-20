package file

import (
	"log"
	"os"
	"time"
	"wgwd/internal/netinfo"
	"wgwd/internal/wireguard"
)

// getNetInfo 从 file 文件获取指定 id 的网络信息
func getNetInfo(filepath string, encryptionKey []byte) (data *netinfo.Data, err error) {
	d, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	return netinfo.GetFromJsonBytes(d, encryptionKey)
}

func ReceiveRequest(filepath string, encryptionKey []byte, remoteInterface string, wgInterface string, wgPeerKey string) (err error) {
	data, err := getNetInfo(filepath, encryptionKey)
	if err != nil {
		return err
	}
	publicIP, err := data.GetPublicIP(remoteInterface)
	if err != nil {
		return err
	}
	return wireguard.UpdateEndpoint(wgInterface, wgPeerKey, publicIP, -1)
}

func ReceiveRequestLoop(filepath string, encryptionKey []byte, remoteInterface string, wgInterface string, wgPeerKey string, interval time.Duration) {
	for {
		err := ReceiveRequest(filepath, encryptionKey, remoteInterface, wgInterface, wgPeerKey)
		if err != nil {
			log.Println(err)
		}
		time.Sleep(interval)
	}
}
