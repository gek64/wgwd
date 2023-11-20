package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"time"
	"wgwd/internal/receive/file"
	"wgwd/internal/receive/nconnect"
	"wgwd/internal/receive/s3"
	"wgwd/internal/receive/webdav"
)

func main() {
	// get mode
	var id string
	var allow_insecure bool
	var encryption_key string
	var interval time.Duration
	var endpoint string
	var username string
	var password string

	// get mode file
	var filepath string

	// get mode s3
	var regin string
	var sts_token string
	var path_style bool
	var bucket string
	var object_path string

	// wireguard
	var remote_interface string
	var wg_interface string
	var wg_peer_key string

	cmds := []*cli.Command{
		{
			Name:  "get",
			Usage: "get wireguard endpoint from network information",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:        "remote_interface",
					Usage:       "set remote interface",
					Required:    true,
					Destination: &remote_interface,
				},
				&cli.StringFlag{
					Name:        "wg_interface",
					Usage:       "set wireguard interface",
					Required:    true,
					Destination: &wg_interface,
				},
				&cli.StringFlag{
					Name:        "wg_peer_key",
					Usage:       "set wireguard peer key",
					Destination: &wg_peer_key,
				},
				&cli.DurationFlag{
					Name:        "interval",
					Usage:       "set send interval",
					Destination: &interval,
				},
			},

			Subcommands: []*cli.Command{
				{
					Name:  "file",
					Usage: "get network information from filesystem",
					Flags: []cli.Flag{
						&cli.StringFlag{
							Name:        "filepath",
							Usage:       "set file path",
							Required:    true,
							Destination: &filepath,
						},
						&cli.StringFlag{
							Name:        "encryption_key",
							Usage:       "set file encryption key",
							Destination: &encryption_key,
						},
					},
					Action: func(ctx *cli.Context) error {
						if interval != 0 {
							file.ReceiveRequestLoop(filepath, []byte(encryption_key), remote_interface, wg_interface, wg_peer_key, interval)
						} else {
							err := file.ReceiveRequest(filepath, []byte(encryption_key), remote_interface, wg_interface, wg_peer_key)
							if err != nil {
								return err
							}
						}
						return nil
					},
				},
				{
					Name:  "s3",
					Usage: "get network information from s3 server",
					Flags: []cli.Flag{
						&cli.BoolFlag{
							Name:        "allow_insecure",
							Usage:       "set allow insecure connect",
							Value:       false,
							Destination: &allow_insecure,
						},
						&cli.StringFlag{
							Name:        "encryption_key",
							Usage:       "set file encryption key",
							Destination: &encryption_key,
						},
						&cli.StringFlag{
							Name:        "endpoint",
							Usage:       "set s3 server endpoint",
							Required:    true,
							Destination: &endpoint,
						},
						&cli.StringFlag{
							Name:        "regin",
							Usage:       "set s3 server regin",
							Value:       "us-east-1",
							Destination: &regin,
						},
						&cli.StringFlag{
							Name:        "access_key_id",
							Usage:       "set s3 server access key id",
							Required:    true,
							Destination: &username,
						},
						&cli.StringFlag{
							Name:        "secret_access_key",
							Usage:       "set s3 server secret access key",
							Required:    true,
							Destination: &password,
						},
						&cli.StringFlag{
							Name:        "sts_token",
							Usage:       "set s3 server sts token",
							Destination: &sts_token,
						},
						&cli.BoolFlag{
							Name:        "path_style",
							Usage:       "set s3 server path style, false: virtual host, true: path",
							Value:       false,
							Destination: &path_style,
						},
						&cli.StringFlag{
							Name:        "bucket",
							Usage:       "set s3 server bucket",
							Required:    true,
							Destination: &bucket,
						},
						&cli.StringFlag{
							Name:        "object_path",
							Usage:       "set s3 server object path",
							Required:    true,
							Destination: &object_path,
						},
					},
					Action: func(ctx *cli.Context) error {
						if interval != 0 {
							s3.ReceiveRequestLoop(endpoint, regin, username, password, sts_token, path_style, allow_insecure, bucket, object_path, []byte(encryption_key), remote_interface, wg_interface, wg_peer_key, interval)
						} else {
							err := s3.ReceiveRequest(endpoint, regin, username, password, sts_token, path_style, allow_insecure, bucket, object_path, []byte(encryption_key), remote_interface, wg_interface, wg_peer_key)
							if err != nil {
								return err
							}
						}
						return nil
					},
				},
				{
					Name:  "webdav",
					Usage: "get network information from webdav server",
					Flags: []cli.Flag{
						&cli.BoolFlag{
							Name:        "allow_insecure",
							Usage:       "set allow insecure connect",
							Value:       false,
							Destination: &allow_insecure,
						},
						&cli.StringFlag{
							Name:        "encryption_key",
							Usage:       "set file encryption key",
							Destination: &encryption_key,
						},
						&cli.StringFlag{
							Name:        "endpoint",
							Usage:       "set webdav server endpoint",
							Required:    true,
							Destination: &endpoint,
						},
						&cli.StringFlag{
							Name:        "username",
							Usage:       "set webdav server username",
							Destination: &username,
						},
						&cli.StringFlag{
							Name:        "password",
							Usage:       "set webdav server password",
							Destination: &password,
						},
						&cli.StringFlag{
							Name:        "filepath",
							Usage:       "set webdav server filepath",
							Required:    true,
							Destination: &filepath,
						},
					},
					Action: func(ctx *cli.Context) error {
						if interval != 0 {
							webdav.ReceiveRequestLoop(endpoint, username, password, allow_insecure, filepath, []byte(encryption_key), remote_interface, wg_interface, wg_peer_key, interval)
						} else {
							err := webdav.ReceiveRequest(endpoint, username, password, allow_insecure, filepath, []byte(encryption_key), remote_interface, wg_interface, wg_peer_key)
							if err != nil {
								return err
							}
						}
						return nil
					},
				},
				{
					Name:  "nconnect",
					Usage: "get network information from nconnect server",
					Flags: []cli.Flag{
						&cli.StringFlag{
							Name:        "id",
							Usage:       "set id",
							Required:    true,
							Destination: &id,
						},
						&cli.BoolFlag{
							Name:        "allow_insecure",
							Usage:       "set allow insecure connect",
							Value:       false,
							Destination: &allow_insecure,
						},
						&cli.StringFlag{
							Name:        "endpoint",
							Usage:       "set nconnect server endpoint",
							Required:    true,
							Destination: &endpoint,
						},
					},
					Action: func(ctx *cli.Context) error {
						if interval != 0 {
							nconnect.ReceiveRequestLoop(id, endpoint, username, password, allow_insecure, remote_interface, wg_interface, wg_peer_key, interval)
						} else {
							err := nconnect.ReceiveRequest(id, endpoint, username, password, allow_insecure, remote_interface, wg_interface, wg_peer_key)
							if err != nil {
								return err
							}
						}
						return nil
					},
				},
			},
		},
	}

	// 打印版本函数
	cli.VersionPrinter = func(cCtx *cli.Context) {
		fmt.Printf("%s", cCtx.App.Version)
	}

	app := &cli.App{
		Usage:    "WireGuard Watchdog",
		Version:  "v1.20",
		Commands: cmds,
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
