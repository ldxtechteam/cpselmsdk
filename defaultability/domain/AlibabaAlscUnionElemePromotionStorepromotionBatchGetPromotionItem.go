package domain


type AlibabaAlscUnionElemePromotionStorepromotionBatchGetPromotionItem struct {
    /*
        标题     */
    Title  *string `json:"title,omitempty" `

    /*
        原价     */
    OriginPrice  *string `json:"origin_price,omitempty" `

    /*
        现价     */
    Price  *string `json:"price,omitempty" `

    /*
        图片     */
    Picture  *string `json:"picture,omitempty" `

}

func (s *AlibabaAlscUnionElemePromotionStorepromotionBatchGetPromotionItem) SetTitle(v string) *AlibabaAlscUnionElemePromotionStorepromotionBatchGetPromotionItem {
    s.Title = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionStorepromotionBatchGetPromotionItem) SetOriginPrice(v string) *AlibabaAlscUnionElemePromotionStorepromotionBatchGetPromotionItem {
    s.OriginPrice = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionStorepromotionBatchGetPromotionItem) SetPrice(v string) *AlibabaAlscUnionElemePromotionStorepromotionBatchGetPromotionItem {
    s.Price = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionStorepromotionBatchGetPromotionItem) SetPicture(v string) *AlibabaAlscUnionElemePromotionStorepromotionBatchGetPromotionItem {
    s.Picture = &v
    return s
}
