package request

import (
	"github.com/ldxtechteam/cpselmsdk/defaultability/domain"
	"github.com/ldxtechteam/cpselmsdk/util"
)

type AlibabaAlscUnionElemePromotionOtherchannelGetRequest struct {
	/*
	   查询request对象     */
	OtherPromotionLinkRequest *domain.AlibabaAlscUnionElemePromotionOtherchannelGetOtherPromotionLinkRequest `json:"other_promotion_link_request" required:"true" `
}

func (s *AlibabaAlscUnionElemePromotionOtherchannelGetRequest) SetOtherPromotionLinkRequest(v domain.AlibabaAlscUnionElemePromotionOtherchannelGetOtherPromotionLinkRequest) *AlibabaAlscUnionElemePromotionOtherchannelGetRequest {
	s.OtherPromotionLinkRequest = &v
	return s
}

func (req *AlibabaAlscUnionElemePromotionOtherchannelGetRequest) ToMap() map[string]interface{} {
	paramMap := make(map[string]interface{})
	if req.OtherPromotionLinkRequest != nil {
		paramMap["other_promotion_link_request"] = util.ConvertStruct(*req.OtherPromotionLinkRequest)
	}
	return paramMap
}

func (req *AlibabaAlscUnionElemePromotionOtherchannelGetRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	return fileMap
}
