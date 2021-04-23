package management

type User struct {
	ID        int
	FirstName string
	LastNamse string
	email     string
	isActive  bool
}

type Group struct {
	Name        string
	Admin       User
	Users       []User
	isAvailable bool
}
