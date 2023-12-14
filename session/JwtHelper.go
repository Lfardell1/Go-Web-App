package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

func generateToken() (string, error) {
	// Generate a random byte slice with 32 bytes
	tokenBytes := make([]byte, 32)
	_, err := rand.Read(tokenBytes)
	if err != nil {
		return "", err
	}

	// Encode the byte slice to base64 string
	token := base64.StdEncoding.EncodeToString(tokenBytes)

	return token, nil
}

func main() {
	token, err := generateToken()
	if err != nil {
		fmt.Println("Error generating token:", err)
		return
	}

	fmt.Println("Generated token:", token)
}
