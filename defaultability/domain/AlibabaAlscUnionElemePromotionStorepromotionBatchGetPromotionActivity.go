package domain


type AlibabaAlscUnionElemePromotionStorepromotionBatchGetPromotionActivity struct {
    /*
        活动Id     */
    ActivityId  *string `json:"activity_id,omitempty" `

    /*
        营销计划服务费（分）     */
    ServiceFeeCent  *int64 `json:"service_fee_cent,omitempty" `

    /*
        奖励金红包面额（分）     */
    BonusCent  *int64 `json:"bonus_cent,omitempty" `

    /*
        活动的日库存     */
    DailyQuantity  *int64 `json:"daily_quantity,omitempty" `

    /*
        活动日剩余库存     */
    DailySellableQuantity  *int64 `json:"daily_sellable_quantity,omitempty" `

    /*
        起始时间（秒）     */
    StartTime  *int64 `json:"start_time,omitempty" `

    /*
        结束时间（秒）     */
    EndTime  *int64 `json:"end_time,omitempty" `

    /*
        奖励金门槛 (分)     */
    BountyMinLimitCent  *int64 `json:"bounty_min_limit_cent,omitempty" `

}

func (s *AlibabaAlscUnionElemePromotionStorepromotionBatchGetPromotionActivity) SetActivityId(v string) *AlibabaAlscUnionElemePromotionStorepromotionBatchGetPromotionActivity {
    s.ActivityId = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionStorepromotionBatchGetPromotionActivity) SetServiceFeeCent(v int64) *AlibabaAlscUnionElemePromotionStorepromotionBatchGetPromotionActivity {
    s.ServiceFeeCent = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionStorepromotionBatchGetPromotionActivity) SetBonusCent(v int64) *AlibabaAlscUnionElemePromotionStorepromotionBatchGetPromotionActivity {
    s.BonusCent = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionStorepromotionBatchGetPromotionActivity) SetDailyQuantity(v int64) *AlibabaAlscUnionElemePromotionStorepromotionBatchGetPromotionActivity {
    s.DailyQuantity = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionStorepromotionBatchGetPromotionActivity) SetDailySellableQuantity(v int64) *AlibabaAlscUnionElemePromotionStorepromotionBatchGetPromotionActivity {
    s.DailySellableQuantity = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionStorepromotionBatchGetPromotionActivity) SetStartTime(v int64) *AlibabaAlscUnionElemePromotionStorepromotionBatchGetPromotionActivity {
    s.StartTime = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionStorepromotionBatchGetPromotionActivity) SetEndTime(v int64) *AlibabaAlscUnionElemePromotionStorepromotionBatchGetPromotionActivity {
    s.EndTime = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionStorepromotionBatchGetPromotionActivity) SetBountyMinLimitCent(v int64) *AlibabaAlscUnionElemePromotionStorepromotionBatchGetPromotionActivity {
    s.BountyMinLimitCent = &v
    return s
}
