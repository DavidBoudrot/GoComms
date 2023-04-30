package models

type User struct {
	id       int
	name     string
	email    string
	password string
}

func NewUser(id int, name string, email string, password string) *User {
	return &User{id: id, name: name, email: email, password: password}
}

func GetUser(id int) *User {
	// TODO: Get user from database
	return nil
}

func (u *User) GetId() int {
	return u.id
}

func (u *User) GetName() string {
	return u.name
}

func (u *User) GetEmail() string {
	return u.email
}

func (u *User) GetPassword() string {
	return u.password
}

func (u *User) SetId(id int) {
	u.id = id
}

func (u *User) SetName(name string) {
	u.name = name
}

func (u *User) SetEmail(email string) {
	u.email = email
}

func (u *User) SetPassword(password string) {
	u.password = password
}

// TODO: Implement save() function to save user to database

//TODO: Implement RegisterUser() and LoginUser() functions
