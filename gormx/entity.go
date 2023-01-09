package entity

import (
	"time"
)

// TUser 用户表
type TUser struct {
	Id        uint64    `gorm:"column:id; comment:ID主键"`
	Name      string    `gorm:"column: name; comment: 用户名"`
	Email     string    `gorm:"index:idx_email; column:email; not null; size:256; default:''; comment: 用户邮箱"`
	Sex       int       `gorm:"column:sex; not_null; size:8; default:0; comment:'用户昵称 0:女 1:男'" `
	Birthday  time.Time `gorm:"column:birthday; not_null; comment:'用户出生日期'" `
	CreatedAt time.Time `gorm:"index:idx_created_at;column:created_at"`
	UpdatedAt time.Time `gorm:"index:idx_updated_at;column:updated_at"`
	DeletedAt time.Time `gorm:"index:idx_deleted_at;column:deleted_at"`
}

// TableName DB表
func (t *TUser) TableName() string {
	return "t_users"
}
