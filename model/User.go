package model

type User struct {
	ID        int    `gorm:"column:id;primaryKey"`
	Email     string `gorm:"column:email;unique"`
	Password  string `gorm:"column:password"`
	CreatedAt string `gorm:"column:created_at;default:CURRENT_TIMESTAMP()"`
	UpdateAt  string `gorm:"column:updated_at;default:CURRENT_TIMESTAMP()"`
}
