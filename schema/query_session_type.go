package schema

import (
	"time"

	"github.com/graphql-go/graphql"
)

type SessionType struct {
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

func (*SessionType) Description() string {
	return "登录会话信息"
}

type SessionTypeArgs struct {
	Token string `graphql:"!" description:"授权token，登录接口返回"`
}

func (*SessionType) Args() *SessionTypeArgs {
	return &SessionTypeArgs{}
}

func (*SessionType) Resolve() graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (i interface{}, err error) {
		var args *SessionTypeArgs
		MapToStruct(p.Args, &args)
		s := NewSession().FindOneByToken(args.Token)
		return &SessionType{
			OrgId:         s.OrgId,
			Uid:           s.Uid,
			LoginId:       s.LoginId,
			Mobile:        s.Mobile,
			Email:         s.Email,
			Nickname:      s.Nickname,
			Avatar:        s.Avatar,
			CreatedAt:     s.CreatedAt.Format(time.RFC3339),
			LastLoginTime: s.LastLoginTime.Format(time.RFC3339),
		}, err
	}
}
