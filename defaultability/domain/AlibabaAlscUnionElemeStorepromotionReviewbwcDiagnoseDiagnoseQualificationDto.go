package domain


type AlibabaAlscUnionElemeStorepromotionReviewbwcDiagnoseDiagnoseQualificationDto struct {
    /*
        饿了么门店ID     */
    ShopId  *string `json:"shop_id,omitempty" `

    /*
        资格领取手机号     */
    ApplyMobile  *string `json:"apply_mobile,omitempty" `

    /*
        领取时间     */
    ApplyTimeText  *string `json:"apply_time_text,omitempty" `

    /*
        归因订单     */
    AscribeOrderId  *string `json:"ascribe_order_id,omitempty" `

}

func (s *AlibabaAlscUnionElemeStorepromotionReviewbwcDiagnoseDiagnoseQualificationDto) SetShopId(v string) *AlibabaAlscUnionElemeStorepromotionReviewbwcDiagnoseDiagnoseQualificationDto {
    s.ShopId = &v
    return s
}
func (s *AlibabaAlscUnionElemeStorepromotionReviewbwcDiagnoseDiagnoseQualificationDto) SetApplyMobile(v string) *AlibabaAlscUnionElemeStorepromotionReviewbwcDiagnoseDiagnoseQualificationDto {
    s.ApplyMobile = &v
    return s
}
func (s *AlibabaAlscUnionElemeStorepromotionReviewbwcDiagnoseDiagnoseQualificationDto) SetApplyTimeText(v string) *AlibabaAlscUnionElemeStorepromotionReviewbwcDiagnoseDiagnoseQualificationDto {
    s.ApplyTimeText = &v
    return s
}
func (s *AlibabaAlscUnionElemeStorepromotionReviewbwcDiagnoseDiagnoseQualificationDto) SetAscribeOrderId(v string) *AlibabaAlscUnionElemeStorepromotionReviewbwcDiagnoseDiagnoseQualificationDto {
    s.AscribeOrderId = &v
    return s
}
