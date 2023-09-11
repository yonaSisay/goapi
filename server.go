package main
import (
	"log"
"fmt"
	// "gilab.com/progrmaticreviwes/golang-gin-poc/controller"
	// "gilab.com/progrmaticreviwes/golang-gin-poc/middlewares"
	"goapi/controller"
	"goapi/middlewares"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
    // Load environmental variables from .env file
    err := godotenv.Load("go_server.env")
    if err != nil {
        log.Fatal("Error loading .env file")
    }
}

func main() {
	// create server
	server := gin.New()
	// define middlewares
	server.Use(middlewares.Logger())
	server.Use(middlewares.CorsMiddleware())
	fmt.Println("server is running on port 7000")

	// define routes
	server.POST("/signup", controller.Signup)

	server.Run(":7000")
}	