package domain


type AlibabaAlscUnionElemePromotionItempromotionQueryItemPromotionQueryRequest struct {
    /*
        会话ID（查询第一页为空，从第二页开始赋值，取值来自第一页返回结果）     */
    SessionId  *string `json:"session_id,omitempty" `

    /*
        商品类型（hoard_coupon-囤券券）     */
    BizType  *string `json:"biz_type,omitempty" `

    /*
        推广位     */
    Pid  *string `json:"pid,omitempty" `

    /*
        城市编码（国标）     */
    CityCode  *string `json:"city_code,omitempty" `

    /*
        排序（normal-佣金倒序，commission_desc-佣金倒序，commission_rate_desc-佣金比例倒序，sell_price_asc-价格正序，sell_price_desc-价格倒序）     */
    SortType  *string `json:"sort_type,omitempty" `

    /*
        请求页（从1开始）     */
    PageNumber  *int64 `json:"page_number,omitempty" `

    /*
        每页数（1~20）     */
    PageSize  *int64 `json:"page_size,omitempty" `

    /*
        会员ID（需要联系运营申请）     */
    Sid  *string `json:"sid,omitempty" `

    /*
        品牌搜索     */
    SearchContent  *string `json:"search_content,omitempty" `

    /*
        商品类型，多值'|'分隔，默认全部（1-商品券；2-代金券）     */
    ItemType  *string `json:"item_type,omitempty" `

}

func (s *AlibabaAlscUnionElemePromotionItempromotionQueryItemPromotionQueryRequest) SetSessionId(v string) *AlibabaAlscUnionElemePromotionItempromotionQueryItemPromotionQueryRequest {
    s.SessionId = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionItempromotionQueryItemPromotionQueryRequest) SetBizType(v string) *AlibabaAlscUnionElemePromotionItempromotionQueryItemPromotionQueryRequest {
    s.BizType = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionItempromotionQueryItemPromotionQueryRequest) SetPid(v string) *AlibabaAlscUnionElemePromotionItempromotionQueryItemPromotionQueryRequest {
    s.Pid = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionItempromotionQueryItemPromotionQueryRequest) SetCityCode(v string) *AlibabaAlscUnionElemePromotionItempromotionQueryItemPromotionQueryRequest {
    s.CityCode = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionItempromotionQueryItemPromotionQueryRequest) SetSortType(v string) *AlibabaAlscUnionElemePromotionItempromotionQueryItemPromotionQueryRequest {
    s.SortType = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionItempromotionQueryItemPromotionQueryRequest) SetPageNumber(v int64) *AlibabaAlscUnionElemePromotionItempromotionQueryItemPromotionQueryRequest {
    s.PageNumber = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionItempromotionQueryItemPromotionQueryRequest) SetPageSize(v int64) *AlibabaAlscUnionElemePromotionItempromotionQueryItemPromotionQueryRequest {
    s.PageSize = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionItempromotionQueryItemPromotionQueryRequest) SetSid(v string) *AlibabaAlscUnionElemePromotionItempromotionQueryItemPromotionQueryRequest {
    s.Sid = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionItempromotionQueryItemPromotionQueryRequest) SetSearchContent(v string) *AlibabaAlscUnionElemePromotionItempromotionQueryItemPromotionQueryRequest {
    s.SearchContent = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionItempromotionQueryItemPromotionQueryRequest) SetItemType(v string) *AlibabaAlscUnionElemePromotionItempromotionQueryItemPromotionQueryRequest {
    s.ItemType = &v
    return s
}
