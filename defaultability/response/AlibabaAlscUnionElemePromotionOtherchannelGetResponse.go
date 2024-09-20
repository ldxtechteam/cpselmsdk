package response

import (
	"github.com/ldxtechteam/cpselmsdk/defaultability/domain"
)

type AlibabaAlscUnionElemePromotionOtherchannelGetResponse struct {

	/*
	   System request id
	*/
	RequestId string `json:"request_id,omitempty" `

	/*
	   System body
	*/
	Body string

	/*
	   特殊推广链接
	*/
	Data domain.AlibabaAlscUnionElemePromotionOtherchannelGetOtherPromotionLink `json:"data,omitempty" `
	/*
	   请求结果码
	*/
	ResultCode int64 `json:"result_code,omitempty" `
	/*
	   返回描述
	*/
	Message string `json:"message,omitempty" `
}
