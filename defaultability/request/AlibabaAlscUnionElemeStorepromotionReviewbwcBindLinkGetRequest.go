package request

import (
	"github.com/ldxtechteam/cpselmsdk/defaultability/domain"
	"github.com/ldxtechteam/cpselmsdk/util"
)

type AlibabaAlscUnionElemeStorepromotionReviewbwcBindLinkGetRequest struct {
	/*
	   查询rquest     */
	QueryRequest *domain.AlibabaAlscUnionElemeStorepromotionReviewbwcBindLinkGetReviewBwcSidBindLinkRequest `json:"query_request" required:"true" `
}

func (s *AlibabaAlscUnionElemeStorepromotionReviewbwcBindLinkGetRequest) SetQueryRequest(v domain.AlibabaAlscUnionElemeStorepromotionReviewbwcBindLinkGetReviewBwcSidBindLinkRequest) *AlibabaAlscUnionElemeStorepromotionReviewbwcBindLinkGetRequest {
	s.QueryRequest = &v
	return s
}

func (req *AlibabaAlscUnionElemeStorepromotionReviewbwcBindLinkGetRequest) ToMap() map[string]interface{} {
	paramMap := make(map[string]interface{})
	if req.QueryRequest != nil {
		paramMap["query_request"] = util.ConvertStruct(*req.QueryRequest)
	}
	return paramMap
}

func (req *AlibabaAlscUnionElemeStorepromotionReviewbwcBindLinkGetRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	return fileMap
}
