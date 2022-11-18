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
	query, err := db.Query(`SELECT * FROM portfolio_builder_schema.Skill_Categories
	WHERE user_id=` + strconv.Itoa(myRequest.ID) + `;`)
	if err != nil {
		panic(err.Error())
	}
	defer query.Close()
	var skillCategory structsPackage.Skill_Category
	var skillCategories []structsPackage.Skill_Category
	var skillQuery *sql.Rows
	var skill structsPackage.Skill
	var skills []structsPackage.Skill
	for query.Next() {
		query.Scan(&skillCategory.ID, &skillCategory.Category, &skillCategory.User_id)
		skillQuery, err = db.Query(`SELECT *
	FROM portfolio_builder_schema.skills WHERE skill_category_id = ` + strconv.Itoa(skillCategory.ID) + `;`)
		skills = nil
		for skillQuery.Next() {
			skillQuery.Scan(&skill.ID, &skill.Skill, &skill.Skill_Category_ID)
			skills = append(skills, skill)
		}
		skillCategory.Skills = skills
		skillCategories = append(skillCategories, skillCategory)
	}
	resp, err := json.Marshal(skillCategories)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}
	return events.APIGatewayProxyResponse{Body: string(resp), StatusCode: 200}, nil

}

func main() {
	lambda.Start(HandleLambdaEvent)
}
