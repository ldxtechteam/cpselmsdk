package domain


type AlibabaAlscUnionElemeStorepromotionReviewbwcStockLockReviewBwcStockLockResult struct {
    /*
        库存锁定ID     */
    LockId  *int64 `json:"lock_id,omitempty" `

    /*
        库存锁定时时间戳（毫秒）     */
    LockTime  *int64 `json:"lock_time,omitempty" `

}

func (s *AlibabaAlscUnionElemeStorepromotionReviewbwcStockLockReviewBwcStockLockResult) SetLockId(v int64) *AlibabaAlscUnionElemeStorepromotionReviewbwcStockLockReviewBwcStockLockResult {
    s.LockId = &v
    return s
}
func (s *AlibabaAlscUnionElemeStorepromotionReviewbwcStockLockReviewBwcStockLockResult) SetLockTime(v int64) *AlibabaAlscUnionElemeStorepromotionReviewbwcStockLockReviewBwcStockLockResult {
    s.LockTime = &v
    return s
}
