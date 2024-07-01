package response

type LoginResponse struct {
	UserID     int    `json:"user_id,omitempty"`
	Email      string `json:"email"`
	CustomerID int    `json:"customer_id,omitempty"`
	Token      string `json:"token"`
	ExpiredAt  string `json:"expired_at"`
}
