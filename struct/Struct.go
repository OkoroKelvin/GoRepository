package main

import "fmt"

func main() {
	type user struct {
		ID        int
		FirstName string
		LastName  string
	}
	var u user
	u.ID = 1
	u.FirstName = "Arthur"
	u.LastName = "Dent"
	fmt.Println(u)

	var u2 user
	u2.ID = 2
	u2.FirstName = "Henry"
	u2.LastName = "Chigozie"
	fmt.Println(u2)

	u3 := user{ID: 1, FirstName: "Arthur", LastName: "Dent"}
	fmt.Println(u3)
}
