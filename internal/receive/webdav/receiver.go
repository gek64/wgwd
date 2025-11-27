package webdav

import (
	"io"
	"log"
	"time"
	"wgwd/internal/netinfo"
	"wgwd/internal/wireguard"

	"github.com/unix755/xtools/xWebDAV"
)

// getNetInfo 从 webdav 服务器获取指定 id 的网络信息
func getNetInfo(endpoint string, username string, password string, allowInsecure bool, filepath string, encryptionKey []byte) (data *netinfo.NetInfo, err error) {
	client, err := xWebDAV.NewClient(endpoint, username, password, allowInsecure)
	response, err := client.Download(filepath)
	if err != nil {
		return nil, err
	}

	// 读取从 webdav 服务器下载的数据流
	d, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return netinfo.FromBytes(d, encryptionKey)
}

func ReceiveRequest(endpoint string, username string, password string, allowInsecure bool, filepath string, encryptionKey []byte, remoteInterface string, wgInterface string, wgPeerKey string) (err error) {
	data, err := getNetInfo(endpoint, username, password, allowInsecure, filepath, encryptionKey)
	if err != nil {
		return err
	}
	publicIP, err := data.GetPublicIP(remoteInterface)
	if err != nil {
		return err
	}
	return wireguard.UpdateEndpoint(wgInterface, wgPeerKey, publicIP, -1)
}

func ReceiveRequestLoop(endpoint string, username string, password string, allowInsecure bool, filepath string, encryptionKey []byte, remoteInterface string, wgInterface string, wgPeerKey string, interval time.Duration) {
	for {
		err := ReceiveRequest(endpoint, username, password, allowInsecure, filepath, encryptionKey, remoteInterface, wgInterface, wgPeerKey)
		if err != nil {
			log.Println(err)
		}
		time.Sleep(interval)
	}
}
