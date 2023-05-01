package repositories

import (
	"messenger/models"
)

type UserRepo interface {
	GetUser(id int) *models.User
	SaveUser(user *models.User)
	DeleteUser(id int)
}

type MessageRepo interface {
	GetMessage(id int) *models.Message
	SaveMessage(message *models.Message)
	DeleteMessage(id int)
}
