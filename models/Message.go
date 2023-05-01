package models

type Message struct {
	ID         int
	SenderID   int
	ReceiverID int
	Message    string
}

func CreateMessage(ID int, senderID int, receiverID int, message string) *Message {
	return &Message{ID: ID, SenderID: senderID, ReceiverID: receiverID, Message: message}
}

func (m *Message) GetId() int {
	return m.ID
}

func (m *Message) GetSenderId() int {
	return m.SenderID
}

func (m *Message) GetReceiverId() int {
	return m.ReceiverID
}

func (m *Message) GetMessage() string {
	return m.Message
}

func (m *Message) SetId(ID int) {
	m.ID = ID
}

func (m *Message) SetSenderId(senderID int) {
	m.SenderID = senderID
}

func (m *Message) SetReceiverId(receiverID int) {
	m.ReceiverID = receiverID
}

func (m *Message) SetMessage(message string) {
	m.Message = message
}
