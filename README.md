# WireGuard watchdog(wgwd)

## Features

- Get IP from remote or a file
- Periodically update WireGuard endpoint IP

## Usage

```sh
# Get local network information from a file
wgwd get -remote_interface="pppoe0" -wg_interface="wg0" file -filepath="./center.json"
## Get local network information from a file and decrypt the file
wgwd get -remote_interface="pppoe0" -wg_interface="wg0" file -filepath="./center.json" -encryption_key="admin123"
## Loop get local network information from a file and decrypt the file
wgwd get -remote_interface="pppoe0" -wg_interface="wg0" -interval="5m" file -filepath="./center.json" -encryption_key="admin123"

# Get local network information from s3 server
wgwd get -remote_interface="pppoe0" -wg_interface="wg0" s3 -endpoint="https://s3.amazonaws.com" -access_key_id="admin" -secret_access_key="adminadmin" -bucket="storage" -object_path="center.json"
## Get local network information from minio s3 server
wgwd get -remote_interface="pppoe0" -wg_interface="wg0" s3 -endpoint="http://192.168.1.185:9000" -path_style -access_key_id="admin" -secret_access_key="adminadmin" -bucket="storage" -object_path="center.json"
## Get local network information from minio s3 server and decrypt the file
wgwd get -remote_interface="pppoe0" -wg_interface="wg0" s3 -endpoint="http://192.168.1.185:9000" -path_style -access_key_id="admin" -secret_access_key="adminadmin" -bucket="storage" -object_path="center.json" -encryption_key="admin123"
## Get Get local network information from minio s3 server and decrypt the file
wgwd get -remote_interface="pppoe0" -wg_interface="wg0" -interval="5m" s3 -endpoint="http://192.168.1.185:9000" -path_style -access_key_id="admin" -secret_access_key="adminadmin" -bucket="storage" -object_path="center.json" -encryption_key="admin123"

# Get local network information from webdav server
wgwd get -remote_interface="pppoe0" -wg_interface="wg0" webdav -endpoint="http://192.168.1.2/" -filepath="/dav/center.json"
## Get local network information from webdav server and decrypt the file
wgwd get -remote_interface="pppoe0" -wg_interface="wg0" webdav -endpoint="http://192.168.1.2/" -filepath="/dav/center.json" -encryption_key="admin123"
## Loop Get local network information from webdav server and decrypt the file
wgwd get -remote_interface="pppoe0" -wg_interface="wg0" -interval="5m" webdav -endpoint="http://192.168.1.2/" -filepath="/dav/center.json" -encryption_key="admin123"

# Get local network information from nconnect server
wgwd get -remote_interface="pppoe0" -wg_interface="wg0" nconnect -id="center" -endpoint="http://localhost:1996/"
## Loop Get local network information from nconnect server
wgwd get -remote_interface="pppoe0" -wg_interface="wg0" -interval="5m" nconnect -id="center" -endpoint="http://localhost:1996/"

# Decrypt a encrypted file
wgwd decrypt -filepath "./center.json" -encryption_key="admin123"
```

## Install

```sh
# system is linux(debian,redhat linux,ubuntu,fedora...) and arch is amd64
curl -Lo /usr/local/bin/wgwd https://github.com/gek64/wgwd/releases/latest/download/wgwd-linux-amd64
chmod +x /usr/local/bin/wgwd

# system is freebsd and arch is amd64
curl -Lo /usr/local/bin/wgwd https://github.com/gek64/wgwd/releases/latest/download/wgwd-freebsd-amd64
chmod +x /usr/local/bin/wgwd
```

## Install Service

### Linux(systemd)

```sh
curl -Lo /etc/systemd/system/xxx.service https://github.com/gek64/wgwd/raw/main/configs/systemd/xxx.service
systemctl enable xxx.service && systemctl restart xxx.service && systemctl status xxx.service
curl -Lo /etc/systemd/system/xxx.timer https://github.com/gek64/wgwd/raw/main/configs/systemd/xxx.timer
systemctl enable xxx.timer && systemctl restart xxx.timer && systemctl status xxx.timer
```

### Alpine Linux(openrc)

```sh
curl -Lo /etc/init.d/xxx https://github.com/gek64/wgwd/raw/main/configs/openrc/xxx
chmod +x /etc/init.d/xxx
rc-update add xxx && rc-service xxx restart && rc-service xxx status
```

### FreeBSD(rc.d)

```sh
mkdir /usr/local/etc/rc.d/
curl -Lo /usr/local/etc/rc.d/xxx https://github.com/gek64/wgwd/raw/main/configs/rc.d/xxx
chmod +x /usr/local/etc/rc.d/xxx
service xxx enable && service xxx restart && service xxx status
```

### OpenWRT(init.d)

```sh
curl -Lo /etc/init.d/xxx https://github.com/gek64/wgwd/raw/main/configs/init.d/xxx
chmod +x /etc/init.d/xxx
service xxx enable && service xxx restart && service xxx status
```

## Compile

### How to compile if prebuilt binaries are not found

```sh
git clone https://github.com/gek64/wgwd.git
cd wgwd
export CGO_ENABLED=0
go build -v -trimpath -ldflags "-s -w"
```

### For mipsle router

```sh
git clone https://github.com/gek64/wgwd.git
cd wgwd
export GOOS=linux
export GOARCH=mipsle
export GOMIPS=softfloat
export CGO_ENABLED=0
go build -v -trimpath -ldflags "-s -w"
```

## License

- **GPL-3.0 License**
- See `LICENSE` for details
