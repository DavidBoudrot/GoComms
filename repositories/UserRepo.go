package repositories

import (
	"database/sql"
	"messenger/models"
)

type userRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) UserRepo {
	return &userRepo{db}
}

func (r *userRepo) GetUser(id int) *models.User {
	// TODO: implement logic to get user from database using r.db
	return nil
}

func (r *userRepo) SaveUser(user *models.User) {
	// TODO: implement logic to save user to database using r.db
}

func (r *userRepo) DeleteUser(id int) {
	// TODO: implement logic to delete user from database using r.db
}
