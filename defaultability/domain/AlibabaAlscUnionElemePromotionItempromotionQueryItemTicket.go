package domain


type AlibabaAlscUnionElemePromotionItempromotionQueryItemTicket struct {
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

func (s *AlibabaAlscUnionElemePromotionItempromotionQueryItemTicket) SetPrice(v int64) *AlibabaAlscUnionElemePromotionItempromotionQueryItemTicket {
    s.Price = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionItempromotionQueryItemTicket) SetQuantity(v int64) *AlibabaAlscUnionElemePromotionItempromotionQueryItemTicket {
    s.Quantity = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionItempromotionQueryItemTicket) SetThreshold(v int64) *AlibabaAlscUnionElemePromotionItempromotionQueryItemTicket {
    s.Threshold = &v
    return s
}
