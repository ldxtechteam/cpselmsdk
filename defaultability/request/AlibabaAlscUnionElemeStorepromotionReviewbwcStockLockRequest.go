package request


type AlibabaAlscUnionElemeStorepromotionReviewbwcStockLockRequest struct {
    /*
        渠道PID     */
    Pid  *string `json:"pid" required:"true" `
    /*
        门店ID（加密）     */
    ShopId  *string `json:"shop_id" required:"true" `
    /*
        活动ID     */
    ActivityId  *string `json:"activity_id" required:"true" `
    /*
        三方扩展id     */
    Sid  *string `json:"sid,omitempty" required:"false" `
    /*
        领取手机号     */
    Mobile  *string `json:"mobile" required:"true" `
    /*
        领取ID（渠道用户领取资格的唯一标识）     */
    OuterOrderId  *string `json:"outer_order_id" required:"true" `
}

func (s *AlibabaAlscUnionElemeStorepromotionReviewbwcStockLockRequest) SetPid(v string) *AlibabaAlscUnionElemeStorepromotionReviewbwcStockLockRequest {
    s.Pid = &v
    return s
}
func (s *AlibabaAlscUnionElemeStorepromotionReviewbwcStockLockRequest) SetShopId(v string) *AlibabaAlscUnionElemeStorepromotionReviewbwcStockLockRequest {
    s.ShopId = &v
    return s
}
func (s *AlibabaAlscUnionElemeStorepromotionReviewbwcStockLockRequest) SetActivityId(v string) *AlibabaAlscUnionElemeStorepromotionReviewbwcStockLockRequest {
    s.ActivityId = &v
    return s
}
func (s *AlibabaAlscUnionElemeStorepromotionReviewbwcStockLockRequest) SetSid(v string) *AlibabaAlscUnionElemeStorepromotionReviewbwcStockLockRequest {
    s.Sid = &v
    return s
}
func (s *AlibabaAlscUnionElemeStorepromotionReviewbwcStockLockRequest) SetMobile(v string) *AlibabaAlscUnionElemeStorepromotionReviewbwcStockLockRequest {
    s.Mobile = &v
    return s
}
func (s *AlibabaAlscUnionElemeStorepromotionReviewbwcStockLockRequest) SetOuterOrderId(v string) *AlibabaAlscUnionElemeStorepromotionReviewbwcStockLockRequest {
    s.OuterOrderId = &v
    return s
}

func (req *AlibabaAlscUnionElemeStorepromotionReviewbwcStockLockRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.Pid != nil) {
        paramMap["pid"] = *req.Pid
    }
    if(req.ShopId != nil) {
        paramMap["shop_id"] = *req.ShopId
    }
    if(req.ActivityId != nil) {
        paramMap["activity_id"] = *req.ActivityId
    }
    if(req.Sid != nil) {
        paramMap["sid"] = *req.Sid
    }
    if(req.Mobile != nil) {
        paramMap["mobile"] = *req.Mobile
    }
    if(req.OuterOrderId != nil) {
        paramMap["outer_order_id"] = *req.OuterOrderId
    }
    return paramMap
}

func (req *AlibabaAlscUnionElemeStorepromotionReviewbwcStockLockRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}