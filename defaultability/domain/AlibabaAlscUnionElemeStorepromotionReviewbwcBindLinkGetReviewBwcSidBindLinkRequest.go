package domain


type AlibabaAlscUnionElemeStorepromotionReviewbwcBindLinkGetReviewBwcSidBindLinkRequest struct {
    /*
        推广位     */
    Pid  *string `json:"pid,omitempty" `

    /*
        会员ID     */
    Sid  *string `json:"sid,omitempty" `

    /*
        活动ID     */
    ActivityId  *string `json:"activity_id,omitempty" `

}

func (s *AlibabaAlscUnionElemeStorepromotionReviewbwcBindLinkGetReviewBwcSidBindLinkRequest) SetPid(v string) *AlibabaAlscUnionElemeStorepromotionReviewbwcBindLinkGetReviewBwcSidBindLinkRequest {
    s.Pid = &v
    return s
}
func (s *AlibabaAlscUnionElemeStorepromotionReviewbwcBindLinkGetReviewBwcSidBindLinkRequest) SetSid(v string) *AlibabaAlscUnionElemeStorepromotionReviewbwcBindLinkGetReviewBwcSidBindLinkRequest {
    s.Sid = &v
    return s
}
func (s *AlibabaAlscUnionElemeStorepromotionReviewbwcBindLinkGetReviewBwcSidBindLinkRequest) SetActivityId(v string) *AlibabaAlscUnionElemeStorepromotionReviewbwcBindLinkGetReviewBwcSidBindLinkRequest {
    s.ActivityId = &v
    return s
}
