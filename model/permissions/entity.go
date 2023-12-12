package permissions

import (
	"strconv"
	"time"
)

type Permissions struct {
	ID        int        `gorm:"primaryKey" json:"id"`
	CreatedAt *time.Time `gorm:"column:created_at;default:now()" json:"created_at"`
	UpdatedAt *time.Time `gorm:"column:updated_at;default:now()" json:"updated_at"`
	Name      string     `json:"name"`
}

func (m *Permissions) TableName() string {
	return "permissions"
}

func (e *Permissions) RedisKey() string {
	if e.ID == 0 {
		return "permissions"
	}

	return "permissions:" + strconv.Itoa(e.ID)
}
