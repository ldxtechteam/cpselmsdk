package domain


type AlibabaAlscUnionElemeStorepromotionReviewbwcQueryPromotionItem struct {
    /*
        小程序appId     */
    WxAppid  *string `json:"wx_appid,omitempty" `

    /*
        微信小程序path链接     */
    WxPath  *string `json:"wx_path,omitempty" `

}

func (s *AlibabaAlscUnionElemeStorepromotionReviewbwcQueryPromotionItem) SetWxAppid(v string) *AlibabaAlscUnionElemeStorepromotionReviewbwcQueryPromotionItem {
    s.WxAppid = &v
    return s
}
func (s *AlibabaAlscUnionElemeStorepromotionReviewbwcQueryPromotionItem) SetWxPath(v string) *AlibabaAlscUnionElemeStorepromotionReviewbwcQueryPromotionItem {
    s.WxPath = &v
    return s
}
