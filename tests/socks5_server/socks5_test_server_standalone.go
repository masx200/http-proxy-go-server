package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"gitee.com/masx200/go-socks5"
)

func main() {
	// 创建SOCKS5配置
	conf := &socks5.Config{
		AuthMethods: []socks5.Authenticator{
			&socks5.UserPassAuthenticator{
				Credentials: socks5.StaticCredentials{
					"g7envpwz14b0u55": "juvytdsdzc225pq",
				},
			},
		},
		Rules:  socks5.PermitAll(),
		Logger: log.New(os.Stdout, "[SOCKS5] ", log.LstdFlags),
	}

	// 创建SOCKS5服务器
	server, err := socks5.New(conf)
	if err != nil {
		log.Fatalf("Failed to create SOCKS5 server: %v", err)
	}

	// 监听端口
	listener, err := net.Listen("tcp", ":44444")
	if err != nil {
		log.Fatalf("Failed to listen on port 44444: %v", err)
	}
	defer listener.Close()

	fmt.Println("SOCKS5 server started on :44444")
	fmt.Println("Username: g7envpwz14b0u55")
	fmt.Println("Password: juvytdsdzc225pq")
	fmt.Println("Press Ctrl+C to stop the server")

	// 接受连接
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Failed to accept connection: %v", err)
			continue
		}

		log.Printf("New connection from: %s", conn.RemoteAddr())

		go func() {
			defer conn.Close()
			defer log.Printf("Connection closed: %s", conn.RemoteAddr())

			err := server.ServeConn(conn)
			if err != nil {
				log.Printf("SOCKS5 connection error: %v", err)
			} else {
				log.Printf("SOCKS5 connection handled successfully: %s", conn.RemoteAddr())
			}
		}()
	}
}
