package main

import (
	"database/sql"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

func HandleLambdaEvent(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	db, err := sql.Open("mysql", "admin:Ks0756454835@tcp(portfolio-builder-database-dev.cbwqxjvaa6sv.us-west-1.rds.amazonaws.com:3306)/portfolio_builder_schema")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	fmt.Println("Success!")
	randomNumber := uuid.New()
	strVar := randomNumber.String()
	query, err := db.Query(`INSERT INTO user_keys (user_secret_key) VALUES (MD5("` + strVar + `"));`)
	if err != nil {
		panic(err.Error())
	}
	defer query.Close()
	return events.APIGatewayProxyResponse{Body: strVar, StatusCode: 200}, nil

}

func main() {
	lambda.Start(HandleLambdaEvent)
}
