package schema

import (
	"os"

	"github.com/mszsgo/hmgdb"
)

func init() {
	// 项目加载设置Mongodb连接字符串
	hmgdb.SetConnectString(hmgdb.DEFAULT, os.Getenv("MS_MONGODB_CONNECT"))
}
