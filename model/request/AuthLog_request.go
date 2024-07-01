package request

type AuthLogRequest struct {
	Name       string         `json:"name"`
	UserID     string         `json:"user_id"`
	Email      string         `json:"email"`
	AuthUserID string         `json:"-"` // authenticated user whose doing this request
	Param      DataTableParam `json:"param"`
}
