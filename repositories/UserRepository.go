package repositories

import (
	"database/sql"
	"messenger/models"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db}
}

func (r *MessageRepository) GetUser(id int) (*models.User, error) {
	userModel := &models.User{}
	err := r.db.QueryRow("SELECT ID, name, email, password FROM messages WHERE id = ?", id).
		Scan(&userModel.ID, &userModel.Name, &userModel.Email, &userModel.Password)

	if err != nil {
		return nil, err
	}
	return userModel, nil
}

func (r *MessageRepository) SaveUser(message *models.Message) error {
	_, err := r.db.Exec("INSERT INTO users(name, email, password) VALUES (?, ?, ?)",
		message.GetSenderId(), message.GetReceiverId(), message.GetReceiverId(), message.GetMessage())
	if err != nil {
		return err
	}
	return nil
}

func (r *MessageRepository) DeleteUser(id int) error {
	_, err := r.db.Exec("DELETE FROM  users WHERE id = ?", id)
	if err != nil {
		return err
	}
	return nil
}
