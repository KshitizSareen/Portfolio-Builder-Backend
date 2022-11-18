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
	query, err := db.Query(`UPDATE portfolio_builder_schema.Projects
	SET
	Project_Name = "` + myRequest.Project_Name + `",
	Position =  "` + myRequest.Position + `",
	Start_Date =  "` + myRequest.StartDate + `",
	End_Date =  "` + myRequest.EndDate + `",
	Project_URL =  "` + myRequest.Project_URL + `"
	WHERE ID = ` + strconv.Itoa(myRequest.ID) + `;`)
	if err != nil {
		panic(err.Error())
	}
	defer query.Close()
	deleteQuery, err := db.Query(`DELETE FROM Project_Description WHERE project_id = ` + strconv.Itoa(myRequest.ID) + `;`)
	if err != nil {
		panic(err.Error())
	}
	defer deleteQuery.Close()
	var insertProjectDescription *sql.Rows
	var Project_Descriptions []structsPackage.Project_Description = myRequest.Project_Descriptions
	var i int
	for i = 0; i < len(Project_Descriptions); i++ {
		insertProjectDescription, err = db.Query(`INSERT INTO portfolio_builder_schema.Project_Description
			(Description,
				project_id)
			VALUES
			("` + Project_Descriptions[i].Description + `",
			` + strconv.Itoa(myRequest.ID) + `);`)
		if err != nil {
			panic(err.Error())
		}
		defer insertProjectDescription.Close()
	}
	return events.APIGatewayProxyResponse{Body: "Updated Project Succesfully", StatusCode: 200}, nil

}

func main() {
	lambda.Start(HandleLambdaEvent)
}
