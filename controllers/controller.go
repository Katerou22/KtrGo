package controllers

import (
	"fmt"
	"ktr/models"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request){
	user := models.User{Username: "Katerou22", Password: "123123"}
	 models.Create(user)
	//fmt.Println(id)
	fmt.Fprintf(w,"Home ")
}
func About(w http.ResponseWriter, r *http.Request){
	fmt.Print(r.Method)

	fmt.Fprintf(w,"About")
}