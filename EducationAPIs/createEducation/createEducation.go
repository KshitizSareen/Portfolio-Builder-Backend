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
	query, err := db.Query(`INSERT INTO portfolio_builder_schema.education
	(Name,
	Degree,
	Location,
	Major,
	Start_Date,
	End_Date,
	GPA,
	user_id)
	VALUES
	("` + myRequest.Name + `",
	"` + myRequest.Degree + `",
	"` + myRequest.Location + `",
	"` + myRequest.Major + `",
	"` + myRequest.StartDate + `",
	"` + myRequest.EndDate + `",
	` + fmt.Sprintf("%.2f", myRequest.GPA) + `,
	` + strconv.Itoa(myRequest.User_id) + `);`)
	if err != nil {
		panic(err.Error())
	}
	defer query.Close()
	res, err := db.Query("SELECT ID FROM education ORDER BY ID DESC LIMIT 1")
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
	var insertCourse *sql.Rows
	var courses []structsPackage.Course = myRequest.Courses
	var i int
	for i = 0; i < len(courses); i++ {
		insertCourse, err = db.Query(`INSERT INTO portfolio_builder_schema.courses
			(CourseName,
			education_id)
			VALUES
			("` + courses[i].CourseName + `",
			` + strconv.Itoa(ID) + `);`)
		if err != nil {
			panic(err.Error())
		}
		defer insertCourse.Close()
	}
	return events.APIGatewayProxyResponse{Body: "Education Created Succesfully", StatusCode: 200}, nil

}

func main() {
	lambda.Start(HandleLambdaEvent)
}
