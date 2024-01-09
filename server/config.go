package main

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

// AppConfig Global Configuration
type AppConfig struct {
	Puerto           int
	CantidadPaquetes int
	Delay            int
}

// Load config
func (c *AppConfig) Load() {

	jsonFile, err := os.Open("config.json")
	if err != nil {
		log.Fatal(err)
	}

	defer jsonFile.Close()
	byteValue, _ := io.ReadAll(jsonFile)
	err2 := json.Unmarshal([]byte(byteValue), &c)
	if err2 != nil {
		log.Fatal(err2)
	}

}
