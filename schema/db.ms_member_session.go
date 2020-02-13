package schema

import (
	"time"

	"github.com/mszsgo/hmgdb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// 使用Mongodb集合存储会话信息
// 集合：ms_member_session  ,创建过期索引，自动清理过期数据
type Session struct {
	Token         primitive.ObjectID `bson:"_id" json:"token"`                   // token值，使用mongodb _id
	Expires       time.Time          `bson:"expires" json:"expires"`             // 过期时间
	Uid           string             `bson:"uid" json:"uid"`                     // 会员账户编号
	LoginId       string             `bson:"loginId" json:"loginId"`             // 登录账号
	Mobile        string             `bson:"mobile" json:"mobile"`               // 用户绑定手机，可用于登录
	Email         string             `bson:"email" json:"email"`                 // 用户绑定邮箱,可用于登录
	Nickname      string             `bson:"nickname" json:"nickname"`           // 用户昵称
	Avatar        string             `bson:"avatar" json:"avatar"`               // 用户头像URL
	CreatedAt     time.Time          `bson:"createdAt" json:"createdAt"`         // 注册时间
	LastLoginTime time.Time          `bson:"lastLoginTime" json:"lastLoginTime"` // 最后登录时间
	Status        MemberStatus       `bson:"status" json:"status"`               // 用户状态
	OrgId         string             `bson:"orgId" json:"orgId"`                 // 当前用户所属的机构编号
}

func NewSession() *Session {
	return &Session{}
}

func (o *Session) Collection() *mongo.Collection {
	return hmgdb.Db().Collection("ms_member_session")
}

func (o *Session) Create(s *Session) (token string) {
	s.Token = primitive.NewObjectID()
	s.Expires = time.Now().Add(2 * time.Hour) // 会话有效期2小时
	s.CreatedAt = time.Now()
	return hmgdb.InsertOne(nil, o.Collection(), &s)
}

func (o *Session) FindOneByToken(token string) *Session {
	_id, err := primitive.ObjectIDFromHex(token)
	if err != nil {
		Panic(E10011, err.Error())
	}
	var rsToken *Session
	hmgdb.FindOne(nil, o.Collection(), bson.M{"_id": _id}, rsToken)
	return rsToken
}
