package domain


type AlibabaAlscUnionElemePromotionStorepromotionBatchGetPromotionLink struct {
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

}

func (s *AlibabaAlscUnionElemePromotionStorepromotionBatchGetPromotionLink) SetWxAppid(v string) *AlibabaAlscUnionElemePromotionStorepromotionBatchGetPromotionLink {
    s.WxAppid = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionStorepromotionBatchGetPromotionLink) SetWxPath(v string) *AlibabaAlscUnionElemePromotionStorepromotionBatchGetPromotionLink {
    s.WxPath = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionStorepromotionBatchGetPromotionLink) SetPicture(v string) *AlibabaAlscUnionElemePromotionStorepromotionBatchGetPromotionLink {
    s.Picture = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionStorepromotionBatchGetPromotionLink) SetReduceWxAppid(v string) *AlibabaAlscUnionElemePromotionStorepromotionBatchGetPromotionLink {
    s.ReduceWxAppid = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionStorepromotionBatchGetPromotionLink) SetReduceWxPath(v string) *AlibabaAlscUnionElemePromotionStorepromotionBatchGetPromotionLink {
    s.ReduceWxPath = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionStorepromotionBatchGetPromotionLink) SetReducePicture(v string) *AlibabaAlscUnionElemePromotionStorepromotionBatchGetPromotionLink {
    s.ReducePicture = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionStorepromotionBatchGetPromotionLink) SetMiniQrcode(v string) *AlibabaAlscUnionElemePromotionStorepromotionBatchGetPromotionLink {
    s.MiniQrcode = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionStorepromotionBatchGetPromotionLink) SetMediaActivityWxAppid(v string) *AlibabaAlscUnionElemePromotionStorepromotionBatchGetPromotionLink {
    s.MediaActivityWxAppid = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionStorepromotionBatchGetPromotionLink) SetMediaActivityWxPath(v string) *AlibabaAlscUnionElemePromotionStorepromotionBatchGetPromotionLink {
    s.MediaActivityWxPath = &v
    return s
}
