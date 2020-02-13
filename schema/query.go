package schema

import (
	"github.com/graphql-go/graphql"
)

type Query struct {
	Member MemberQuery `description:"会员服务"`
}

type MemberQuery struct {
	Session SessionType `description:"会员登录会话信息"`
}

func (*MemberQuery) Resolve() graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (i interface{}, err error) {
		return "", err
	}
}
