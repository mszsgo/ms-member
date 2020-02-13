package schema

import (
	"time"

	"github.com/graphql-go/graphql"
)

type LoginType struct {
	Token         string `description:"授权token"`
	OrgId         string `description:"会员所属机构编号"`
	Uid           string `description:"会员用户编号"`
	LoginId       string `description:"登录账号"`
	Mobile        string `description:"用户绑定手机，可用于登录"`
	Email         string `description:"用户绑定邮箱,可用于登录"`
	Nickname      string `description:"用户昵称"`
	Avatar        string `description:"用户头像URL"`
	CreatedAt     string `description:"注册时间"`
	LastLoginTime string `description:"最后登录时间"`
}

func (*LoginType) Description() string {
	return "会员用户登录"
}

type LoginTypeArgs struct {
	MallId   string `graphql:"!" description:"商城编号，由前端根据入口参数自动获取"`
	LoginId  string `graphql:"!" description:"登录账号，可以是账号、手机、邮箱，用户输入"`
	Password string `graphql:"!" description:"登录密码，用户输入"`
}

func (*LoginType) Args() *LoginTypeArgs {
	return &LoginTypeArgs{}
}

func (*LoginType) Resolve() graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (i interface{}, err error) {
		var args *LoginTypeArgs
		MapToStruct(p.Args, &args)
		m := NewMember().Login(args)
		return &LoginType{
			Token:         m.CreateSession(),
			OrgId:         m.OrgId,
			Uid:           m.Uid,
			LoginId:       m.LoginId,
			Mobile:        m.Mobile,
			Email:         m.Email,
			Nickname:      m.Nickname,
			Avatar:        m.Avatar,
			CreatedAt:     m.CreatedAt.Format(time.RFC3339),
			LastLoginTime: m.LastLoginTime.Format(time.RFC3339),
		}, err
	}
}
