package nconnect

import (
	"fmt"
	"github.com/imroc/req/v3"
	"log"
	"time"
	"wgwd/internal/netinfo"
	"wgwd/internal/wireguard"
)

// getNetInfo 从 nconnect 服务器获取指定 id 的网络信息
func getNetInfo(id string, endpoint string, username string, password string, allowInsecure bool) (netInfoInMemoryData *netinfo.NetInfo, err error) {
	client := req.C()

	// 默认不启用跳过TLS证书检测
	if allowInsecure {
		client.EnableInsecureSkipVerify()
	}

	// 发送GET请求
	resp, err := client.R().
		SetQueryParam("id", id).
		SetSuccessResult(&netInfoInMemoryData).
		SetRetryCount(3).
		SetRetryBackoffInterval(1*time.Second, 5*time.Second).
		SetBasicAuth(username, password).
		Get(endpoint)
	if err != nil {
		return nil, err
	}

	// 返回值
	if resp.IsSuccessState() {
		return netInfoInMemoryData, nil
	} else {
		return nil, fmt.Errorf(resp.ToString())
	}
}

func ReceiveRequest(id string, endpoint string, username string, password string, allowInsecure bool, remoteInterface string, wgInterface string, wgPeerKey string) (err error) {
	inMemoryData, err := getNetInfo(id, endpoint, username, password, allowInsecure)
	if err != nil {
		return err
	}
	publicIP, err := inMemoryData.GetPublicIP(remoteInterface)
	if err != nil {
		return err
	}
	return wireguard.UpdateEndpoint(wgInterface, wgPeerKey, publicIP, -1)
}

func ReceiveRequestLoop(id string, endpoint string, username string, password string, allowInsecure bool, remoteInterface string, wgInterface string, wgPeerKey string, interval time.Duration) {
	for {
		err := ReceiveRequest(id, endpoint, username, password, allowInsecure, remoteInterface, wgInterface, wgPeerKey)
		if err != nil {
			log.Println(err)
		}
		time.Sleep(interval)
	}
}
