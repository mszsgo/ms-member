package schema

import (
	"github.com/graphql-go/graphql"
)

type SmsResetPasswordType struct {
	Token  string `description:"授权Token"`
	Uid    string `description:"用户编号"`
	Mobile string `description:"手机号码"`
}

func (*SmsResetPasswordType) Description() string {
	return "短信重置密码"
}

type SmsResetPasswordTypeArgs struct {
	MallId      string `graphql:"!" description:"商城编号"`
	Mobile      string `graphql:"!" description:"手机号码"`
	Code        string `graphql:"!" description:"短信验证码"`
	NewPassword string `graphql:"!" description:"新密码"`
}

func (*SmsResetPasswordType) Args() *SmsResetPasswordTypeArgs {
	return &SmsResetPasswordTypeArgs{}
}

func (*SmsResetPasswordType) Resolve() graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (i interface{}, err error) {
		var args *SmsResetPasswordTypeArgs
		MapToStruct(p.Args, &args)
		m := NewMember().ResetPassword(args)
		return &SmsResetPasswordType{
			Mobile: m.Mobile,
			Uid:    m.Uid,
			Token:  m.CreateSession(),
		}, err
	}
}
