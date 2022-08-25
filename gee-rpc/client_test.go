package main

import (
	"gee-rpc/client"
	"gee-rpc/server"
	"net"
	"testing"
	"time"
)

func TestClient_dialTimeout(t *testing.T) {
	t.Parallel()
	l, _ := net.Listen("tcp", ":0")

	f := func(conn net.Conn, opt *server.Option) (client *client.Client, err error) {
		_ = conn.Close()
		time.Sleep(time.Second * 2)
		return nil, nil
	}
	t.Run("timeout", func(t *testing.T) {
		_, err := client.DialTimeout(f, "tcp", l.Addr().String(), &server.Option{ConnectTimeout: time.Second})
		if err != nil {

		}
		//_assert(err != nil && strings.Contains(err.Error(), "connect timeout"), "expect a timeout error")
	})
	t.Run("0", func(t *testing.T) {
		_, err := client.DialTimeout(f, "tcp", l.Addr().String(), &server.Option{ConnectTimeout: 0})
		if err != nil {

		}
		//_assert(err == nil, "0 means no limit")
	})
}
