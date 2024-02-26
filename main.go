package main

import (
	"flag"
	"fmt"

	"github.com/masx200/http-proxy-go-server/auth"
	"github.com/masx200/http-proxy-go-server/simple"
	"github.com/masx200/http-proxy-go-server/tls"
	tls_auth "github.com/masx200/http-proxy-go-server/tls+auth"
)

var (
	hostname    = flag.String("hostname", "0.0.0.0", "an String value for hostname")
	port        = flag.Int("port", 8080, "TCP port to listen on")
	server_cert = flag.String("server_cert", "", "tls server cert")
	server_key  = flag.String("server_key", "", "tls server key")
	username    = flag.String("username", "", "username")
	password    = flag.String("password", "", "password")
)

func main() {
	flag.Parse()
	//parse cmd flags
	fmt.Println(
		"hostname:", *hostname)
	fmt.Println(
		"port:", *port)
	fmt.Println(
		"server_cert:", *server_cert)
	fmt.Println(
		"server_key:", *server_key)
	fmt.Println(
		"username:", *username)
	fmt.Println(
		"password:", *password)

	if len(*username) > 0 && len(*password) > 0 && len(*server_cert) > 0 && len(*server_key) > 0 {
		tls_auth.Tls_auth(*server_cert, *server_key, *hostname, *port, *username, *password)
		return
	}
	if len(*username) > 0 && len(*password) > 0 {
		auth.Auth(*hostname, *port, *username, *password)
		return
	}
	if len(*server_cert) > 0 && len(*server_key) > 0 {
		tls.Tls(*server_cert, *server_key, *hostname, *port)
		return
	}
	if len(*username) == 0 && len(*password) == 0 && len(*server_cert) == 0 && len(*server_key) == 0 {
		simple.Simple(*hostname, *port)
		return
	}

}
