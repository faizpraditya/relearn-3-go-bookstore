package main

import (
	"fmt"
	"go-bookstore/pkg/config" // Replace with the actual path to your config package
)

func main() {
	config.Connect()
	defer config.Disconnect() // Ensure that the connection is closed when main function exits

	// Check if the database connection is successful
	if err := config.GetDB().DB().Ping(); err != nil {
		fmt.Println("Failed to connect to the database:", err)
	} else {
		fmt.Println("Successfully connected to the database.")
	}
}
