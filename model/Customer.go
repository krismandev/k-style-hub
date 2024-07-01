package model

type Customer struct {
	ID          int     `gorm:"id;primaryKey"`
	UserID      int     `gorm:"user_id"`
	FirstName   string  `gorm:"first_name"`
	LastName    *string `gorm:"last_name"`
	PhoneNumber *string `gorm:"phone_number"`
	Gender      string  `gorm:"gender"`
	CreatedAt   *string `gorm:"created_at"`
	UpdatedAt   *string `gorm:"updated_at"`
	DeletedAt   *string `gorm:"deleted_at"`
	User        User    `gorm:"foreignKey:UserID;references:ID"`
}
