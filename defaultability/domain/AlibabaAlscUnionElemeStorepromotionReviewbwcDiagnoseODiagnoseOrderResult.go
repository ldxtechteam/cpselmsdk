package domain


type AlibabaAlscUnionElemeStorepromotionReviewbwcDiagnoseODiagnoseOrderResult struct {
    /*
        下单手机号（后2位）     */
    OrderMobile  *string `json:"order_mobile,omitempty" `

    /*
        饿了么门店ID     */
    ShopId  *string `json:"shop_id,omitempty" `

    /*
        店铺名称     */
    ShopName  *string `json:"shop_name,omitempty" `

}

func (s *AlibabaAlscUnionElemeStorepromotionReviewbwcDiagnoseODiagnoseOrderResult) SetOrderMobile(v string) *AlibabaAlscUnionElemeStorepromotionReviewbwcDiagnoseODiagnoseOrderResult {
    s.OrderMobile = &v
    return s
}
func (s *AlibabaAlscUnionElemeStorepromotionReviewbwcDiagnoseODiagnoseOrderResult) SetShopId(v string) *AlibabaAlscUnionElemeStorepromotionReviewbwcDiagnoseODiagnoseOrderResult {
    s.ShopId = &v
    return s
}
func (s *AlibabaAlscUnionElemeStorepromotionReviewbwcDiagnoseODiagnoseOrderResult) SetShopName(v string) *AlibabaAlscUnionElemeStorepromotionReviewbwcDiagnoseODiagnoseOrderResult {
    s.ShopName = &v
    return s
}
