package main

import (
	"fmt"
	"log"
	"net"
)

// Get preferred outbound ip of this machine
func GetOutboundIP() (net.IP, string) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP, localAddr.IP.String()
}

func main() {
	ip, ipString := GetOutboundIP()
	fmt.Println(ip)
	fmt.Println(ipString)
}
