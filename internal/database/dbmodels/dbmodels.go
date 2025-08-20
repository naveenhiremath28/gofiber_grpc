package dbmodels

import "time"

type User struct {
    ID        int       `gorm:"primaryKey;autoIncrement" json:"id"`
    Username  string    `gorm:"size:50;not null;unique" json:"username"`
    Email     string    `gorm:"size:100;not null;unique" json:"email"`
    FullName  string    `gorm:"size:100" json:"full_name"`
    CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
    UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (User) TableName() string {
    return "users"
}