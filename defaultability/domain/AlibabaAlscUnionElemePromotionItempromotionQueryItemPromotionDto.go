package domain


type AlibabaAlscUnionElemePromotionItempromotionQueryItemPromotionDto struct {
    /*
        商品ID     */
    ItemId  *string `json:"item_id,omitempty" `

    /*
        标题     */
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
    Link  *AlibabaAlscUnionElemePromotionItempromotionQueryPromotionLink `json:"link,omitempty" `

    /*
        商品类型（1-商品券；2-代金券）     */
    ItemType  *int64 `json:"item_type,omitempty" `

    /*
        凭证信息     */
    Ticket  *AlibabaAlscUnionElemePromotionItempromotionQueryItemTicket `json:"ticket,omitempty" `

}

func (s *AlibabaAlscUnionElemePromotionItempromotionQueryItemPromotionDto) SetItemId(v string) *AlibabaAlscUnionElemePromotionItempromotionQueryItemPromotionDto {
    s.ItemId = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionItempromotionQueryItemPromotionDto) SetItemName(v string) *AlibabaAlscUnionElemePromotionItempromotionQueryItemPromotionDto {
    s.ItemName = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionItempromotionQueryItemPromotionDto) SetPicture(v string) *AlibabaAlscUnionElemePromotionItempromotionQueryItemPromotionDto {
    s.Picture = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionItempromotionQueryItemPromotionDto) SetBizType(v string) *AlibabaAlscUnionElemePromotionItempromotionQueryItemPromotionDto {
    s.BizType = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionItempromotionQueryItemPromotionDto) SetOriginalPriceCent(v int64) *AlibabaAlscUnionElemePromotionItempromotionQueryItemPromotionDto {
    s.OriginalPriceCent = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionItempromotionQueryItemPromotionDto) SetSellPriceCent(v int64) *AlibabaAlscUnionElemePromotionItempromotionQueryItemPromotionDto {
    s.SellPriceCent = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionItempromotionQueryItemPromotionDto) SetDiscount(v string) *AlibabaAlscUnionElemePromotionItempromotionQueryItemPromotionDto {
    s.Discount = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionItempromotionQueryItemPromotionDto) SetStartIme(v int64) *AlibabaAlscUnionElemePromotionItempromotionQueryItemPromotionDto {
    s.StartIme = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionItempromotionQueryItemPromotionDto) SetEndTime(v int64) *AlibabaAlscUnionElemePromotionItempromotionQueryItemPromotionDto {
    s.EndTime = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionItempromotionQueryItemPromotionDto) SetCommissionRate(v string) *AlibabaAlscUnionElemePromotionItempromotionQueryItemPromotionDto {
    s.CommissionRate = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionItempromotionQueryItemPromotionDto) SetCommission(v int64) *AlibabaAlscUnionElemePromotionItempromotionQueryItemPromotionDto {
    s.Commission = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionItempromotionQueryItemPromotionDto) SetStock(v int64) *AlibabaAlscUnionElemePromotionItempromotionQueryItemPromotionDto {
    s.Stock = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionItempromotionQueryItemPromotionDto) SetApplyCityCount(v int64) *AlibabaAlscUnionElemePromotionItempromotionQueryItemPromotionDto {
    s.ApplyCityCount = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionItempromotionQueryItemPromotionDto) SetApplyShopCount(v int64) *AlibabaAlscUnionElemePromotionItempromotionQueryItemPromotionDto {
    s.ApplyShopCount = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionItempromotionQueryItemPromotionDto) SetLink(v AlibabaAlscUnionElemePromotionItempromotionQueryPromotionLink) *AlibabaAlscUnionElemePromotionItempromotionQueryItemPromotionDto {
    s.Link = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionItempromotionQueryItemPromotionDto) SetItemType(v int64) *AlibabaAlscUnionElemePromotionItempromotionQueryItemPromotionDto {
    s.ItemType = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionItempromotionQueryItemPromotionDto) SetTicket(v AlibabaAlscUnionElemePromotionItempromotionQueryItemTicket) *AlibabaAlscUnionElemePromotionItempromotionQueryItemPromotionDto {
    s.Ticket = &v
    return s
}
