package domain


type AlibabaAlscUnionElemePromotionOtherchannelGetOtherPromotionLinkRequest struct {
    /*
        链接类型 1 抖音schema     */
    Type  *int64 `json:"type,omitempty" `

    /*
        用户扩展id     */
    Sid  *string `json:"sid,omitempty" `

}

func (s *AlibabaAlscUnionElemePromotionOtherchannelGetOtherPromotionLinkRequest) SetType(v int64) *AlibabaAlscUnionElemePromotionOtherchannelGetOtherPromotionLinkRequest {
    s.Type = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionOtherchannelGetOtherPromotionLinkRequest) SetSid(v string) *AlibabaAlscUnionElemePromotionOtherchannelGetOtherPromotionLinkRequest {
    s.Sid = &v
    return s
}
