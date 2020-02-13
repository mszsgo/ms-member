// Map常用操作函数
package schema

import (
	"errors"
	"fmt"

	"github.com/mszsgo/hmap"
)

// Map转Struct
// 输入参数为Map类型值，输出为结构体
func MapToStruct(input interface{}, output interface{}) {
	err := hmap.Decode(input, output)
	if err != nil {
		panic(errors.New(fmt.Sprintf("MapToStruct error %s", err.Error())))
	}
}
