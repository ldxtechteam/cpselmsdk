package domain


type AlibabaAlscUnionElemeStorepromotionReviewbwcDiagnoseOAscribeDiagnoseResult struct {
    /*
        资格信息     */
    Qualification  *AlibabaAlscUnionElemeStorepromotionReviewbwcDiagnoseDiagnoseQualificationDto `json:"qualification,omitempty" `

    /*
        订单信息     */
    Order  *AlibabaAlscUnionElemeStorepromotionReviewbwcDiagnoseODiagnoseOrderResult `json:"order,omitempty" `

    /*
        诊断状态（1-未归因；2-已归因，未评价；3-已评价，待返现；4-已返现；5-不返现）     */
    DiagnoseState  *int64 `json:"diagnose_state,omitempty" `

    /*
        诊断code（NOT_ASCRIBE_ORDER_NOT_PAY-订单未支付；NOT_ASCRIBE_MOBILE_NOT_MATCH-资格领取手机号与订单手机号不一致；NOT_ASCRIBE_SHOP_NOT_MATCH-资格领取门店与订单门店不一致；NOT_ASCRIBE_ORDER_CANCEL-订单已取消；NOT_ASCRIBE_QUALIFICATION_ALREADY_USED-资格已被使用；NOT_ASCRIBE_ORDER_USE_BWC-订单使用叠红包；NOT_ASCRIBE_ORDER_AFTER_ASCRIBE_WINDOW-订单时间晚于资格有效期；NOT_SETTLE_NO_REVIEW：不满足评价有礼结算规则-资格领取次日24点前该订单无评价；NOT_SETTLE_ORDER_PAY_AMOUNT_NOT_MATCH-不满足用户实付门槛，确认收货金额不满足实付门槛；NOT_SETTLE_ORDER_CANCEL_PAY_AMOUNT_NOT_MATCH-不满足用户实付门槛，订单售后退，不满足实付门槛）     */
    DiagnoseCode  *string `json:"diagnose_code,omitempty" `

    /*
        诊断描述（文案会变，仅作为参考，接入方可以根据code自己定义文案）     */
    DiagnoseDesc  *string `json:"diagnose_desc,omitempty" `

}

func (s *AlibabaAlscUnionElemeStorepromotionReviewbwcDiagnoseOAscribeDiagnoseResult) SetQualification(v AlibabaAlscUnionElemeStorepromotionReviewbwcDiagnoseDiagnoseQualificationDto) *AlibabaAlscUnionElemeStorepromotionReviewbwcDiagnoseOAscribeDiagnoseResult {
    s.Qualification = &v
    return s
}
func (s *AlibabaAlscUnionElemeStorepromotionReviewbwcDiagnoseOAscribeDiagnoseResult) SetOrder(v AlibabaAlscUnionElemeStorepromotionReviewbwcDiagnoseODiagnoseOrderResult) *AlibabaAlscUnionElemeStorepromotionReviewbwcDiagnoseOAscribeDiagnoseResult {
    s.Order = &v
    return s
}
func (s *AlibabaAlscUnionElemeStorepromotionReviewbwcDiagnoseOAscribeDiagnoseResult) SetDiagnoseState(v int64) *AlibabaAlscUnionElemeStorepromotionReviewbwcDiagnoseOAscribeDiagnoseResult {
    s.DiagnoseState = &v
    return s
}
func (s *AlibabaAlscUnionElemeStorepromotionReviewbwcDiagnoseOAscribeDiagnoseResult) SetDiagnoseCode(v string) *AlibabaAlscUnionElemeStorepromotionReviewbwcDiagnoseOAscribeDiagnoseResult {
    s.DiagnoseCode = &v
    return s
}
func (s *AlibabaAlscUnionElemeStorepromotionReviewbwcDiagnoseOAscribeDiagnoseResult) SetDiagnoseDesc(v string) *AlibabaAlscUnionElemeStorepromotionReviewbwcDiagnoseOAscribeDiagnoseResult {
    s.DiagnoseDesc = &v
    return s
}
