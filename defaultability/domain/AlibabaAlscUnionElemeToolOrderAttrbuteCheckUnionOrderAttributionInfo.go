package domain


type AlibabaAlscUnionElemeToolOrderAttrbuteCheckUnionOrderAttributionInfo struct {
    /*
        订单基本信息     */
    OrderInfo  *AlibabaAlscUnionElemeToolOrderAttrbuteCheckOrderInfo `json:"order_info,omitempty" `

    /*
        归因结果数据     */
    DetailInfos  *[]AlibabaAlscUnionElemeToolOrderAttrbuteCheckAttributionDetailInfo `json:"detail_infos,omitempty" `

}

func (s *AlibabaAlscUnionElemeToolOrderAttrbuteCheckUnionOrderAttributionInfo) SetOrderInfo(v AlibabaAlscUnionElemeToolOrderAttrbuteCheckOrderInfo) *AlibabaAlscUnionElemeToolOrderAttrbuteCheckUnionOrderAttributionInfo {
    s.OrderInfo = &v
    return s
}
func (s *AlibabaAlscUnionElemeToolOrderAttrbuteCheckUnionOrderAttributionInfo) SetDetailInfos(v []AlibabaAlscUnionElemeToolOrderAttrbuteCheckAttributionDetailInfo) *AlibabaAlscUnionElemeToolOrderAttrbuteCheckUnionOrderAttributionInfo {
    s.DetailInfos = &v
    return s
}
