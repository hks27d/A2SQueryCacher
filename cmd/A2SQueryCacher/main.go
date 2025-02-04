package main

import (
	"A2SQueryCacher/internal/a2sqc_cache"
	"A2SQueryCacher/internal/a2sqc_cli"
	"A2SQueryCacher/internal/a2sqc_json"
	"A2SQueryCacher/internal/a2sqc_listener"
	"fmt"
	"log"
	"time"

	"github.com/natefinch/lumberjack"
)

func commandlineArgsOrJSON() (string, string, int, int, int, int) {
	cli, bindIP, gameServerIP, bindPort, gameServerPort, cacheTTL, threads := a2sqc_cli.CommandlineArgs()
	if !cli {
		fmt.Println("No CLI arguments found. Trying with JSON file...")
	} else {
		fmt.Println("CLI arguments successfully loaded!")
		return bindIP, gameServerIP, bindPort, gameServerPort, cacheTTL, threads
	}

	json, bindIP, gameServerIP, bindPort, gameServerPort, cacheTTL, threads := a2sqc_json.JSONSettings()
	if !json {
		fmt.Println("Failed to load the settings from the JSON file.")
		log.Fatalln("Failed to load the settings from CLI args and JSON.")
	}

	fmt.Println("JSON settings successfully loaded!")
	return bindIP, gameServerIP, bindPort, gameServerPort, cacheTTL, threads
}

func init() {
	log.SetOutput(&lumberjack.Logger{
		Filename:   "log.log",
		MaxSize:    1, // MB
		MaxBackups: 3,
		MaxAge:     90, // days
		Compress:   false,
	})
}

func main() {
	bindIP, gameServerIP, bindPort, gameServerPort, cacheTTL, threads := commandlineArgsOrJSON()
	proxyAddr := fmt.Sprintf("%s:%d", bindIP, bindPort)
	serverAddr := fmt.Sprintf("%s:%d", gameServerIP, gameServerPort)

	fmt.Println("")
	fmt.Println("Cache address:", proxyAddr)
	fmt.Println("Game server:", serverAddr)
	fmt.Println("CacheTTL:", cacheTTL, "seconds")
	fmt.Println("Threads:", threads)

	cache := a2sqc_cache.NewCache(time.Duration(cacheTTL) * time.Second)

	err := a2sqc_listener.StartListener(bindIP, bindPort, gameServerIP, gameServerPort, threads, cache)
	if err != nil {
		fmt.Printf("Error starting listener: %v\n", err)
		time.Sleep(time.Second)
		log.Fatalf("Error starting listener: %v\n", err)
	}
}
