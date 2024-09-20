package domain


type AlibabaAlscUnionElemePromotionOfficialactivityGetTopAlipayPromotionDto struct {
    /*
        appId     */
    AppId  *string `json:"app_id,omitempty" `

    /*
        appPath     */
    AppPath  *string `json:"app_path,omitempty" `

    /*
        scheme唤端地址     */
    AlipaySchemeUrl  *string `json:"alipay_scheme_url,omitempty" `

    /*
        h5地址     */
    H5Url  *string `json:"h5_url,omitempty" `

    /*
        支付宝口令     */
    AlipayWatchword  *string `json:"alipay_watchword,omitempty" `

    /*
        支付宝完整口令     */
    AlipayWatchwordSuggest  *string `json:"alipay_watchword_suggest,omitempty" `

    /*
        海报地址     */
    ImgUrl  *string `json:"img_url,omitempty" `

    /*
        二维码地址     */
    AlipayQrCode  *string `json:"alipay_qr_code,omitempty" `

    /*
        h5短连接     */
    ShortLink  *string `json:"short_link,omitempty" `

}

func (s *AlibabaAlscUnionElemePromotionOfficialactivityGetTopAlipayPromotionDto) SetAppId(v string) *AlibabaAlscUnionElemePromotionOfficialactivityGetTopAlipayPromotionDto {
    s.AppId = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionOfficialactivityGetTopAlipayPromotionDto) SetAppPath(v string) *AlibabaAlscUnionElemePromotionOfficialactivityGetTopAlipayPromotionDto {
    s.AppPath = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionOfficialactivityGetTopAlipayPromotionDto) SetAlipaySchemeUrl(v string) *AlibabaAlscUnionElemePromotionOfficialactivityGetTopAlipayPromotionDto {
    s.AlipaySchemeUrl = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionOfficialactivityGetTopAlipayPromotionDto) SetH5Url(v string) *AlibabaAlscUnionElemePromotionOfficialactivityGetTopAlipayPromotionDto {
    s.H5Url = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionOfficialactivityGetTopAlipayPromotionDto) SetAlipayWatchword(v string) *AlibabaAlscUnionElemePromotionOfficialactivityGetTopAlipayPromotionDto {
    s.AlipayWatchword = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionOfficialactivityGetTopAlipayPromotionDto) SetAlipayWatchwordSuggest(v string) *AlibabaAlscUnionElemePromotionOfficialactivityGetTopAlipayPromotionDto {
    s.AlipayWatchwordSuggest = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionOfficialactivityGetTopAlipayPromotionDto) SetImgUrl(v string) *AlibabaAlscUnionElemePromotionOfficialactivityGetTopAlipayPromotionDto {
    s.ImgUrl = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionOfficialactivityGetTopAlipayPromotionDto) SetAlipayQrCode(v string) *AlibabaAlscUnionElemePromotionOfficialactivityGetTopAlipayPromotionDto {
    s.AlipayQrCode = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionOfficialactivityGetTopAlipayPromotionDto) SetShortLink(v string) *AlibabaAlscUnionElemePromotionOfficialactivityGetTopAlipayPromotionDto {
    s.ShortLink = &v
    return s
}
