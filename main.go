package main

import (
	"ktr/migrations"
	"ktr/router"
)

func main(){
	migrations.Migrate()


	router.Serve()
}