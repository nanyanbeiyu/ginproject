/*
@author: NanYan
*/
package common

import (
	"fmt"
	"math/rand"
	"time"
)

// 生成随机验证码
func GenerateRandomCode(codeLength int) string {
	// 随机数生成器的种子
	rand.Seed(time.Now().UnixNano())

	// 生成codeLength位随机数
	code := ""
	for i := 0; i < codeLength; i++ {
		// 生成0-9之间的随机数
		num := rand.Intn(10)
		code += fmt.Sprint(num)
	}

	return code
}
