package schema

import (
	"github.com/graphql-go/graphql"
)

type SmsUpdateMobileType struct {
	Uid    string `description:"用户编号"`
	Mobile string `description:"新手机号码"`
	Token  string `description:"授权Token"`
}

func (*SmsUpdateMobileType) Description() string {
	return "会员用户短信注册"
}

type SmsUpdateMobileTypeArgs struct {
	Uid       string `graphql:"!" description:"会员用户ID"`
	NewMobile string `graphql:"!" description:"新手机号码"`
	Code      string `graphql:"!" description:"新号码短信验证码"`
}

func (*SmsUpdateMobileType) Args() *SmsUpdateMobileTypeArgs {
	return &SmsUpdateMobileTypeArgs{}
}

func (*SmsUpdateMobileType) Resolve() graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (i interface{}, err error) {
		var args *SmsUpdateMobileTypeArgs
		MapToStruct(p.Args, &args)
		m := NewMember().UpdateMobile(args)
		return &SmsUpdateMobileType{
			Token:  m.CreateSession(),
			Uid:    args.Uid,
			Mobile: args.NewMobile,
		}, err
	}
}
