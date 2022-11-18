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
	query, err := db.Query(`UPDATE portfolio_builder_schema.Skill_Categories
	SET
	Category = "` + myRequest.Category + `"
	WHERE ID = ` + strconv.Itoa(myRequest.ID) + `;`)
	if err != nil {
		panic(err.Error())
	}
	defer query.Close()
	deleteQuery, err := db.Query(`DELETE FROM skills WHERE skill_category_id = ` + strconv.Itoa(myRequest.ID) + `;`)
	if err != nil {
		panic(err.Error())
	}
	defer deleteQuery.Close()
	var insertSkill *sql.Rows
	var skills []structsPackage.Skill = myRequest.Skills
	var i int
	for i = 0; i < len(skills); i++ {
		insertSkill, err = db.Query(`INSERT INTO portfolio_builder_schema.skills
			(skill,
				skill_category_id)
			VALUES
			("` + skills[i].Skill + `",
			` + strconv.Itoa(myRequest.ID) + `);`)
		if err != nil {
			panic(err.Error())
		}
		defer insertSkill.Close()
	}
	return events.APIGatewayProxyResponse{Body: "Edited Skills Succesfully", StatusCode: 200}, nil

}

func main() {
	lambda.Start(HandleLambdaEvent)
}
