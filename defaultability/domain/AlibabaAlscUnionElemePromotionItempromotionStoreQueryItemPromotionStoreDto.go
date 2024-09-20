package domain


type AlibabaAlscUnionElemePromotionItempromotionStoreQueryItemPromotionStoreDto struct {
    /*
        门店ID     */
    StoreId  *string `json:"store_id,omitempty" `

}

func (s *AlibabaAlscUnionElemePromotionItempromotionStoreQueryItemPromotionStoreDto) SetStoreId(v string) *AlibabaAlscUnionElemePromotionItempromotionStoreQueryItemPromotionStoreDto {
    s.StoreId = &v
    return s
}
