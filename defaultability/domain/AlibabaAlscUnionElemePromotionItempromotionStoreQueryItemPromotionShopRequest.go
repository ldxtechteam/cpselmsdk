package domain


type AlibabaAlscUnionElemePromotionItempromotionStoreQueryItemPromotionShopRequest struct {
    /*
        商品ID     */
    ItemId  *string `json:"item_id,omitempty" `

    /*
        请求页（从1开始）     */
    PageNumber  *int64 `json:"page_number,omitempty" `

    /*
        每页数（1~20）     */
    PageSize  *int64 `json:"page_size,omitempty" `

}

func (s *AlibabaAlscUnionElemePromotionItempromotionStoreQueryItemPromotionShopRequest) SetItemId(v string) *AlibabaAlscUnionElemePromotionItempromotionStoreQueryItemPromotionShopRequest {
    s.ItemId = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionItempromotionStoreQueryItemPromotionShopRequest) SetPageNumber(v int64) *AlibabaAlscUnionElemePromotionItempromotionStoreQueryItemPromotionShopRequest {
    s.PageNumber = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionItempromotionStoreQueryItemPromotionShopRequest) SetPageSize(v int64) *AlibabaAlscUnionElemePromotionItempromotionStoreQueryItemPromotionShopRequest {
    s.PageSize = &v
    return s
}
