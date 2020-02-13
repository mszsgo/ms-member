package schema

import (
	"encoding/json"
	"errors"
	"log"

	"github.com/mszsgo/hgraph"
)

// 调用服务
func ServiceCall(query string) *FeignResponseData {
	reqModel := hgraph.GraphRequestModel{
		RequestId:     GenRequestId(),
		Token:         "",
		OperationName: "",
		Query:         query,
		Variables:     nil,
	}
	reqbytes, _ := json.Marshal(reqModel)
	log.Printf("Request：%s", string(reqbytes))
	resModel := hgraph.Feign(&reqModel)
	resbytes, _ := json.Marshal(resModel)
	log.Printf("Response：%s", string(resbytes))
	if len(resModel.Errors) > 0 {
		Panic(errors.New(resModel.Errors[0]["message"].(string)))
	}
	// orgId = resModel.Data["mall"].(map[string]interface{})["info"].(map[string]interface{})["orgId"].(string)
	var resData *FeignResponseData
	MapToStruct(resModel.Data, &resData)
	return resData
}

// 调用其它服务的响应数据结果集
type FeignResponseData struct {
	Mall MallServiceFeign
}

type MallServiceFeign struct {
	Info MallInfoFeign
}

type MallInfoFeign struct {
	OrgId string
}
