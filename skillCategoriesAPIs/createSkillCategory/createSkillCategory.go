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

	var myRequest structsPackage.Skill_Category
	json.Unmarshal([]byte(req.Body), &myRequest)

	db, err := sql.Open("mysql", "admin:Ks0756454835@tcp(portfolio-builder-database-dev.cbwqxjvaa6sv.us-west-1.rds.amazonaws.com:3306)/portfolio_builder_schema")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	fmt.Println("Success!")
	query, err := db.Query(`INSERT INTO portfolio_builder_schema.Skill_Categories
	(
		Category,
		user_id)
	VALUES
	("` + myRequest.Category + `",
	` + strconv.Itoa(myRequest.User_id) + `);`)
	if err != nil {
		panic(err.Error())
	}
	defer query.Close()
	res, err := db.Query("SELECT ID FROM Skill_Categories ORDER BY ID DESC LIMIT 1")
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
	var insertSkills *sql.Rows
	var skills []structsPackage.Skill = myRequest.Skills
	var i int
	for i = 0; i < len(skills); i++ {
		insertSkills, err = db.Query(`INSERT INTO portfolio_builder_schema.skills
			(skill,
				skill_category_id)
			VALUES
			("` + skills[i].Skill + `",
			` + strconv.Itoa(ID) + `);`)
		if err != nil {
			panic(err.Error())
		}
		defer insertSkills.Close()
	}
	return events.APIGatewayProxyResponse{Body: "Skill Category Created Succesfully", StatusCode: 200}, nil

}

func main() {
	lambda.Start(HandleLambdaEvent)
}
