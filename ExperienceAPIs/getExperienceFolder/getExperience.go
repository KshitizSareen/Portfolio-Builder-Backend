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

	var myRequest structsPackage.User
	json.Unmarshal([]byte(req.Body), &myRequest)

	db, err := sql.Open("mysql", "admin:Ks0756454835@tcp(portfolio-builder-database-dev.cbwqxjvaa6sv.us-west-1.rds.amazonaws.com:3306)/portfolio_builder_schema")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	fmt.Println("Success!")
	query, err := db.Query(`SELECT * FROM portfolio_builder_schema.Experience
	WHERE user_id=` + strconv.Itoa(myRequest.ID) + `;`)
	if err != nil {
		panic(err.Error())
	}
	defer query.Close()
	var experience structsPackage.Experience
	var experiences []structsPackage.Experience
	var experienceDescriptionQuery *sql.Rows
	var experienceDescription structsPackage.Experience_Description
	var experienceDescriptions []structsPackage.Experience_Description
	for query.Next() {
		query.Scan(&experience.ID, &experience.Company_Name, &experience.Position, &experience.Location, &experience.StartDate, &experience.EndDate, &experience.User_id)
		experienceDescriptionQuery, err = db.Query(`SELECT *
	FROM portfolio_builder_schema.Experience_Description WHERE experience_id = ` + strconv.Itoa(experience.ID) + `;`)
		experienceDescriptions = nil
		for experienceDescriptionQuery.Next() {
			experienceDescriptionQuery.Scan(&experienceDescription.ID, &experienceDescription.Description, &experienceDescription.Experience_id)
			experienceDescriptions = append(experienceDescriptions, experienceDescription)
		}
		experience.Experience_Descriptions = experienceDescriptions
		experiences = append(experiences, experience)
	}
	resp, err := json.Marshal(experiences)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}
	return events.APIGatewayProxyResponse{Body: string(resp), StatusCode: 200}, nil

}

func main() {
	lambda.Start(HandleLambdaEvent)
}
