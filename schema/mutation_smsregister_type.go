package schema

import (
	"github.com/graphql-go/graphql"
)

type MallSmsRegisterType struct {
	Uid   string `description:"用户ID"`
	Token string `description:"授权token，可实现注册完成默认登录成功"`
}

func (*MallSmsRegisterType) Description() string {
	return "会员用户短信注册"
}

type SmsRegisterTypeArgs struct {
	MallId   string `graphql:"!" description:"商城编号"`
	Mobile   string `graphql:"!" description:"手机号码"`
	Password string `graphql:"!" description:"密码，前端Md5值"`
	Code     string `graphql:"!" description:"短信验证码"`
}

func (*MallSmsRegisterType) Args() *SmsRegisterTypeArgs {
	return &SmsRegisterTypeArgs{}
}

func (*MallSmsRegisterType) Resolve() graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (i interface{}, err error) {
		var args *SmsRegisterTypeArgs
		MapToStruct(p.Args, &args)
		member := NewMember().SmsRegister(args)
		return &MallSmsRegisterType{
			Uid:   member.Uid,
			Token: member.CreateSession(),
		}, err
	}
}
