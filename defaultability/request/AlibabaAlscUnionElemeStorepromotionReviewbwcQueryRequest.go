package request


type AlibabaAlscUnionElemeStorepromotionReviewbwcQueryRequest struct {
    /*
        每页数量（1~20，默认20）     */
    PageSize  *int64 `json:"page_size,omitempty" required:"false" `
    /*
        会话ID（分页场景首次请求结果返回，后续请求必须携带，服务根据同一个session_id的请求次数自动叠加翻页返回数据，直至分页结束，返回空）     */
    SessionId  *string `json:"session_id,omitempty" required:"false" `
    /*
        渠道PID     */
    Pid  *string `json:"pid" required:"true" `
    /*
        经度     */
    Longitude  *string `json:"longitude" required:"true" `
    /*
        纬度     */
    Latitude  *string `json:"latitude" required:"true" `
    /*
        排序类型，默认normal，排序规则包括:{"normal":"佣金倒序","distance_asc":"距离由近到远","commission_desc":"佣金倒序","month_sales_desc":"月销量从高到低","commission_rate_desc":"佣金比例倒序", "activity_reward_desc":"返现金额倒序"}     */
    SortType  *string `json:"sort_type,omitempty" required:"false" `
    /*
        店铺佣金比例下限，代表筛选店铺全店佣金大于等于0.01的店铺     */
    MinCommissionRate  *string `json:"min_commission_rate,omitempty" required:"false" `
    /*
        三方扩展id     */
    Sid  *string `json:"sid,omitempty" required:"false" `
    /*
        以一级类目进行类目限定，以,或者|进行类目分隔     */
    FilterFirstCategories  *string `json:"filter_first_categories,omitempty" required:"false" `
    /*
        1.5级类目查询，以"|"分隔     */
    FilterOnePointFiveCategories  *string `json:"filter_one_point_five_categories,omitempty" required:"false" `
    /*
        城市ID（经纬度范围覆盖多个城市时，精准召回）     */
    FilterCityId  *string `json:"filter_city_id,omitempty" required:"false" `
    /*
        搜索内容（店铺名）     */
    SearchContent  *string `json:"search_content,omitempty" required:"false" `
    /*
        false-返回链接，true-不返回链接；8.1号及以后该属性移除，默认不返回链接 defalutValue��true    */
    ExcludeLink  *bool `json:"exclude_link,omitempty" required:"false" `
}

func (s *AlibabaAlscUnionElemeStorepromotionReviewbwcQueryRequest) SetPageSize(v int64) *AlibabaAlscUnionElemeStorepromotionReviewbwcQueryRequest {
    s.PageSize = &v
    return s
}
func (s *AlibabaAlscUnionElemeStorepromotionReviewbwcQueryRequest) SetSessionId(v string) *AlibabaAlscUnionElemeStorepromotionReviewbwcQueryRequest {
    s.SessionId = &v
    return s
}
func (s *AlibabaAlscUnionElemeStorepromotionReviewbwcQueryRequest) SetPid(v string) *AlibabaAlscUnionElemeStorepromotionReviewbwcQueryRequest {
    s.Pid = &v
    return s
}
func (s *AlibabaAlscUnionElemeStorepromotionReviewbwcQueryRequest) SetLongitude(v string) *AlibabaAlscUnionElemeStorepromotionReviewbwcQueryRequest {
    s.Longitude = &v
    return s
}
func (s *AlibabaAlscUnionElemeStorepromotionReviewbwcQueryRequest) SetLatitude(v string) *AlibabaAlscUnionElemeStorepromotionReviewbwcQueryRequest {
    s.Latitude = &v
    return s
}
func (s *AlibabaAlscUnionElemeStorepromotionReviewbwcQueryRequest) SetSortType(v string) *AlibabaAlscUnionElemeStorepromotionReviewbwcQueryRequest {
    s.SortType = &v
    return s
}
func (s *AlibabaAlscUnionElemeStorepromotionReviewbwcQueryRequest) SetMinCommissionRate(v string) *AlibabaAlscUnionElemeStorepromotionReviewbwcQueryRequest {
    s.MinCommissionRate = &v
    return s
}
func (s *AlibabaAlscUnionElemeStorepromotionReviewbwcQueryRequest) SetSid(v string) *AlibabaAlscUnionElemeStorepromotionReviewbwcQueryRequest {
    s.Sid = &v
    return s
}
func (s *AlibabaAlscUnionElemeStorepromotionReviewbwcQueryRequest) SetFilterFirstCategories(v string) *AlibabaAlscUnionElemeStorepromotionReviewbwcQueryRequest {
    s.FilterFirstCategories = &v
    return s
}
func (s *AlibabaAlscUnionElemeStorepromotionReviewbwcQueryRequest) SetFilterOnePointFiveCategories(v string) *AlibabaAlscUnionElemeStorepromotionReviewbwcQueryRequest {
    s.FilterOnePointFiveCategories = &v
    return s
}
func (s *AlibabaAlscUnionElemeStorepromotionReviewbwcQueryRequest) SetFilterCityId(v string) *AlibabaAlscUnionElemeStorepromotionReviewbwcQueryRequest {
    s.FilterCityId = &v
    return s
}
func (s *AlibabaAlscUnionElemeStorepromotionReviewbwcQueryRequest) SetSearchContent(v string) *AlibabaAlscUnionElemeStorepromotionReviewbwcQueryRequest {
    s.SearchContent = &v
    return s
}
func (s *AlibabaAlscUnionElemeStorepromotionReviewbwcQueryRequest) SetExcludeLink(v bool) *AlibabaAlscUnionElemeStorepromotionReviewbwcQueryRequest {
    s.ExcludeLink = &v
    return s
}

func (req *AlibabaAlscUnionElemeStorepromotionReviewbwcQueryRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.PageSize != nil) {
        paramMap["page_size"] = *req.PageSize
    }
    if(req.SessionId != nil) {
        paramMap["session_id"] = *req.SessionId
    }
    if(req.Pid != nil) {
        paramMap["pid"] = *req.Pid
    }
    if(req.Longitude != nil) {
        paramMap["longitude"] = *req.Longitude
    }
    if(req.Latitude != nil) {
        paramMap["latitude"] = *req.Latitude
    }
    if(req.SortType != nil) {
        paramMap["sort_type"] = *req.SortType
    }
    if(req.MinCommissionRate != nil) {
        paramMap["min_commission_rate"] = *req.MinCommissionRate
    }
    if(req.Sid != nil) {
        paramMap["sid"] = *req.Sid
    }
    if(req.FilterFirstCategories != nil) {
        paramMap["filter_first_categories"] = *req.FilterFirstCategories
    }
    if(req.FilterOnePointFiveCategories != nil) {
        paramMap["filter_one_point_five_categories"] = *req.FilterOnePointFiveCategories
    }
    if(req.FilterCityId != nil) {
        paramMap["filter_city_id"] = *req.FilterCityId
    }
    if(req.SearchContent != nil) {
        paramMap["search_content"] = *req.SearchContent
    }
    if(req.ExcludeLink != nil) {
        paramMap["exclude_link"] = *req.ExcludeLink
    }
    return paramMap
}

func (req *AlibabaAlscUnionElemeStorepromotionReviewbwcQueryRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}