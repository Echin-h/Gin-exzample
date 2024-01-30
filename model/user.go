package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

// 注意：所有字段的零值, 比如0, "",false或者其它零值，都不会保存到数据库内，但会使用他们的默认值。
//如果你想避免这种情况，可以考虑使用指针或实现 Scanner/Valuer接口

type User struct {
	gorm.Model
	Name      string    `gorm:"varchar(20);not null" json:"name"`     //姓名
	Gender    *string   `gorm:"char" json:"gender"`                   //性别
	Age       int       `json:"age"`                                  //年龄
	Birthday  time.Time `gorm:"not null" json:"birthday"`             //出生日期
	Telephone *string   `gorm:"varchar(20)" json:"telephone"`         //电话
	Password  string    `gorm:"varchar(20);not null" json:"password"` //密码
	Email     string    `gorm:"varchar(30);not null" json:"email"`    //邮箱
}
