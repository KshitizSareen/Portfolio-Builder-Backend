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

	var myRequest structsPackage.Header
	json.Unmarshal([]byte(req.Body), &myRequest)

	db, err := sql.Open("mysql", "admin:Ks0756454835@tcp(portfolio-builder-database-dev.cbwqxjvaa6sv.us-west-1.rds.amazonaws.com:3306)/portfolio_builder_schema")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	fmt.Println("Success!")
	query, err := db.Query(`INSERT INTO portfolio_builder_schema.header
	(Name,
	Email,
	LinkedIn,
	Profile_URL,
	Summary,
	user_id)
	VALUES
	("` + myRequest.Name + `",
	"` + myRequest.Email + `",
	"` + myRequest.LinkedIn + `",
	"` + myRequest.Profile_URL + `",
	"` + myRequest.Summary + `",
	` + myRequest.User_id + `);`)
	if err != nil {
		/*panic(`INSERT INTO portfolio_builder_schema.header
		(Name,
		Email,
		LinkedIn,
		Profile_URL,
		Summary,
		user_id)
		VALUES
		("` + myRequest.Name + `",
		"` + myRequest.Email + `",
		"` + myRequest.LinkedIn + `",
		"` + myRequest.Profile_URL + `",
		"` + myRequest.Summary + `",
		"` + myRequest.User_id + `");`)*/
		panic(err.Error())
	}
	defer query.Close()
	return events.APIGatewayProxyResponse{Body: "Header Created Succesfully", StatusCode: 200}, nil

}

func main() {
	lambda.Start(HandleLambdaEvent)
}
