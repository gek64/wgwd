package main

import (
    "flag"
    "fmt"
    "github.com/gek64/gek/gToolbox"
    "log"
    "netinfo-wireguard/internal/client"
    "os"
    "time"
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
    flag.DurationVar(&cliInterval, "interval", time.Hour, "-interval 1h")
    flag.BoolVar(&cliSkipCertificateVerify, "skip-certificate-verify", false, "-skip-certificate-verify")

    flag.BoolVar(&cliHelp, "h", false, "show help")
    flag.BoolVar(&cliVersion, "v", false, "show version")
    flag.Parse()

    // 重写显示用法函数
    flag.Usage = func() {
        var helpInfo = `Usage:
netinfo-wireguard [Command] {Server Option} [Other Option]
	
Command:
  -h                : show help
  -v                : show version

Server Option:
  -id            <ID>          : set server id
  -url           <Url>         : set server url
  -interface     <Name>        : set server interface name
  -wg_interface  <Name>        : set wireguard interface name
  -wg_peer_key   <Key>         : set wireguard peer key

Other Option:
  -username      <Username>    : set client basic auth username
  -password      <Password>    : set client password
  -interval      <Time>        : set client interval
  -skip-certificate-verify     : skip tls certificate verification for http requests
	
Example:
  1) netinfo-wireguard -id 80000000-4000-4000-4000-120000000000 -url http://localhost:1996/record/ -interface pppoe0 -wg_interface wg0 -wg_peer_key aInIXfBwnwrfr/oc8oW2Vhyhh/5v3mvS5MmYQbMiXm4=
  2) netinfo-wireguard -h
  3) netinfo-wireguard -v`

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
    if cliId == "" || cliUrl == "" || cliInterface == "" || cliWgInterface == "" || cliWgPeerKey == "" {
        log.Panicln("incomplete input parameters")
    }
}

func main() {
    client.UpdateWireGuardEndpointLoop(cliId, cliUrl, cliInterface, cliWgInterface, cliWgPeerKey, cliUsername, cliPassword, cliSkipCertificateVerify, cliInterval)
}
