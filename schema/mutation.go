package schema

import (
	"github.com/graphql-go/graphql"
)

type Mutation struct {
	Member MemberMutation `description:"会员服务"`
}

type MemberMutation struct {
	Login            LoginType            `description:"会员账号密码登录"`
	SmsRegister      MallSmsRegisterType  `description:"用户短信注册商城用户"`
	SmsUpdateMobile  SmsUpdateMobileType  `description:"用户短信修改手机号"`
	SmsResetPassword SmsResetPasswordType `description:"用户短信重置密码"`
	AddMember        AddMemberType        `description:"添加会员，用于后台运营操作"`
}

func (*MemberMutation) Resolve() graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (i interface{}, err error) {
		return "", err
	}
}
