package request

import (
	"github.com/ldxtechteam/cpselmsdk/defaultability/domain"
	"github.com/ldxtechteam/cpselmsdk/util"
)

type AlibabaAlscUnionElemePromotionOfficialactivityGetRequest struct {
	/*
	   查询rquest     */
	QueryRequest *domain.AlibabaAlscUnionElemePromotionOfficialactivityGetActivityRequest `json:"query_request" required:"true" `
}

func (s *AlibabaAlscUnionElemePromotionOfficialactivityGetRequest) SetQueryRequest(v domain.AlibabaAlscUnionElemePromotionOfficialactivityGetActivityRequest) *AlibabaAlscUnionElemePromotionOfficialactivityGetRequest {
	s.QueryRequest = &v
	return s
}

func (req *AlibabaAlscUnionElemePromotionOfficialactivityGetRequest) ToMap() map[string]interface{} {
	paramMap := make(map[string]interface{})
	if req.QueryRequest != nil {
		paramMap["query_request"] = util.ConvertStruct(*req.QueryRequest)
	}
	return paramMap
}

func (req *AlibabaAlscUnionElemePromotionOfficialactivityGetRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	return fileMap
}
