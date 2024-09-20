package request

import (
	"github.com/ldxtechteam/cpselmsdk/defaultability/domain"
	"github.com/ldxtechteam/cpselmsdk/util"
)

type AlibabaAlscUnionElemePromotionItempromotionQueryRequest struct {
	/*
	   查询rquest     */
	QueryRequest *domain.AlibabaAlscUnionElemePromotionItempromotionQueryItemPromotionQueryRequest `json:"query_request" required:"true" `
}

func (s *AlibabaAlscUnionElemePromotionItempromotionQueryRequest) SetQueryRequest(v domain.AlibabaAlscUnionElemePromotionItempromotionQueryItemPromotionQueryRequest) *AlibabaAlscUnionElemePromotionItempromotionQueryRequest {
	s.QueryRequest = &v
	return s
}

func (req *AlibabaAlscUnionElemePromotionItempromotionQueryRequest) ToMap() map[string]interface{} {
	paramMap := make(map[string]interface{})
	if req.QueryRequest != nil {
		paramMap["query_request"] = util.ConvertStruct(*req.QueryRequest)
	}
	return paramMap
}

func (req *AlibabaAlscUnionElemePromotionItempromotionQueryRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	return fileMap
}
