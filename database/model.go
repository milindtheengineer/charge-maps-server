package database

type User struct {
	Email string
	Name  string
}

type UserRow struct {
	Id int
	User
}
