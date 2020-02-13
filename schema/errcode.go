package schema

import (
	"errors"
	"fmt"
	"log"
)

/* 全局接口错误码定义 99开头的为系统级错误 */

// 会员服务错误码10***
var (
	// 通用错误码
	SUCCESS     = errors.New("00000:ok")
	FAIL        = errors.New("99999:%s")
	MONGO_ERROR = errors.New("99100:mongo error -> %s")

	// 业务错误码

	// 会话 100
	E10011 = errors.New("10011:Token字符串格式错误 %s")
	E10012 = errors.New("10012:生成token出错 %s")

	// 短信验证
	E10020 = errors.New("10020:短信验证码错误")

	// 用户 101
	E10110 = errors.New("10110:手机号已注册")
	E10111 = errors.New("10111:账号不存在")
	E10112 = errors.New("10112:登录密码错误")
	E10113 = errors.New("10113:用户编号不存在")
)

func Panic(userDefinedErr error, args ...interface{}) {
	if userDefinedErr == nil {
		Panic(FAIL, "Panic is nil")
	}
	e := userDefinedErr.Error()
	errMsg := fmt.Sprintf(e, args...)
	log.Printf("Panic Exception： %s", errMsg)
	panic(errors.New(errMsg))
}
