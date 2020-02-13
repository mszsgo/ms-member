package schema

import (
	"log"
	"math/rand"
	"time"

	"github.com/mszsgo/snowflake"
)

var (
	_genNode *snowflake.Node
)

func IdNode() *snowflake.Node {
	if _genNode != nil {
		return _genNode
	}
	rand.Seed(time.Now().UnixNano())
	_genNode, err := snowflake.NewNode(rand.Int63n(31))
	if err != nil {
		log.Fatal(err)
	}
	return _genNode
}

// 会员用户ID
func GenUid() string {
	return IdNode().Generate().String()
}

// 机构ID
func GenOrgId() string {
	return IdNode().Generate().String()
}

// 商城ID
func GenMallId() string {
	return IdNode().Generate().String()
}

// 请求ID
func GenRequestId() string {
	return IdNode().Generate().String()
}

// 交易ID
func GenTradeId() string {
	return IdNode().Generate().String()
}
