package s3

import (
	"github.com/gek64/gek/gS3"
	"io"
	"log"
	"time"
	"wgwd/internal/netinfo"
	"wgwd/internal/wireguard"
)

// getNetInfo 从 s3 服务器获取指定 id 的网络信息
func getNetInfo(endpoint string, region string, accessKeyId string, secretAccessKey string, stsToken string, pathStyle bool, allowInsecure bool, bucket string, objectPath string, encryptionKey []byte) (data *netinfo.Data, err error) {
	s := gS3.NewS3Session(endpoint, region, accessKeyId, secretAccessKey, stsToken, pathStyle, allowInsecure)
	response, err := s.GetObject(bucket, objectPath)
	if err != nil {
		return nil, err
	}

	// 读取从 s3 服务器下载的数据流
	d, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return netinfo.GetFromJsonBytes(d, encryptionKey)
}

func ReceiveRequest(endpoint string, region string, accessKeyId string, secretAccessKey string, stsToken string, pathStyle bool, allowInsecure bool, bucket string, objectPath string, encryptionKey []byte, remoteInterface string, wgInterface string, wgPeerKey string) (err error) {
	data, err := getNetInfo(endpoint, region, accessKeyId, secretAccessKey, stsToken, pathStyle, allowInsecure, bucket, objectPath, encryptionKey)
	if err != nil {
		return err
	}
	publicIP, err := data.GetPublicIP(remoteInterface)
	if err != nil {
		return err
	}
	return wireguard.UpdateEndpoint(wgInterface, wgPeerKey, publicIP, -1)
}

func ReceiveRequestLoop(endpoint string, region string, accessKeyId string, secretAccessKey string, stsToken string, pathStyle bool, allowInsecure bool, bucket string, objectPath string, encryptionKey []byte, remoteInterface string, wgInterface string, wgPeerKey string, interval time.Duration) {
	for {
		err := ReceiveRequest(endpoint, region, accessKeyId, secretAccessKey, stsToken, pathStyle, allowInsecure, bucket, objectPath, encryptionKey, remoteInterface, wgInterface, wgPeerKey)
		if err != nil {
			log.Println(err)
		}
		time.Sleep(interval)
	}
}
