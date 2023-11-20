package nconnect

import (
	"fmt"
	"github.com/gek64/gek/gNet"
	"github.com/imroc/req/v3"
	"log"
	"net/netip"
	"time"
	"wgwd/internal/wireguard"
)

type NetInfoInMemoryData struct {
	ID            string         `json:"id"`
	UpdatedAt     time.Time      `json:"updatedAt"`
	RequestIP     netip.Addr     `json:"requestIP"`
	NetInterfaces []NetInterface `json:"netInterfaces"`
}

type NetInterface struct {
	Name string       `json:"name"`
	IPs  []netip.Addr `json:"ips"`
	Mac  string       `json:"mac,omitempty"`
}

// getNetInfoInMemoryData 从 nconnect 服务获取指定 id 的网络信息
func getNetInfoInMemoryData(id string, endpoint string, username string, password string, allowInsecure bool) (netInfoInMemoryData *NetInfoInMemoryData, err error) {
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

// getPublicIP 从网络信息中获取公共 IP
func (r *NetInfoInMemoryData) getPublicIP(remoteInterface string) (ip string, err error) {
	for _, netInterface := range r.NetInterfaces {
		if netInterface.Name == remoteInterface {
			for _, ip := range netInterface.IPs {
				isPublic, _ := gNet.IsPublic(ip.String())
				if isPublic {
					return ip.String(), nil
				}
			}
		}
	}
	return "", fmt.Errorf("no valid public IP found")
}

func ReceiveRequest(id string, endpoint string, username string, password string, allowInsecure bool, remoteInterface string, wgInterface string, wgPeerKey string) (err error) {
	inMemoryData, err := getNetInfoInMemoryData(id, endpoint, username, password, allowInsecure)
	if err != nil {
		return err
	}
	publicIP, err := inMemoryData.getPublicIP(remoteInterface)
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
