package main

import (
	"fmt"

	"github.com/intetunder-temp/go-webservice/models"
)

func main() {
	u := models.User{
		ID:        2,
		FirstName: "Danstar",
		LastName:  "Scarpa",
	}
	fmt.Println(u)
}
