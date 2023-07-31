package client

import (
    "log"
    "netinfo-wireguard/internal/wireguard"
    "strconv"
    "time"
)

func UpdateWireGuardEndpoint(id string, url string, netInterface string, wgInterface string, wgPeerKey string, username string, password string, skipCertificateVerify bool) (err error) {
    record, err := GetRecord(id, url, username, password, skipCertificateVerify)
    if err != nil {
        return err
    }

    newEndpointAddr, err := record.GetNewEndpointAddr(netInterface)
    if err != nil {
        return err
    }

    addrPort, err := wireguard.GetEndpointAddrPort(wgInterface, wgPeerKey)
    if err != nil {
        return err
    }

    if addrPort.Addr().String() != newEndpointAddr {
        c := wireguard.NewEndpointConfig(wgInterface, wgPeerKey, newEndpointAddr+":"+strconv.Itoa(int(addrPort.Port())))
        return c.ChangeEndpoint()
    }

    return nil
}

func UpdateWireGuardEndpointLoop(id string, url string, netInterface string, wgInterface string, wgPeerKey string, username string, password string, skipCertificateVerify bool, interval time.Duration) {
    for {
        err := UpdateWireGuardEndpoint(id, url, netInterface, wgInterface, wgPeerKey, username, password, skipCertificateVerify)
        if err != nil {
            log.Println(err)
        } else {
            log.Println("update completed")
        }

        time.Sleep(interval)
    }
}
