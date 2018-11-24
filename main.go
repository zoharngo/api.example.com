package main

import (
	"flag"
	"log"
	"os"

	"api.example.com/src/app"
	DB "api.example.com/src/db"
	"github.com/joho/godotenv"
)

var port string

func init() {
	flag.StringVar(&port, "port", "8000", "Assign server port")
	flag.Parse()

	if err := godotenv.Load("config.ini"); err != nil {
		panic(err)
	}

	envPort := os.Getenv("PORT")
	if len(envPort) > 0 {
		log.Println("Overwrite port:" + port + " with config.ini port:" + envPort)
		port = envPort
	}
}

func main() {
	db, err := DB.Connect()
	if err != nil {
		panic(err)
	}
	s := app.NewServer()

	s.Init(port, db)
	s.Start()
}
