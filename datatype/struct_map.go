package main

import (
	"fmt"

	//	"encoding/json"

	"github.com/fatih/structs"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/mitchellh/mapstructure"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type User struct {
	Name string
	Age  int
	Date *timestamp.Timestamp
}

func ConvertStructToMap(s interface{}) (map[string]interface{}, error) {
	//data, _ := json.Marshal(s)
	//mt.Println(data)
	//user := User{}
	//json.Unmarshal(data, &user)
	//fmt.Println(user)
	return structs.Map(s), nil
}

func ConvertMapToStruct(m map[string]interface{}) (interface{}, error) {
	user := User{}
	mapstructure.Decode(m, &user)

	return user, nil
}
func main() {
	user := User{"minh", 18, timestamppb.Now()}
	m, err := ConvertStructToMap(user)
	fmt.Println(fmt.Sprintf("%T", m["Age"]))
	fmt.Println(fmt.Sprintf("%T", m["Name"]))
	fmt.Println(fmt.Sprintf("%T", m["Date"]))
	if err != nil {
		fmt.Print("Converting Error!")
	}
	//fmt.Println(m)

	fmt.Println(ConvertMapToStruct(m))
}
