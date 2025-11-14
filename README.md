# WireGuard watchdog(wgwd)

## Features

- Get IP from remote or a file
- Periodically update WireGuard endpoint IP

## Usage

```sh
# Get local network information from a file
wgwd get -remote_interface="pppoe-wan" -wg_interface="wg0" file -filepath="./home.json"
## Get local network information from a file and decrypt the file
wgwd get -remote_interface="pppoe-wan" -wg_interface="wg0" file -filepath="./home.json" -encryption_key="admin123"
## Loop get local network information from a file and decrypt the file
wgwd get -remote_interface="pppoe-wan" -wg_interface="wg0" -interval="5m" file -filepath="./home.json" -encryption_key="admin123"

# Get local network information from s3 server
wgwd get -remote_interface="pppoe-wan" -wg_interface="wg0" s3 -endpoint="https://s3.amazonaws.com" -access_key_id="admin" -secret_access_key="adminadmin" -bucket="storage" -object_path="home.json"
## Get local network information from minio s3 server
wgwd get -remote_interface="pppoe-wan" -wg_interface="wg0" s3 -endpoint="http://192.168.1.185:9000" -path_style -access_key_id="admin" -secret_access_key="adminadmin" -bucket="storage" -object_path="home.json"
## Get local network information from minio s3 server and decrypt the file
wgwd get -remote_interface="pppoe-wan" -wg_interface="wg0" s3 -endpoint="http://192.168.1.185:9000" -path_style -access_key_id="admin" -secret_access_key="adminadmin" -bucket="storage" -object_path="home.json" -encryption_key="admin123"
## Get Get local network information from minio s3 server and decrypt the file
wgwd get -remote_interface="pppoe-wan" -wg_interface="wg0" -interval="5m" s3 -endpoint="http://192.168.1.185:9000" -path_style -access_key_id="admin" -secret_access_key="adminadmin" -bucket="storage" -object_path="home.json" -encryption_key="admin123"

# Get local network information from webdav server
wgwd get -remote_interface="pppoe-wan" -wg_interface="wg0" webdav -endpoint="http://192.168.1.2/" -filepath="/dav/home.json"
## Get local network information from webdav server and decrypt the file
wgwd get -remote_interface="pppoe-wan" -wg_interface="wg0" webdav -endpoint="http://192.168.1.2/" -filepath="/dav/home.json" -encryption_key="admin123"
## Loop Get local network information from webdav server and decrypt the file
wgwd get -remote_interface="pppoe-wan" -wg_interface="wg0" -interval="5m" webdav -endpoint="http://192.168.1.2/" -filepath="/dav/home.json" -encryption_key="admin123"

# Decrypt a encrypted file
wgwd decrypt -filepath "./home.json" -encryption_key="admin123"
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

## Install Service(WebDAV usage examples)

### Linux(systemd)

```sh
curl -Lo "/etc/systemd/system/wgwd.service" "https://github.com/gek64/wgwd/raw/main/configs/systemd/wgwd_webdav.service"
systemctl enable wgwd.service && systemctl restart wgwd.service && systemctl status wgwd.service
curl -Lo "/etc/systemd/system/wgwd.timer" "https://github.com/gek64/wgwd/raw/main/configs/systemd/wgwd.timer"
systemctl enable wgwd.timer && systemctl restart wgwd.timer && systemctl status wgwd.timer
```

### Alpine Linux(openrc)

```sh
curl -Lo "/etc/init.d/wgwd" "https://github.com/gek64/wgwd/raw/main/configs/openrc/wgwd_webdav"
chmod +x /etc/init.d/wgwd
rc-update add wgwd && rc-service wgwd restart && rc-service wgwd status
```

### FreeBSD(rc.d)

```sh
mkdir /usr/local/etc/rc.d/
curl -Lo "/usr/local/etc/rc.d/wgwd" "https://github.com/gek64/wgwd/raw/main/configs/rc.d/wgwd_webdav"
chmod +x /usr/local/etc/rc.d/wgwd
service wgwd enable && service wgwd restart && service wgwd status
```

### OpenWRT(init.d)

```sh
curl -Lo "/etc/init.d/wgwd" "https://github.com/gek64/wgwd/raw/main/configs/init.d/wgwd_webdav"
chmod +x /etc/init.d/wgwd
service wgwd enable && service wgwd restart && service wgwd status
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
