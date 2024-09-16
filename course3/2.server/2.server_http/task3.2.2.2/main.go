package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load() // 👈 load .env file
	if err != nil {
		log.Fatal(err)
	}

	port := os.Getenv("PORT")

	fmt.Println("server zapushen na portu", port)
	http.ListenAndServe(":"+port, nil)
}
