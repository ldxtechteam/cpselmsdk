package domain


type AlibabaAlscUnionKbcpxRefundOrderGetRefundOrderDetailReportDTO struct {
    /*
        商品标题     */
    Title  *string `json:"title,omitempty" `

    /*
        商品图片url     */
    PicUrl  *string `json:"pic_url,omitempty" `

    /*
        店铺名称     */
    ShopName  *string `json:"shop_name,omitempty" `

    /*
        付款金额     */
    PayAmount  *string `json:"pay_amount,omitempty" `

    /*
        订单结算时间     */
    SettleTime  *string `json:"settle_time,omitempty" `

    /*
        维权创建时间     */
    ExplainStartTime  *string `json:"explain_start_time,omitempty" `

    /*
        维权结束时间     */
    ExplainEndTime  *string `json:"explain_end_time,omitempty" `

    /*
        淘宝子单号     */
    BizOrderId  *string `json:"biz_order_id,omitempty" `

    /*
        淘宝主单号     */
    ParentOrderId  *string `json:"parent_order_id,omitempty" `

    /*
        主商品ID，针对CPS订单     */
    MainItemId  *string `json:"main_item_id,omitempty" `

    /*
        主商品信息，针对CPS订单     */
    MainItemTitle  *string `json:"main_item_title,omitempty" `

    /*
        维权状态，0-维权成功 1-维权创建 2-维权关闭 3-维权失败     */
    ExplainState  *int64 `json:"explain_state,omitempty" `

    /*
        维权成功后的佣金返回状态 0-待返回 1-已返回     */
    ReturnCommissionState  *int64 `json:"return_commission_state,omitempty" `

    /*
        维权退款金额     */
    RefundAmount  *string `json:"refund_amount,omitempty" `

    /*
        应返回商家金额     */
    Settle  *string `json:"settle,omitempty" `

    /*
        更新时间     */
    GmtModified  *string `json:"gmt_modified,omitempty" `

    /*
        会员标识（开通SID权限后下单返回）     */
    Sid  *string `json:"sid,omitempty" `

    /*
        1口碑，2饿了么     */
    PlatformType  *int64 `json:"platform_type,omitempty" `

    /*
        活动ID     */
    ActivityId  *int64 `json:"activity_id,omitempty" `

    /*
        pid     */
    Pid  *string `json:"pid,omitempty" `

    /*
        卡券订单号     */
    RelationOrderId  *int64 `json:"relation_order_id,omitempty" `

    /*
        场景值，7卡券订单，8卡券核销订单。16/17囤券券订单，18囤券券外卖订单     */
    FlowType  *int64 `json:"flow_type,omitempty" `

    /*
        活动信息明细     */
    ActivityInfoRemarkList  *string `json:"activity_info_remark_list,omitempty" `

    /*
        权益ID（flow_type=10或11时为媒体出资活动ID）     */
    ChannelRightId  *string `json:"channel_right_id,omitempty" `

    /*
        订单扩展信息     */
    ExtInfo  *string `json:"ext_info,omitempty" `

}

func (s *AlibabaAlscUnionKbcpxRefundOrderGetRefundOrderDetailReportDTO) SetTitle(v string) *AlibabaAlscUnionKbcpxRefundOrderGetRefundOrderDetailReportDTO {
    s.Title = &v
    return s
}
func (s *AlibabaAlscUnionKbcpxRefundOrderGetRefundOrderDetailReportDTO) SetPicUrl(v string) *AlibabaAlscUnionKbcpxRefundOrderGetRefundOrderDetailReportDTO {
    s.PicUrl = &v
    return s
}
func (s *AlibabaAlscUnionKbcpxRefundOrderGetRefundOrderDetailReportDTO) SetShopName(v string) *AlibabaAlscUnionKbcpxRefundOrderGetRefundOrderDetailReportDTO {
    s.ShopName = &v
    return s
}
func (s *AlibabaAlscUnionKbcpxRefundOrderGetRefundOrderDetailReportDTO) SetPayAmount(v string) *AlibabaAlscUnionKbcpxRefundOrderGetRefundOrderDetailReportDTO {
    s.PayAmount = &v
    return s
}
func (s *AlibabaAlscUnionKbcpxRefundOrderGetRefundOrderDetailReportDTO) SetSettleTime(v string) *AlibabaAlscUnionKbcpxRefundOrderGetRefundOrderDetailReportDTO {
    s.SettleTime = &v
    return s
}
func (s *AlibabaAlscUnionKbcpxRefundOrderGetRefundOrderDetailReportDTO) SetExplainStartTime(v string) *AlibabaAlscUnionKbcpxRefundOrderGetRefundOrderDetailReportDTO {
    s.ExplainStartTime = &v
    return s
}
func (s *AlibabaAlscUnionKbcpxRefundOrderGetRefundOrderDetailReportDTO) SetExplainEndTime(v string) *AlibabaAlscUnionKbcpxRefundOrderGetRefundOrderDetailReportDTO {
    s.ExplainEndTime = &v
    return s
}
func (s *AlibabaAlscUnionKbcpxRefundOrderGetRefundOrderDetailReportDTO) SetBizOrderId(v string) *AlibabaAlscUnionKbcpxRefundOrderGetRefundOrderDetailReportDTO {
    s.BizOrderId = &v
    return s
}
func (s *AlibabaAlscUnionKbcpxRefundOrderGetRefundOrderDetailReportDTO) SetParentOrderId(v string) *AlibabaAlscUnionKbcpxRefundOrderGetRefundOrderDetailReportDTO {
    s.ParentOrderId = &v
    return s
}
func (s *AlibabaAlscUnionKbcpxRefundOrderGetRefundOrderDetailReportDTO) SetMainItemId(v string) *AlibabaAlscUnionKbcpxRefundOrderGetRefundOrderDetailReportDTO {
    s.MainItemId = &v
    return s
}
func (s *AlibabaAlscUnionKbcpxRefundOrderGetRefundOrderDetailReportDTO) SetMainItemTitle(v string) *AlibabaAlscUnionKbcpxRefundOrderGetRefundOrderDetailReportDTO {
    s.MainItemTitle = &v
    return s
}
func (s *AlibabaAlscUnionKbcpxRefundOrderGetRefundOrderDetailReportDTO) SetExplainState(v int64) *AlibabaAlscUnionKbcpxRefundOrderGetRefundOrderDetailReportDTO {
    s.ExplainState = &v
    return s
}
func (s *AlibabaAlscUnionKbcpxRefundOrderGetRefundOrderDetailReportDTO) SetReturnCommissionState(v int64) *AlibabaAlscUnionKbcpxRefundOrderGetRefundOrderDetailReportDTO {
    s.ReturnCommissionState = &v
    return s
}
func (s *AlibabaAlscUnionKbcpxRefundOrderGetRefundOrderDetailReportDTO) SetRefundAmount(v string) *AlibabaAlscUnionKbcpxRefundOrderGetRefundOrderDetailReportDTO {
    s.RefundAmount = &v
    return s
}
func (s *AlibabaAlscUnionKbcpxRefundOrderGetRefundOrderDetailReportDTO) SetSettle(v string) *AlibabaAlscUnionKbcpxRefundOrderGetRefundOrderDetailReportDTO {
    s.Settle = &v
    return s
}
func (s *AlibabaAlscUnionKbcpxRefundOrderGetRefundOrderDetailReportDTO) SetGmtModified(v string) *AlibabaAlscUnionKbcpxRefundOrderGetRefundOrderDetailReportDTO {
    s.GmtModified = &v
    return s
}
func (s *AlibabaAlscUnionKbcpxRefundOrderGetRefundOrderDetailReportDTO) SetSid(v string) *AlibabaAlscUnionKbcpxRefundOrderGetRefundOrderDetailReportDTO {
    s.Sid = &v
    return s
}
func (s *AlibabaAlscUnionKbcpxRefundOrderGetRefundOrderDetailReportDTO) SetPlatformType(v int64) *AlibabaAlscUnionKbcpxRefundOrderGetRefundOrderDetailReportDTO {
    s.PlatformType = &v
    return s
}
func (s *AlibabaAlscUnionKbcpxRefundOrderGetRefundOrderDetailReportDTO) SetActivityId(v int64) *AlibabaAlscUnionKbcpxRefundOrderGetRefundOrderDetailReportDTO {
    s.ActivityId = &v
    return s
}
func (s *AlibabaAlscUnionKbcpxRefundOrderGetRefundOrderDetailReportDTO) SetPid(v string) *AlibabaAlscUnionKbcpxRefundOrderGetRefundOrderDetailReportDTO {
    s.Pid = &v
    return s
}
func (s *AlibabaAlscUnionKbcpxRefundOrderGetRefundOrderDetailReportDTO) SetRelationOrderId(v int64) *AlibabaAlscUnionKbcpxRefundOrderGetRefundOrderDetailReportDTO {
    s.RelationOrderId = &v
    return s
}
func (s *AlibabaAlscUnionKbcpxRefundOrderGetRefundOrderDetailReportDTO) SetFlowType(v int64) *AlibabaAlscUnionKbcpxRefundOrderGetRefundOrderDetailReportDTO {
    s.FlowType = &v
    return s
}
func (s *AlibabaAlscUnionKbcpxRefundOrderGetRefundOrderDetailReportDTO) SetActivityInfoRemarkList(v string) *AlibabaAlscUnionKbcpxRefundOrderGetRefundOrderDetailReportDTO {
    s.ActivityInfoRemarkList = &v
    return s
}
func (s *AlibabaAlscUnionKbcpxRefundOrderGetRefundOrderDetailReportDTO) SetChannelRightId(v string) *AlibabaAlscUnionKbcpxRefundOrderGetRefundOrderDetailReportDTO {
    s.ChannelRightId = &v
    return s
}
func (s *AlibabaAlscUnionKbcpxRefundOrderGetRefundOrderDetailReportDTO) SetExtInfo(v string) *AlibabaAlscUnionKbcpxRefundOrderGetRefundOrderDetailReportDTO {
    s.ExtInfo = &v
    return s
}
