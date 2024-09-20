package request

import (
	"github.com/ldxtechteam/cpselmsdk/defaultability/domain"
	"github.com/ldxtechteam/cpselmsdk/util"
)

type AlibabaAlscUnionElemeToolOrderAttrbuteCheckRequest struct {
	/*
	   订单归因排查工具请求对象     */
	QueryRequest *domain.AlibabaAlscUnionElemeToolOrderAttrbuteCheckOrderCheckRequest `json:"query_request" required:"true" `
}

func (s *AlibabaAlscUnionElemeToolOrderAttrbuteCheckRequest) SetQueryRequest(v domain.AlibabaAlscUnionElemeToolOrderAttrbuteCheckOrderCheckRequest) *AlibabaAlscUnionElemeToolOrderAttrbuteCheckRequest {
	s.QueryRequest = &v
	return s
}

func (req *AlibabaAlscUnionElemeToolOrderAttrbuteCheckRequest) ToMap() map[string]interface{} {
	paramMap := make(map[string]interface{})
	if req.QueryRequest != nil {
		paramMap["query_request"] = util.ConvertStruct(*req.QueryRequest)
	}
	return paramMap
}

func (req *AlibabaAlscUnionElemeToolOrderAttrbuteCheckRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	return fileMap
}
