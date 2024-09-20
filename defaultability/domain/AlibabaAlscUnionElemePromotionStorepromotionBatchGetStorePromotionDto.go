package domain


type AlibabaAlscUnionElemePromotionStorepromotionBatchGetStorePromotionDto struct {
    /*
        门店名称     */
    Title  *string `json:"title,omitempty" `

    /*
        门店logo     */
    ShopLogo  *string `json:"shop_logo,omitempty" `

    /*
        模糊销量     */
    IndistinctMonthlySales  *string `json:"indistinct_monthly_sales,omitempty" `

    /*
        佣金比例     */
    CommissionRate  *string `json:"commission_rate,omitempty" `

    /*
        店铺类型（"activityCps":活动cps,"ordinaryCps":基础cps）     */
    BizType  *string `json:"biz_type,omitempty" `

    /*
        活动数据     */
    Activity  *AlibabaAlscUnionElemePromotionStorepromotionBatchGetPromotionActivity `json:"activity,omitempty" `

    /*
        推广链接     */
    Link  *AlibabaAlscUnionElemePromotionStorepromotionBatchGetPromotionLink `json:"link,omitempty" `

    /*
        一级类目ID，高级字段     */
    Category1Id  *string `json:"category_1_id,omitempty" `

    /*
        起送价（元），高级字段     */
    DeliveryPrice  *string `json:"delivery_price,omitempty" `

    /*
        推荐理由，高级字段     */
    RecommendReasons  *[]string `json:"recommend_reasons,omitempty" `

    /*
        服务评级，高级字段     */
    ServiceRating  *string `json:"service_rating,omitempty" `

    /*
        推荐商品，高级字段     */
    Items  *[]AlibabaAlscUnionElemePromotionStorepromotionBatchGetPromotionItem `json:"items,omitempty" `

    /*
        一级类目名称，高级字段     */
    Category1Name  *string `json:"category_1_name,omitempty" `

    /*
        满减标签，高级字段     */
    DiscountTags  *[]string `json:"discount_tags,omitempty" `

    /*
        配送费（元）     */
    DeliveryFee  *string `json:"delivery_fee,omitempty" `

    /*
        配送距离（米），高级字段     */
    DeliveryDistance  *int64 `json:"delivery_distance,omitempty" `

    /*
        配送时间（分），高级字段     */
    DeliveryTime  *int64 `json:"delivery_time,omitempty" `

    /*
        店铺ID（加密）     */
    ShopId  *string `json:"shop_id,omitempty" `

    /*
        预估佣金（分）     */
    Commission  *int64 `json:"commission,omitempty" `

}

func (s *AlibabaAlscUnionElemePromotionStorepromotionBatchGetStorePromotionDto) SetTitle(v string) *AlibabaAlscUnionElemePromotionStorepromotionBatchGetStorePromotionDto {
    s.Title = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionStorepromotionBatchGetStorePromotionDto) SetShopLogo(v string) *AlibabaAlscUnionElemePromotionStorepromotionBatchGetStorePromotionDto {
    s.ShopLogo = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionStorepromotionBatchGetStorePromotionDto) SetIndistinctMonthlySales(v string) *AlibabaAlscUnionElemePromotionStorepromotionBatchGetStorePromotionDto {
    s.IndistinctMonthlySales = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionStorepromotionBatchGetStorePromotionDto) SetCommissionRate(v string) *AlibabaAlscUnionElemePromotionStorepromotionBatchGetStorePromotionDto {
    s.CommissionRate = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionStorepromotionBatchGetStorePromotionDto) SetBizType(v string) *AlibabaAlscUnionElemePromotionStorepromotionBatchGetStorePromotionDto {
    s.BizType = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionStorepromotionBatchGetStorePromotionDto) SetActivity(v AlibabaAlscUnionElemePromotionStorepromotionBatchGetPromotionActivity) *AlibabaAlscUnionElemePromotionStorepromotionBatchGetStorePromotionDto {
    s.Activity = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionStorepromotionBatchGetStorePromotionDto) SetLink(v AlibabaAlscUnionElemePromotionStorepromotionBatchGetPromotionLink) *AlibabaAlscUnionElemePromotionStorepromotionBatchGetStorePromotionDto {
    s.Link = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionStorepromotionBatchGetStorePromotionDto) SetCategory1Id(v string) *AlibabaAlscUnionElemePromotionStorepromotionBatchGetStorePromotionDto {
    s.Category1Id = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionStorepromotionBatchGetStorePromotionDto) SetDeliveryPrice(v string) *AlibabaAlscUnionElemePromotionStorepromotionBatchGetStorePromotionDto {
    s.DeliveryPrice = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionStorepromotionBatchGetStorePromotionDto) SetRecommendReasons(v []string) *AlibabaAlscUnionElemePromotionStorepromotionBatchGetStorePromotionDto {
    s.RecommendReasons = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionStorepromotionBatchGetStorePromotionDto) SetServiceRating(v string) *AlibabaAlscUnionElemePromotionStorepromotionBatchGetStorePromotionDto {
    s.ServiceRating = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionStorepromotionBatchGetStorePromotionDto) SetItems(v []AlibabaAlscUnionElemePromotionStorepromotionBatchGetPromotionItem) *AlibabaAlscUnionElemePromotionStorepromotionBatchGetStorePromotionDto {
    s.Items = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionStorepromotionBatchGetStorePromotionDto) SetCategory1Name(v string) *AlibabaAlscUnionElemePromotionStorepromotionBatchGetStorePromotionDto {
    s.Category1Name = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionStorepromotionBatchGetStorePromotionDto) SetDiscountTags(v []string) *AlibabaAlscUnionElemePromotionStorepromotionBatchGetStorePromotionDto {
    s.DiscountTags = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionStorepromotionBatchGetStorePromotionDto) SetDeliveryFee(v string) *AlibabaAlscUnionElemePromotionStorepromotionBatchGetStorePromotionDto {
    s.DeliveryFee = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionStorepromotionBatchGetStorePromotionDto) SetDeliveryDistance(v int64) *AlibabaAlscUnionElemePromotionStorepromotionBatchGetStorePromotionDto {
    s.DeliveryDistance = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionStorepromotionBatchGetStorePromotionDto) SetDeliveryTime(v int64) *AlibabaAlscUnionElemePromotionStorepromotionBatchGetStorePromotionDto {
    s.DeliveryTime = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionStorepromotionBatchGetStorePromotionDto) SetShopId(v string) *AlibabaAlscUnionElemePromotionStorepromotionBatchGetStorePromotionDto {
    s.ShopId = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionStorepromotionBatchGetStorePromotionDto) SetCommission(v int64) *AlibabaAlscUnionElemePromotionStorepromotionBatchGetStorePromotionDto {
    s.Commission = &v
    return s
}
