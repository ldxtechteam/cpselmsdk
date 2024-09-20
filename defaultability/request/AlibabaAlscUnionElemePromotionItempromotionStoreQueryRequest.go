package request

import (
	"github.com/ldxtechteam/cpselmsdk/defaultability/domain"
	"github.com/ldxtechteam/cpselmsdk/util"
)

type AlibabaAlscUnionElemePromotionItempromotionStoreQueryRequest struct {
	/*
	   商品门店关系查询rquest     */
	QueryRequest *domain.AlibabaAlscUnionElemePromotionItempromotionStoreQueryItemPromotionShopRequest `json:"query_request" required:"true" `
}

func (s *AlibabaAlscUnionElemePromotionItempromotionStoreQueryRequest) SetQueryRequest(v domain.AlibabaAlscUnionElemePromotionItempromotionStoreQueryItemPromotionShopRequest) *AlibabaAlscUnionElemePromotionItempromotionStoreQueryRequest {
	s.QueryRequest = &v
	return s
}

func (req *AlibabaAlscUnionElemePromotionItempromotionStoreQueryRequest) ToMap() map[string]interface{} {
	paramMap := make(map[string]interface{})
	if req.QueryRequest != nil {
		paramMap["query_request"] = util.ConvertStruct(*req.QueryRequest)
	}
	return paramMap
}

func (req *AlibabaAlscUnionElemePromotionItempromotionStoreQueryRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	return fileMap
}
