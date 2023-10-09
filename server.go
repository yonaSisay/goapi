package main
import (
	"goapi/utilService"
	"log"
"fmt"
	"os"
	"goapi/controller"
	"goapi/middlewares"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	// "github.com/hasura/go-graphql-client"
)

func init() {
    // Load environmental variables from .env file
    err := godotenv.Load("go_server.env")
    if err != nil {
        log.Fatal("Error loading .env file")
    }
}

func main() {

	fmt.Println("admin_secret", os.Getenv("HASURA_GRAPHQL_ADMIN_SECRET"))
	client :=  utilService.Client()
	log.Print(client)
	fmt.Println("newClient object", os.Getenv("HASURA_GRAPHQL_ENDPOINT"))
	// create server
	server := gin.New()
	// define middlewares
	server.Use(middlewares.Logger())
	server.Use(middlewares.CorsMiddleware())
	fmt.Println("server is running on port 7000")

	// define routes
	server.POST("/signup", controller.Signup)
	server.POST("/login", controller.Login)
	server.POST("/updateUser",controller.UpdateUser)
	server.POST("/uploadImage", controller.UploadImage)

	server.Run(":7000")
}	