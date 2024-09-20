package domain


type AlibabaAlscUnionElemePromotionItempromotionGetItemPromotionDto struct {
    /*
        商品ID     */
    ItemId  *string `json:"item_id,omitempty" `

    /*
        名称     */
    ItemName  *string `json:"item_name,omitempty" `

    /*
        图片     */
    Picture  *string `json:"picture,omitempty" `

    /*
        商品类型     */
    BizType  *string `json:"biz_type,omitempty" `

    /*
        原价（分）     */
    OriginalPriceCent  *int64 `json:"original_price_cent,omitempty" `

    /*
        售价（分）     */
    SellPriceCent  *int64 `json:"sell_price_cent,omitempty" `

    /*
        折扣     */
    Discount  *string `json:"discount,omitempty" `

    /*
        起始时间（秒）     */
    StartIme  *int64 `json:"start_ime,omitempty" `

    /*
        结束时间（秒）     */
    EndTime  *int64 `json:"end_time,omitempty" `

    /*
        佣金比例     */
    CommissionRate  *string `json:"commission_rate,omitempty" `

    /*
        预估佣金（分）     */
    Commission  *int64 `json:"commission,omitempty" `

    /*
        库存     */
    Stock  *int64 `json:"stock,omitempty" `

    /*
        适用城市数量     */
    ApplyCityCount  *int64 `json:"apply_city_count,omitempty" `

    /*
        适用门店数量     */
    ApplyShopCount  *int64 `json:"apply_shop_count,omitempty" `

    /*
        分享链接     */
    Link  *AlibabaAlscUnionElemePromotionItempromotionGetPromotionLink `json:"link,omitempty" `

    /*
        商品类型（1-商品券；2-代金券）     */
    ItemType  *int64 `json:"item_type,omitempty" `

    /*
        凭证信息     */
    Ticket  *AlibabaAlscUnionElemePromotionItempromotionGetItemTicket `json:"ticket,omitempty" `

}

func (s *AlibabaAlscUnionElemePromotionItempromotionGetItemPromotionDto) SetItemId(v string) *AlibabaAlscUnionElemePromotionItempromotionGetItemPromotionDto {
    s.ItemId = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionItempromotionGetItemPromotionDto) SetItemName(v string) *AlibabaAlscUnionElemePromotionItempromotionGetItemPromotionDto {
    s.ItemName = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionItempromotionGetItemPromotionDto) SetPicture(v string) *AlibabaAlscUnionElemePromotionItempromotionGetItemPromotionDto {
    s.Picture = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionItempromotionGetItemPromotionDto) SetBizType(v string) *AlibabaAlscUnionElemePromotionItempromotionGetItemPromotionDto {
    s.BizType = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionItempromotionGetItemPromotionDto) SetOriginalPriceCent(v int64) *AlibabaAlscUnionElemePromotionItempromotionGetItemPromotionDto {
    s.OriginalPriceCent = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionItempromotionGetItemPromotionDto) SetSellPriceCent(v int64) *AlibabaAlscUnionElemePromotionItempromotionGetItemPromotionDto {
    s.SellPriceCent = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionItempromotionGetItemPromotionDto) SetDiscount(v string) *AlibabaAlscUnionElemePromotionItempromotionGetItemPromotionDto {
    s.Discount = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionItempromotionGetItemPromotionDto) SetStartIme(v int64) *AlibabaAlscUnionElemePromotionItempromotionGetItemPromotionDto {
    s.StartIme = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionItempromotionGetItemPromotionDto) SetEndTime(v int64) *AlibabaAlscUnionElemePromotionItempromotionGetItemPromotionDto {
    s.EndTime = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionItempromotionGetItemPromotionDto) SetCommissionRate(v string) *AlibabaAlscUnionElemePromotionItempromotionGetItemPromotionDto {
    s.CommissionRate = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionItempromotionGetItemPromotionDto) SetCommission(v int64) *AlibabaAlscUnionElemePromotionItempromotionGetItemPromotionDto {
    s.Commission = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionItempromotionGetItemPromotionDto) SetStock(v int64) *AlibabaAlscUnionElemePromotionItempromotionGetItemPromotionDto {
    s.Stock = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionItempromotionGetItemPromotionDto) SetApplyCityCount(v int64) *AlibabaAlscUnionElemePromotionItempromotionGetItemPromotionDto {
    s.ApplyCityCount = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionItempromotionGetItemPromotionDto) SetApplyShopCount(v int64) *AlibabaAlscUnionElemePromotionItempromotionGetItemPromotionDto {
    s.ApplyShopCount = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionItempromotionGetItemPromotionDto) SetLink(v AlibabaAlscUnionElemePromotionItempromotionGetPromotionLink) *AlibabaAlscUnionElemePromotionItempromotionGetItemPromotionDto {
    s.Link = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionItempromotionGetItemPromotionDto) SetItemType(v int64) *AlibabaAlscUnionElemePromotionItempromotionGetItemPromotionDto {
    s.ItemType = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionItempromotionGetItemPromotionDto) SetTicket(v AlibabaAlscUnionElemePromotionItempromotionGetItemTicket) *AlibabaAlscUnionElemePromotionItempromotionGetItemPromotionDto {
    s.Ticket = &v
    return s
}
