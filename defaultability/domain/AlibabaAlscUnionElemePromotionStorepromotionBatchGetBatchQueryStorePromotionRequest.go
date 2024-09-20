package domain


type AlibabaAlscUnionElemePromotionStorepromotionBatchGetBatchQueryStorePromotionRequest struct {
    /*
        渠道PID     */
    Pid  *string `json:"pid,omitempty" `

    /*
        门店ID，支持多值，'|'分隔     */
    ShopId  *string `json:"shop_id,omitempty" `

    /*
        活动ID     */
    ActivityId  *string `json:"activity_id,omitempty" `

    /*
        媒体出资活动ID     */
    MediaActivityId  *string `json:"media_activity_id,omitempty" `

    /*
        三方扩展id     */
    Sid  *string `json:"sid,omitempty" `

}

func (s *AlibabaAlscUnionElemePromotionStorepromotionBatchGetBatchQueryStorePromotionRequest) SetPid(v string) *AlibabaAlscUnionElemePromotionStorepromotionBatchGetBatchQueryStorePromotionRequest {
    s.Pid = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionStorepromotionBatchGetBatchQueryStorePromotionRequest) SetShopId(v string) *AlibabaAlscUnionElemePromotionStorepromotionBatchGetBatchQueryStorePromotionRequest {
    s.ShopId = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionStorepromotionBatchGetBatchQueryStorePromotionRequest) SetActivityId(v string) *AlibabaAlscUnionElemePromotionStorepromotionBatchGetBatchQueryStorePromotionRequest {
    s.ActivityId = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionStorepromotionBatchGetBatchQueryStorePromotionRequest) SetMediaActivityId(v string) *AlibabaAlscUnionElemePromotionStorepromotionBatchGetBatchQueryStorePromotionRequest {
    s.MediaActivityId = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionStorepromotionBatchGetBatchQueryStorePromotionRequest) SetSid(v string) *AlibabaAlscUnionElemePromotionStorepromotionBatchGetBatchQueryStorePromotionRequest {
    s.Sid = &v
    return s
}
