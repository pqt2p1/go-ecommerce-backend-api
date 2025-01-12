package po

import (
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	ID       int64  `gorm:"column:id; type:varchar(255); not null; primaryKey; autoIncrement;"`
	RoleName string `gorm:"column:role_name"`
	RoleNote string `gorm:"column:role_note; type:text;"`
}

func (r *Role) TableName() string {
	return "go_db_role"
}
