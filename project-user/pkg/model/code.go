/*
@author: NanYan
*/
package model

import common "carrygpc.com/project-common"

const (
	NoLegalMobile common.BusinessCode = 2001 // 手机号不合法

	NoLegalCaptcha    common.BusinessCode = 3000 // 验证码错误
	VerifyCodeSuccess common.BusinessCode = 3001 // 验证码发送成功

	NoLegalEmail common.BusinessCode = 4001 // 邮箱不合法
	EmailExist   common.BusinessCode = 4002 // 邮箱已存在
	NameExist    common.BusinessCode = 4003 // 昵称已存在
	MobileExist  common.BusinessCode = 4004 // 手机号已存在

	ServerError common.BusinessCode = 5000 // 服务器错误
	DBError     common.BusinessCode = 5001 // 数据库错误
)

var CodeMsg = map[common.BusinessCode]string{
	NoLegalMobile:     "手机号不合法",
	ServerError:       "服务出错了",
	VerifyCodeSuccess: "验证码发送成功",
	NoLegalCaptcha:    "验证码错误",
	NoLegalEmail:      "邮箱不合法",
	EmailExist:        "邮箱已存在",
	DBError:           "数据库错误",
	NameExist:         "昵称已存在",
	MobileExist:       "手机号已存在",
}

func GetMsg(code common.BusinessCode) string {
	return CodeMsg[code]
}
