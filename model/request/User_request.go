package request

type UserRequest struct {
	ID        int    `json:"id"`
	FullName  string `json:"full_name"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	CreatedAt string `json:"created_at"`
	UpdateAt  string `json:"update_at"`
}

type RegisterUserRequest struct {
	FirstName  string `json:"first_name" validate:"required,max=100,min=1"`
	LastName   string `json:"last_name" validate:"max=100"`
	Email      string `json:"email" validate:"required,email"`
	Gender     string `json:"gender" validate:"required,oneof=L P"`
	Password   string `json:"password" validate:"required,min=6"`
	PhoneNumer string `json:"phone_numer"`
}
