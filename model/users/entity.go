package users

import (
	"strconv"
	"time"
)

type Users struct {
	ID              int        `gorm:"primaryKey" json:"id"`
	CreatedAt       *time.Time `gorm:"column:created_at;default:now()" json:"created_at"`
	UpdatedAt       *time.Time `gorm:"column:updated_at;default:now()" json:"updated_at"`
	RoleId          *int       `json:"role_id"`
	Name            string     `json:"name"`
	Username        string     `json:"username"`
	Email           string     `json:"email"`
	EmailVerifiedAt *time.Time `gorm:"column:email_verified_at;default:now()" json:"email_verified_at"`
	Password        string     `json:"password"`
	RememberToken   string     `json:"remember_token"`
	FotoProfil      string     `json:"foto_profil"`
}

func (m *Users) TableName() string {
	return "users"
}

func (e *Users) RedisKey() string {
	if e.ID == 0 {
		return "users"
	}

	return "users:" + strconv.Itoa(e.ID)
}
