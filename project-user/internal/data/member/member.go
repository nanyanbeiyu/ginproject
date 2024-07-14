/*
@author: NanYan
*/
package member

import (
	"gorm.io/gorm"
	"time"
)

// Member 结构体表示一个成员实体，包含了详细的成员信息
type Member struct {
	gorm.Model
	Account         int64     `json:"account" gorm:"comment:用户账号"`                // 用户账号
	Password        string    `json:"password" gorm:"comment:用户密码"`               // 用户密码
	Name            string    `json:"name" gorm:"comment:用户名称"`                   // 用户名称
	Mobile          string    `json:"mobile" gorm:"comment:用户手机号"`                // 用户手机号
	Realname        string    `json:"realname" gorm:"comment:用户真实姓名"`             // 用户真实姓名
	Status          int       `json:"status" gorm:"comment:用户状态，指示账户是否激活"`        // 用户状态，指示账户是否激活
	LastLoginTime   time.Time `json:"last_login_time" gorm:"comment:最近登录时间"`      // 最近登录时间
	Sex             int       `json:"sex" gorm:"comment:用户性别"`                    // 用户性别
	Avatar          string    `json:"avatar" gorm:"comment:用户头像URL"`              // 用户头像URL
	Idcard          string    `json:"idcard" gorm:"comment:用户身份证号码"`              // 用户身份证号码
	Province        int       `json:"province" gorm:"comment:用户所在省份代码"`           // 用户所在省份代码
	City            int       `json:"city" gorm:"comment:用户所在城市代码"`               // 用户所在城市代码
	Area            int       `json:"area" gorm:"comment:用户所在地区代码"`               // 用户所在地区代码
	Address         string    `json:"address" gorm:"comment:用户详细地址"`              // 用户详细地址
	Description     string    `json:"description" gorm:"comment:用户个人描述"`          // 用户个人描述
	Email           string    `json:"email" gorm:"comment:用户邮箱地址"`                // 用户邮箱地址
	DingtalkOpenid  string    `json:"dingtalk_openid" gorm:"comment:钉钉开放ID"`      // 钉钉开放ID
	DingtalkUnionid string    `json:"dingtalk_unionid" gorm:"comment:钉钉Union ID"` // 钉钉Union ID
	DingtalkUserid  string    `json:"dingtalk_userid" gorm:"comment:钉钉用户ID"`      // 钉钉用户ID
}

func (*Member) TableName() string {
	return "member"
}
