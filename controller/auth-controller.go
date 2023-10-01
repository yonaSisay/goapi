package controller

// imports
import (
	"context"
	"fmt"
	"net/http"
	"goapi/utilService"
	"github.com/gin-gonic/gin"
	"strconv"

)


type AuthResponse struct {
	ID string `json:"id"`
	Role string `json:"role"`
	Token string `json:"token"`
}

func sendToken(ctx *gin.Context, role string, response AuthResponse){
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

func Login( ctx *gin.Context){
	//1.bind the json data to the struct
	var input struct {
		Email string `json: "email"`
		Password string `json: "password"`
		
	}
	
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	fmt.Println("login controller", input.Email)
	//2.creating the graphql query and then query the db

	
	var query struct {
		Users[] struct {
			ID     int `json:"id"`
			Email string `json:"email"`
			Password string `json:"password"`
			Role string `json:"role"`

		} `graphql:"users(where: {email: {_eq: $email}})"`
	}

	variables := map[string]interface{}{
		"email":  input.Email,
	}
	err := utilService.Client().Query(context.Background(), &query, variables)
	// 3.check if user exist with that email 
	if err != nil {
		fmt.Println(err.Error(), "when querying user to login")
		ctx.JSON(400, err)
		return
	}
	// 4.check if the password is correct if it is send token
	if len(query.Users) > 0  && utilService.ComparePasswords(query.Users[0].Password, input.Password){
		var response AuthResponse
		response.ID = strconv.Itoa(query.Users[0].ID)
		response.Role = query.Users[0].Role
		sendToken(ctx, query.Users[0].Role, response)
		return
	}else{
		ctx.JSON(400, gin.H{"message": "Invalid credentials"})
		return
	}

}

// Signup Controller
func Signup(ctx *gin.Context){
	// get the user input from the request body
	type inputUser struct {
		Name string `json:"name"`
		Email string `json:"email"`
		Password string `json:"password"`
	}

	var newUser inputUser

	if err := ctx.ShouldBindJSON(&newUser); err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Printf(newUser.Email)
	fmt.Printf(newUser.Name)
	fmt.Printf(newUser.Password)

	// define the graphql mutation string
	var mutation struct {
		InsertUsers struct {
			Returning []struct {
				ID int `json:"id"`
			} `json:"returning"`
		} `graphql:"insert_users(objects: $objects)"`
	}
	
   type users_insert_input map[string]interface{}
	// hash password

	password, err4 := utilService.HashPassword(newUser.Password)
	if err4 != nil {
		fmt.Println(err4.Error(), "when hashing password")
		ctx.JSON(400, gin.H{"error": err4.Error()})
		return
	}
	fmt.Println(password)
	// construct graphql variable

	variables := map[string]interface{}{
		"objects": []users_insert_input{
			{
				"name":     newUser.Name,
				"email":    newUser.Email,
				"password": password,
			},
		},
	}
	// execute the request 
	client := utilService.Client();

	err := client.Mutate(context.Background(), &mutation, variables)
	if err != nil {
		ctx.JSON(400, err)
		return
	}

	fmt.Println(mutation.InsertUsers.Returning[0].ID)

	// if data stored successfully call sendtoken function with response object
	var response AuthResponse
	response.ID = strconv.Itoa(mutation.InsertUsers.Returning[0].ID)
	sendToken(ctx, "user", response)
}