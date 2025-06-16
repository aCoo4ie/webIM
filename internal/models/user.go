package models

import (
	"time"
)

const (
	SEX_MAX    = "男"
	SEX_WOMAN  = "女"
	SEX_UNKOWN = "未知"
)

type User struct {
	ID        int64     `xorm:"pk autoincr 'id'" json:"id"`   // 主键，自增
	Username  string    `xorm:"varchar(64)" json:"username"`  // 用户名，唯一且不能为空
	Password  string    `xorm:"varchar(128)" json:"password"` // 密码，不能为空
	Mobile    string    `xorm:"varchar(20)" json:"mobile"`    // 手机，唯一
	Avatar    string    `xorm:"varchar(255)" json:"avatar"`   // 头像
	Sex       string    `xorm:"varchar(10)" json:"sex"`       // 性别
	Nickname  string    `xorm:"varchar(64)" json:"nickname"`  // 昵称
	Salt      string    `xorm:"varchar(32)" json:"salt"`      // 密码加盐
	Token     string    `xorm:"varchar(255)" json:"token"`    // 认证令牌
	Online    int       `xorm:"int" json:"online"`            // 在线状态
	Memo      string    `xorm:"varchar(255)" json:"memo"`     // 备注
	CreatedAt time.Time `xorm:"created" json:"created_at"`    // 创建时间
	UpdatedAt time.Time `xorm:"updated" json:"updated_at"`    // 更新时间
	DeletedAt time.Time `xorm:"deleted" json:"deleted_at"`    // 删除时间（软删除）
}

type UserRegister struct {
	Mobile   string `json:"mobile"`
	Nickname string `json:"nickname"`
	PlainPwd string `json:"plainPwd"`
	RePwd    string `json:"rePwd"`
	Avatar   string `json:"avatar"`
	Sex      string `json:"sex"`
}

type UserLogin struct {
	Mobile   string `json:"mobile"`
	PlainPwd string `json:"plainPwd"` // 也可以加上: `json:"plainPwd,password"`
}
