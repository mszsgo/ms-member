package schema

import (
	"github.com/graphql-go/graphql"
)

type AddMemberType struct {
	Uid string `description:"用户编号"`
}

func (*AddMemberType) Description() string {
	return "后台添加会员用户"
}

type AddMemberTypeArgs struct {
	OrgId    string `graphql:"!" description:"商城编号"`
	LoginId  string `graphql:"!" description:"登录账号"`
	Password string `graphql:"!" description:"登录密码，MD5值"`
	Mobile   string `graphql:"" description:"手机号码"`
	Email    string `graphql:"" description:"邮箱"`
}

func (*AddMemberType) Args() *AddMemberTypeArgs {
	return &AddMemberTypeArgs{}
}

func (*AddMemberType) Resolve() graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (i interface{}, err error) {
		var args *AddMemberTypeArgs
		MapToStruct(p.Args, &args)
		uid := NewMember().AddMember(args)
		return &AddMemberType{
			Uid: uid,
		}, err
	}
}
