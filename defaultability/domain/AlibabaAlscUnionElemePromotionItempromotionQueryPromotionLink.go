package domain


type AlibabaAlscUnionElemePromotionItempromotionQueryPromotionLink struct {
    /*
        微信小程序appId     */
    WxAppid  *string `json:"wx_appid,omitempty" `

    /*
        微信小程序path链接     */
    WxPath  *string `json:"wx_path,omitempty" `

    /*
        微信小程序二维码     */
    WxMiniQrcode  *string `json:"wx_mini_qrcode,omitempty" `

    /*
        支付宝schemal链接     */
    AlipaySchemeUrl  *string `json:"alipay_scheme_url,omitempty" `

}

func (s *AlibabaAlscUnionElemePromotionItempromotionQueryPromotionLink) SetWxAppid(v string) *AlibabaAlscUnionElemePromotionItempromotionQueryPromotionLink {
    s.WxAppid = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionItempromotionQueryPromotionLink) SetWxPath(v string) *AlibabaAlscUnionElemePromotionItempromotionQueryPromotionLink {
    s.WxPath = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionItempromotionQueryPromotionLink) SetWxMiniQrcode(v string) *AlibabaAlscUnionElemePromotionItempromotionQueryPromotionLink {
    s.WxMiniQrcode = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionItempromotionQueryPromotionLink) SetAlipaySchemeUrl(v string) *AlibabaAlscUnionElemePromotionItempromotionQueryPromotionLink {
    s.AlipaySchemeUrl = &v
    return s
}
