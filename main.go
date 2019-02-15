package main

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"restgo_framework/routers"
)

func main() {
	//gin.SetMode(gin.ReleaseMode)
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	r := routers.SetupRouter()

	_ = r.Run(os.Getenv("APP_HOST") + ":" + os.Getenv("APP_PORT"))
}
