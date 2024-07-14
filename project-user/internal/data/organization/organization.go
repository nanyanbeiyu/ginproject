/*
@author: NanYan
*/
package organization

import "gorm.io/gorm"

// Organization 组织机构
type Organization struct {
	gorm.Model
	Name        string `gorm:"comment: 组织的名称"`                        // Name: 组织的名称
	Avatar      string `gorm:"comment: 组织的头像链接"`                      // Avatar: 组织的头像链接
	Description string `gorm:"comment: 对组织的简要描述"`                     // Description: 对组织的简要描述
	MemberId    int64  `gorm:"comment: 组织创建者的用户ID"`                   // MemberId: 组织创建者的用户ID
	Personal    int32  `gorm:"comment: 组织类型，个人组织的标志（0：非个人组织，1：个人组织）"` // Personal: 组织类型，个人组织的标志（0：非个人组织，1：个人组织）
	Address     string `gorm:"comment: 组织的所在地地址"`                     // Address: 组织的所在地地址
	Province    int32  `gorm:"comment: 所在省份的ID"`                      // Province: 所在省份的ID
	City        int32  `gorm:"comment: 所在城市的ID"`                      // City: 所在城市的ID
	Area        int32  `gorm:"comment: 所在区域的ID"`                      // Area: 所在区域的ID
}

func (*Organization) TableName() string {
	return "organization"
}
