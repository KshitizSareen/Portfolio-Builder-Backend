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
	query, err := db.Query(`SELECT * FROM portfolio_builder_schema.education
	WHERE user_id=` + strconv.Itoa(myRequest.ID) + `;`)
	if err != nil {
		panic(err.Error())
	}
	defer query.Close()
	var education structsPackage.Education
	var educations []structsPackage.Education
	var courseQuery *sql.Rows
	var course structsPackage.Course
	var courses []structsPackage.Course
	for query.Next() {
		query.Scan(&education.ID, &education.Name, &education.Degree, &education.Location, &education.Major, &education.StartDate, &education.EndDate, &education.GPA, &education.User_id)
		courseQuery, err = db.Query(`SELECT *
	FROM portfolio_builder_schema.courses WHERE education_id = ` + strconv.Itoa(education.ID) + `;`)
		courses = nil
		for courseQuery.Next() {
			courseQuery.Scan(&course.ID, &course.CourseName, &course.Education_ID)
			courses = append(courses, course)
		}
		education.Courses = courses
		educations = append(educations, education)
	}
	resp, err := json.Marshal(educations)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}
	return events.APIGatewayProxyResponse{Body: string(resp), StatusCode: 200}, nil

}

func main() {
	lambda.Start(HandleLambdaEvent)
}
