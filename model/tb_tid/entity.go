package tb_tid

import (
	"strconv"
	"time"
)

type TbTid struct {
	ID         int       `gorm:"primaryKey" json:"id"`
	CreatedAt  time.Time `gorm:"column:created_at;default:now()" json:"created_at"`
	UpdatedAt  time.Time `gorm:"column:updated_at;default:now()" json:"updated_at"`
	Tid        string    `json:"tid"`
	IpAddress  string    `json:"ip_address"`
	SnMiniPc   string    `json:"sn_mini_pc"`
	LocationId *int      `json:"location_id"`
}

func (m *TbTid) TableName() string {
	return "tb_tid"
}

func (e *TbTid) RedisKey() string {
	if e.ID == 0 {
		return "tb_tid"
	}

	return "tb_tid:" + strconv.Itoa(e.ID)
}
