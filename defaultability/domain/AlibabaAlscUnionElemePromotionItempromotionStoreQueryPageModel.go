package domain


type AlibabaAlscUnionElemePromotionItempromotionStoreQueryPageModel struct {
    /*
        分页记录     */
    Records  *[]AlibabaAlscUnionElemePromotionItempromotionStoreQueryItemPromotionStoreDto `json:"records,omitempty" `

    /*
        总数     */
    TotalCount  *int64 `json:"total_count,omitempty" `

}

func (s *AlibabaAlscUnionElemePromotionItempromotionStoreQueryPageModel) SetRecords(v []AlibabaAlscUnionElemePromotionItempromotionStoreQueryItemPromotionStoreDto) *AlibabaAlscUnionElemePromotionItempromotionStoreQueryPageModel {
    s.Records = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionItempromotionStoreQueryPageModel) SetTotalCount(v int64) *AlibabaAlscUnionElemePromotionItempromotionStoreQueryPageModel {
    s.TotalCount = &v
    return s
}
