package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/santoshmohan35/tickets-data-api/internal/configuration"
	"github.com/santoshmohan35/tickets-data-api/internal/router"
	"os"
)

func main() {
	//load env variables
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	myEnvVar := os.Getenv("MY_ENV")
	fmt.Println("Env variables loaded MY_ENV:", myEnvVar)

	// load configurations
	config, err := configuration.LoadConfig()
	if err != nil {
		fmt.Println("Error loading config:", err)
	}

	// setup router
	r := router.SetupRouter(config)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
