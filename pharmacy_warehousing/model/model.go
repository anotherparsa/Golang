package model

import "errors"

type Staff struct {
	Id       string
	Name     string
	Family   string
	Staffid  string
	Userid   string
	Position string
	Password string
}

type Drug struct {
	Id      string
	Name    string
	Drugid  string
	Company string
	Price   string
	Stock   string
}

var User_is_ot_authorized error = errors.New("User is not authorized")
var Cookie_doesnt_exist error = errors.New("Cookie doesn't exist")

