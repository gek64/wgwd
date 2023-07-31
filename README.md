# netinfo-wireguard

## Features

- Get IP from netinfo
- Update WireGuard Endpoint IP

## Usage

```
Usage:
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
  3) netinfo-wireguard -v
```

## Install

```sh
# system is linux(debian,redhat linux,ubuntu,fedora...) and arch is amd64
curl -Lo /usr/local/bin/netinfo-wireguard https://github.com/gek64/netinfo-wireguard/releases/latest/download/netinfo-wireguard-linux-386
chmod +x /usr/local/bin/netinfo-wireguard

# system is freebsd and arch is amd64
curl -Lo /usr/local/bin/netinfo-wireguard https://github.com/gek64/netinfo-wireguard/releases/latest/download/netinfo-wireguard-freebsd-amd64
chmod +x /usr/local/bin/netinfo-wireguard
```

## Install Service

### Linux(systemd)

```sh
curl -Lo /etc/systemd/system/netinfo.service https://github.com/gek64/netinfo/raw/main/configs/netinfo.service
systemctl enable netinfo && systemctl restart netinfo && systemctl status netinfo
```

### Linux(openrc)

```sh
curl -Lo /etc/init.d/netinfo https://github.com/gek64/netinfo/raw/main/configs/netinfo.openrc
chmod +x /etc/init.d/netinfo
rc-update add netinfo && rc-service netinfo restart && rc-service netinfo status
```

### FreeBSD(rc.d)

```sh
mkdir /usr/local/etc/rc.d/
curl -Lo /usr/local/etc/rc.d/netinfo https://github.com/gek64/netinfo/raw/main/configs/netinfo.rcd
chmod +x /usr/local/etc/rc.d/netinfo
service netinfo enable && service netinfo restart && service netinfo status
```

## Compile

### How to compile if prebuilt binaries are not found

```sh
git clone https://github.com/gek64/netinfo.git
cd netinfo
go build -v -trimpath -ldflags "-s -w"
```

## Test

```sh
# start netinfo server at 127.0.0.1:1996
netinfo -server localhost:1996

# start netinfo client
netinfo -client http://localhost:1996/record -interval 15m -description main

# check info
curl http://localhost:1996/record/all
```

## License

- **GPL-3.0 License**
- See `LICENSE` for details
