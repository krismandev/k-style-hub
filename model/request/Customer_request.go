package request

type RegisterCustomerRequest struct {
	FirstName   string `json:"first_name" validate:"required,max=100,min=1"`
	LastName    string `json:"last_name" validate:"max=100"`
	Email       string `json:"email" validate:"required,email"`
	Gender      string `json:"gender" validate:"required,oneof=L P"`
	Password    string `json:"password" validate:"required,min=6"`
	PhoneNumber string `json:"phone_number"`
}

type UpdateCustomerRequest struct {
	CustomerID  int    `json:"customer_id" validate:"required"`
	FirstName   string `json:"first_name" validate:"required,max=100,min=1"`
	LastName    string `json:"last_name" validate:"max=100"`
	Gender      string `json:"gender" validate:"required,oneof=L P"`
	PhoneNumber string `json:"phone_number"`
}

type DeleteCustomerRequest struct {
	CustomerID int `json:"customer_id" validate:"required"`
}

type GetCustomerRequest struct {
	Name  string         `json:"name"`
	Email string         `json:"email"`
	Param DataTableParam `json:"param"`
}
