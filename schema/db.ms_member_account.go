// 会员信息数据库操作
package schema

import (
	"log"
	"time"

	"github.com/mszsgo/hmgdb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

/*
会员服务提供功能：
- 会员登录
- 会员注册
- 会员信息维护等

## Mongodb 索引记录
unique index [uid]
unique index [orgId,mobile]
index [email,loginId]
*/

// 平台账户信息 (ms_member_account)
// 用于平台用户登录授权信息
type MemberAccount struct {
	Uid           string       `bson:"uid" json:"uid"`                     // 会员账户编号
	OrgId         string       `bson:"orgId" json:"orgId"`                 // 用户属于某机构的会员，必须有一个所属机构
	LoginId       string       `bson:"loginId" json:"loginId"`             // 登录账号
	Password      string       `bson:"password" json:"-"`                  // 密码，保存md5值
	Mobile        string       `bson:"mobile" json:"mobile"`               // 用户绑定手机，可用于登录
	Email         string       `bson:"email" json:"email"`                 // 用户绑定邮箱,可用于登录
	Nickname      string       `bson:"nickname" json:"nickname"`           // 用户昵称
	Avatar        string       `bson:"avatar" json:"avatar"`               // 用户头像URL
	CreatedAt     time.Time    `bson:"createdAt" json:"createdAt"`         // 注册时间
	UpdatedAt     time.Time    `bson:"updatedAt" json:"updatedAt"`         // 最后更新时间
	LastLoginTime time.Time    `bson:"lastLoginTime" json:"lastLoginTime"` // 最后登录时间
	Status        MemberStatus `bson:"status" json:"status"`               // 用户状态
	User          MemberUser   `bson:"user" json:"user"`                   // 用户资料
}

// 平台普通用户
// 登录用户系统时，账号身份就是用户
type MemberUser struct {
	RealName   string     `bson:"realName" json:"realName"`     // 真实姓名
	Province   string     `bson:"province" json:"province"`     // 省
	City       string     `bson:"city" json:"city"`             // 市
	Address    string     `bson:"address" json:"address"`       // 联系地址
	IcType     UserIcType `bson:"icType" json:"icType"`         // 证件类型
	IcNumber   string     `bson:"icNumber" json:icNumber`       // 证件号码
	IcAddress  string     `bson:"icAddress" json:"icAddress"`   // 证件地址
	IcBegTime  time.Time  `bson:"icBegTime" json:"icBegTime"`   // 有效期开始
	IcEndTime  time.Time  `bson:"icEndTime" json:"icEndTime"`   // 有效期结束
	IcImgFront string     `bson:"icImgFront" json:"icImgFront"` // 证件正面图片
	IcImgBack  string     `bson:"icImgBack" json:"icImgBack"`   // 证件反面图片
}

func NewMember() *MemberAccount {
	return &MemberAccount{}
}

func (o *MemberAccount) Collection() *mongo.Collection {
	return hmgdb.Db().Collection("ms_member_account")
}

func (member *MemberAccount) CreateSession() (token string) {
	return NewSession().Create(&Session{
		Uid:           member.Uid,
		LoginId:       member.LoginId,
		Mobile:        member.Mobile,
		Email:         member.Email,
		Nickname:      member.Email,
		Avatar:        member.Avatar,
		Status:        member.Status,
		OrgId:         member.OrgId,
		LastLoginTime: member.LastLoginTime,
	})
}

func (m *MemberAccount) Exists(orgId, mobile string) bool {
	return hmgdb.Exists(nil, m.Collection(), bson.M{"orgId": orgId, "mobile": mobile})
}

func (m *MemberAccount) FindOneByMobile(orgId, mobile string) *MemberAccount {
	var member *MemberAccount
	hmgdb.FindOne(nil, m.Collection(), bson.M{"orgId": orgId, "mobile": mobile}, &member)
	return member
}

func (m *MemberAccount) FindOneByUid(uid string) *MemberAccount {
	var member *MemberAccount
	hmgdb.FindOne(nil, m.Collection(), bson.M{"uid": uid}, &member)
	return member
}

func (m *MemberAccount) CheckSmsCode(mobile, code string) {
	// TODO 调用短信服务验证短信验证码是否正确

	// 验证短信验证码
	if code != "214365" {
		Panic(E10020)
	}
}

// 账号密码登录
func (o *MemberAccount) Login(args *LoginTypeArgs) *MemberAccount {
	orgId := FindOrgIdByMallId(args.MallId)
	var member *MemberAccount
	err := o.Collection().FindOne(nil, bson.M{
		"orgId": orgId,
		"$or":   bson.A{bson.M{"mobile": args.LoginId}, bson.M{"email": args.LoginId}, bson.M{"loginId": args.LoginId}},
	}).Decode(&member)
	if err != nil {
		log.Printf(err.Error())
	}
	// 用户不存在
	if member == nil {
		Panic(E10111)
	}
	// 密码错误
	if member.Password != args.Password {
		Panic(E10112)
	}
	// 更新登录时间
	hmgdb.UpdateOne(nil, o.Collection(), bson.M{"uid": member.Uid}, bson.M{"$set": bson.M{"lastLoginTime": time.Now()}})
	return member
}

func (o *MemberAccount) SmsRegister(args *SmsRegisterTypeArgs) *MemberAccount {
	orgId := FindOrgIdByMallId(args.MallId)
	// 验证短信验证码
	o.CheckSmsCode(args.Mobile, args.Code)

	// 验证手机加机构，是否已注册
	if o.Exists(orgId, args.Mobile) {
		Panic(E10110)
	}
	// 会员用户对象
	member := &MemberAccount{
		Uid:           GenUid(),
		OrgId:         orgId,
		LoginId:       "",
		Password:      args.Password,
		Mobile:        args.Mobile,
		Email:         "",
		Nickname:      "",
		Avatar:        "",
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
		LastLoginTime: time.Now(),
		Status:        MEMBER_STATUS_NORMAL,
		User:          MemberUser{},
	}
	hmgdb.InsertOne(nil, o.Collection(), member)
	return member
}

func (o *MemberAccount) UpdateMobile(args *SmsUpdateMobileTypeArgs) *MemberAccount {
	// 验证短信验证码
	o.CheckSmsCode(args.NewMobile, args.Code)

	// 更新手机号
	u := hmgdb.UpdateOne(nil, o.Collection(), bson.M{"uid": args.Uid}, bson.M{"$set": bson.M{"mobile": args.NewMobile, "updatedAt": time.Now()}})
	if u.ModifiedCount != 1 {
		Panic(E10113)
	}
	return o.FindOneByUid(args.Uid)
}

func (o *MemberAccount) ResetPassword(args *SmsResetPasswordTypeArgs) *MemberAccount {
	// 验证短信验证码
	o.CheckSmsCode(args.Mobile, args.Code)
	orgId := FindOrgIdByMallId(args.MallId)
	// 更新密码
	u := hmgdb.UpdateOne(nil, o.Collection(), bson.M{"orgId": orgId, "mobile": args.Mobile}, bson.M{"$set": bson.M{"password": args.NewPassword, "updatedAt": time.Now()}})
	if u.ModifiedCount != 1 {
		Panic(E10113)
	}
	return o.FindOneByMobile(orgId, args.Mobile)
}

// 查询会员用户信息
func (o *MemberAccount) findOrgMemberByUid(orgId string, uid string) *MemberAccount {
	var member *MemberAccount
	hmgdb.FindOne(nil, o.Collection(), bson.M{"orgId": orgId, "uid": uid}, &member)
	return member
}

func (o *MemberAccount) AddMember(args *AddMemberTypeArgs) (uid string) {
	uid = GenUid()
	// TODO 验证邮箱是否已经注册

	// TODO 验证手机是否已经注册

	hmgdb.InsertOne(nil, o.Collection(), &MemberAccount{
		Uid:           uid,
		OrgId:         args.OrgId,
		LoginId:       args.LoginId,
		Password:      args.Password,
		Mobile:        args.Mobile,
		Email:         args.Email,
		Nickname:      "",
		Avatar:        "",
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
		LastLoginTime: time.Now(),
		Status:        "",
		User:          MemberUser{},
	})
	return uid
}
