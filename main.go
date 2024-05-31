package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"optician-rest-api/routes"
)

func main() {
	routes.ListenAndServe()
	initENV()
}

func initENV() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file!")
		log.Fatal(err)
	}
}
