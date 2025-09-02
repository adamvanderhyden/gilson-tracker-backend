package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

const (
	TABLE_TRACKER = "Tracker"
	PAR_SERIAL    = "serial"
)

var HEADERS = map[string]string{
	// "Access-Control-Allow-Headers": "Content-Type,X-Api-Key",
	"Access-Control-Allow-Origin": "*",
	// "Access-Control-Allow-Methods": "POST",
}
var db *dynamodb.Client

type Resp struct {
	Error        string         `json:"error"`
	TrackerJson  []TrackerJson  `json:"trackers,omitempty"`
	EmployeeJson []EmployeeJson `json:"employees,omitempty"`
}
type Data struct {
	Room         string         `json:"room"`
	Type         string         `json:"type"`
	Pin          string         `json:"pin"`
	TrackerJson  []TrackerJson  `json:"trackers,omitempty"`
	EmployeeJson []EmployeeJson `json:"employees,omitempty"`
}

func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	fmt.Println("Request Body:", string(request.Body))

	var d Data
	var r Resp
	r.Error = "good"

	err := json.Unmarshal([]byte(request.Body), &d)

	if err != nil {

		r.Error = fmt.Sprintf("Invalid JSON: %v", err)

		resp, _ := json.MarshalIndent(r, "", "  ")

		return events.APIGatewayProxyResponse{
			StatusCode: 400,
			Headers:    HEADERS,
			Body:       string(resp),
		}, nil
	}

	if d.Type == "list" {

		switch d.Room {

		case "hold":

			expr, bErr := expression.NewBuilder().WithFilter(
				expression.And(
					expression.AttributeNotExists(expression.Name("is_deleted")),
					expression.AttributeNotExists(expression.Name("time_finish")),
					expression.AttributeNotExists(expression.Name("time_start")))).Build()

			if bErr != nil {

				r.Error = fmt.Sprintf("Bad Filter Build: %v", err)

				resp, _ := json.MarshalIndent(r, "", "  ")

				return events.APIGatewayProxyResponse{
					StatusCode: 400,
					Headers:    HEADERS,
					Body:       string(resp),
				}, nil
			}

			return getTrackerRoom(expr)

		case "side":

			expr, bErr := expression.NewBuilder().WithFilter(
				expression.And(
					expression.AttributeNotExists(expression.Name("is_deleted")),
					expression.AttributeNotExists(expression.Name("time_finish")),
					expression.AttributeExists(expression.Name("time_start")),
					expression.Or(
						expression.AttributeNotExists(expression.Name("side_cnc")),
						expression.AttributeNotExists(expression.Name("side_pour"))))).Build()

			if bErr != nil {

				r.Error = fmt.Sprintf("Bad Filter Build: %v", err)

				resp, _ := json.MarshalIndent(r, "", "  ")

				return events.APIGatewayProxyResponse{
					StatusCode: 400,
					Headers:    HEADERS,
					Body:       string(resp),
				}, nil
			}

			return getTrackerRoom(expr)

		case "wood":
		case "sub":
		case "ed":
		case "dr":
		case "lay":
		case "fin":
		case "wax":
		default:

			r.Error = "Bad Room"

			resp, _ := json.MarshalIndent(r, "", "  ")

			return events.APIGatewayProxyResponse{
				StatusCode: 400,
				Headers:    HEADERS,
				Body:       string(resp),
			}, nil
		}
	} else if d.Type == "update" {
		switch d.Room {
		case "hold":
			return updateRoom()
		case "side":
			return updateRoom()
		case "wood":
		case "sub":
		case "ed":
		case "dr":
		case "lay":
		case "fin":
		case "wax":
		default:

			r.Error = "Bad Room"

			resp, _ := json.MarshalIndent(r, "", "  ")

			return events.APIGatewayProxyResponse{
				StatusCode: 400,
				Headers:    HEADERS,
				Body:       string(resp),
			}, nil
		}
	}

	r.Error = "Bad Type"

	resp, _ := json.MarshalIndent(r, "", "  ")

	return events.APIGatewayProxyResponse{
		StatusCode: 400,
		Headers:    HEADERS,
		Body:       string(resp),
	}, nil

}

func updateRoom() (events.APIGatewayProxyResponse, error) {

	var r Resp
	r.Error = "good"

	resp, _ := json.MarshalIndent(r, "", "  ")

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers:    HEADERS,
		Body:       string(resp),
	}, nil
}

func getTrackerRoom(inExpr expression.Expression) (events.APIGatewayProxyResponse, error) {

	var r Resp
	r.Error = "good"

	out, err := db.Scan(context.TODO(), &dynamodb.ScanInput{
		TableName:                aws.String(TABLE_TRACKER),
		ExpressionAttributeNames: inExpr.Names(),
		FilterExpression:         inExpr.Filter(),
	})

	if err != nil {

		r.Error = fmt.Sprintf("Bad Scan (%s): %v", TABLE_TRACKER, err)

		resp, _ := json.MarshalIndent(r, "", "  ")

		return events.APIGatewayProxyResponse{
			StatusCode: 400,
			Headers:    HEADERS,
			Body:       string(resp),
		}, nil
	}

	var td []TrackerDynamo

	err = attributevalue.UnmarshalListOfMaps(out.Items, &td)

	if err != nil {

		r.Error = fmt.Sprintf("Bad Unmarshal: %v", err)

		resp, _ := json.MarshalIndent(r, "", "  ")

		return events.APIGatewayProxyResponse{
			StatusCode: 400,
			Headers:    HEADERS,
			Body:       string(resp),
		}, nil
	}

	for _, t := range td {
		r.TrackerJson = append(r.TrackerJson, TrackerJson(t))
	}

	resp, _ := json.MarshalIndent(r, "", "  ")
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers:    HEADERS,
		Body:       string(resp),
	}, nil
}

func main() {

	cfg, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		panic("unable to load AWS SDK config, " + err.Error())
	}

	// Initialize DynamoDB client
	db = dynamodb.NewFromConfig(cfg)

	lambda.Start(Handler)
}
