package main

import (
	"log"

	"github.com/fvnis/golang-crud/src/controller/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
  err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file")
  }

  router := gin.Default()

  routes.InitiRoutes(&router.RouterGroup)

  if err := router.Run(":8080"); err != nil {
    log.Fatal(err)
  }
}
