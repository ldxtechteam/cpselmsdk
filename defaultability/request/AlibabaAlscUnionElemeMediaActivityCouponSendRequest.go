package request

import (
	"github.com/ldxtechteam/cpselmsdk/defaultability/domain"
	"github.com/ldxtechteam/cpselmsdk/util"
)

type AlibabaAlscUnionElemeMediaActivityCouponSendRequest struct {
	/*
	   请求对象     */
	SendRequest *domain.AlibabaAlscUnionElemeMediaActivityCouponSendMediaActivityCouponSendRequest `json:"send_request" required:"true" `
}

func (s *AlibabaAlscUnionElemeMediaActivityCouponSendRequest) SetSendRequest(v domain.AlibabaAlscUnionElemeMediaActivityCouponSendMediaActivityCouponSendRequest) *AlibabaAlscUnionElemeMediaActivityCouponSendRequest {
	s.SendRequest = &v
	return s
}

func (req *AlibabaAlscUnionElemeMediaActivityCouponSendRequest) ToMap() map[string]interface{} {
	paramMap := make(map[string]interface{})
	if req.SendRequest != nil {
		paramMap["send_request"] = util.ConvertStruct(*req.SendRequest)
	}
	return paramMap
}

func (req *AlibabaAlscUnionElemeMediaActivityCouponSendRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	return fileMap
}
