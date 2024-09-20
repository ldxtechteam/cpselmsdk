package domain


type AlibabaAlscUnionElemePromotionOfficialactivityGetTopTaobaoPromotionDto struct {
    /*
        推广链接     */
    H5Url  *string `json:"h5_url,omitempty" `

    /*
        短连接     */
    H5ShortUrl  *string `json:"h5_short_url,omitempty" `

    /*
        二维码     */
    TbQrCode  *string `json:"tb_qr_code,omitempty" `

    /*
        海报     */
    ImgUrl  *string `json:"img_url,omitempty" `

    /*
        appid     */
    AppId  *string `json:"app_id,omitempty" `

    /*
        apppath     */
    AppPath  *string `json:"app_path,omitempty" `

    /*
        schemeurl     */
    SchemeUrl  *string `json:"scheme_url,omitempty" `

}

func (s *AlibabaAlscUnionElemePromotionOfficialactivityGetTopTaobaoPromotionDto) SetH5Url(v string) *AlibabaAlscUnionElemePromotionOfficialactivityGetTopTaobaoPromotionDto {
    s.H5Url = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionOfficialactivityGetTopTaobaoPromotionDto) SetH5ShortUrl(v string) *AlibabaAlscUnionElemePromotionOfficialactivityGetTopTaobaoPromotionDto {
    s.H5ShortUrl = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionOfficialactivityGetTopTaobaoPromotionDto) SetTbQrCode(v string) *AlibabaAlscUnionElemePromotionOfficialactivityGetTopTaobaoPromotionDto {
    s.TbQrCode = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionOfficialactivityGetTopTaobaoPromotionDto) SetImgUrl(v string) *AlibabaAlscUnionElemePromotionOfficialactivityGetTopTaobaoPromotionDto {
    s.ImgUrl = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionOfficialactivityGetTopTaobaoPromotionDto) SetAppId(v string) *AlibabaAlscUnionElemePromotionOfficialactivityGetTopTaobaoPromotionDto {
    s.AppId = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionOfficialactivityGetTopTaobaoPromotionDto) SetAppPath(v string) *AlibabaAlscUnionElemePromotionOfficialactivityGetTopTaobaoPromotionDto {
    s.AppPath = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionOfficialactivityGetTopTaobaoPromotionDto) SetSchemeUrl(v string) *AlibabaAlscUnionElemePromotionOfficialactivityGetTopTaobaoPromotionDto {
    s.SchemeUrl = &v
    return s
}
