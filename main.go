package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

type Config struct {
	Babyoil    int            `json:"babyoil"`
	OiledUsers map[string]int `json:"oiled_users"`
}

const configFilePath = ".config/pdiddy/config.json"

func main() {
	if len(os.Args) < 2 {
		showHelp()
		return
	}

	command := os.Args[1]
	config := loadConfig()

	switch command {
	case "guestlist":
		listGuestlist()
	case "party":
		listParty()
	case "babyoil":
		if len(os.Args) < 3 {
			fmt.Println("Usage: pdiddy babyoil <user>")
			return
		}
		babyoilUser(config, os.Args[2])
	case "touch":
		if len(os.Args) < 3 {
			fmt.Println("Usage: pdiddy touch <user>")
			return
		}
		touchUser(config, os.Args[2])
	case "babyoilhouse":
		showBabyoil(config)
	case "buy":
		if len(os.Args) < 3 {
			fmt.Println("Usage: pdiddy buy <amount>")
			return
		}
		amount, err := strconv.Atoi(os.Args[2])
		if err != nil || amount <= 0 {
			fmt.Println("Invalid amount. Please specify a positive integer.")
			return
		}
		buyBabyoil(config, amount)
	default:
		fmt.Println("Unknown command.")
	}
}

func showHelp() {
	fmt.Println("Usage: pdiddy <command> [arguments]")
	fmt.Println("Commands:")
	fmt.Println("  guestlist         List all system users.")
	fmt.Println("  party             List all currently logged-in users.")
	fmt.Println("  babyoil <user>    Apply babyoil to a user. Reduces the house's babyoil by 1.")
	fmt.Println("  touch <user>      Reset a user's babyoil usage.")
	fmt.Println("  babyoilhouse      Show the remaining amount of babyoil in the house.")
	fmt.Println("  buy <amount>      Buy more babyoil. Specify the amount to add.")
	fmt.Println("  help              Show this help message.")
}

func getConfigPath() string {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("Failed to get home directory: %v", err)
	}
	return filepath.Join(home, configFilePath)
}

func loadConfig() Config {
	path := getConfigPath()
	config := Config{
		Babyoil:    0,
		OiledUsers: make(map[string]int),
	}

	file, err := os.Open(path)
	if err != nil {
		if os.IsNotExist(err) {
			saveConfig(config)
			return config
		}
		log.Fatalf("Failed to open config file: %v", err)
	}
	defer file.Close()

	if err := json.NewDecoder(file).Decode(&config); err != nil {
		log.Fatalf("Failed to parse config file: %v", err)
	}

	return config
}

func saveConfig(config Config) {
	path := getConfigPath()
	dir := filepath.Dir(path)

	if err := os.MkdirAll(dir, 0755); err != nil {
		log.Fatalf("Failed to create config directory: %v", err)
	}

	file, err := os.Create(path)
	if err != nil {
		log.Fatalf("Failed to create config file: %v", err)
	}
	defer file.Close()

	if err := json.NewEncoder(file).Encode(config); err != nil {
		log.Fatalf("Failed to write config file: %v", err)
	}
}

func listParty() {
	cmd := exec.Command("who")
	output, err := cmd.Output()
	if err != nil {
		log.Fatalf("Failed to list party: %v", err)
	}

	fmt.Println("The current party has:")
	fmt.Println(strings.TrimSpace(string(output)))
}

func listGuestlist() {
	users, err := os.ReadDir("/home")
	if err != nil {
		log.Fatalf("Failed to list guestlist: %v", err)
	}

	fmt.Println("All system users:")
	for _, u := range users {
		fmt.Println(u.Name())
	}
}

func babyoilUser(config Config, username string) {
	if config.Babyoil <= 0 {
		fmt.Println("No babyoil left in the house! Please buy more babyoil.")
		return
	}

	config.OiledUsers[username]++
	config.Babyoil--
	saveConfig(config)
	fmt.Printf("%s has been oiled. Remaining babyoil: %d\n", username, config.Babyoil)
}

func touchUser(config Config, username string) {
	config.OiledUsers[username] = 0
	saveConfig(config)
	fmt.Printf("%s has been touched. Babyoil usage reset.\n", username)
}

func showBabyoil(config Config) {
	fmt.Printf("Babyoil left in the house: %d\n", config.Babyoil)
}

func buyBabyoil(config Config, amount int) {
	config.Babyoil += amount
	saveConfig(config)
	fmt.Printf("Bought %d babyoil. Total babyoil in the house: %d\n", amount, config.Babyoil)
}
