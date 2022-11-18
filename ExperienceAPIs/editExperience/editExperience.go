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

	var myRequest structsPackage.Experience
	json.Unmarshal([]byte(req.Body), &myRequest)

	db, err := sql.Open("mysql", "admin:Ks0756454835@tcp(portfolio-builder-database-dev.cbwqxjvaa6sv.us-west-1.rds.amazonaws.com:3306)/portfolio_builder_schema")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	fmt.Println("Success!")
	query, err := db.Query(`UPDATE portfolio_builder_schema.Experience
	SET
	Company_Name = "` + myRequest.Company_Name + `",
	Position =  "` + myRequest.Position + `",
	Location =  "` + myRequest.Location + `",
	Start_Date =  "` + myRequest.StartDate + `",
	End_Date =  "` + myRequest.EndDate + `"
	WHERE ID = ` + strconv.Itoa(myRequest.ID) + `;`)
	if err != nil {
		panic(err.Error())
	}
	defer query.Close()
	deleteQuery, err := db.Query(`DELETE FROM Experience_Description WHERE experience_id = ` + strconv.Itoa(myRequest.ID) + `;`)
	if err != nil {
		panic(err.Error())
	}
	defer deleteQuery.Close()
	var insertExperienceDescription *sql.Rows
	var Experience_Descriptions []structsPackage.Experience_Description = myRequest.Experience_Descriptions
	var i int
	for i = 0; i < len(Experience_Descriptions); i++ {
		insertExperienceDescription, err = db.Query(`INSERT INTO portfolio_builder_schema.Experience_Description
			(Description,
			experience_id)
			VALUES
			("` + Experience_Descriptions[i].Description + `",
			` + strconv.Itoa(myRequest.ID) + `);`)
		if err != nil {
			panic(err.Error())
		}
		defer insertExperienceDescription.Close()
	}
	return events.APIGatewayProxyResponse{Body: "Updated Experience Succesfully", StatusCode: 200}, nil

}

func main() {
	lambda.Start(HandleLambdaEvent)
}
