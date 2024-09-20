package domain


type AlibabaAlscUnionElemeStorepromotionReviewbwcDetailGetReviewBwcActivity struct {
    /*
        活动ID     */
    ActivityId  *string `json:"activity_id,omitempty" `

    /*
        开始时间（秒）     */
    StartTime  *int64 `json:"start_time,omitempty" `

    /*
        结束时间（秒）     */
    EndTime  *int64 `json:"end_time,omitempty" `

    /*
        活动的日库存     */
    DailyStock  *int64 `json:"daily_stock,omitempty" `

    /*
        活动日剩余库存     */
    DailyRemainStock  *int64 `json:"daily_remain_stock,omitempty" `

    /*
        返现金额（分）     */
    RewardCent  *string `json:"reward_cent,omitempty" `

    /*
        返现门槛（分）     */
    RewardThresholdCent  *string `json:"reward_threshold_cent,omitempty" `

    /*
        佣金比例     */
    CommissionRate  *string `json:"commission_rate,omitempty" `

    /*
        预估佣金（分）     */
    Commission  *string `json:"commission,omitempty" `

}

func (s *AlibabaAlscUnionElemeStorepromotionReviewbwcDetailGetReviewBwcActivity) SetActivityId(v string) *AlibabaAlscUnionElemeStorepromotionReviewbwcDetailGetReviewBwcActivity {
    s.ActivityId = &v
    return s
}
func (s *AlibabaAlscUnionElemeStorepromotionReviewbwcDetailGetReviewBwcActivity) SetStartTime(v int64) *AlibabaAlscUnionElemeStorepromotionReviewbwcDetailGetReviewBwcActivity {
    s.StartTime = &v
    return s
}
func (s *AlibabaAlscUnionElemeStorepromotionReviewbwcDetailGetReviewBwcActivity) SetEndTime(v int64) *AlibabaAlscUnionElemeStorepromotionReviewbwcDetailGetReviewBwcActivity {
    s.EndTime = &v
    return s
}
func (s *AlibabaAlscUnionElemeStorepromotionReviewbwcDetailGetReviewBwcActivity) SetDailyStock(v int64) *AlibabaAlscUnionElemeStorepromotionReviewbwcDetailGetReviewBwcActivity {
    s.DailyStock = &v
    return s
}
func (s *AlibabaAlscUnionElemeStorepromotionReviewbwcDetailGetReviewBwcActivity) SetDailyRemainStock(v int64) *AlibabaAlscUnionElemeStorepromotionReviewbwcDetailGetReviewBwcActivity {
    s.DailyRemainStock = &v
    return s
}
func (s *AlibabaAlscUnionElemeStorepromotionReviewbwcDetailGetReviewBwcActivity) SetRewardCent(v string) *AlibabaAlscUnionElemeStorepromotionReviewbwcDetailGetReviewBwcActivity {
    s.RewardCent = &v
    return s
}
func (s *AlibabaAlscUnionElemeStorepromotionReviewbwcDetailGetReviewBwcActivity) SetRewardThresholdCent(v string) *AlibabaAlscUnionElemeStorepromotionReviewbwcDetailGetReviewBwcActivity {
    s.RewardThresholdCent = &v
    return s
}
func (s *AlibabaAlscUnionElemeStorepromotionReviewbwcDetailGetReviewBwcActivity) SetCommissionRate(v string) *AlibabaAlscUnionElemeStorepromotionReviewbwcDetailGetReviewBwcActivity {
    s.CommissionRate = &v
    return s
}
func (s *AlibabaAlscUnionElemeStorepromotionReviewbwcDetailGetReviewBwcActivity) SetCommission(v string) *AlibabaAlscUnionElemeStorepromotionReviewbwcDetailGetReviewBwcActivity {
    s.Commission = &v
    return s
}
