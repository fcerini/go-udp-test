package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

var gloClients map[string]time.Time

func main() {
	go udpListenLoop()
	for {
		time.Sleep(1 * time.Second)
		fmt.Println("Press 'Enter' to send...")
		bufio.NewReader(os.Stdin).ReadBytes('\n')

		send()
	}
}

// Listen for and handle UDP packets.
func udpListenLoop() {

	gloClients = make(map[string]time.Time)

	host := "0.0.0.0"
	port := 64749

	// Setup our UDP listener
	udpconn, _ := net.ListenUDP("udp",
		&net.UDPAddr{
			IP:   net.ParseIP(host),
			Port: port})

	fmt.Println("UDP SERVER Listening in ", udpconn.LocalAddr())

	buf := make([]byte, 1024)
	for {
		_, remote, err := udpconn.ReadFrom(buf)
		if err != nil {
			fmt.Println("ERR udpListenLoop:", err)
			return

		}

		udpaddr, ok := remote.(*net.UDPAddr)
		if !ok {
			fmt.Println("No UDPAddr in read packet. (Windows?)")
			return
		}

		_, ok = gloClients[udpaddr.String()]
		if !ok {
			fmt.Println("Nuevo Cliente ", udpaddr.String())
			return
		}

		gloClients[udpaddr.String()] = time.Now()
	}
}

func send() {

	var buf []byte
	for i := 0; i < 100; i++ {

		buf = []byte{byte(i), 9, 9, 9, 9, 9, 9, 9, 9, 9,
			9, 9, 9, 9, 9, 9, 9, 9, 9, 9,
			9, 9, 9, 9, 9, 9, 9, 9, 9, 9,
			9, 9, 9, 9, 9, 9, 9, 9, 9, 9,
			9, 9, 9, 9, 9, 9, 9, 9, 9, 9,
			9, 9, 9, 9, 9, 9, 9, 9, 9, 9,
		}
		time.Sleep(25 * time.Millisecond)

		for k := range gloClients {

			s, err := net.ResolveUDPAddr("udp4", k)
			if err != nil {
				log.Fatal("ERR Broadcast ResolveUDPAddr", err)
				return
			}
			auxConn, err := net.DialUDP("udp4", nil, s)
			if err != nil {
				log.Fatal("ERR Broadcast DialUDP", err)
				return
			}

			defer auxConn.Close()

			_, err = auxConn.Write(buf)
			if err != nil {
				fmt.Println("ERR APP.Broadcasts..Conn.Write: ", err)
			}

		}
	}
}
