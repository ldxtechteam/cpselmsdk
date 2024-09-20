package request

import (
	"github.com/ldxtechteam/cpselmsdk/defaultability/domain"
	"github.com/ldxtechteam/cpselmsdk/util"
)

type AlibabaAlscUnionElemePromotionStorepromotionBatchGetRequest struct {
	/*
	   查询rquest     */
	QueryRequest *domain.AlibabaAlscUnionElemePromotionStorepromotionBatchGetBatchQueryStorePromotionRequest `json:"query_request" required:"true" `
}

func (s *AlibabaAlscUnionElemePromotionStorepromotionBatchGetRequest) SetQueryRequest(v domain.AlibabaAlscUnionElemePromotionStorepromotionBatchGetBatchQueryStorePromotionRequest) *AlibabaAlscUnionElemePromotionStorepromotionBatchGetRequest {
	s.QueryRequest = &v
	return s
}

func (req *AlibabaAlscUnionElemePromotionStorepromotionBatchGetRequest) ToMap() map[string]interface{} {
	paramMap := make(map[string]interface{})
	if req.QueryRequest != nil {
		paramMap["query_request"] = util.ConvertStruct(*req.QueryRequest)
	}
	return paramMap
}

func (req *AlibabaAlscUnionElemePromotionStorepromotionBatchGetRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	return fileMap
}
