package request

import (
	"github.com/ldxtechteam/cpselmsdk/defaultability/domain"
	"github.com/ldxtechteam/cpselmsdk/util"
)

type AlibabaAlscUnionPromotionLinkAnalyzeRequest struct {
	/*
	   查询request对象     */
	PromotionLinkAnalyzeRequest *domain.AlibabaAlscUnionPromotionLinkAnalyzePromotionLinkAnalyzeRequest `json:"promotion_link_analyze_request" required:"true" `
}

func (s *AlibabaAlscUnionPromotionLinkAnalyzeRequest) SetPromotionLinkAnalyzeRequest(v domain.AlibabaAlscUnionPromotionLinkAnalyzePromotionLinkAnalyzeRequest) *AlibabaAlscUnionPromotionLinkAnalyzeRequest {
	s.PromotionLinkAnalyzeRequest = &v
	return s
}

func (req *AlibabaAlscUnionPromotionLinkAnalyzeRequest) ToMap() map[string]interface{} {
	paramMap := make(map[string]interface{})
	if req.PromotionLinkAnalyzeRequest != nil {
		paramMap["promotion_link_analyze_request"] = util.ConvertStruct(*req.PromotionLinkAnalyzeRequest)
	}
	return paramMap
}

func (req *AlibabaAlscUnionPromotionLinkAnalyzeRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	return fileMap
}
