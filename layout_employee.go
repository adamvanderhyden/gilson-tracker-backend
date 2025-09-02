package main

type EmployeeJson struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
	Pin  string `json:"pin"`
}
type EmployeeDynamo struct {
	Id   int64  `dynamodbav:"id"`
	Name string `dynamodbav:"name"`
	Pin  string `dynamodbav:"pin"`
}
