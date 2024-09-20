package request


type AlibabaAlscUnionElemeStorepromotionReviewbwcDiagnoseRequest struct {
    /*
        资格ID     */
    LockId  *int64 `json:"lock_id" required:"true" `
    /*
        饿了么订单ID     */
    OrderId  *string `json:"order_id" required:"true" `
}

func (s *AlibabaAlscUnionElemeStorepromotionReviewbwcDiagnoseRequest) SetLockId(v int64) *AlibabaAlscUnionElemeStorepromotionReviewbwcDiagnoseRequest {
    s.LockId = &v
    return s
}
func (s *AlibabaAlscUnionElemeStorepromotionReviewbwcDiagnoseRequest) SetOrderId(v string) *AlibabaAlscUnionElemeStorepromotionReviewbwcDiagnoseRequest {
    s.OrderId = &v
    return s
}

func (req *AlibabaAlscUnionElemeStorepromotionReviewbwcDiagnoseRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.LockId != nil) {
        paramMap["lock_id"] = *req.LockId
    }
    if(req.OrderId != nil) {
        paramMap["order_id"] = *req.OrderId
    }
    return paramMap
}

func (req *AlibabaAlscUnionElemeStorepromotionReviewbwcDiagnoseRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}