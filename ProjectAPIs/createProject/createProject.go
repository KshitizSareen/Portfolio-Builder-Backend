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

	var myRequest structsPackage.Project
	json.Unmarshal([]byte(req.Body), &myRequest)

	db, err := sql.Open("mysql", "admin:Ks0756454835@tcp(portfolio-builder-database-dev.cbwqxjvaa6sv.us-west-1.rds.amazonaws.com:3306)/portfolio_builder_schema")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	fmt.Println("Success!")
	query, err := db.Query(`INSERT INTO portfolio_builder_schema.Projects
	(
		Project_Name,
		Position,
		Start_Date,
		End_Date,
		Project_URL,
		user_id)
	VALUES
	("` + myRequest.Project_Name + `",
	"` + myRequest.Position + `",
	"` + myRequest.StartDate + `",
	"` + myRequest.EndDate + `",
	"` + myRequest.Project_URL + `",
	` + strconv.Itoa(myRequest.User_id) + `);`)
	if err != nil {
		panic(err.Error())
	}
	defer query.Close()
	res, err := db.Query("SELECT ID FROM Projects ORDER BY ID DESC LIMIT 1")
	if err != nil {
		panic(err.Error())
	}
	defer res.Close()
	var ID int
	for res.Next() {
		err := res.Scan(&ID)
		if err != nil {
			panic(err.Error())
		}
	}
	var insertProjectDescription *sql.Rows
	var Project_Descriptions []structsPackage.Project_Description = myRequest.Project_Descriptions
	var i int
	for i = 0; i < len(Project_Descriptions); i++ {
		insertProjectDescription, err = db.Query(`INSERT INTO portfolio_builder_schema.Project_Description
			(Description,
				project_id)
			VALUES
			("` + Project_Descriptions[i].Description + `",
			` + strconv.Itoa(ID) + `);`)
		if err != nil {
			panic(err.Error())
		}
		defer insertProjectDescription.Close()
	}
	return events.APIGatewayProxyResponse{Body: "Project Created Succesfully", StatusCode: 200}, nil

}

func main() {
	lambda.Start(HandleLambdaEvent)
}
