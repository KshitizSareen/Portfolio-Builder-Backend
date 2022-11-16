package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type MyEvent struct {
	Name string `json:"What is your name?"`
	Age  int    `json:"How old are you?"`
}

type MyResponse struct {
	Message string `json:"test"`
}

type InputResponse struct {
	ID int `json:"Test"`
}

/*
	func HandleLambdaEvent(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
		var myResponse MyResponse
		json.Unmarshal([]byte(req.Body), &myResponse)
		return events.APIGatewayProxyResponse{Body: myResponse.Message, StatusCode: 200}, nil
	}
*/
func main() {
	db, err := sql.Open("mysql", "admin:Ks0756454835@tcp(portfolio-builder-database-dev.cbwqxjvaa6sv.us-west-1.rds.amazonaws.com:3306)/portfolio_builder_schema")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	fmt.Println("Success!")
	query, err := db.Query("INSERT INTO test_tables VALUES ()")
	if err != nil {
		panic(err.Error())
	}
	defer query.Close()
}
