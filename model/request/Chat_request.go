package request

type GetChatRequest struct {
	ID         int    `json:"id"`
	Message    string `json:"message"`
	From       string `json:"from"`
	To         string `json:"to"`
	IsChatroom int    `json:"is_chatroom"`
	CreatedAt  string `json:"created_at"`
}

type CreateChatRequest struct {
	Message    string `json:"message"`
	From       string `json:"from"`
	To         string `json:"to"`
	IsChatroom int    `json:"is_chatroom"`
}
