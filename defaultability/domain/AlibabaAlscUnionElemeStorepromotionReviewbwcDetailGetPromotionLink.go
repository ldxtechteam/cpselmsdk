package domain


type AlibabaAlscUnionElemeStorepromotionReviewbwcDetailGetPromotionLink struct {
    /*
        小程序appId     */
    WxAppid  *string `json:"wx_appid,omitempty" `

    /*
        微信小程序path链接     */
    WxPath  *string `json:"wx_path,omitempty" `

}

func (s *AlibabaAlscUnionElemeStorepromotionReviewbwcDetailGetPromotionLink) SetWxAppid(v string) *AlibabaAlscUnionElemeStorepromotionReviewbwcDetailGetPromotionLink {
    s.WxAppid = &v
    return s
}
func (s *AlibabaAlscUnionElemeStorepromotionReviewbwcDetailGetPromotionLink) SetWxPath(v string) *AlibabaAlscUnionElemeStorepromotionReviewbwcDetailGetPromotionLink {
    s.WxPath = &v
    return s
}
