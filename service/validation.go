package service

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

func Digit(fl validator.FieldLevel) bool {
	field := fl.Field().String()
	for _, ch := range field {
		if ch >= '0' && ch <= '9' {
			return true
		}
	}
	return false
}

func LowerCase(fl validator.FieldLevel) bool {
	field := fl.Field().String()
	for _, ch := range field {
		if ch >= 'a' && ch <= 'z' {
			return true
		}
	}
	return false
}

func UpperCase(fl validator.FieldLevel) bool {
	field := fl.Field().String()
	for _, ch := range field {
		if ch >= 'A' && ch <= 'Z' {
			return true
		}
	}
	return false
}

func SpecialCase(fl validator.FieldLevel) bool {
	field := fl.Field().String()
	specail := "!@#$%^&*"
	for _, ch := range field {
		for _, sp := range specail {
			if ch == sp {
				return true
			}
		}
	}
	return false
}

func IsValid(user interface{}) bool {
	v := validator.New()
	e1 := v.RegisterValidation("digit", Digit)
	if e1 != nil {
		fmt.Println(e1.Error())
	}
	e2 := v.RegisterValidation("lowercase", LowerCase)
	if e2 != nil {
		fmt.Println(e2.Error())
	}
	e3 := v.RegisterValidation("uppercase", UpperCase)
	if e3 != nil {
		fmt.Println(e3.Error())
	}
	e4 := v.RegisterValidation("special", SpecialCase)
	if e4 != nil {
		fmt.Println(e3.Error())
	}
	err := v.Struct(user)
	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			fmt.Println(e)
		}
		return false
	}
	return true
}
