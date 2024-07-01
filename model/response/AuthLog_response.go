package response

import "k-style-test/model"

type AuthLogResponse struct {
	UserID     int    `json:"user_id,omitempty"`
	Email      string `json:"email"`
	LoggedInAt string `json:"logged_in_at"`
}

func ToAuthLogResponse(data *model.AuthLog) AuthLogResponse {
	var resp AuthLogResponse

	resp.UserID = data.UserID
	resp.Email = data.User.Email
	resp.LoggedInAt = *data.LoggedInAt
	return resp
}
