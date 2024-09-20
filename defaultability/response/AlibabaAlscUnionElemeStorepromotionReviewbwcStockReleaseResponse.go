package response

import (
	"github.com/ldxtechteam/cpselmsdk/defaultability/domain"
)

type AlibabaAlscUnionElemeStorepromotionReviewbwcStockReleaseResponse struct {

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
	Data domain.AlibabaAlscUnionElemeStorepromotionReviewbwcStockReleaseReviewBwcStockReleaseResult `json:"data,omitempty" `
	/*
	   错误码
	*/
	BizErrorCode string `json:"biz_error_code,omitempty" `
	/*
	   错误消息
	*/
	BizErrorMessage string `json:"biz_error_message,omitempty" `
}
