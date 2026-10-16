package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"moka-agent/internal/agent"
	"moka-agent/internal/memory"
	"moka-agent/internal/provider"

	"github.com/joho/godotenv"
	"github.com/stanj98/moka-agent/internal/agent"
	"github.com/stanj98/moka-agent/internal/memory"
	"github.com/stanj98/moka-agent/internal/provider"
)

func main() {
	_ = godotenv.Load()

	store, err := memory.NewStore()
	if err != nil {
		log.Fatalf("Failed to initialize local SQLite memory: %v", err)
	}

	client := provider.NewClient()
	ag := agent.NewAgent(store, client)
	fmt.Println("Go Moka Agent")
	fmt.Println("Type your message below. Type 'exit' to quit.\n")

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("moka> ")
		if !scanner.Scan() {
			break
		}
		text := strings.TrimSpace(scanner.Text())
		if text == "exit" || text == "quit" {
			break
		}
		if text == "" {
			continue
		}

		reply, err := ag.Run(text)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			continue
		}

		fmt.Printf("\n%s\n\n", &reply)
	}
}
