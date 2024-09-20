package response

import (
	"github.com/ldxtechteam/cpselmsdk/defaultability/domain"
)

type AlibabaAlscUnionElemeStorepromotionReviewbwcQueryResponse struct {

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
	Data domain.AlibabaAlscUnionElemeStorepromotionReviewbwcQueryPageModel `json:"data,omitempty" `
	/*
	   错误码
	*/
	BizErrorCode string `json:"biz_error_code,omitempty" `
	/*
	   错误消息
	*/
	BizErrorMessage string `json:"biz_error_message,omitempty" `
}
