package db

import (
	"database/sql"
)
import _ "github.com/go-sql-driver/mysql"




type Database struct {
	q *sql.DB
}


func (db *Database) open(){
	connected,err := sql.Open("mysql", "ktr_user:123123@/ktr")
	if err != nil{
		panic(err)
	}
	db.q = connected

}

func (db *Database) close(){
	_ = db.q.Close()

}


func Query(statement string) *sql.Rows{

	db := Database{}
	db.open()
	query,err := db.q.Query(statement)
	db.close()

	if err != nil{
		panic(err)
	}
	return query


}


