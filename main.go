package main

import (
	"GoVenice/models"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"io"
	"log"
	"net/http"
	"os"
)

const BASE_URL = "https://api.venice.ai/api/v1"
const COMPLETIONS_URL = "https://api.venice.ai/api/v1/chat/completions"

type Messages struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type VenicePayload struct {
	Model    string     `json:"model"`
	Messages []Messages `json:"messages"`
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	prompt := Messages{
		Role:    "user",
		Content: "Why did the South Sea company collapse so late?",
	}

	payload := VenicePayload{
		Model:    models.DeepseekR1_Llama_70b.ID,
		Messages: []Messages{prompt}, // Note the slice initialization
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		panic(err)
	}

	result := getResponse(jsonPayload)

	log.Println("This is the response: \n\n\n", result)
}

func getResponse(payload []byte) string {
	req, _ := http.NewRequest("POST", COMPLETIONS_URL, bytes.NewBuffer(payload))

	bearer := os.Getenv("BEARER_TOKEN")

	req.Header.Add("Authorization", "Bearer "+bearer)
	req.Header.Set("Content-Type", "application/json")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

	return string(body)
}
