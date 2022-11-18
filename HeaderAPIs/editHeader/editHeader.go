package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"portfolio-builder-backend/structsPackage"
	"strconv"

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
	query, err := db.Query(`UPDATE portfolio_builder_schema.header
	SET
	Name="` + myRequest.Name + `",
	Email="` + myRequest.Email + `",
	LinkedIn="` + myRequest.LinkedIn + `",
	Profile_URL="` + myRequest.Profile_URL + `",
	Summary="` + myRequest.Summary + `"
	WHERE ID=` + strconv.Itoa(myRequest.ID) + `;`)
	if err != nil {
		panic(err.Error())
	}
	defer query.Close()
	return events.APIGatewayProxyResponse{Body: string("HEader Edited Succesfully"), StatusCode: 200}, nil

}

func main() {
	lambda.Start(HandleLambdaEvent)
}
