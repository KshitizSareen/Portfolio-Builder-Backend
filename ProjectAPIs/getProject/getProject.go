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
	query, err := db.Query(`SELECT * FROM portfolio_builder_schema.Projects
	WHERE user_id=` + strconv.Itoa(myRequest.ID) + `;`)
	if err != nil {
		panic(err.Error())
	}
	defer query.Close()
	var project structsPackage.Project
	var projects []structsPackage.Project
	var projectDescriptionQuery *sql.Rows
	var projectDescription structsPackage.Project_Description
	var projectDescriptions []structsPackage.Project_Description
	for query.Next() {
		query.Scan(&project.ID, &project.Project_Name, &project.Position, &project.StartDate, &project.EndDate, &project.Project_URL, &project.User_id)
		projectDescriptionQuery, err = db.Query(`SELECT *
	FROM portfolio_builder_schema.Project_Description WHERE project_id = ` + strconv.Itoa(project.ID) + `;`)
		projectDescriptions = nil
		for projectDescriptionQuery.Next() {
			projectDescriptionQuery.Scan(&projectDescription.ID, &projectDescription.Description, &projectDescription.Project_id)
			projectDescriptions = append(projectDescriptions, projectDescription)
		}
		project.Project_Descriptions = projectDescriptions
		projects = append(projects, project)
	}
	resp, err := json.Marshal(projects)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}
	return events.APIGatewayProxyResponse{Body: string(resp), StatusCode: 200}, nil

}

func main() {
	lambda.Start(HandleLambdaEvent)
}
