package models

type UserModel struct {
	FirstName string `json:"firstName,omitempty" dynamodbav:"firstName,omitempty"`
	LastName  string `json:"lastName,omitempty" dynamodbav:"lastName,omitempty"`
}
