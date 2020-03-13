package migrations

import (
	"fmt"
	"ktr/db"
	"reflect"
	"strings"
)

func Migrate(){

	models := make(map[string]map[string]string)
	// Models add like the example, drop and alter tables not added yet!
	models["users"] = make(map[string]string)
	models["users"]= map[string]string{"id": "int AUTO_INCREMENT", "username":"text","password":"text"}

	var tables []string
	
	query := db.Query("SHOW TABLES;")
	for query.Next(){
		var table string
		_ = query.Scan(&table)
		tables = append(tables,table)
	}
	
	for key, value := range models {
		if ! existsIn(tables,key) {
			createTable(key,value)
		}
	}

}


func createTable(table string,columns map[string]string){
	fmt.Println("Please wait, Migrating table: " + table)

	columnsString := ""

	for key, value := range columns {
		columnsString += key + " " + value + ","
	}

	//columnsString = trim(columnsString,",")

	var schema = "CREATE TABLE "+ table +" ("+ columnsString +" PRIMARY KEY (id));"
	db.Query(schema)
	fmt.Println("Successfully Migrated table: " + table)
}


func trim(s, suffix string) string {
	if strings.HasSuffix(s, suffix) {
		s = s[:len(s)-len(suffix)]
	}
	return s
}

func existsIn(slice interface{}, item interface{}) bool {
	s := reflect.ValueOf(slice)

	if s.Kind() != reflect.Slice {
		panic("Invalid data-type")
	}

	for i := 0; i < s.Len(); i++ {
		if s.Index(i).Interface() == item {
			return true
		}
	}

	return false
}
