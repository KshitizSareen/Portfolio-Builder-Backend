package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"portfolio-builder-backend/structsPackage"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	_ "github.com/go-sql-driver/mysql"
)

func HandleLambdaEvent(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	db, err := sql.Open("mysql", "admin:Ks0756454835@tcp(portfolio-builder-database-dev.cbwqxjvaa6sv.us-west-1.rds.amazonaws.com:3306)/portfolio_builder_schema")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	fmt.Println("Success!")
	var myRequest structsPackage.UserCode
	json.Unmarshal([]byte(req.Body), &myRequest)
	query, err := db.Query(`SELECT user_keys.ID
FROM portfolio_builder_schema.user_keys WHERE user_secret_key=MD5("` + myRequest.User_Code + `");`)
	if err != nil {
		panic(err.Error())
	}
	defer query.Close()
	var User structsPackage.User
	for query.Next() {
		query.Scan(&User.ID)
	}
	resp, err := json.Marshal(User)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}
	return events.APIGatewayProxyResponse{Body: string(resp), StatusCode: 200}, nil

}

func main() {
	lambda.Start(HandleLambdaEvent)
}
