package client

import (
	"fmt"
	"github.com/gek64/gek/gNet"
	"github.com/imroc/req/v3"
	"net/netip"
	"time"
)

type NetInterface struct {
	Name string       `json:"name"`
	IPs  []netip.Addr `json:"ips"`
	Mac  string       `json:"mac,omitempty"`
}

type Record struct {
	ID            string         `json:"id"`
	Description   string         `json:"description,omitempty"`
	RequestIP     string         `json:"requestIP,omitempty"`
	NetInterfaces []NetInterface `json:"netInterfaces"`
}

// GetRecord 从服务端 netinfo 服务获取指定 id 的记录
func GetRecord(id string, serverUrl string, username string, password string, skipCertVerify bool) (record Record, err error) {
	client := req.C()

	// 默认不启用跳过TLS证书检测
	if skipCertVerify {
		client.EnableInsecureSkipVerify()
	}

	// 发送GET请求
	resp, err := client.R().
		SetQueryParam("id", id).
		SetSuccessResult(&record).
		SetRetryCount(3).
		SetRetryBackoffInterval(1*time.Second, 5*time.Second).
		SetBasicAuth(username, password).
		Get(serverUrl)
	if err != nil {
		return Record{}, err
	}

	// 返回值
	if resp.IsSuccessState() {
		return record, nil
	} else {
		return Record{}, fmt.Errorf(resp.ToString())
	}
}

// GetNewEndpointAddr 从记录中获取 endpoint 的地址
func (r Record) GetNewEndpointAddr(targetInterfaceName string) (ip string, err error) {
	for _, netInterface := range r.NetInterfaces {
		if netInterface.Name == targetInterfaceName {
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
