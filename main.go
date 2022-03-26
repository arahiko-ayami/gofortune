package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"time"
)

type Quote struct {
	Text   string `json:"text"`
	Author string `json:"author"`
}

func main() {
	quotes := readQuoteFromFile()
	quote := randomQuote(quotes)

	if quote.Author == "" {
		quote.Author = "Unknown"
	}

	fmt.Printf("\"%s\"\n", quote.Text)
	fmt.Printf("\t-- %s\n", quote.Author)
}

func readQuoteFromFile() []Quote {
	jsonFile, err := os.ReadFile("quotes.json")
	if err != nil {
		fmt.Println(err)
	}

	var quotes []Quote
	_ = json.Unmarshal(jsonFile, &quotes)
	return quotes
}

func randomQuote(quotes []Quote) Quote {
	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(len(quotes))
	return quotes[index]
}
