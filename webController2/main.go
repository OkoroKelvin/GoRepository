package main

import (
	"awesomeProject/controllers"
	"net/http"
)

func main() {
	controllers.RegisterController()
	http.ListenAndServe(":3000", nil)

}
