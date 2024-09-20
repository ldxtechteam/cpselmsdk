package domain


type AlibabaAlscUnionElemeStorepromotionReviewbwcStockReleaseReviewBwcStockReleaseResult struct {
    /*
        库存锁定ID     */
    LockId  *int64 `json:"lock_id,omitempty" `

}

func (s *AlibabaAlscUnionElemeStorepromotionReviewbwcStockReleaseReviewBwcStockReleaseResult) SetLockId(v int64) *AlibabaAlscUnionElemeStorepromotionReviewbwcStockReleaseReviewBwcStockReleaseResult {
    s.LockId = &v
    return s
}
