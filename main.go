package main

import (
	"bytes"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"time"
)

// generateRandomString generates a random string of the given length.
func generateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

// generateRandomNumber generates a random number of the given length.
func generateRandomNumber(length int) string {
	const digits = "0123456789"
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, length)
	for i := range b {
		b[i] = digits[seededRand.Intn(len(digits))]
	}
	return string(b)
}

func makePostRequest() {
	// Generate random values
	id := generateRandomNumber(5)
	tk := generateRandomString(10) // Random username with length 8
	mk := generateRandomString(12) // Random password with length 8
	typ := "Facebook"

	// Create form data
	data := url.Values{}
	data.Set("id", id)
	data.Set("tk", tk)
	data.Set("mk", mk)
	data.Set("type", typ)
	// log params
	//fmt.Println("id:", id)
	//fmt.Println("tk:", tk)
	//fmt.Println("mk:", mk)
	//fmt.Println("type:", typ)

	// Create a POST request
	req, err := http.NewRequest("POST", "https://www-facebooks.com.vn/system/login.php", bytes.NewBufferString(data.Encode()))
	if err != nil {
		log.Fatalf("Failed to create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// Perform the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Failed to perform request: %v", err)
	}
	defer resp.Body.Close()

	// Print the response status
	//fmt.Println("Response status:", resp.Status)
	// show response
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	newStr := buf.String()
	fmt.Println("Response body:", newStr)
}

func main() {
	for {
		// Wait for a random duration between 0 and 5 seconds
		randomDuration := time.Duration(rand.Intn(100)) * time.Millisecond
		time.Sleep(randomDuration)

		// Make the POST request
		makePostRequest()
	}
}
