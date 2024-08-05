package po

import (
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	ID       int64  `gorm:"column:id; type:int; not null; primary_key; auto_increment; comment:'Primary Key is ID"`
	RoleName string `gorm:"column:role_name; type:varchar(255);"`
	RoleDesc string `gorm:"column:role_desc; type:text;"`
}

func (r *Role) TableName() string {
	return "roles"
}
