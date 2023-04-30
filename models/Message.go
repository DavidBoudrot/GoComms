package models

type Message struct {
	id         int
	senderId   int
	receiverId int
	message    string
}

func CreateMessage(id int, senderId int, receiverId int, message string) *Message {
	return &Message{id: id, senderId: senderId, receiverId: receiverId, message: message}
}

func (m *Message) GetId() int {
	return m.id
}

func (m *Message) GetSenderId() int {
	return m.senderId
}

func (m *Message) GetReceiverId() int {
	return m.receiverId
}

func (m *Message) GetMessage() string {
	return m.message
}

func (m *Message) SetId(id int) {
	m.id = id
}

func (m *Message) SetSenderId(senderId int) {
	m.senderId = senderId
}

func (m *Message) SetReceiverId(receiverId int) {
	m.receiverId = receiverId
}

func (m *Message) SetMessage(message string) {
	m.message = message
}
