package config

import "github.com/gorilla/websocket"

type WebSocketConnection struct {
	*websocket.Conn
	Username string
	Chatroom string
}

type MessagePayload struct {
	From       string `json:"from"`
	To         string `json:"to"`
	Message    string `json:"message"`
	IsChatroom bool   `json:"is_chatroom"`
}
