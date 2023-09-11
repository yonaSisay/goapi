package controller

// imports
import (
	"context"
	"fmt"
	"net/http"
	// "gilab.com/progrmaticreviwes/golang-gin-poc/utilService"
	"goapi/utilService"
	"github.com/gin-gonic/gin"
)


type AuthResponse struct {
	ID string `json: "id"`
	Role string `json: "role"`
	Token string `json: "token"`
}

func sendtoken(ctx *gin.Context, role string, response AuthResponse){
	// generate jwt token
	token, err := utilService.GetToken(response.ID, role)

	if err != nil {
		fmt.Println(err.Error(), "when generating token")
		ctx.JSON(400, gin.H{"message": "something went wrong when creating token"})
		return
	}
	response.Role = role
	response.Token = token
	ctx.JSON(200, response)
}

func Signup(ctx *gin.Context){
	// get the user input from the request body
	type inputUser struct {
		name string `json: "name"`
		Email string `json: "email"`
		Password string `json: "password"`
	}

	var newUser inputUser

	if err := ctx.ShouldBindJSON(&newUser); err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Printf("%+v\n", newUser,"new user coming")

	// define the graphql mutation string
	var mutation struct {
		InsertUsers struct {
			Returning []struct{
				ID string `json: "id"`
			} `json: "returning'`
		} `graphql: "insert_users(objects: {name:$name, email:$email, password:$password})"`
	}

	// hash password
	password, err4 := utilService.HashPassword(newUser.Password)
	if err4 != nil {
		fmt.Println(err4.Error(), "when hashing password")
		ctx.JSON(400, gin.H{"error": err4.Error()})
		return
	}

	// construct graphql variable

	variables := map[string]interface{}{
		"name": newUser.name,
		"email": newUser.Email,
		"password": password,
	}
	// execute the request 
	err := utilService.Client.Mutate(context.Background(), &mutation, variables)
	if err != nil {
		fmt.Println(err.Error(), "when executing mutation")
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	// if data stored successfully call sendtoken function with response object
	var response AuthResponse
	response.ID = mutation.InsertUsers.Returning[0].ID
	sendtoken(ctx, "user", response)
}