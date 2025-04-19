package main

import (
	"github.com/joho/godotenv"
	"github.com/ucrnsemin/EnesGoland/config"
	"log"
)

func main() {
	// .env dosyasını yükle
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err.Error())
		return
	}
	config.InitDB()

}
