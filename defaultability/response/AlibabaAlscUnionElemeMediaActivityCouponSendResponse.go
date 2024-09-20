package response

import (
)

type AlibabaAlscUnionElemeMediaActivityCouponSendResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        结果码
    */
    ResultCode  int64 `json:"result_code,omitempty" `
    /*
        错误消息
    */
    Message  string `json:"message,omitempty" `
}
