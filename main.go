package main

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

const (
	TABLE_TRACKER = "Tracker"
	PAR_SERIAL    = "serial"
)

var HEADERS = map[string]string{
	// "Access-Control-Allow-Headers": "Content-Type,X-Api-Key",
	// "Access-Control-Allow-Methods": "POST",
	"Access-Control-Allow-Origin": "*",
}
var db *dynamodb.Client

type Resp struct {
	Error        string         `json:"error"`
	TrackerJson  []TrackerJson  `json:"trackers,omitempty"`
	EmployeeJson []EmployeeJson `json:"employees,omitempty"`
	Note         string         `json:"note,omitempty"`
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
		fmt.Println(r.Error)

		resp, _ := json.MarshalIndent(r, "", "  ")

		return events.APIGatewayProxyResponse{
			StatusCode: 400,
			Headers:    HEADERS,
			Body:       string(resp),
		}, nil
	}

	if d.Type == "list" {

		switch d.Room {

		// TODO - sort everything

		case "hold":

			expr, bErr := expression.NewBuilder().WithFilter(
				expression.And(
					expression.AttributeNotExists(expression.Name("is_deleted")),
					expression.AttributeNotExists(expression.Name("time_finish")),
					expression.AttributeNotExists(expression.Name("time_start")))).Build()

			// expression.And(
			// expression.Or(
			// 	expression.AttributeNotExists(expression.Name("is_deleted")),
			// 	expression.AttributeNotExists(expression.Name("is_deleted"))),
			// expression.Or(
			// 	expression.AttributeNotExists(expression.Name("time_finish")),
			// 	expression.AttributeNotExists(expression.Name("time_finish"))),
			// expression.Or(
			// 	expression.AttributeNotExists(expression.Name("time_start")),
			// 	expression.AttributeNotExists(expression.Name("time_start"))))).Build()

			if bErr != nil {

				r.Error = fmt.Sprintf("Bad Filter Build: %v", bErr)
				fmt.Println(r.Error)

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

				r.Error = fmt.Sprintf("Bad Filter Build: %v", bErr)
				fmt.Println(r.Error)

				resp, _ := json.MarshalIndent(r, "", "  ")

				return events.APIGatewayProxyResponse{
					StatusCode: 400,
					Headers:    HEADERS,
					Body:       string(resp),
				}, nil
			}

			return getTrackerRoom(expr)

		case "wood":

			expr, bErr := expression.NewBuilder().WithFilter(
				expression.And(
					expression.AttributeNotExists(expression.Name("is_deleted")),
					expression.AttributeNotExists(expression.Name("time_finish")),
					expression.AttributeExists(expression.Name("time_start")),
					expression.AttributeExists(expression.Name("side_cnc")),
					expression.AttributeExists(expression.Name("side_pour")),
					expression.Or(
						expression.AttributeNotExists(expression.Name("wood_core")),
						expression.AttributeNotExists(expression.Name("wood_cart"))))).Build()

			if bErr != nil {

				r.Error = fmt.Sprintf("Bad Filter Build: %v", bErr)
				fmt.Println(r.Error)

				resp, _ := json.MarshalIndent(r, "", "  ")

				return events.APIGatewayProxyResponse{
					StatusCode: 400,
					Headers:    HEADERS,
					Body:       string(resp),
				}, nil
			}

			return getTrackerRoom(expr)

		case "sub":

			expr, bErr := expression.NewBuilder().WithFilter(
				expression.And(
					expression.AttributeNotExists(expression.Name("is_deleted")),
					expression.AttributeNotExists(expression.Name("time_finish")),
					expression.AttributeExists(expression.Name("time_start")),
					expression.AttributeExists(expression.Name("side_cnc")),
					expression.AttributeExists(expression.Name("side_pour")),
					expression.Or(
						expression.AttributeNotExists(expression.Name("sub_base")),
						expression.AttributeNotExists(expression.Name("sub_top")),
						expression.AttributeNotExists(expression.Name("sub_cart"))))).Build()

			if bErr != nil {

				r.Error = fmt.Sprintf("Bad Filter Build: %v", bErr)
				fmt.Println(r.Error)

				resp, _ := json.MarshalIndent(r, "", "  ")

				return events.APIGatewayProxyResponse{
					StatusCode: 400,
					Headers:    HEADERS,
					Body:       string(resp),
				}, nil
			}

			return getTrackerRoom(expr)

		case "ed":

			expr, bErr := expression.NewBuilder().WithFilter(
				expression.And(
					expression.AttributeNotExists(expression.Name("is_deleted")),
					expression.AttributeNotExists(expression.Name("time_finish")),
					expression.AttributeExists(expression.Name("time_start")),
					expression.AttributeExists(expression.Name("sub_base")),
					expression.AttributeExists(expression.Name("sub_top")),
					expression.AttributeExists(expression.Name("sub_cart")),
					expression.Or(
						expression.AttributeNotExists(expression.Name("ed_glue")),
						expression.AttributeNotExists(expression.Name("ed_cart"))))).Build()

			if bErr != nil {

				r.Error = fmt.Sprintf("Bad Filter Build: %v", bErr)
				fmt.Println(r.Error)

				resp, _ := json.MarshalIndent(r, "", "  ")

				return events.APIGatewayProxyResponse{
					StatusCode: 400,
					Headers:    HEADERS,
					Body:       string(resp),
				}, nil
			}

			return getTrackerRoom(expr)

		case "dr":

			expr, bErr := expression.NewBuilder().WithFilter(
				expression.And(
					expression.AttributeNotExists(expression.Name("is_deleted")),
					expression.AttributeNotExists(expression.Name("time_finish")),
					expression.AttributeExists(expression.Name("time_start")),
					expression.Or(
						expression.AttributeExists(expression.Name("wood_cart")),
						expression.AttributeExists(expression.Name("sub_cart")),
						expression.AttributeExists(expression.Name("ed_cart"))),
					expression.Or(
						expression.AttributeNotExists(expression.Name("wood_cart")),
						expression.AttributeNotExists(expression.Name("sub_cart")),
						expression.AttributeNotExists(expression.Name("ed_cart"))))).Build()

			if bErr != nil {

				r.Error = fmt.Sprintf("Bad Filter Build: %v", bErr)
				fmt.Println(r.Error)

				resp, _ := json.MarshalIndent(r, "", "  ")

				return events.APIGatewayProxyResponse{
					StatusCode: 400,
					Headers:    HEADERS,
					Body:       string(resp),
				}, nil
			}

			return getTrackerRoom(expr)

		case "lay":

			expr, bErr := expression.NewBuilder().WithFilter(
				expression.And(
					expression.AttributeNotExists(expression.Name("is_deleted")),
					expression.AttributeNotExists(expression.Name("time_finish")),
					expression.AttributeExists(expression.Name("time_start")),
					expression.AttributeExists(expression.Name("wood_cart")),
					expression.AttributeExists(expression.Name("sub_cart")),
					expression.AttributeExists(expression.Name("ed_cart")),
					expression.Or(
						expression.AttributeNotExists(expression.Name("lay_press")),
						expression.AttributeNotExists(expression.Name("lay_inspect"))))).Build()

			if bErr != nil {

				r.Error = fmt.Sprintf("Bad Filter Build: %v", bErr)
				fmt.Println(r.Error)

				resp, _ := json.MarshalIndent(r, "", "  ")

				return events.APIGatewayProxyResponse{
					StatusCode: 400,
					Headers:    HEADERS,
					Body:       string(resp),
				}, nil
			}

			return getTrackerRoom(expr)

		case "fin":

			expr, bErr := expression.NewBuilder().WithFilter(
				expression.And(
					expression.AttributeNotExists(expression.Name("is_deleted")),
					expression.AttributeNotExists(expression.Name("time_finish")),
					expression.AttributeExists(expression.Name("time_start")),
					expression.AttributeExists(expression.Name("lay_press")),
					expression.AttributeExists(expression.Name("lay_inspect")),
					expression.Or(
						expression.AttributeNotExists(expression.Name("fin_tune")),
						expression.AttributeNotExists(expression.Name("fin_inspect"))))).Build()

			if bErr != nil {

				r.Error = fmt.Sprintf("Bad Filter Build: %v", bErr)
				fmt.Println(r.Error)

				resp, _ := json.MarshalIndent(r, "", "  ")

				return events.APIGatewayProxyResponse{
					StatusCode: 400,
					Headers:    HEADERS,
					Body:       string(resp),
				}, nil
			}

			return getTrackerRoom(expr)

		case "wax":

			expr, bErr := expression.NewBuilder().WithFilter(
				expression.And(
					expression.AttributeNotExists(expression.Name("is_deleted")),
					expression.AttributeNotExists(expression.Name("time_finish")),
					expression.AttributeExists(expression.Name("time_start")),
					expression.AttributeExists(expression.Name("fin_tune")),
					expression.AttributeExists(expression.Name("fin_inspect")),
					expression.Or(
						expression.AttributeNotExists(expression.Name("wax_wax")),
						expression.AttributeNotExists(expression.Name("wax_inspect"))))).Build()

			if bErr != nil {

				r.Error = fmt.Sprintf("Bad Filter Build: %v", bErr)
				fmt.Println(r.Error)

				resp, _ := json.MarshalIndent(r, "", "  ")

				return events.APIGatewayProxyResponse{
					StatusCode: 400,
					Headers:    HEADERS,
					Body:       string(resp),
				}, nil
			}

			return getTrackerRoom(expr)

		default:

			r.Error = "Bad Room"
			fmt.Println(r.Error)

			resp, _ := json.MarshalIndent(r, "", "  ")

			return events.APIGatewayProxyResponse{
				StatusCode: 400,
				Headers:    HEADERS,
				Body:       string(resp),
			}, nil
		}
	} else if d.Type == "update" {

		// get employee
		employee, bErr := getEmployeeFromPin(d.Pin)

		if employee == "" || bErr != nil {

			r.Error = "Bad PIN"
			fmt.Println(r.Error)

			resp, _ := json.MarshalIndent(r, "", "  ")

			return events.APIGatewayProxyResponse{
				StatusCode: 200,
				Headers:    HEADERS,
				Body:       string(resp),
			}, nil
		}

		//*
		//*
		//* UPDATES
		//*
		//*

		for _, tj := range d.TrackerJson {

			if len(tj.Notes) > 0 {

				r.Note, err = updateNotes(employee, tj.Serial, tj.Notes)

				if err != nil {
					r.Error = "Error writing notes"
					fmt.Println(r.Error)

					resp, _ := json.MarshalIndent(r, "", "  ")
					return events.APIGatewayProxyResponse{
						StatusCode: 400,
						Headers:    HEADERS,
						Body:       string(resp),
					}, nil
				}
			}

			switch d.Room {
			case "hold":

				// TODO - per serial in TrackerJson
				// build query using time_start, deleted, and who_deleted (employee)

				// **** if time_start or deleted is present only update if the value in the table
				// is not present.  if time_start or deleted/who_deleted is blank, update dynamodb
				// item to be removed

				return updateRoom()
			case "side":
			case "wood":
			case "sub":
			case "ed":
			case "dr":
			case "lay":
			case "fin":
			case "wax":
			default:

				r.Error = "Bad Room"
				fmt.Println(r.Error)

				resp, _ := json.MarshalIndent(r, "", "  ")

				return events.APIGatewayProxyResponse{
					StatusCode: 400,
					Headers:    HEADERS,
					Body:       string(resp),
				}, nil
			}
		}
	}

	r.Error = "Bad Type"
	fmt.Println(r.Error)

	resp, _ := json.MarshalIndent(r, "", "  ")

	return events.APIGatewayProxyResponse{
		StatusCode: 400,
		Headers:    HEADERS,
		Body:       string(resp),
	}, nil

}

func updateNotes(inEmployee string, inSerial string, inNotes []string) (string, error) {

	// there should never be more than one note, but it's array based to jive with the
	// rest of the "update" process to handle many
	firstNote := ""

	t := time.Now()
	hh, min, _ := t.Clock()
	yy, mon, dd := t.Date()

	for _, n := range inNotes {

		if firstNote == "" {
			firstNote = n
		}

		note := fmt.Sprintf("%02d/%02d/%d %02d:%02d (%s): %s", int(mon), dd, yy, hh, min, inEmployee, n)
		newNotes := []types.AttributeValue{&types.AttributeValueMemberS{Value: note}}

		// Execute the UpdateItem operation
		_, err := db.UpdateItem(context.TODO(), &dynamodb.UpdateItemInput{
			TableName: aws.String(TABLE_TRACKER),
			Key: map[string]types.AttributeValue{
				"Serial": &types.AttributeValueMemberS{Value: inSerial},
			},
			UpdateExpression: aws.String("SET #listField = list_append(#listField, :newNote)"),
			ExpressionAttributeNames: map[string]string{
				"#listField": "notes",
			},
			ExpressionAttributeValues: map[string]types.AttributeValue{
				":newNote": &types.AttributeValueMemberL{Value: newNotes},
			},
			// ReturnValues: "ALL_NEW", // Return the updated item
		})

		if err != nil {
			return "", err
		}
	}

	return firstNote, nil
}

func getEmployeeFromPin(inPin string) (string, error) {

	employee := ""

	expr, err := expression.NewBuilder().WithFilter(
		expression.Contains(expression.Name("pin"), inPin)).Build()

	if err != nil {
		return employee, err
	}

	out, err := db.Scan(context.TODO(), &dynamodb.ScanInput{
		TableName:                 aws.String(TABLE_TRACKER),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		FilterExpression:          expr.Filter(),
	})

	var em EmployeeDynamo

	if len(out.Items) == 1 {
		err = attributevalue.UnmarshalListOfMaps(out.Items, &em)
		if err != nil {
			return employee, err
		}
		employee = em.Name
	}
	return employee, err
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
