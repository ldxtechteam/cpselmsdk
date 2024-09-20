package domain


type AlibabaAlscUnionElemeStorepromotionReviewbwcBindLinkGetReviewBwcSidBindLinkResult struct {
    /*
        微信小程序appId     */
    WxAppid  *string `json:"wx_appid,omitempty" `

    /*
        微信小程序path链接     */
    WxPath  *string `json:"wx_path,omitempty" `

}

func (s *AlibabaAlscUnionElemeStorepromotionReviewbwcBindLinkGetReviewBwcSidBindLinkResult) SetWxAppid(v string) *AlibabaAlscUnionElemeStorepromotionReviewbwcBindLinkGetReviewBwcSidBindLinkResult {
    s.WxAppid = &v
    return s
}
func (s *AlibabaAlscUnionElemeStorepromotionReviewbwcBindLinkGetReviewBwcSidBindLinkResult) SetWxPath(v string) *AlibabaAlscUnionElemeStorepromotionReviewbwcBindLinkGetReviewBwcSidBindLinkResult {
    s.WxPath = &v
    return s
}
