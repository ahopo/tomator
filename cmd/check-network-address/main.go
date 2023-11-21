package checknetworkaddress

import (
	"fmt"
	"net"
	"net/http"
	"strings"
	"time"
)

func checkReachability(address Model) (bool, int, error) {
	startTime := time.Now()

	// If the port is not specified, use the default port for the protocol.
	var conn net.Conn
	var err error
	if address.port == 0 {
		switch net.ParseIP(address.hostname).To4() {
		case nil:
			conn, err = net.Dial("tcp", address.hostname+":80")
		default:
			conn, err = net.Dial("tcp", address.hostname+":443")
		}
	} else {
		conn, err = net.Dial("tcp", fmt.Sprintf("%s:%d", address.hostname, address.port))
	}

	if err != nil {
		endTime := time.Now()
		duration := endTime.Sub(startTime)
		return false, 0, fmt.Errorf("address %s is not reachable: %w (%s)", addressInfo(address), err, duration)
	}

	defer conn.Close()

	// Try HTTPS GET first
	resp, err := http.Get("https://" + address.hostname)
	if err == nil {
		defer resp.Body.Close()

		// Check the HTTP status code
		statusCode := resp.StatusCode
		return true, statusCode, nil
	}

	// If HTTPS GET fails, try HTTP GET
	resp, err = http.Get("http://" + address.hostname)
	if err == nil {
		defer resp.Body.Close()

		// Check the HTTP status code
		statusCode := resp.StatusCode
		return true, statusCode, nil
	}

	// If both HTTPS and HTTP GET fail, the address is not reachable
	return false, 0, fmt.Errorf("address %s is not reachable: %t", addressInfo(address), err)
}

func Init() {
	addresses := []Model{
		{hostname: "www.google.com"},
		{hostname: "127.0.0.1"},
		{hostname: "api.github.com"},
		{hostname: "120.28.178.247"},
		{hostname: "loadbalancer.example.com", port: 443},
	}

	for _, address := range addresses {
		reachable, statusCode, err := checkReachability(address)
		if err != nil {
			fmt.Println(err)
		} else {
			if reachable {
				fmt.Printf("Address %s is reachable with status code %d\n", address.hostname, statusCode)
			} else {
				fmt.Printf("Address %s is not reachable\n", address.hostname)
			}
		}
	}
}
func addressInfo(a Model) string {
	if a.port == 0 {
		return a.hostname
	}
	return strings.Join([]string{a.hostname, fmt.Sprint(a.port)}, ":")
}
