package router

import (
	"ktr/controllers"
	"net/http"
)

func init(){
	route("/", controllers.Home)

	route("/about", controllers.About)


}
func Serve(){

	_ = http.ListenAndServe(":8000",nil)
}
func route(route string,callback func(w http.ResponseWriter ,r *http.Request)){


	http.HandleFunc(route,callback)

}