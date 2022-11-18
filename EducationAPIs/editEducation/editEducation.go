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

	var myRequest structsPackage.Education
	json.Unmarshal([]byte(req.Body), &myRequest)

	db, err := sql.Open("mysql", "admin:Ks0756454835@tcp(portfolio-builder-database-dev.cbwqxjvaa6sv.us-west-1.rds.amazonaws.com:3306)/portfolio_builder_schema")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	fmt.Println("Success!")
	query, err := db.Query(`UPDATE portfolio_builder_schema.education
	SET
	Name = "` + myRequest.Name + `",
	Degree =  "` + myRequest.Degree + `",
	Location =  "` + myRequest.Location + `",
	Major =  "` + myRequest.Major + `",
	Start_Date =  "` + myRequest.StartDate + `",
	End_Date =  "` + myRequest.EndDate + `",
	GPA =  "` + fmt.Sprintf("%.2f", myRequest.GPA) + `"
	WHERE ID = ` + strconv.Itoa(myRequest.ID) + `;`)
	if err != nil {
		panic(err.Error())
	}
	defer query.Close()
	deleteQuery, err := db.Query(`DELETE FROM courses WHERE education_id = ` + strconv.Itoa(myRequest.ID) + `;`)
	if err != nil {
		panic(err.Error())
	}
	defer deleteQuery.Close()
	var insertCourse *sql.Rows
	var courses []structsPackage.Course = myRequest.Courses
	var i int
	for i = 0; i < len(courses); i++ {
		insertCourse, err = db.Query(`INSERT INTO portfolio_builder_schema.courses
			(CourseName,
			education_id)
			VALUES
			("` + courses[i].CourseName + `",
			` + strconv.Itoa(myRequest.ID) + `);`)
		if err != nil {
			panic(err.Error())
		}
		defer insertCourse.Close()
	}
	return events.APIGatewayProxyResponse{Body: "Updated Education Succesfully", StatusCode: 200}, nil

}

func main() {
	lambda.Start(HandleLambdaEvent)
}
