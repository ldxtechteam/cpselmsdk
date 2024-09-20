package domain


type AlibabaAlscUnionElemeToolOrderAttrbuteCheckOrderInfo struct {
    /*
        订单号     */
    OrderNo  *string `json:"order_no,omitempty" `

    /*
        店铺名称     */
    ShopName  *string `json:"shop_name,omitempty" `

    /*
        订单类型     */
    TypeName  *string `json:"type_name,omitempty" `

    /*
        订单创建时间     */
    CreateTime  *string `json:"create_time,omitempty" `

    /*
        订单完成时间     */
    FinishTime  *string `json:"finish_time,omitempty" `

    /*
        订单结算时间     */
    SettleDate  *string `json:"settle_date,omitempty" `

}

func (s *AlibabaAlscUnionElemeToolOrderAttrbuteCheckOrderInfo) SetOrderNo(v string) *AlibabaAlscUnionElemeToolOrderAttrbuteCheckOrderInfo {
    s.OrderNo = &v
    return s
}
func (s *AlibabaAlscUnionElemeToolOrderAttrbuteCheckOrderInfo) SetShopName(v string) *AlibabaAlscUnionElemeToolOrderAttrbuteCheckOrderInfo {
    s.ShopName = &v
    return s
}
func (s *AlibabaAlscUnionElemeToolOrderAttrbuteCheckOrderInfo) SetTypeName(v string) *AlibabaAlscUnionElemeToolOrderAttrbuteCheckOrderInfo {
    s.TypeName = &v
    return s
}
func (s *AlibabaAlscUnionElemeToolOrderAttrbuteCheckOrderInfo) SetCreateTime(v string) *AlibabaAlscUnionElemeToolOrderAttrbuteCheckOrderInfo {
    s.CreateTime = &v
    return s
}
func (s *AlibabaAlscUnionElemeToolOrderAttrbuteCheckOrderInfo) SetFinishTime(v string) *AlibabaAlscUnionElemeToolOrderAttrbuteCheckOrderInfo {
    s.FinishTime = &v
    return s
}
func (s *AlibabaAlscUnionElemeToolOrderAttrbuteCheckOrderInfo) SetSettleDate(v string) *AlibabaAlscUnionElemeToolOrderAttrbuteCheckOrderInfo {
    s.SettleDate = &v
    return s
}
