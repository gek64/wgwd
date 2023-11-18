package main

import (
	"flag"
	"fmt"
	"github.com/gek64/gek/gToolbox"
	"log"
	"os"
	"time"
	"wgwd/internal/client"
)

var (
	cliId        string
	cliUrl       string
	cliInterface string

	cliWgInterface string
	cliWgPeerKey   string

	cliUsername              string
	cliPassword              string
	cliInterval              time.Duration
	cliSkipCertificateVerify bool

	cliHelp    bool
	cliVersion bool
)

func init() {
	flag.StringVar(&cliId, "id", "", "-id 80000000-4000-4000-4000-120000000000")
	flag.StringVar(&cliUrl, "url", "", "-url http://localhost:1996/record/")
	flag.StringVar(&cliInterface, "interface", "", "-interface pppoe0")

	flag.StringVar(&cliWgInterface, "wg_interface", "", "-wg_interface wg0")
	flag.StringVar(&cliWgPeerKey, "wg_peer_key", "", "-wg_peer_key aInIXfBwnwrfr/oc8oW2Vhyhh/5v3mvS5MmYQbMiXm4=")

	flag.StringVar(&cliUsername, "username", "", "-username bob")
	flag.StringVar(&cliPassword, "password", "", "-password 123456")
	flag.DurationVar(&cliInterval, "interval", 0, "-interval 1h")
	flag.BoolVar(&cliSkipCertificateVerify, "skip-certificate-verify", false, "-skip-certificate-verify")

	flag.BoolVar(&cliHelp, "h", false, "show help")
	flag.BoolVar(&cliVersion, "v", false, "show version")
	flag.Parse()

	// 重写显示用法函数
	flag.Usage = func() {
		var helpInfo = `Usage:
wgwd [Command] {Server Option} [Other Option]
	
Command:
  -h                : show help
  -v                : show version

Server Option:
  -id            <ID>          : set server id
  -url           <Url>         : set server url
  -interface     <Name>        : set server interface name
  -wg_interface  <Name>        : set wireguard interface name

Other Option:
  -wg_peer_key   <Key>         : set wireguard peer key
  -interval      <Time>        : set client interval
  -skip-certificate-verify     : skip tls certificate verification for http requests
  -username      <Username>    : set client basic auth username
  -password      <Password>    : set client basic auth password
	
Example:
  1) wgwd -id center -url http://localhost:1996/ -interface pppoe0 -wg_interface wg0
  2) wgwd -id center -url http://localhost:1996/ -interface pppoe0 -wg_interface wg0 -interval 5m
  3) wgwd -h
  4) wgwd -v`

		fmt.Println(helpInfo)
	}

	// 打印帮助信息
	if len(os.Args) == 1 || cliHelp {
		flag.Usage()
		os.Exit(0)
	}

	// 打印版本信息
	if cliVersion {
		fmt.Println("v1.00")
		os.Exit(0)
	}

	// 检查运行库是否完整
	err := gToolbox.CheckToolbox([]string{"wg"})
	if err != nil {
		log.Panicln(err)
	}

	// 输入参数检测
	if cliId == "" || cliUrl == "" || cliInterface == "" || cliWgInterface == "" {
		log.Panicln("incomplete input parameters")
	}
}

func main() {
	// cliInterval 为 0, 则只执行一次就停止, 若 cliInterval 不为 0,则定时循环运行
	if cliInterval != 0 {
		client.UpdateWireGuardEndpointLoop(cliId, cliUrl, cliInterface, cliWgInterface, cliWgPeerKey, cliUsername, cliPassword, cliSkipCertificateVerify, cliInterval)
	} else {
		err := client.UpdateWireGuardEndpoint(cliId, cliUrl, cliInterface, cliWgInterface, cliWgPeerKey, cliUsername, cliPassword, cliSkipCertificateVerify)
		if err != nil {
			log.Println(err)
		} else {
			log.Println("updating wireguard endpoint succeeded")
		}
	}
}
