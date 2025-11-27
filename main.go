package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"
	"wgwd/internal/decrypt"
	"wgwd/internal/receive/file"
	"wgwd/internal/receive/s3"
	"wgwd/internal/receive/webdav"

	"github.com/urfave/cli/v3"
)

func main() {
	// get mode
	var allowInsecure bool
	var encryptionKey string
	var interval time.Duration
	var endpoint string
	var username string
	var password string

	// get mode file
	var filepath string

	// get mode s3
	var regin string
	var stsToken string
	var pathStyle bool
	var bucket string
	var objectPath string

	// wireguard
	var remoteInterface string
	var wgInterface string
	var wgPeerKey string

	cmds := []*cli.Command{
		{
			Name:    "get",
			Aliases: []string{"g"},
			Usage:   "get wireguard endpoint from network information",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:        "remote_interface",
					Aliases:     []string{"r"},
					Usage:       "set remote interface",
					Required:    true,
					Destination: &remoteInterface,
				},
				&cli.StringFlag{
					Name:        "wg_interface",
					Aliases:     []string{"wi"},
					Usage:       "set wireguard interface",
					Required:    true,
					Destination: &wgInterface,
				},
				&cli.StringFlag{
					Name:        "wg_peer_key",
					Aliases:     []string{"wk"},
					Usage:       "set wireguard peer key",
					Destination: &wgPeerKey,
				},
				&cli.DurationFlag{
					Name:        "interval",
					Aliases:     []string{"i"},
					Usage:       "set send interval",
					Destination: &interval,
				},
			},

			Commands: []*cli.Command{
				{
					Name:  "file",
					Usage: "get network information from filesystem",
					Flags: []cli.Flag{
						&cli.StringFlag{
							Name:        "filepath",
							Aliases:     []string{"f"},
							Usage:       "set file path",
							Required:    true,
							Destination: &filepath,
						},
						&cli.StringFlag{
							Name:        "encryption_key",
							Aliases:     []string{"e"},
							Usage:       "set file encryption key",
							Destination: &encryptionKey,
						},
					},
					Action: func(ctx context.Context, cmd *cli.Command) (err error) {
						if interval != 0 {
							file.ReceiveRequestLoop(filepath, []byte(encryptionKey), remoteInterface, wgInterface, wgPeerKey, interval)
						} else {
							err := file.ReceiveRequest(filepath, []byte(encryptionKey), remoteInterface, wgInterface, wgPeerKey)
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
							Destination: &allowInsecure,
						},
						&cli.StringFlag{
							Name:        "encryption_key",
							Aliases:     []string{"e"},
							Usage:       "set file encryption key",
							Destination: &encryptionKey,
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
							Destination: &stsToken,
						},
						&cli.BoolFlag{
							Name:        "path_style",
							Usage:       "set s3 server path style, false: virtual host, true: path",
							Value:       false,
							Destination: &pathStyle,
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
							Destination: &objectPath,
						},
					},
					Action: func(ctx context.Context, cmd *cli.Command) (err error) {
						if interval != 0 {
							s3.ReceiveRequestLoop(endpoint, regin, username, password, stsToken, pathStyle, allowInsecure, bucket, objectPath, []byte(encryptionKey), remoteInterface, wgInterface, wgPeerKey, interval)
						} else {
							err = s3.ReceiveRequest(endpoint, regin, username, password, stsToken, pathStyle, allowInsecure, bucket, objectPath, []byte(encryptionKey), remoteInterface, wgInterface, wgPeerKey)
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
							Destination: &allowInsecure,
						},
						&cli.StringFlag{
							Name:        "encryption_key",
							Aliases:     []string{"e"},
							Usage:       "set file encryption key",
							Destination: &encryptionKey,
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
							Aliases:     []string{"f"},
							Usage:       "set webdav server filepath",
							Required:    true,
							Destination: &filepath,
						},
					},
					Action: func(ctx context.Context, cmd *cli.Command) (err error) {
						if interval != 0 {
							webdav.ReceiveRequestLoop(endpoint, username, password, allowInsecure, filepath, []byte(encryptionKey), remoteInterface, wgInterface, wgPeerKey, interval)
						} else {
							err = webdav.ReceiveRequest(endpoint, username, password, allowInsecure, filepath, []byte(encryptionKey), remoteInterface, wgInterface, wgPeerKey)
							if err != nil {
								return err
							}
						}
						return nil
					},
				},
			},
		},
		{
			Name:    "decrypt",
			Aliases: []string{"d"},
			Usage:   "decrypt a file",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:        "filepath",
					Aliases:     []string{"f"},
					Usage:       "set file path",
					Required:    true,
					Destination: &filepath,
				},
				&cli.StringFlag{
					Name:        "encryption_key",
					Aliases:     []string{"e"},
					Usage:       "set file encryption key",
					Required:    true,
					Destination: &encryptionKey,
				},
			},
			Action: func(ctx context.Context, cmd *cli.Command) (err error) {
				plaintext, err := decrypt.FromFile(filepath, []byte(encryptionKey))
				if err != nil {
					return err
				}
				fmt.Println(string(plaintext))
				return nil
			},
		},
	}

	// 打印版本函数
	cli.VersionPrinter = func(cmd *cli.Command) {
		fmt.Printf("%s\n", cmd.Root().Version)
	}

	cmd := &cli.Command{
		Usage:    "WireGuard Watchdog",
		Version:  "v1.40",
		Commands: cmds,
	}

	err := cmd.Run(context.Background(), os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
