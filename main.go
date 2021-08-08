package main

import (
	"fmt"
	"net/http"

	"github.com/golangwebapp/controllers"
)

type HTTPRequest struct {
	Method string
}

func main() {
	fmt.Println("Starting web server...")
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}
