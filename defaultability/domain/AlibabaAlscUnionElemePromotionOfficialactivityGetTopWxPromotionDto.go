package domain


type AlibabaAlscUnionElemePromotionOfficialactivityGetTopWxPromotionDto struct {
    /*
        appid     */
    WxAppId  *string `json:"wx_app_id,omitempty" `

    /*
        apppath     */
    WxPath  *string `json:"wx_path,omitempty" `

    /*
        二维码     */
    WxQrCode  *string `json:"wx_qr_code,omitempty" `

    /*
        海报     */
    ImgUrl  *string `json:"img_url,omitempty" `

    /*
        立减的appid     */
    ReduceWxAppId  *string `json:"reduce_wx_app_id,omitempty" `

    /*
        立减的apppath     */
    ReduceWxPath  *string `json:"reduce_wx_path,omitempty" `

    /*
        立减的二维码     */
    ReduceWxQrCode  *string `json:"reduce_wx_qr_code,omitempty" `

    /*
        立减的海报     */
    ReduceImgUrl  *string `json:"reduce_img_url,omitempty" `

    /*
        h5短链     */
    H5ShortLink  *string `json:"h5_short_link,omitempty" `

    /*
        媒体出资appid     */
    MarketWxAppId  *string `json:"market_wx_app_id,omitempty" `

    /*
        媒体出资apppath     */
    MarketWxAppPath  *string `json:"market_wx_app_path,omitempty" `

}

func (s *AlibabaAlscUnionElemePromotionOfficialactivityGetTopWxPromotionDto) SetWxAppId(v string) *AlibabaAlscUnionElemePromotionOfficialactivityGetTopWxPromotionDto {
    s.WxAppId = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionOfficialactivityGetTopWxPromotionDto) SetWxPath(v string) *AlibabaAlscUnionElemePromotionOfficialactivityGetTopWxPromotionDto {
    s.WxPath = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionOfficialactivityGetTopWxPromotionDto) SetWxQrCode(v string) *AlibabaAlscUnionElemePromotionOfficialactivityGetTopWxPromotionDto {
    s.WxQrCode = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionOfficialactivityGetTopWxPromotionDto) SetImgUrl(v string) *AlibabaAlscUnionElemePromotionOfficialactivityGetTopWxPromotionDto {
    s.ImgUrl = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionOfficialactivityGetTopWxPromotionDto) SetReduceWxAppId(v string) *AlibabaAlscUnionElemePromotionOfficialactivityGetTopWxPromotionDto {
    s.ReduceWxAppId = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionOfficialactivityGetTopWxPromotionDto) SetReduceWxPath(v string) *AlibabaAlscUnionElemePromotionOfficialactivityGetTopWxPromotionDto {
    s.ReduceWxPath = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionOfficialactivityGetTopWxPromotionDto) SetReduceWxQrCode(v string) *AlibabaAlscUnionElemePromotionOfficialactivityGetTopWxPromotionDto {
    s.ReduceWxQrCode = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionOfficialactivityGetTopWxPromotionDto) SetReduceImgUrl(v string) *AlibabaAlscUnionElemePromotionOfficialactivityGetTopWxPromotionDto {
    s.ReduceImgUrl = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionOfficialactivityGetTopWxPromotionDto) SetH5ShortLink(v string) *AlibabaAlscUnionElemePromotionOfficialactivityGetTopWxPromotionDto {
    s.H5ShortLink = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionOfficialactivityGetTopWxPromotionDto) SetMarketWxAppId(v string) *AlibabaAlscUnionElemePromotionOfficialactivityGetTopWxPromotionDto {
    s.MarketWxAppId = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionOfficialactivityGetTopWxPromotionDto) SetMarketWxAppPath(v string) *AlibabaAlscUnionElemePromotionOfficialactivityGetTopWxPromotionDto {
    s.MarketWxAppPath = &v
    return s
}
