package domain


type AlibabaAlscUnionElemePromotionOfficialactivityGetPromotionLink struct {
    /*
        小程序appId     */
    WxAppid  *string `json:"wx_appid,omitempty" `

    /*
        微信小程序path链接     */
    WxPath  *string `json:"wx_path,omitempty" `

    /*
        推广图片地址     */
    Picture  *string `json:"picture,omitempty" `

    /*
        支付宝小程序推广链接     */
    AlipayMiniUrl  *string `json:"alipay_mini_url,omitempty" `

    /*
        h5推广地址     */
    H5Url  *string `json:"h5_url,omitempty" `

    /*
        淘宝二维码图片地址     */
    TbQrCode  *string `json:"tb_qr_code,omitempty" `

    /*
        微信独立二维码     */
    MiniQrcode  *string `json:"mini_qrcode,omitempty" `

    /*
        淘宝独立二维码     */
    TbMiniQrcode  *string `json:"tb_mini_qrcode,omitempty" `

    /*
        饿了么唤端链接     */
    EleSchemeUrl  *string `json:"ele_scheme_url,omitempty" `

    /*
        h5推广地址短链     */
    H5ShortLink  *string `json:"h5_short_link,omitempty" `

    /*
        H5二维码图片地址     */
    H5MiniQrcode  *string `json:"h5_mini_qrcode,omitempty" `

    /*
        饿了么口令（有效期30天，建议到期前重新获取）     */
    ElemeWord  *string `json:"eleme_word,omitempty" `

    /*
        淘宝唤端链接     */
    TbSchemeUrl  *string `json:"tb_scheme_url,omitempty" `

    /*
        淘口令（有效期30天，建议到期前重新获取）     */
    TaobaoWord  *string `json:"taobao_word,omitempty" `

    /*
        带文案淘口令（有效期30天，建议到期前重新获取）     */
    FullTaobaoWord  *string `json:"full_taobao_word,omitempty" `

    /*
        推广链接     */
    H5Promotion  *AlibabaAlscUnionElemePromotionOfficialactivityGetTopH5PromotionDto `json:"h5_promotion,omitempty" `

    /*
        淘宝链接     */
    TaobaoPromotion  *AlibabaAlscUnionElemePromotionOfficialactivityGetTopTaobaoPromotionDto `json:"taobao_promotion,omitempty" `

    /*
        微信链接     */
    WxPromotion  *AlibabaAlscUnionElemePromotionOfficialactivityGetTopWxPromotionDto `json:"wx_promotion,omitempty" `

    /*
        支付宝链接     */
    AlipayPromotion  *AlibabaAlscUnionElemePromotionOfficialactivityGetTopAlipayPromotionDto `json:"alipay_promotion,omitempty" `

    /*
        app链接     */
    AppPromotion  *AlibabaAlscUnionElemePromotionOfficialactivityGetTopAppPromotionDto `json:"app_promotion,omitempty" `

}

func (s *AlibabaAlscUnionElemePromotionOfficialactivityGetPromotionLink) SetWxAppid(v string) *AlibabaAlscUnionElemePromotionOfficialactivityGetPromotionLink {
    s.WxAppid = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionOfficialactivityGetPromotionLink) SetWxPath(v string) *AlibabaAlscUnionElemePromotionOfficialactivityGetPromotionLink {
    s.WxPath = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionOfficialactivityGetPromotionLink) SetPicture(v string) *AlibabaAlscUnionElemePromotionOfficialactivityGetPromotionLink {
    s.Picture = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionOfficialactivityGetPromotionLink) SetAlipayMiniUrl(v string) *AlibabaAlscUnionElemePromotionOfficialactivityGetPromotionLink {
    s.AlipayMiniUrl = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionOfficialactivityGetPromotionLink) SetH5Url(v string) *AlibabaAlscUnionElemePromotionOfficialactivityGetPromotionLink {
    s.H5Url = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionOfficialactivityGetPromotionLink) SetTbQrCode(v string) *AlibabaAlscUnionElemePromotionOfficialactivityGetPromotionLink {
    s.TbQrCode = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionOfficialactivityGetPromotionLink) SetMiniQrcode(v string) *AlibabaAlscUnionElemePromotionOfficialactivityGetPromotionLink {
    s.MiniQrcode = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionOfficialactivityGetPromotionLink) SetTbMiniQrcode(v string) *AlibabaAlscUnionElemePromotionOfficialactivityGetPromotionLink {
    s.TbMiniQrcode = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionOfficialactivityGetPromotionLink) SetEleSchemeUrl(v string) *AlibabaAlscUnionElemePromotionOfficialactivityGetPromotionLink {
    s.EleSchemeUrl = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionOfficialactivityGetPromotionLink) SetH5ShortLink(v string) *AlibabaAlscUnionElemePromotionOfficialactivityGetPromotionLink {
    s.H5ShortLink = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionOfficialactivityGetPromotionLink) SetH5MiniQrcode(v string) *AlibabaAlscUnionElemePromotionOfficialactivityGetPromotionLink {
    s.H5MiniQrcode = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionOfficialactivityGetPromotionLink) SetElemeWord(v string) *AlibabaAlscUnionElemePromotionOfficialactivityGetPromotionLink {
    s.ElemeWord = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionOfficialactivityGetPromotionLink) SetTbSchemeUrl(v string) *AlibabaAlscUnionElemePromotionOfficialactivityGetPromotionLink {
    s.TbSchemeUrl = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionOfficialactivityGetPromotionLink) SetTaobaoWord(v string) *AlibabaAlscUnionElemePromotionOfficialactivityGetPromotionLink {
    s.TaobaoWord = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionOfficialactivityGetPromotionLink) SetFullTaobaoWord(v string) *AlibabaAlscUnionElemePromotionOfficialactivityGetPromotionLink {
    s.FullTaobaoWord = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionOfficialactivityGetPromotionLink) SetH5Promotion(v AlibabaAlscUnionElemePromotionOfficialactivityGetTopH5PromotionDto) *AlibabaAlscUnionElemePromotionOfficialactivityGetPromotionLink {
    s.H5Promotion = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionOfficialactivityGetPromotionLink) SetTaobaoPromotion(v AlibabaAlscUnionElemePromotionOfficialactivityGetTopTaobaoPromotionDto) *AlibabaAlscUnionElemePromotionOfficialactivityGetPromotionLink {
    s.TaobaoPromotion = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionOfficialactivityGetPromotionLink) SetWxPromotion(v AlibabaAlscUnionElemePromotionOfficialactivityGetTopWxPromotionDto) *AlibabaAlscUnionElemePromotionOfficialactivityGetPromotionLink {
    s.WxPromotion = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionOfficialactivityGetPromotionLink) SetAlipayPromotion(v AlibabaAlscUnionElemePromotionOfficialactivityGetTopAlipayPromotionDto) *AlibabaAlscUnionElemePromotionOfficialactivityGetPromotionLink {
    s.AlipayPromotion = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionOfficialactivityGetPromotionLink) SetAppPromotion(v AlibabaAlscUnionElemePromotionOfficialactivityGetTopAppPromotionDto) *AlibabaAlscUnionElemePromotionOfficialactivityGetPromotionLink {
    s.AppPromotion = &v
    return s
}
