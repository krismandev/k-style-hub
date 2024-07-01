package model

type AuthLog struct {
	ID         int     `gorm:"id;primaryKey"`
	UserID     int     `gorm:"user_id"`
	LoggedInAt *string `gorm:"logged_in_at"`
	User       User    `gorm:"foreignKey:UserID;references:ID"`
}
