package main

import (
	"fmt"
	"reflect"
)

type order struct {
	ordId      int
	customerId int
}

type employee struct {
	name    string
	id      int
	address string
	salary  int
	country string
}

func createQuery(q interface{}) {
	//fmt.Println("Struct datas:", reflect.ValueOf(q))
	//fmt.Println("Struct name type:", reflect.TypeOf(q).Name())
	//Check the type of argument is struct
	if reflect.ValueOf(q).Kind() == reflect.Struct {
		//get the type of struct
		t := reflect.TypeOf(q).Name()
		//init query statement
		query := fmt.Sprintf("insert into %s(", t)
		v := reflect.ValueOf(q)
		var fieldName string
		//generate query field
		for j := 0; j < v.NumField(); j++ {
			fieldName = v.Type().Field(j).Name
			if j == 0 {
				query = fmt.Sprintf("%s%s", query, fieldName)
			} else {
				query = fmt.Sprintf("%s, %s", query, fieldName)
			}
		}
		query = fmt.Sprintf("%s) values(", query)
		//generate query field value
		for i := 0; i < v.NumField(); i++ {
			switch v.Field(i).Kind() {
			case reflect.Int:
				if i == 0 {
					query = fmt.Sprintf("%s%d", query, v.Field(i).Int())
				} else {
					query = fmt.Sprintf("%s, %d", query, v.Field(i).Int())
				}
			case reflect.String:
				if i == 0 {
					query = fmt.Sprintf("%s\"%s\"", query, v.Field(i).String())
				} else {
					query = fmt.Sprintf("%s, \"%s\"", query, v.Field(i).String())
				}
			default:
				fmt.Println("Unsupported type")
				return
			}
		}
		query = fmt.Sprintf("%s)", query)
		fmt.Println("Full query:", query)
		return
	}
	fmt.Println("Unsupported type")
}

func main() {
	o := order{
		ordId:      123,
		customerId: 54,
	}
	createQuery(o)

	e := employee{
		name:    "hieunguyen",
		id:      543,
		address: "HaNoi",
		salary:  439232,
		country: "VietNam",
	}
	createQuery(e)
	createQuery(90)
}
