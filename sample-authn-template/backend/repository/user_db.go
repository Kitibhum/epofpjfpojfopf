package repository

import (
	"database/sql"
	"fmt"
)

type userRepositoryDB struct {
	db *sql.DB
}

func NewRepositoryDB(db *sql.DB) userRepositoryDB { // รับ instance database ที่เราจะใช้มาแล้วออกมาเป็น "Adapter" ของ repository
	return userRepositoryDB{db: db}
}

func (u userRepositoryDB) CreateUser(email string, password string, secret string) (*User, error) {
	// implement me
	fmt.Println("User Inserted")
	return nil, nil
}
func (u userRepositoryDB) CheckUser(email string) (*User, error) {

	var user User
	err := u.db.QueryRow("SELECT id, email, password, secret FROM users where email = ?", email).Scan(&user)
	
	fmt.Println("User Fetched")
	return nil, nil
}

func (u userRepositoryDB) GetUsers() ([]*User, error) {
	// implement me
	fmt.Println("Users Fetched")
	return nil, nil
}
