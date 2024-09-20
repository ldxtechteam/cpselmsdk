package domain


type AlibabaAlscUnionPromotionLinkAnalyzePromotionLinkAnalyzeRequest struct {
    /*
        链接类型（1-h5；2-h5短链；3-微信小程序；4-饿了么APP；5-淘口令）     */
    Type  *int64 `json:"type,omitempty" `

    /*
        推广链接或者口令     */
    Link  *string `json:"link,omitempty" `

}

func (s *AlibabaAlscUnionPromotionLinkAnalyzePromotionLinkAnalyzeRequest) SetType(v int64) *AlibabaAlscUnionPromotionLinkAnalyzePromotionLinkAnalyzeRequest {
    s.Type = &v
    return s
}
func (s *AlibabaAlscUnionPromotionLinkAnalyzePromotionLinkAnalyzeRequest) SetLink(v string) *AlibabaAlscUnionPromotionLinkAnalyzePromotionLinkAnalyzeRequest {
    s.Link = &v
    return s
}
