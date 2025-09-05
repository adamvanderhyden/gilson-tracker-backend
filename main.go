package main

import (
	"context"
	"encoding/json"
	"fmt"
	"sort"
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
	TABLE_TRACKER  = "Tracker"
	TABLE_EMPLOYEE = "Employee"
	PAR_SERIAL     = "serial"
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

		case "hold":

			expr, bErr := expression.NewBuilder().WithFilter(
				expression.And(
					expression.Name("deleted").Equal(expression.Value("")),
					expression.Name("finished").Equal(expression.Value("")),
					expression.Name("started").Equal(expression.Value("")))).Build()

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

			return getTrackerRoom(expr, d.Room)

		case "sidewalls":

			expr, bErr := expression.NewBuilder().WithFilter(
				expression.And(
					expression.Name("deleted").Equal(expression.Value("")),
					expression.Name("finished").Equal(expression.Value("")),
					expression.Name("started").NotEqual(expression.Value("")),
					expression.Or( // finished holds a date
						expression.Name("side_cnc").Equal(expression.Value("")),
						expression.Name("side_pour").Equal(expression.Value(""))))).Build()

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

			return getTrackerRoom(expr, d.Room)

		case "woodshop":

			expr, bErr := expression.NewBuilder().WithFilter(
				expression.And(
					expression.Name("deleted").Equal(expression.Value("")),
					expression.Name("finished").Equal(expression.Value("")),
					expression.Name("started").NotEqual(expression.Value("")),
					expression.Name("side_cnc").NotEqual(expression.Value("")),
					expression.Name("side_pour").NotEqual(expression.Value("")),
					expression.Or( // finished holds a date
						expression.Name("wood_core").Equal(expression.Value("")),
						expression.Name("wood_cart").Equal(expression.Value(""))))).Build()

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

			return getTrackerRoom(expr, d.Room)

		case "sublimation":

			expr, bErr := expression.NewBuilder().WithFilter(
				expression.And(
					expression.Name("deleted").Equal(expression.Value("")),
					expression.Name("finished").Equal(expression.Value("")),
					expression.Name("started").NotEqual(expression.Value("")),
					expression.Name("side_cnc").NotEqual(expression.Value("")),
					expression.Name("side_pour").NotEqual(expression.Value("")),
					expression.Or( // finished holds a date
						expression.Name("sub_base").Equal(expression.Value("")),
						expression.Name("sub_top").Equal(expression.Value("")),
						expression.Name("sub_cart").Equal(expression.Value(""))))).Build()

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

			return getTrackerRoom(expr, d.Room)

		case "edges":

			expr, bErr := expression.NewBuilder().WithFilter(
				expression.And(
					expression.Name("deleted").Equal(expression.Value("")),
					expression.Name("finished").Equal(expression.Value("")),
					expression.Name("started").NotEqual(expression.Value("")),
					expression.Name("sub_base").NotEqual(expression.Value("")),
					expression.Name("sub_top").NotEqual(expression.Value("")),
					expression.Name("sub_cart").NotEqual(expression.Value("")),
					expression.Or( // finished holds a date
						expression.Name("ed_glue").Equal(expression.Value("")),
						expression.Name("ed_cart").Equal(expression.Value(""))))).Build()

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

			return getTrackerRoom(expr, d.Room)

		case "cart":

			expr, bErr := expression.NewBuilder().WithFilter(
				expression.And(
					expression.Name("deleted").Equal(expression.Value("")),
					expression.Name("finished").Equal(expression.Value("")),
					expression.Name("started").NotEqual(expression.Value("")),
					expression.Name("sub_base").NotEqual(expression.Value("")),
					expression.Name("sub_top").NotEqual(expression.Value("")),
					expression.Name("sub_cart").NotEqual(expression.Value("")),
					expression.Or( // finished holds a date
						expression.Name("wood_cart").Equal(expression.Value("")),
						expression.Name("sub_cart").Equal(expression.Value("")),
						expression.Name("ed_cart").Equal(expression.Value(""))),
					expression.Or(
						expression.Name("wood_cart").NotEqual(expression.Value("")),
						expression.Name("sub_cart").NotEqual(expression.Value("")),
						expression.Name("ed_cart").NotEqual(expression.Value(""))))).Build()

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

			return getTrackerRoom(expr, d.Room)

		case "layup":

			expr, bErr := expression.NewBuilder().WithFilter(
				expression.And(
					expression.Name("deleted").Equal(expression.Value("")),
					expression.Name("finished").Equal(expression.Value("")),
					expression.Name("started").NotEqual(expression.Value("")),
					expression.Name("wood_cart").NotEqual(expression.Value("")),
					expression.Name("sub_cart").NotEqual(expression.Value("")),
					expression.Name("ed_cart").NotEqual(expression.Value("")),
					expression.Or( // finished holds a date
						expression.Name("lay_press").Equal(expression.Value("")),
						expression.Name("lay_inspect").Equal(expression.Value(""))))).Build()

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

			return getTrackerRoom(expr, d.Room)

		case "finishing":

			expr, bErr := expression.NewBuilder().WithFilter(
				expression.And(
					expression.Name("deleted").Equal(expression.Value("")),
					expression.Name("finished").Equal(expression.Value("")),
					expression.Name("started").NotEqual(expression.Value("")),
					expression.Name("lay_press").NotEqual(expression.Value("")),
					expression.Name("lay_inspect").NotEqual(expression.Value("")),
					expression.Or( // finished holds a date
						expression.Name("fin_tune").Equal(expression.Value("")),
						expression.Name("fin_inspect").Equal(expression.Value(""))))).Build()

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

			return getTrackerRoom(expr, d.Room)

		case "wax":

			expr, bErr := expression.NewBuilder().WithFilter(
				expression.And(
					expression.Name("deleted").Equal(expression.Value("")),
					expression.Name("finished").Equal(expression.Value("")),
					expression.Name("started").NotEqual(expression.Value("")),
					expression.Name("fin_tune").NotEqual(expression.Value("")),
					expression.Name("fin_inspect").NotEqual(expression.Value("")),
					expression.Or( // finished holds a date
						expression.Name("wax_wax").Equal(expression.Value("")),
						expression.Name("wax_inspect").Equal(expression.Value(""))))).Build()

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

			return getTrackerRoom(expr, d.Room)

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

			// pass reponse 'r' into all updates so it possibly retains the new note.
			switch d.Room {
			case "hold":

				// TODO - [rush] could have set/unset/or be blank
				//      - set, make it "rush" for sorting purposes
				//      - unset, make it blank
				//      - blank, do not update
				// TODO - [artist] could have <artist name>/unset/or be blank
				//      - <artist name> (not "unset" and not ""), set input to artist field
				//      - unset, make it blank
				//      - blank, do not update
				// TODO - [started] could have set/blank
				//      - set, make it CCYY-MM-DDTHH:MM:SSZ (EMPLOYEE)
				//      - blank, do not update
				// TODO - [deleted] could have set/blank
				//      - set, make it CCYY-MM-DDTHH:MM:SSZ (EMPLOYEE)
				//      - blank, do not update

				return updateRoom(r)
			case "sidewalls":
			case "woodshop":
			case "sublimation":
			case "edges":
			case "cart":
			case "layup":
			case "finishing":
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

	firstNote := ""

	t := time.Now()
	hh, min, _ := t.Clock()
	yy, mon, dd := t.Date()

	// there should never be more than one note, but it's array based to jive with the
	// rest of the "update" process to handle many
	for _, n := range inNotes {

		// we have notes, check if it's not blank
		if n != "" {

			note := fmt.Sprintf("%02d/%02d/%d %02d:%02d (%s): %s", int(mon), dd, yy, hh, min, inEmployee, n)
			newNotes := []types.AttributeValue{&types.AttributeValueMemberS{Value: note}}

			// save the "first" one... there should only be one.
			if firstNote == "" {
				firstNote = note
			}

			// Execute
			_, err := db.UpdateItem(context.TODO(), &dynamodb.UpdateItemInput{
				TableName: aws.String(TABLE_TRACKER),
				Key: map[string]types.AttributeValue{
					"serial": &types.AttributeValueMemberS{Value: inSerial},
				},
				UpdateExpression: aws.String("SET #listField = list_append(if_not_exists(#listField, :emptyArr), :newNote)"),
				ExpressionAttributeNames: map[string]string{
					"#listField": "notes",
				},
				ExpressionAttributeValues: map[string]types.AttributeValue{
					":emptyArr": &types.AttributeValueMemberL{Value: []types.AttributeValue{}},
					":newNote":  &types.AttributeValueMemberL{Value: newNotes},
				},
			})

			if err != nil {
				fmt.Printf("blew up<%v>\n", err)
				return "", err
			}
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
		TableName:                 aws.String(TABLE_EMPLOYEE),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		FilterExpression:          expr.Filter(),
	})

	var em []EmployeeDynamo

	if len(out.Items) == 1 {

		err = attributevalue.UnmarshalListOfMaps(out.Items, &em)

		if err != nil {
			return employee, err
		}

		employee = em[0].Name
	}
	return employee, err
}

func updateRoom(r Resp) (events.APIGatewayProxyResponse, error) {

	resp, _ := json.MarshalIndent(r, "", "  ")

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers:    HEADERS,
		Body:       string(resp),
	}, nil
}

func getTrackerRoom(inExpr expression.Expression, inRoom string) (events.APIGatewayProxyResponse, error) {

	var r Resp
	r.Error = "good"

	out, err := db.Scan(context.TODO(), &dynamodb.ScanInput{
		TableName:                 aws.String(TABLE_TRACKER),
		ExpressionAttributeNames:  inExpr.Names(),
		ExpressionAttributeValues: inExpr.Values(),
		FilterExpression:          inExpr.Filter(),
	})

	if err != nil {

		r.Error = fmt.Sprintf("Bad Scan (%s): %v", TABLE_TRACKER, err)
		fmt.Println(r.Error)

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
		fmt.Println(r.Error)

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

	if inRoom == "hold" {
		r.TrackerJson = sortRushCreated(r.TrackerJson)
	} else {
		r.TrackerJson = sortRushStarted(r.TrackerJson)
	}

	resp, _ := json.MarshalIndent(r, "", "  ")
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers:    HEADERS,
		Body:       string(resp),
	}, nil
}

func sortRushCreated(tj []TrackerJson) []TrackerJson {

	sort.Slice(tj, func(i, j int) bool {
		if tj[i].Rush != tj[j].Rush {
			return tj[i].Rush > tj[j].Rush
		}
		return tj[i].Created < tj[j].Created
	})
	return tj
}

func sortRushStarted(tj []TrackerJson) []TrackerJson {

	sort.Slice(tj, func(i, j int) bool {
		if tj[i].Rush != tj[j].Rush {
			return tj[i].Rush > tj[j].Rush
		}
		return tj[i].Started < tj[j].Started
	})
	return tj
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
