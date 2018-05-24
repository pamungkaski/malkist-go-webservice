package main

import (
	"github.com/joho/godotenv"
	"github.com/pamungkaski/malkist-go-webservice/service"
	"log"
	"os"
)

// Main function will create the service using environment variable from .env file.
func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file :", err.Error())
	}
	app := service.NewService(os.Getenv("KEY"))

	app.RegisterService()

	app.StartService(":" + os.Getenv("PORT"))
}
