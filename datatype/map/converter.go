package main

import (
	"fmt"
	"reflect"
	"strings"
)

type GirlFriend struct {
	Name string
	Sex  string
}

type User struct {
	Name       string
	Age        int
	Address    string
	sex        string
	School     string
	GirlFriend GirlFriend
}

func structToMap(obj interface{}) (*map[string]interface{}, error) {
	v := reflect.ValueOf(obj)
	f := v.Type()
	result := make(map[string]interface{})
	for i := 0; i < v.NumField(); i++ {
		fieldName := f.Field(i).Name
		if string(fieldName[0]) == strings.ToLower(string(fieldName[0])) {
			continue
		}
		value := v.Field(i).Interface()
		if value == nil || value == "" {
			continue
		}
		fmt.Printf("field: %s value: %v \n", f.Field(i).Name, v.Field(i).Interface())
		result[strings.ToLower(fieldName)] = value
	}

	return &result, nil
}

func main() {
	gf := GirlFriend{
		Name: "huhu",
		Sex:  "female",
	}
	user := User{
		Name:       "kumin",
		Age:        25,
		Address:    "HCM",
		GirlFriend: gf,
	}

	fmt.Println(structToMap(user))
}
