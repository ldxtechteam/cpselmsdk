package response

import (
	"github.com/ldxtechteam/cpselmsdk/defaultability/domain"
)

type AlibabaAlscUnionElemeStorepromotionReviewbwcBindLinkGetResponse struct {

	/*
	   System request id
	*/
	RequestId string `json:"request_id,omitempty" `

	/*
	   System body
	*/
	Body string

	/*
	   数据
	*/
	Data domain.AlibabaAlscUnionElemeStorepromotionReviewbwcBindLinkGetReviewBwcSidBindLinkResult `json:"data,omitempty" `
	/*
	   返回码
	*/
	ResultCode int64 `json:"result_code,omitempty" `
	/*
	   返回消息
	*/
	Message string `json:"message,omitempty" `
	/*
	   错误消息
	*/
	ErrorMessage string `json:"error_message,omitempty" `
}
