package request

import (
	"github.com/ldxtechteam/cpselmsdk/defaultability/domain"
	"github.com/ldxtechteam/cpselmsdk/util"
)

type AlibabaAlscUnionElemePromotionItempromotionGetRequest struct {
	/*
	   查询rquest     */
	QueryRequest *domain.AlibabaAlscUnionElemePromotionItempromotionGetSingleItemPromotionRequest `json:"query_request" required:"true" `
}

func (s *AlibabaAlscUnionElemePromotionItempromotionGetRequest) SetQueryRequest(v domain.AlibabaAlscUnionElemePromotionItempromotionGetSingleItemPromotionRequest) *AlibabaAlscUnionElemePromotionItempromotionGetRequest {
	s.QueryRequest = &v
	return s
}

func (req *AlibabaAlscUnionElemePromotionItempromotionGetRequest) ToMap() map[string]interface{} {
	paramMap := make(map[string]interface{})
	if req.QueryRequest != nil {
		paramMap["query_request"] = util.ConvertStruct(*req.QueryRequest)
	}
	return paramMap
}

func (req *AlibabaAlscUnionElemePromotionItempromotionGetRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	return fileMap
}
