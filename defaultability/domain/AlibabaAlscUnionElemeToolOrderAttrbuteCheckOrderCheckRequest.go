package domain


type AlibabaAlscUnionElemeToolOrderAttrbuteCheckOrderCheckRequest struct {
    /*
        饿了么订单id     */
    OrderId  *string `json:"order_id,omitempty" `

}

func (s *AlibabaAlscUnionElemeToolOrderAttrbuteCheckOrderCheckRequest) SetOrderId(v string) *AlibabaAlscUnionElemeToolOrderAttrbuteCheckOrderCheckRequest {
    s.OrderId = &v
    return s
}
