package domain


type AlibabaAlscUnionElemePromotionOfficialactivityGetTopH5PromotionDto struct {
    /*
        长连接     */
    H5Url  *string `json:"h5_url,omitempty" `

    /*
        短连接     */
    ShortLink  *string `json:"short_link,omitempty" `

    /*
        推广二维码     */
    H5QrCode  *string `json:"h5_qr_code,omitempty" `

    /*
        推广海报     */
    H5ImgUrl  *string `json:"h5_img_url,omitempty" `

    /*
        推荐有奖拉新链接     */
    TjH5Url  *string `json:"tj_h5_url,omitempty" `

}

func (s *AlibabaAlscUnionElemePromotionOfficialactivityGetTopH5PromotionDto) SetH5Url(v string) *AlibabaAlscUnionElemePromotionOfficialactivityGetTopH5PromotionDto {
    s.H5Url = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionOfficialactivityGetTopH5PromotionDto) SetShortLink(v string) *AlibabaAlscUnionElemePromotionOfficialactivityGetTopH5PromotionDto {
    s.ShortLink = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionOfficialactivityGetTopH5PromotionDto) SetH5QrCode(v string) *AlibabaAlscUnionElemePromotionOfficialactivityGetTopH5PromotionDto {
    s.H5QrCode = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionOfficialactivityGetTopH5PromotionDto) SetH5ImgUrl(v string) *AlibabaAlscUnionElemePromotionOfficialactivityGetTopH5PromotionDto {
    s.H5ImgUrl = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionOfficialactivityGetTopH5PromotionDto) SetTjH5Url(v string) *AlibabaAlscUnionElemePromotionOfficialactivityGetTopH5PromotionDto {
    s.TjH5Url = &v
    return s
}
