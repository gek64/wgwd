# Netinfo WireGuard(nwg)

## Features

- Get IP from netinfo
- Update WireGuard Endpoint IP

## Usage

```
Usage:
nwg [Command] {Server Option} [Other Option]
	
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
  1) nwg -id DEVICE_UUID -url http://localhost:1996/record/ -interface pppoe0 -wg_interface wg0
  2) nwg -id DEVICE_UUID -url http://localhost:1996/record/ -interface pppoe0 -wg_interface wg0 -interval 30m
  3) nwg -h
  4) nwg -v
```

## Install

```sh
# system is linux(debian,redhat linux,ubuntu,fedora...) and arch is amd64
curl -Lo /usr/local/bin/nwg https://github.com/gek64/nwg/releases/latest/download/nwg-linux-386
chmod +x /usr/local/bin/nwg

# system is freebsd and arch is amd64
curl -Lo /usr/local/bin/nwg https://github.com/gek64/nwg/releases/latest/download/nwg-freebsd-amd64
chmod +x /usr/local/bin/nwg
```

## Install Service

### Linux(systemd)

```sh
curl -Lo /etc/systemd/system/nwg.service https://github.com/gek64/nwg/raw/main/configs/nwg.service
systemctl enable nwg && systemctl restart nwg && systemctl status nwg
```

### Linux(openrc)

```sh
curl -Lo /etc/init.d/nwg https://github.com/gek64/nwg/raw/main/configs/nwg.openrc
chmod +x /etc/init.d/nwg
rc-update add nwg && rc-service nwg restart && rc-service nwg status
```

### FreeBSD(rc.d)

```sh
mkdir /usr/local/etc/rc.d/
curl -Lo /usr/local/etc/rc.d/nwg https://github.com/gek64/nwg/raw/main/configs/nwg.rcd
chmod +x /usr/local/etc/rc.d/nwg
service nwg enable && service nwg restart && service nwg status
```

## Compile

### How to compile if prebuilt binaries are not found

```sh
git clone https://github.com/gek64/nwg.git
cd nwg
go build -v -trimpath -ldflags "-s -w"
```

## License

- **GPL-3.0 License**
- See `LICENSE` for details
