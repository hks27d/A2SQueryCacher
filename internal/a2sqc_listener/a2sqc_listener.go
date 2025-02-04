package a2sqc_listener

import (
	"fmt"
	"net"

	"A2SQueryCacher/internal/a2sqc_cache"
	"A2SQueryCacher/internal/a2sqc_handler"
	"A2SQueryCacher/internal/a2sqc_types"
)

func StartListener(bindAddr string, bindPort int, gameServerIP string, gameServerPort int, threads int, cache *a2sqc_cache.Cache) error {
	proxyAddr := fmt.Sprintf("%s:%d", bindAddr, bindPort)
	serverAddr := fmt.Sprintf("%s:%d", gameServerIP, gameServerPort)

	proxyConn, err := net.ListenPacket("udp", proxyAddr)
	if err != nil {
		return fmt.Errorf("error starting UDP listener: %w", err)
	}
	defer proxyConn.Close()

	fmt.Println("")
	fmt.Println("UDP listener started on address", proxyAddr)
	fmt.Println("Forwarding A2S cached responses to UDP address", serverAddr)

	requests := make(chan a2sqc_types.Request, threads)

	for i := 0; i < threads; i++ {
		go a2sqc_handler.HandleRequests(requests, proxyConn, serverAddr, cache)
	}

	buffer := make([]byte, 1024)

	for {
		n, clientAddr, err := proxyConn.ReadFrom(buffer)
		if err != nil {
			fmt.Println("Read error:", err)
			continue
		}

		requests <- a2sqc_types.Request{Data: buffer[:n], Addr: clientAddr}
	}
}
