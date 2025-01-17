package request


type AlibabaAlscUnionKbcpxRefundOrderGetRequest struct {
    /*
        时间维度，1.订单结算时间 2.维权创建时间 3.维权完成时间 4更新时间 defalutValue��1    */
    DateType  *int64 `json:"date_type" required:"true" `
    /*
        查询截止时间。开始和结束时间不能超过31天     */
    EndDate  *string `json:"end_date,omitempty" required:"false" `
    /*
        1-CPA 2-CPS defalutValue��1    */
    BizUnit  *int64 `json:"biz_unit" required:"true" `
    /*
        每页返回数据大小，默认10，最大返回50 defalutValue��10    */
    PageSize  *int64 `json:"page_size" required:"true" `
    /*
        页码，默认第一页，取值范围1~50 defalutValue��1    */
    PageNumber  *int64 `json:"page_number" required:"true" `
    /*
        查询起始时间。开始和结束时间不能超过31天     */
    StartDate  *string `json:"start_date" required:"true" `
    /*
        场景值，支持多场景（英文逗号分隔）查询7卡券订单，8卡券核销订单，10-媒体出资CPS红包，11-媒体出资霸王餐加码红包     */
    FlowType  *string `json:"flow_type,omitempty" required:"false" `
    /*
        推广位pid     */
    Pid  *string `json:"pid,omitempty" required:"false" `
    /*
        淘宝子订单号或饿了么订单号     */
    OrderId  *string `json:"order_id,omitempty" required:"false" `
}

func (s *AlibabaAlscUnionKbcpxRefundOrderGetRequest) SetDateType(v int64) *AlibabaAlscUnionKbcpxRefundOrderGetRequest {
    s.DateType = &v
    return s
}
func (s *AlibabaAlscUnionKbcpxRefundOrderGetRequest) SetEndDate(v string) *AlibabaAlscUnionKbcpxRefundOrderGetRequest {
    s.EndDate = &v
    return s
}
func (s *AlibabaAlscUnionKbcpxRefundOrderGetRequest) SetBizUnit(v int64) *AlibabaAlscUnionKbcpxRefundOrderGetRequest {
    s.BizUnit = &v
    return s
}
func (s *AlibabaAlscUnionKbcpxRefundOrderGetRequest) SetPageSize(v int64) *AlibabaAlscUnionKbcpxRefundOrderGetRequest {
    s.PageSize = &v
    return s
}
func (s *AlibabaAlscUnionKbcpxRefundOrderGetRequest) SetPageNumber(v int64) *AlibabaAlscUnionKbcpxRefundOrderGetRequest {
    s.PageNumber = &v
    return s
}
func (s *AlibabaAlscUnionKbcpxRefundOrderGetRequest) SetStartDate(v string) *AlibabaAlscUnionKbcpxRefundOrderGetRequest {
    s.StartDate = &v
    return s
}
func (s *AlibabaAlscUnionKbcpxRefundOrderGetRequest) SetFlowType(v string) *AlibabaAlscUnionKbcpxRefundOrderGetRequest {
    s.FlowType = &v
    return s
}
func (s *AlibabaAlscUnionKbcpxRefundOrderGetRequest) SetPid(v string) *AlibabaAlscUnionKbcpxRefundOrderGetRequest {
    s.Pid = &v
    return s
}
func (s *AlibabaAlscUnionKbcpxRefundOrderGetRequest) SetOrderId(v string) *AlibabaAlscUnionKbcpxRefundOrderGetRequest {
    s.OrderId = &v
    return s
}

func (req *AlibabaAlscUnionKbcpxRefundOrderGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.DateType != nil) {
        paramMap["date_type"] = *req.DateType
    }
    if(req.EndDate != nil) {
        paramMap["end_date"] = *req.EndDate
    }
    if(req.BizUnit != nil) {
        paramMap["biz_unit"] = *req.BizUnit
    }
    if(req.PageSize != nil) {
        paramMap["page_size"] = *req.PageSize
    }
    if(req.PageNumber != nil) {
        paramMap["page_number"] = *req.PageNumber
    }
    if(req.StartDate != nil) {
        paramMap["start_date"] = *req.StartDate
    }
    if(req.FlowType != nil) {
        paramMap["flow_type"] = *req.FlowType
    }
    if(req.Pid != nil) {
        paramMap["pid"] = *req.Pid
    }
    if(req.OrderId != nil) {
        paramMap["order_id"] = *req.OrderId
    }
    return paramMap
}

func (req *AlibabaAlscUnionKbcpxRefundOrderGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}