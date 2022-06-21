package main

import (
	//gin-Gin allows you to build web applications and microservices in Go. It contains a set of commonly used functionalities (e.g., routing, middleware support, rendering, etc.)
	//gin helps in creating rest API but we need to get this dependancy inside our application.
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	aman    = student{"Aman Narayan Singh", "ISE", "4NI19IS119"}
	swathi  = student{"Swathi BJ", "ISE", "4NI19IS103"}
	aditya  = student{"Aditya Mishra", "ISE", "4NI19IS008"}
	shreesh = student{"Shreesh Kulkarni", "ISE", "4NI19CS108"}
)

type todo struct {
	Id           string    `json:"id"`
	Item         string    `json:"title"`
	Completed    bool      `json:"completed"`
	Group        string    `json:"group id"`
	Project_Name string    `json:"Project Name"`
	Team_Members []student `json:"Team Members of Group 2"`
}
type student struct {
	Name   string
	Branch string
	USN    string
}

var todos = []todo{
	{Id: "1", Item: "wake up", Completed: true},
	{Id: "2", Item: "take bath", Completed: true},
	{Id: "3", Item: "eat food", Completed: false},
	{Id: "4", Item: "dress up", Completed: true},
	{Id: "5", Item: "college", Completed: false},
	{Id: "6", Item: "HPE CTY Program", Group: "2", Project_Name: "GoLang SDK package to use for iLO Restful Operations", Team_Members: []student{aman, swathi, aditya, shreesh}, Completed: false},
}

// client and server can communicate to each other through JSON
// so when server decides to send the data back to the client then server has to get the whole todos data structure and convert it to JSON
// and when client gets this json data structure then it has to be coverted back to the data structure which GO can understand
func getTodos(context *gin.Context) {

	//context will contain a bunch of information regarding the HTTP request like we can have data inside of the request body, header,etc. and when we need to extract the data we can do that with the help of context
	//now with context we need to transform the todos data array of data structure into JSON
	context.IndentedJSON(http.StatusOK, todos) //converts data structure to JSON using StatusOK
	//http tatus code will return the status code for the execution of the code
	//200 OK This HTTP status code shows request was successfully completed and includes a representation in its body.

}
func postTodos(c *gin.Context) {
	var newTodos todo
	if err := c.BindJSON(&newTodos); err != nil {
		return
	}
	todos = append(todos, newTodos)
	c.IndentedJSON(http.StatusCreated, newTodos)
}
func main() {
	//we want to create our server by assigning a variable named router
	router := gin.Default()
	router.GET("/todos", getTodos) // todos to get is that we need to get multiple todos and another function returns JSON data to the client.
	router.POST("/todos", postTodos)
	router.Run("localhost:5000") //our application should be running on port 9090
}
