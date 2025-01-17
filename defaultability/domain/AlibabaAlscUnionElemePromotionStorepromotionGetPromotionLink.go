package domain


type AlibabaAlscUnionElemePromotionStorepromotionGetPromotionLink struct {
    /*
        小程序appId     */
    WxAppid  *string `json:"wx_appid,omitempty" `

    /*
        微信小程序path链接     */
    WxPath  *string `json:"wx_path,omitempty" `

    /*
        推广图片地址，图片上展示店铺小程序二维码     */
    Picture  *string `json:"picture,omitempty" `

    /*
        小程序appId-立减活动     */
    ReduceWxAppid  *string `json:"reduce_wx_appid,omitempty" `

    /*
        微信小程序path链接-立减活动     */
    ReduceWxPath  *string `json:"reduce_wx_path,omitempty" `

    /*
        推广图片地址-立减活动，图片上展示店铺小程序二维码     */
    ReducePicture  *string `json:"reduce_picture,omitempty" `

    /*
        独立微信二维码     */
    MiniQrcode  *string `json:"mini_qrcode,omitempty" `

    /*
        小程序appId-媒体出资活动     */
    MediaActivityWxAppid  *string `json:"media_activity_wx_appid,omitempty" `

    /*
        微信小程序path链接-媒体出资活动     */
    MediaActivityWxPath  *string `json:"media_activity_wx_path,omitempty" `

    /*
        淘宝饿了么小程序推广页     */
    TaobaoSchemeUrl  *string `json:"taobao_scheme_url,omitempty" `

    /*
        淘宝二维码     */
    TaobaoQrcode  *string `json:"taobao_qrcode,omitempty" `

}

func (s *AlibabaAlscUnionElemePromotionStorepromotionGetPromotionLink) SetWxAppid(v string) *AlibabaAlscUnionElemePromotionStorepromotionGetPromotionLink {
    s.WxAppid = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionStorepromotionGetPromotionLink) SetWxPath(v string) *AlibabaAlscUnionElemePromotionStorepromotionGetPromotionLink {
    s.WxPath = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionStorepromotionGetPromotionLink) SetPicture(v string) *AlibabaAlscUnionElemePromotionStorepromotionGetPromotionLink {
    s.Picture = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionStorepromotionGetPromotionLink) SetReduceWxAppid(v string) *AlibabaAlscUnionElemePromotionStorepromotionGetPromotionLink {
    s.ReduceWxAppid = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionStorepromotionGetPromotionLink) SetReduceWxPath(v string) *AlibabaAlscUnionElemePromotionStorepromotionGetPromotionLink {
    s.ReduceWxPath = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionStorepromotionGetPromotionLink) SetReducePicture(v string) *AlibabaAlscUnionElemePromotionStorepromotionGetPromotionLink {
    s.ReducePicture = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionStorepromotionGetPromotionLink) SetMiniQrcode(v string) *AlibabaAlscUnionElemePromotionStorepromotionGetPromotionLink {
    s.MiniQrcode = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionStorepromotionGetPromotionLink) SetMediaActivityWxAppid(v string) *AlibabaAlscUnionElemePromotionStorepromotionGetPromotionLink {
    s.MediaActivityWxAppid = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionStorepromotionGetPromotionLink) SetMediaActivityWxPath(v string) *AlibabaAlscUnionElemePromotionStorepromotionGetPromotionLink {
    s.MediaActivityWxPath = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionStorepromotionGetPromotionLink) SetTaobaoSchemeUrl(v string) *AlibabaAlscUnionElemePromotionStorepromotionGetPromotionLink {
    s.TaobaoSchemeUrl = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionStorepromotionGetPromotionLink) SetTaobaoQrcode(v string) *AlibabaAlscUnionElemePromotionStorepromotionGetPromotionLink {
    s.TaobaoQrcode = &v
    return s
}
