package domain


type AlibabaAlscUnionElemeStorepromotionReviewbwcQueryPageModel struct {
    /*
        会话ID（下次请求作为请求参数，用于标记分页会话自动翻页）     */
    SessionId  *string `json:"session_id,omitempty" `

    /*
        分页记录（记录空时表示当前sessionId对应分页记录已全部返回，分页结束）     */
    Records  *[]AlibabaAlscUnionElemeStorepromotionReviewbwcQueryReviewBwcStorePromotionDto `json:"records,omitempty" `

}

func (s *AlibabaAlscUnionElemeStorepromotionReviewbwcQueryPageModel) SetSessionId(v string) *AlibabaAlscUnionElemeStorepromotionReviewbwcQueryPageModel {
    s.SessionId = &v
    return s
}
func (s *AlibabaAlscUnionElemeStorepromotionReviewbwcQueryPageModel) SetRecords(v []AlibabaAlscUnionElemeStorepromotionReviewbwcQueryReviewBwcStorePromotionDto) *AlibabaAlscUnionElemeStorepromotionReviewbwcQueryPageModel {
    s.Records = &v
    return s
}
