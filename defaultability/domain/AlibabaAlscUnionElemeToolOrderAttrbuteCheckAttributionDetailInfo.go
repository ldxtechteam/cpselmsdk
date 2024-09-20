package domain


type AlibabaAlscUnionElemeToolOrderAttrbuteCheckAttributionDetailInfo struct {
    /*
        归因类型     */
    BizType  *string `json:"biz_type,omitempty" `

    /*
        归因结果，1.归因成功 2.归因失败     */
    Result  *int64 `json:"result,omitempty" `

    /*
        归因失败原因     */
    Reason  *string `json:"reason,omitempty" `

    /*
        NO_NEED_SETTLE_NO_COMMENT - 无需结算:未评价     NO_NEED_SETTLE_LOW_RATING - 无需结算:未符合评价结算标准     NO_NEED_SETTLE_BELOW_ORDER_AMOUNT - 无需结算:不满足订单门槛     NO_NEED_SETTLE_NO_VALID_COMMENT_TIME - 无需结算:评价时间超过次日24小时     UN_ATTRIBUTE_NO_USE_COUPON - 订单未使用推荐有奖券     UN_ATTRIBUTE_MISS_SHORT_CODE - 推荐有奖券未携带有效短码信息     UN_ATTRIBUTE_NO_QUALIFICATION - 用户无有效评价资格     COMMON_UN_SElF_ORDER - 非自己渠道推广订单，已被其他渠道优先归因     COMMON_UN_CPS_ATTRIBUTE - 未匹配到联盟所有归因逻辑（点击、券、活动等）UN_ATTRIBUTE_NO_USE_COUPON_NEW - 订单未使用推荐有奖券或用户不满足端新归因资格     */
    ReasonCode  *string `json:"reason_code,omitempty" `

}

func (s *AlibabaAlscUnionElemeToolOrderAttrbuteCheckAttributionDetailInfo) SetBizType(v string) *AlibabaAlscUnionElemeToolOrderAttrbuteCheckAttributionDetailInfo {
    s.BizType = &v
    return s
}
func (s *AlibabaAlscUnionElemeToolOrderAttrbuteCheckAttributionDetailInfo) SetResult(v int64) *AlibabaAlscUnionElemeToolOrderAttrbuteCheckAttributionDetailInfo {
    s.Result = &v
    return s
}
func (s *AlibabaAlscUnionElemeToolOrderAttrbuteCheckAttributionDetailInfo) SetReason(v string) *AlibabaAlscUnionElemeToolOrderAttrbuteCheckAttributionDetailInfo {
    s.Reason = &v
    return s
}
func (s *AlibabaAlscUnionElemeToolOrderAttrbuteCheckAttributionDetailInfo) SetReasonCode(v string) *AlibabaAlscUnionElemeToolOrderAttrbuteCheckAttributionDetailInfo {
    s.ReasonCode = &v
    return s
}
