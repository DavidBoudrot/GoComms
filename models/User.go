package models

type User struct {
	ID       int
	Name     string
	Email    string
	Password string
}

func NewUser(ID int, name string, email string, password string) *User {
	return &User{ID: ID, Name: name, Email: email, Password: password}
}

func GetUser(ID int) *User {
	// TODO: Get user from database
	return nil
}

func (u *User) GetID() int {
	return u.ID
}

func (u *User) GetName() string {
	return u.Name
}

func (u *User) GetEmail() string {
	return u.Email
}

func (u *User) GetPassword() string {
	return u.Password
}

func (u *User) SetId(ID int) {
	u.ID = ID
}

func (u *User) SetName(name string) {
	u.Name = name
}

func (u *User) SetEmail(email string) {
	u.Email = email
}

func (u *User) SetPassword(password string) {
	u.Password = password
}

// TODO: Implement save() function to save user to database

//TODO: Implement RegisterUser() and LoginUser() functions
