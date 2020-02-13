package schema

import (
	"fmt"
)

// 会员服务业务方法

// 根据商城编号查询机构编号
func FindOrgIdByMallId(mallId string) (orgId string) {
	resData := ServiceCall(fmt.Sprintf(`
	query FindOrgIdByMallId {
	  mall {
		mallInfo(mallId: "%s") {
		  orgId
		}
	  }
	}
	`, mallId))
	return resData.Mall.Info.OrgId
}
