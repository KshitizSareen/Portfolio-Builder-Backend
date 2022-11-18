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
	query, err := db.Query(`INSERT INTO portfolio_builder_schema.Experience
	(
		Company_Name,
		Position,
		Location,
		Start_Date,
		End_Date,
		user_id)
	VALUES
	("` + myRequest.Company_Name + `",
	"` + myRequest.Position + `",
	"` + myRequest.Location + `",
	"` + myRequest.StartDate + `",
	"` + myRequest.EndDate + `",
	` + strconv.Itoa(myRequest.User_id) + `);`)
	if err != nil {
		panic(err.Error())
	}
	defer query.Close()
	res, err := db.Query("SELECT ID FROM Experience ORDER BY ID DESC LIMIT 1")
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
	var insertExperienceDescription *sql.Rows
	var Experience_Descriptions []structsPackage.Experience_Description = myRequest.Experience_Descriptions
	var i int
	for i = 0; i < len(Experience_Descriptions); i++ {
		insertExperienceDescription, err = db.Query(`INSERT INTO portfolio_builder_schema.Experience_Description
			(Description,
			experience_id)
			VALUES
			("` + Experience_Descriptions[i].Description + `",
			` + strconv.Itoa(ID) + `);`)
		if err != nil {
			panic(err.Error())
		}
		defer insertExperienceDescription.Close()
	}
	return events.APIGatewayProxyResponse{Body: "Experience Created Succesfully", StatusCode: 200}, nil

}

func main() {
	lambda.Start(HandleLambdaEvent)
}
