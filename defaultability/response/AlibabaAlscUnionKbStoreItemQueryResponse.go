package response

import (
	"github.com/ldxtechteam/cpselmsdk/defaultability/domain"
)

type AlibabaAlscUnionKbStoreItemQueryResponse struct {

	/*
	   System request id
	*/
	RequestId string `json:"request_id,omitempty" `

	/*
	   System body
	*/
	Body string

	/*
	   分页数据
	*/
	Data domain.AlibabaAlscUnionKbStoreItemQueryPageModel `json:"data,omitempty" `
	/*
	   返回码
	*/
	ResultCode int64 `json:"result_code,omitempty" `
	/*
	   返回消息
	*/
	Message string `json:"message,omitempty" `
	/*
	   错误描述
	*/
	ErrorMessage string `json:"error_message,omitempty" `
}
