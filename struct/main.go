package main

import (
	"struct/management"
)

// func (group Group) DisplayGroup() {
// 	fmt.Printf("Name: %s", group.Name)
// 	fmt.Println("")
// 	fmt.Printf("Member Count: %d", len(group.Users))

// 	fmt.Println("Users Name: ")

// 	for _, user := range group.Users {
// 		fmt.Println(user.FirstName)
// 	}
// }

// func (user User) display() string { //! Struct method
// 	result := fmt.Sprintf("Name: %s %s, Email: %s", user.FirstName, user.LastNamse, user.email)
// 	return result
// }

// ! perbedaan Function dengan method
// ? Function melekat dengan package dan tidak di miliki olah siapa siapa, function memiliki parameter
// ? Method di miliki olah suatu struct dan tidak memiliki parameter

func main() {
	// user := User{
	// 	isActive:  true,
	// 	ID:        1,
	// 	FirstName: "Titan",
	// 	LastNamse: "Titann",
	// 	email:     "tasst@mail.com",
	// }

	user := management.User{ //! Harus terurut seperti di Struct
		1,
		"Titan",
		"Titann",
		"tasst@mail.com",
		true,
	}

	// result := user.display()
	// fmt.Println(result)
	user2 := management.User{ //! Harus terurut seperti di Struct
		3,
		"Titanio",
		"yudista",
		"tasst@mail.com",
		true,
	}

	// displayUser1 := displayUser(user)
	// displayUser2 := displayUser(user2)
	// user.ID = 1
	// user.FirstName = "Titanio"
	// user.LastNamse = "Yudista"
	// user.email = "ttt@mail.com"
	// user.isActive = true

	// user2 := User{}
	// user2.ID = 2
	// user2.FirstName = "bui"
	// user2.LastNamse = "budi"
	// user2.email = "ttt@mail.com"
	// user2.isActive = true
	// users := []management.User{user, user2} //!embedded struct
	// group := management.Group{"Gamer", user, users, true}
	// group.DisplayGroup()

	// fmt.Println(displayUser1)
	// fmt.Println(displayUser2)
	// fmt.Println(user2.LastNamse)
	// displayGroup(group)
}

// func displayUser(user User) string {
// 	result := fmt.Sprintf("Name: %s %s, Email: %s", user.FirstName, user.LastNamse, user.email)
// 	return result
// }
