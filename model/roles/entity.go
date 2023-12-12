package roles

import (
	"strconv"
	"time"
)

type Roles struct {
	ID        int        `gorm:"primaryKey" json:"id"`
	CreatedAt *time.Time `gorm:"column:created_at;default:now()" json:"created_at"`
	UpdatedAt *time.Time `gorm:"column:updated_at;default:now()" json:"updated_at"`
	Name      string     `json:"name"`
}

func (m *Roles) TableName() string {
	return "roles"
}

func (e *Roles) RedisKey() string {
	if e.ID == 0 {
		return "roles"
	}

	return "roles:" + strconv.Itoa(e.ID)
}
