package request


type AlibabaAlscUnionElemeStorepromotionReviewbwcDetailGetRequest struct {
    /*
        渠道PID     */
    Pid  *string `json:"pid" required:"true" `
    /*
        三方扩展id     */
    Sid  *string `json:"sid,omitempty" required:"false" `
    /*
        门店ID（加密）     */
    ShopId  *string `json:"shop_id" required:"true" `
    /*
        活动ID     */
    ActivityId  *string `json:"activity_id" required:"true" `
}

func (s *AlibabaAlscUnionElemeStorepromotionReviewbwcDetailGetRequest) SetPid(v string) *AlibabaAlscUnionElemeStorepromotionReviewbwcDetailGetRequest {
    s.Pid = &v
    return s
}
func (s *AlibabaAlscUnionElemeStorepromotionReviewbwcDetailGetRequest) SetSid(v string) *AlibabaAlscUnionElemeStorepromotionReviewbwcDetailGetRequest {
    s.Sid = &v
    return s
}
func (s *AlibabaAlscUnionElemeStorepromotionReviewbwcDetailGetRequest) SetShopId(v string) *AlibabaAlscUnionElemeStorepromotionReviewbwcDetailGetRequest {
    s.ShopId = &v
    return s
}
func (s *AlibabaAlscUnionElemeStorepromotionReviewbwcDetailGetRequest) SetActivityId(v string) *AlibabaAlscUnionElemeStorepromotionReviewbwcDetailGetRequest {
    s.ActivityId = &v
    return s
}

func (req *AlibabaAlscUnionElemeStorepromotionReviewbwcDetailGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.Pid != nil) {
        paramMap["pid"] = *req.Pid
    }
    if(req.Sid != nil) {
        paramMap["sid"] = *req.Sid
    }
    if(req.ShopId != nil) {
        paramMap["shop_id"] = *req.ShopId
    }
    if(req.ActivityId != nil) {
        paramMap["activity_id"] = *req.ActivityId
    }
    return paramMap
}

func (req *AlibabaAlscUnionElemeStorepromotionReviewbwcDetailGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}