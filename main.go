package main

import (
	"log"
	"os"

	"expense-tracker/cli"
	"expense-tracker/storage"
)

func main() {
	// Initialize the storage system (JSON file-based in this case)
	store, err := storage.NewJSONStorage("expenses.json")
	if err != nil {
		log.Fatalf("Failed to initialize storage: %v", err)
	}

	// Parse and handle CLI commands
	if err := cli.Run(os.Args, store); err != nil {
		log.Fatalf("Command failed: %v", err)
	}
}
