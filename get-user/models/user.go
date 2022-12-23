package models

type UserModel struct {
	Uuid         string   `json:"uuid" dynamodbav:"uuid"`
	FirstName    string   `json:"firstName" dynamodbav:"firstName"`
	LastName     string   `json:"lastName" dynamodbav:"lastName"`
	Hobbies      []string `json:"hobbies" dynamodbav:"hobbies"`
	CreatedEntry string   `json:"createdEntry" dynamodbav:"createdEntry"`
}
