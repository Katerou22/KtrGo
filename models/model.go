package models

import (
	"fmt"
	"ktr/db"
	"reflect"
	"strings"
)

type Model interface {
	Create()
}

type User struct {
	Id       string // for some reasons cant take it to uint64
	Username string
	Password string

}

func Create(model interface{}) interface{}{
fmt.Println("Working")
	columns := ""
	values := ""

	attributes := getAttributes(model)

	for column,value := range  attributes {
		columns += column + ","
		values += "'" + value.(string) + "'" + ","
	}
	columns = trim(columns,",")
	values = trim(values,",")

	schema := "INSERT INTO " +  getNameOfStruct(model) + " ("+ columns +") VALUES ("+values+")"
	query := db.Query(schema)
	for query.Next(){
		var (
			id int
			username string
			password string
		)
		err := query.Scan(&id,&username,&password)
		if err != nil {
			panic(err)
		}
		fmt.Println(id,username,password)
	}

	return model


}



func getAttributes(model interface{}) map[string] interface{} {
	e := reflect.Indirect(reflect.ValueOf(model))


	attributes := make(map[string]interface{})



	for i := 0; i < e.NumField(); i++ {
		name := e.Type().Field(i).Name
		value := e.Field(i).Interface()
		if name != "Id" {
			attributes[strings.ToLower(name)] = value

		}


	}
	return attributes
}

func getNameOfStruct(model interface{}) string {
 	name := ""
	if t := reflect.TypeOf(model); t.Kind() == reflect.Ptr {
		name = "*" + t.Elem().Name()
	} else {
		name = t.Name()
	}
	return  strings.ToLower(name + "s") // Not the right plural table name, but it okay with simple models.
}

func trim(s, suffix string) string {
	if strings.HasSuffix(s, suffix) {
		s = s[:len(s)-len(suffix)]
	}
	return s
}