package schema

// 会员状态（1=正常 8=禁用 9=注销）
type MemberStatus string

const (
	MEMBER_STATUS_NORMAL    MemberStatus = "1"
	MEMBER_STATUS_FORBIDDEN MemberStatus = "8"
	MEMBER_STATUS_INVALID   MemberStatus = "9"
)

// 证件类型（1=身份证）
type UserIcType string

const (
	USER_IC_TYPE_SFZ UserIcType = "1"
)
