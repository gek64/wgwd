# WireGuard watchdog(wgwd)

## Features

- Get IP from netinfo
- Update WireGuard Endpoint IP

## Usage

```
Usage:
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
  4) wgwd -v
```

## Install

```sh
# system is linux(debian,redhat linux,ubuntu,fedora...) and arch is amd64
curl -Lo /usr/local/bin/wgwd https://github.com/gek64/wgwd/releases/latest/download/wgwd-linux-386
chmod +x /usr/local/bin/wgwd

# system is freebsd and arch is amd64
curl -Lo /usr/local/bin/wgwd https://github.com/gek64/wgwd/releases/latest/download/wgwd-freebsd-amd64
chmod +x /usr/local/bin/wgwd
```

## Install Service

### Linux(systemd)

```sh
curl -Lo /etc/systemd/system/wgwd.service https://github.com/gek64/wgwd/raw/main/configs/wgwd.service
systemctl enable wgwd && systemctl restart wgwd && systemctl status wgwd
```

### Linux(openrc)

```sh
curl -Lo /etc/init.d/wgwd https://github.com/gek64/wgwd/raw/main/configs/wgwd.openrc
chmod +x /etc/init.d/wgwd
rc-update add wgwd && rc-service wgwd restart && rc-service wgwd status
```

### FreeBSD(rc.d)

```sh
mkdir /usr/local/etc/rc.d/
curl -Lo /usr/local/etc/rc.d/wgwd https://github.com/gek64/wgwd/raw/main/configs/wgwd.rcd
chmod +x /usr/local/etc/rc.d/wgwd
service wgwd enable && service wgwd restart && service wgwd status
```

### OpenWRT(init.d)

```sh
curl -Lo /etc/init.d/wgwd https://github.com/gek64/wgwd/raw/main/configs/wgwd.initd
chmod +x /etc/init.d/wgwd
service wgwd enable && service wgwd restart && service wgwd status
```

## Compile

```sh
git clone https://github.com/gek64/wgwd.git
cd wgwd
go build -v -trimpath -ldflags "-s -w"
```

## For openwrt on mipsle router

```sh
git clone https://github.com/gek64/wgwd.git
cd wgwd
export GOOS=linux
export GOARCH=mipsle
export GOMIPS=softfloat
go build -v -trimpath -ldflags "-s -w"
```

## License

- **GPL-3.0 License**
- See `LICENSE` for details
