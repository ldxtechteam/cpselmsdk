package request


type AlibabaAlscUnionElemeStorepromotionReviewbwcStockReleaseRequest struct {
    /*
        库存锁ID     */
    LockId  *int64 `json:"lock_id" required:"true" `
}

func (s *AlibabaAlscUnionElemeStorepromotionReviewbwcStockReleaseRequest) SetLockId(v int64) *AlibabaAlscUnionElemeStorepromotionReviewbwcStockReleaseRequest {
    s.LockId = &v
    return s
}

func (req *AlibabaAlscUnionElemeStorepromotionReviewbwcStockReleaseRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.LockId != nil) {
        paramMap["lock_id"] = *req.LockId
    }
    return paramMap
}

func (req *AlibabaAlscUnionElemeStorepromotionReviewbwcStockReleaseRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}