package main

import (
	"fmt"
	"struct/management"
)

func main() {

	user1 := management.User{}
	user1.Id = 1
	user1.FirstName = "Aldi"
	user1.LastName = "Arif Setiawan"
	user1.Email = "aldhykohar@stimednp.ac.id"
	user1.IsActive = true

	user2 := management.User{
		Id:        2,
		FirstName: "Aldhy",
		LastName:  "Kohar",
		Email:     "aldhykohar@gmailcom",
		IsActive:  false,
	}

	fmt.Println(user1)
	fmt.Println(user2)

	fmt.Println("=================================")

	//struct as parameter
	displayUser1 := displayUser(user1)
	displayUser2 := displayUser(user2)

	fmt.Println(displayUser1)
	fmt.Println(displayUser2)

	fmt.Println("=================================")

	// Embedded struct
	users := []management.User{user1, user2}
	group := management.Group{"Game", user1, users, true}
	displayGroup(group)

	fmt.Println("=================================")

	//using embedded
	result := user1.Display()
	fmt.Println(result)
	fmt.Println(user2.Display())
	group.Display()
}

// struct as parameter
func displayUser(user management.User) string {
	return fmt.Sprintf("Name : %s %s, Email : %s", user.FirstName, user.LastName, user.Email)

}

func displayGroup(group management.Group) {
	fmt.Printf("Name : %s", group.Name)
	fmt.Println("")
	fmt.Printf("Member Count : %d", len(group.Users))
	fmt.Println("")

	fmt.Println("Users name")
	for _, user := range group.Users {
		fmt.Println(user.FirstName)
	}
}
