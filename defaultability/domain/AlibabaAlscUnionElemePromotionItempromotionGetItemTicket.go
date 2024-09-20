package domain


type AlibabaAlscUnionElemePromotionItempromotionGetItemTicket struct {
    /*
        价格（分）     */
    Price  *int64 `json:"price,omitempty" `

    /*
        数量     */
    Quantity  *int64 `json:"quantity,omitempty" `

    /*
        使用门槛（分）     */
    Threshold  *int64 `json:"threshold,omitempty" `

}

func (s *AlibabaAlscUnionElemePromotionItempromotionGetItemTicket) SetPrice(v int64) *AlibabaAlscUnionElemePromotionItempromotionGetItemTicket {
    s.Price = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionItempromotionGetItemTicket) SetQuantity(v int64) *AlibabaAlscUnionElemePromotionItempromotionGetItemTicket {
    s.Quantity = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionItempromotionGetItemTicket) SetThreshold(v int64) *AlibabaAlscUnionElemePromotionItempromotionGetItemTicket {
    s.Threshold = &v
    return s
}
