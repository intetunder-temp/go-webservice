package main

import (
	"net/http"

	"github.com/intetunder-temp/go-webservice/controllers"
)

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe("127.0.0.1:3000", nil)
}
