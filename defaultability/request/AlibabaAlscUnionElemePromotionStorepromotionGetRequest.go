package request

import (
	"github.com/ldxtechteam/cpselmsdk/defaultability/domain"
	"github.com/ldxtechteam/cpselmsdk/util"
)

type AlibabaAlscUnionElemePromotionStorepromotionGetRequest struct {
	/*
	   查询rquest     */
	QueryRequest *domain.AlibabaAlscUnionElemePromotionStorepromotionGetSingleStorePromotionRequest `json:"query_request" required:"true" `
}

func (s *AlibabaAlscUnionElemePromotionStorepromotionGetRequest) SetQueryRequest(v domain.AlibabaAlscUnionElemePromotionStorepromotionGetSingleStorePromotionRequest) *AlibabaAlscUnionElemePromotionStorepromotionGetRequest {
	s.QueryRequest = &v
	return s
}

func (req *AlibabaAlscUnionElemePromotionStorepromotionGetRequest) ToMap() map[string]interface{} {
	paramMap := make(map[string]interface{})
	if req.QueryRequest != nil {
		paramMap["query_request"] = util.ConvertStruct(*req.QueryRequest)
	}
	return paramMap
}

func (req *AlibabaAlscUnionElemePromotionStorepromotionGetRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	return fileMap
}
