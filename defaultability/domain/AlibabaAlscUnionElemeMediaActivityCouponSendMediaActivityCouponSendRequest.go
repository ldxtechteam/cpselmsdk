package domain


type AlibabaAlscUnionElemeMediaActivityCouponSendMediaActivityCouponSendRequest struct {
    /*
        领券手机号     */
    Mobile  *string `json:"mobile,omitempty" `

    /*
        媒体出资活动ID     */
    MediaActivityId  *string `json:"media_activity_id,omitempty" `

}

func (s *AlibabaAlscUnionElemeMediaActivityCouponSendMediaActivityCouponSendRequest) SetMobile(v string) *AlibabaAlscUnionElemeMediaActivityCouponSendMediaActivityCouponSendRequest {
    s.Mobile = &v
    return s
}
func (s *AlibabaAlscUnionElemeMediaActivityCouponSendMediaActivityCouponSendRequest) SetMediaActivityId(v string) *AlibabaAlscUnionElemeMediaActivityCouponSendMediaActivityCouponSendRequest {
    s.MediaActivityId = &v
    return s
}
