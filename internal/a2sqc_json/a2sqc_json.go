package a2sqc_json

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
)

const a2sqcConfigFile string = "config.json"

type configFile struct {
	BindIP         string `json:"bindIP"`
	GameServerIP   string `json:"gameServerIP"`
	BindPort       int    `json:"bindPort"`
	GameServerPort int    `json:"gameServerPort"`
	CacheTTL       int    `json:"cacheTTL"`
	Threads        int    `json:"threads"`
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if err != nil {
		return false
	}
	return !info.IsDir()
}

func JSONSettings() (bool, string, string, int, int, int, int) {
	if !fileExists(a2sqcConfigFile) {
		fmt.Println("No JSON config file found. Creating " + a2sqcConfigFile + ".")
		log.Println("No JSON config file found. Creating " + a2sqcConfigFile + ".")
		file, err := os.Create(a2sqcConfigFile)
		if err != nil {
			fmt.Println("Error creating the file " + a2sqcConfigFile + ": " + err.Error())
			log.Println("Error creating the file " + a2sqcConfigFile + ": " + err.Error())
			return false, "", "", 0, 0, 0, 0
		}
		defer file.Close()
		fmt.Println("Created the file " + a2sqcConfigFile + ".")
		log.Println("Created the file " + a2sqcConfigFile + ".")

		config := configFile{
			BindIP:         "0.0.0.0",
			GameServerIP:   "127.0.0.1",
			BindPort:       9110,
			GameServerPort: 27015,
			CacheTTL:       10,
			Threads:        4,
		}

		encoder := json.NewEncoder(file)
		encoder.SetIndent("", "  ")

		err = encoder.Encode(config)
		if err != nil {
			fmt.Println("Error encoding " + a2sqcConfigFile + ": " + err.Error())
			log.Println("Error encoding " + a2sqcConfigFile + ": " + err.Error())
			return false, "", "", 0, 0, 0, 0
		}
		fmt.Println("Encoded default values in " + a2sqcConfigFile + ".")
		log.Println("Encoded default values in " + a2sqcConfigFile + ".")
		fmt.Println("Restart the program with proper settings in " + a2sqcConfigFile + ".")
		time.Sleep(time.Second)
		log.Fatalln("Restart the program with proper settings in " + a2sqcConfigFile + ".")
	}

	file, err := os.Open(a2sqcConfigFile)
	if err != nil {
		fmt.Println("Error opening the file " + a2sqcConfigFile + ": " + err.Error())
		log.Fatalln("Error opening the file " + a2sqcConfigFile + ": " + err.Error())
	}
	defer file.Close()

	jsonSettings := configFile{}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&jsonSettings)
	if err != nil {
		log.Fatalln("Error decoding " + a2sqcConfigFile + ": " + err.Error())
	}

	if jsonSettings.BindIP == "0.0.0.0" || jsonSettings.GameServerIP == "127.0.0.1" {
		fmt.Println("You must run the software with non-default config values.")
		time.Sleep(time.Second)
		log.Fatalln("You must run the software with non-default config values.")
	}

	if jsonSettings.CacheTTL < 1 || jsonSettings.CacheTTL > 30 {
		fmt.Println("cacheTTL must be between 1 and 30")
		log.Fatalln("cacheTTL must be between 1 and 30")
	}

	return true, jsonSettings.BindIP, jsonSettings.GameServerIP, jsonSettings.BindPort, jsonSettings.GameServerPort, jsonSettings.CacheTTL, jsonSettings.Threads
}
