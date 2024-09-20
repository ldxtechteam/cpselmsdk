package domain


type AlibabaAlscUnionElemePromotionItempromotionGetSingleItemPromotionRequest struct {
    /*
        商品类型（hoard_coupon-囤券券）     */
    BizType  *string `json:"biz_type,omitempty" `

    /*
        推广位     */
    Pid  *string `json:"pid,omitempty" `

    /*
        商品ID     */
    ItemId  *string `json:"item_id,omitempty" `

    /*
        会员ID（需要联系运营申请）     */
    Sid  *string `json:"sid,omitempty" `

    /*
        是否返回微信推广图片     */
    IncludeWxImg  *bool `json:"include_wx_img,omitempty" `

}

func (s *AlibabaAlscUnionElemePromotionItempromotionGetSingleItemPromotionRequest) SetBizType(v string) *AlibabaAlscUnionElemePromotionItempromotionGetSingleItemPromotionRequest {
    s.BizType = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionItempromotionGetSingleItemPromotionRequest) SetPid(v string) *AlibabaAlscUnionElemePromotionItempromotionGetSingleItemPromotionRequest {
    s.Pid = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionItempromotionGetSingleItemPromotionRequest) SetItemId(v string) *AlibabaAlscUnionElemePromotionItempromotionGetSingleItemPromotionRequest {
    s.ItemId = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionItempromotionGetSingleItemPromotionRequest) SetSid(v string) *AlibabaAlscUnionElemePromotionItempromotionGetSingleItemPromotionRequest {
    s.Sid = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionItempromotionGetSingleItemPromotionRequest) SetIncludeWxImg(v bool) *AlibabaAlscUnionElemePromotionItempromotionGetSingleItemPromotionRequest {
    s.IncludeWxImg = &v
    return s
}
