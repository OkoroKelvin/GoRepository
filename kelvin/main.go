package main

import (
	"awesomeProject/models"
	"fmt"
)

func main() {
	u := models.User{
		ID:        2,
		FirstName: "Tricia",
		LastName:  "McMilan",
	}
	fmt.Println(u)
}
