package main

type User struct {
	ID        int
	FirstName string
	LastName  string
}

func main() {
	u1 := User{
		ID:        1,
		FirstName: "Arthur",
		LastName:  "Dent",
	}
	u2 := User{
		ID:        2,
		FirstName: "Ford",
		LastName:  "Prefect",
	}
	if u1.ID == u2.ID {
		print("Not same user!")
	} else {
		println("Different Users!")
	}
}
