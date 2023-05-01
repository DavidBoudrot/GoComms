package repositories

import "messenger/models"

type MessageRepository struct {
	// Add any necessary dependencies here
}

func (r *MessageRepository) GetMessage(id int) *models.Message {
	// TODO: Retrieve message with specified ID from database
	return nil
}

func (r *MessageRepository) SaveMessage(message *models.Message) {
	// TODO: Save message to database
}

func (r *MessageRepository) DeleteMessage(id int) {
	// TODO: Delete message with specified ID from database
}
