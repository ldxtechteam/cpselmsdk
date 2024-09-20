package domain


type AlibabaAlscUnionElemePromotionOfficialactivityGetActivityRequest struct {
    /*
        渠道PID     */
    Pid  *string `json:"pid,omitempty" `

    /*
        活动ID     */
    ActivityId  *string `json:"activity_id,omitempty" `

    /*
        三方会员id。长度限制50     */
    Sid  *string `json:"sid,omitempty" `

    /*
        是否返回微信推广图片     */
    IncludeWxImg  *bool `json:"include_wx_img,omitempty" `

    /*
        是否包含二维码，如果为false，不返回二维码和图片，只有链接     */
    IncludeQrCode  *bool `json:"include_qr_code,omitempty" `

}

func (s *AlibabaAlscUnionElemePromotionOfficialactivityGetActivityRequest) SetPid(v string) *AlibabaAlscUnionElemePromotionOfficialactivityGetActivityRequest {
    s.Pid = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionOfficialactivityGetActivityRequest) SetActivityId(v string) *AlibabaAlscUnionElemePromotionOfficialactivityGetActivityRequest {
    s.ActivityId = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionOfficialactivityGetActivityRequest) SetSid(v string) *AlibabaAlscUnionElemePromotionOfficialactivityGetActivityRequest {
    s.Sid = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionOfficialactivityGetActivityRequest) SetIncludeWxImg(v bool) *AlibabaAlscUnionElemePromotionOfficialactivityGetActivityRequest {
    s.IncludeWxImg = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionOfficialactivityGetActivityRequest) SetIncludeQrCode(v bool) *AlibabaAlscUnionElemePromotionOfficialactivityGetActivityRequest {
    s.IncludeQrCode = &v
    return s
}
