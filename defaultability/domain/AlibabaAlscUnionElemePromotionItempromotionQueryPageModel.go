package domain


type AlibabaAlscUnionElemePromotionItempromotionQueryPageModel struct {
    /*
        会话ID     */
    SessionId  *string `json:"session_id,omitempty" `

    /*
        数据     */
    Records  *[]AlibabaAlscUnionElemePromotionItempromotionQueryItemPromotionDto `json:"records,omitempty" `

    /*
        页码     */
    PageNumber  *int64 `json:"page_number,omitempty" `

    /*
        每页数目     */
    PageSize  *int64 `json:"page_size,omitempty" `

    /*
        总数     */
    Total  *int64 `json:"total,omitempty" `

}

func (s *AlibabaAlscUnionElemePromotionItempromotionQueryPageModel) SetSessionId(v string) *AlibabaAlscUnionElemePromotionItempromotionQueryPageModel {
    s.SessionId = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionItempromotionQueryPageModel) SetRecords(v []AlibabaAlscUnionElemePromotionItempromotionQueryItemPromotionDto) *AlibabaAlscUnionElemePromotionItempromotionQueryPageModel {
    s.Records = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionItempromotionQueryPageModel) SetPageNumber(v int64) *AlibabaAlscUnionElemePromotionItempromotionQueryPageModel {
    s.PageNumber = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionItempromotionQueryPageModel) SetPageSize(v int64) *AlibabaAlscUnionElemePromotionItempromotionQueryPageModel {
    s.PageSize = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionItempromotionQueryPageModel) SetTotal(v int64) *AlibabaAlscUnionElemePromotionItempromotionQueryPageModel {
    s.Total = &v
    return s
}
