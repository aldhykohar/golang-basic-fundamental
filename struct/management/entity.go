package management

import "fmt"

type User struct {
	Id        int
	FirstName string
	LastName  string
	Email     string
	IsActive  bool
}

// using method
func (user User) Display() string {
	return fmt.Sprintf("Name : %s %s, Email : %s", user.FirstName, user.LastName, user.Email)
}

type Group struct {
	Name        string
	Admin       User
	Users       []User
	IsAvailable bool
}

// using method
func (group Group) Display() {
	fmt.Printf("Name : %s", group.Name)
	fmt.Println("")
	fmt.Printf("Member Count : %d", len(group.Users))
	fmt.Println("")

	fmt.Println("Users name")
	for _, user := range group.Users {
		fmt.Println(user.FirstName)
	}
}
