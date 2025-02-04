package a2sqc_handler

import (
	"A2SQueryCacher/internal/a2sqc_cache"
	"A2SQueryCacher/internal/a2sqc_types"
	"fmt"
	"net"
	"time"
)

func HandleRequests(requests chan a2sqc_types.Request, proxyConn net.PacketConn, serverAddr string, cache *a2sqc_cache.Cache) {
	for req := range requests {
		handleRequest(req.Data, req.Addr, proxyConn, serverAddr, cache)
	}
}

func handleRequest(data []byte, addr net.Addr, proxyConn net.PacketConn, serverAddr string, cache *a2sqc_cache.Cache) {
	queryKey := string(data)

	if cachedResponse, found := cache.Get(queryKey); found {
		proxyConn.WriteTo(cachedResponse, addr)
		fmt.Println("Cache hit: Responding from cache")
		return
	}

	serverConn, err := net.Dial("udp", serverAddr)
	if err != nil {
		fmt.Println("Failed to connect to server:", err)
		return
	}
	defer serverConn.Close()

	_, err = serverConn.Write(data)
	if err != nil {
		fmt.Println("Failed to send to server:", err)
		return
	}

	serverConn.SetReadDeadline(time.Now().Add(2 * time.Second))
	respBuffer := make([]byte, 1024)
	respLen, err := serverConn.Read(respBuffer)

	if err != nil {
		fmt.Println("Failed to receive response:", err)
		return
	}

	cache.Set(queryKey, respBuffer[:respLen])

	proxyConn.WriteTo(respBuffer[:respLen], addr)
	fmt.Println("Forwarded request and cached response")
}
