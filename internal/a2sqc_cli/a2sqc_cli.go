package a2sqc_cli

import (
	"flag"
	"fmt"
	"log"
)

func CommandlineArgs() (bool, string, string, int, int, int, int) {
	bindIP := flag.String("bindip", "0.0.0.0", "Local IP address to bind")
	gameServerIP := flag.String("gameserverip", "127.0.0.1", "Game server IP address")
	bindPort := flag.Int("bindport", 9110, "Local port to bind")
	gameServerPort := flag.Int("gameserverport", 27015, "Game server port")
	cacheTTL := flag.Int("cacheTTL", 10, "Cache TTL in seconds")
	threads := flag.Int("threads", 4, "Number of worker threads")

	flag.Parse()

	if *bindIP == "0.0.0.0" || *gameServerIP == "127.0.0.1" {
		log.Println("If you intended to use CLI args, you must run the software with non-default values.")
		return false, "", "", 0, 0, 0, 0
	}

	if *cacheTTL < 1 || *cacheTTL > 30 {
		fmt.Println("cacheTTL must be between 1 and 30")
		log.Fatalln("cacheTTL must be between 1 and 30")
	}

	return true, *bindIP, *gameServerIP, *bindPort, *gameServerPort, *cacheTTL, *threads
}
