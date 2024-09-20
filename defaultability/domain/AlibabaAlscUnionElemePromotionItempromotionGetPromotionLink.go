package domain


type AlibabaAlscUnionElemePromotionItempromotionGetPromotionLink struct {
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
        微信推广图片     */
    WxPicture  *string `json:"wx_picture,omitempty" `

    /*
        支付宝小程序链接     */
    AlipaySchemeUrl  *string `json:"alipay_scheme_url,omitempty" `

}

func (s *AlibabaAlscUnionElemePromotionItempromotionGetPromotionLink) SetWxAppid(v string) *AlibabaAlscUnionElemePromotionItempromotionGetPromotionLink {
    s.WxAppid = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionItempromotionGetPromotionLink) SetWxPath(v string) *AlibabaAlscUnionElemePromotionItempromotionGetPromotionLink {
    s.WxPath = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionItempromotionGetPromotionLink) SetWxMiniQrcode(v string) *AlibabaAlscUnionElemePromotionItempromotionGetPromotionLink {
    s.WxMiniQrcode = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionItempromotionGetPromotionLink) SetWxPicture(v string) *AlibabaAlscUnionElemePromotionItempromotionGetPromotionLink {
    s.WxPicture = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionItempromotionGetPromotionLink) SetAlipaySchemeUrl(v string) *AlibabaAlscUnionElemePromotionItempromotionGetPromotionLink {
    s.AlipaySchemeUrl = &v
    return s
}
