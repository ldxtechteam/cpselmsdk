package domain


type AlibabaAlscUnionElemeStorepromotionReviewbwcQueryReviewBwcActivity struct {
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

func (s *AlibabaAlscUnionElemeStorepromotionReviewbwcQueryReviewBwcActivity) SetActivityId(v string) *AlibabaAlscUnionElemeStorepromotionReviewbwcQueryReviewBwcActivity {
    s.ActivityId = &v
    return s
}
func (s *AlibabaAlscUnionElemeStorepromotionReviewbwcQueryReviewBwcActivity) SetStartTime(v int64) *AlibabaAlscUnionElemeStorepromotionReviewbwcQueryReviewBwcActivity {
    s.StartTime = &v
    return s
}
func (s *AlibabaAlscUnionElemeStorepromotionReviewbwcQueryReviewBwcActivity) SetEndTime(v int64) *AlibabaAlscUnionElemeStorepromotionReviewbwcQueryReviewBwcActivity {
    s.EndTime = &v
    return s
}
func (s *AlibabaAlscUnionElemeStorepromotionReviewbwcQueryReviewBwcActivity) SetDailyStock(v int64) *AlibabaAlscUnionElemeStorepromotionReviewbwcQueryReviewBwcActivity {
    s.DailyStock = &v
    return s
}
func (s *AlibabaAlscUnionElemeStorepromotionReviewbwcQueryReviewBwcActivity) SetDailyRemainStock(v int64) *AlibabaAlscUnionElemeStorepromotionReviewbwcQueryReviewBwcActivity {
    s.DailyRemainStock = &v
    return s
}
func (s *AlibabaAlscUnionElemeStorepromotionReviewbwcQueryReviewBwcActivity) SetRewardCent(v string) *AlibabaAlscUnionElemeStorepromotionReviewbwcQueryReviewBwcActivity {
    s.RewardCent = &v
    return s
}
func (s *AlibabaAlscUnionElemeStorepromotionReviewbwcQueryReviewBwcActivity) SetRewardThresholdCent(v string) *AlibabaAlscUnionElemeStorepromotionReviewbwcQueryReviewBwcActivity {
    s.RewardThresholdCent = &v
    return s
}
func (s *AlibabaAlscUnionElemeStorepromotionReviewbwcQueryReviewBwcActivity) SetCommissionRate(v string) *AlibabaAlscUnionElemeStorepromotionReviewbwcQueryReviewBwcActivity {
    s.CommissionRate = &v
    return s
}
func (s *AlibabaAlscUnionElemeStorepromotionReviewbwcQueryReviewBwcActivity) SetCommission(v string) *AlibabaAlscUnionElemeStorepromotionReviewbwcQueryReviewBwcActivity {
    s.Commission = &v
    return s
}
