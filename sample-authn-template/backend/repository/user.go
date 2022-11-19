package repository

type User struct{
	Id int64 `json"id"`
	Email string `json:"email"`
	password string `json:"password"`
	Secret string `json:"secret"`
}

type UserRepository interface { // repository Port
	CreateUser(email string, password string, secret string) (*User, error) //as pointer
	CheckUser(email string) (*User, error)
	GetUsers() ([]*User, error)
}