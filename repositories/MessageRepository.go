package repositories

import (
	"database/sql"
	"messenger/models"
)

type MessageRepository struct {
	db *sql.DB
}

func NewMessageRepository(db *sql.DB) *MessageRepository {
	return &MessageRepository{db}
}

func (r *MessageRepository) GetMessage(id int) (*models.Message, error) {
	messageModel := &models.Message{}

	err := r.db.QueryRow("SELECT ID, senderID, receiverID, message FROM messages WHERE id = ?", id).
		Scan(&messageModel.ID, &messageModel.SenderID, &messageModel.ReceiverID, &messageModel.Message)

	if err != nil {
		return nil, err
	}
	return messageModel, nil
}

func (r *MessageRepository) SaveMessage(message *models.Message) error {
	_, err := r.db.Exec("INSERT INTO messages(senderID, receiverID, message) VALUES (?, ?, ?, ?)",
		message.GetSenderId(), message.GetReceiverId(), message.GetReceiverId(), message.GetMessage())
	if err != nil {
		return err
	}
	return nil
}

func (r *MessageRepository) DeleteMessage(id int) error {
	_, err := r.db.Exec("DELETE FROM messages WHERE id = ?", id)
	if err != nil {
		return err
	}
	return nil
}
