/*
@author: NanYan
*/
package model

import common "carrygpc.com/project-common"

const (
	NoLegalMobile common.BusinessCode = 2001 // 手机号不合法

	ServerError common.BusinessCode = 5000 // 服务器错误

	VerifyCodeSuccess common.BusinessCode = 2002 // 验证码发送成功
)

var CodeMsg = map[common.BusinessCode]string{
	NoLegalMobile:     "手机号不合法",
	ServerError:       "服务出错了",
	VerifyCodeSuccess: "验证码发送成功",
}

func GetMsg(code common.BusinessCode) string {
	return CodeMsg[code]
}
