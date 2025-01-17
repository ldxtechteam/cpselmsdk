package domain


type AlibabaAlscUnionElemePromotionStorepromotionQueryStorePromotionDto struct {
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
        奖励金活动数据     */
    Activity  *AlibabaAlscUnionElemePromotionStorepromotionQueryPromotionActivity `json:"activity,omitempty" `

    /*
        推广链接     */
    Link  *AlibabaAlscUnionElemePromotionStorepromotionQueryPromotionLink `json:"link,omitempty" `

    /*
        一级类目ID，高级字段     */
    Category1Id  *string `json:"category_1_id,omitempty" `

    /*
        配送距离（米），高级字段     */
    DeliveryDistance  *int64 `json:"delivery_distance,omitempty" `

    /*
        配送时间（分），高级字段     */
    DeliveryTime  *int64 `json:"delivery_time,omitempty" `

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
    Items  *[]AlibabaAlscUnionElemePromotionStorepromotionQueryPromotionItem `json:"items,omitempty" `

    /*
        店铺ID（静态加密）     */
    ShopId  *string `json:"shop_id,omitempty" `

    /*
        一级类目名称，高级字段     */
    Category1Name  *string `json:"category_1_name,omitempty" `

    /*
        预估佣金（分）     */
    Commission  *int64 `json:"commission,omitempty" `

    /*
        满减标签，高级字段     */
    DiscountTags  *[]string `json:"discount_tags,omitempty" `

    /*
        配送费（元），高级字段     */
    DeliveryFee  *string `json:"delivery_fee,omitempty" `

}

func (s *AlibabaAlscUnionElemePromotionStorepromotionQueryStorePromotionDto) SetTitle(v string) *AlibabaAlscUnionElemePromotionStorepromotionQueryStorePromotionDto {
    s.Title = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionStorepromotionQueryStorePromotionDto) SetShopLogo(v string) *AlibabaAlscUnionElemePromotionStorepromotionQueryStorePromotionDto {
    s.ShopLogo = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionStorepromotionQueryStorePromotionDto) SetIndistinctMonthlySales(v string) *AlibabaAlscUnionElemePromotionStorepromotionQueryStorePromotionDto {
    s.IndistinctMonthlySales = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionStorepromotionQueryStorePromotionDto) SetCommissionRate(v string) *AlibabaAlscUnionElemePromotionStorepromotionQueryStorePromotionDto {
    s.CommissionRate = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionStorepromotionQueryStorePromotionDto) SetBizType(v string) *AlibabaAlscUnionElemePromotionStorepromotionQueryStorePromotionDto {
    s.BizType = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionStorepromotionQueryStorePromotionDto) SetActivity(v AlibabaAlscUnionElemePromotionStorepromotionQueryPromotionActivity) *AlibabaAlscUnionElemePromotionStorepromotionQueryStorePromotionDto {
    s.Activity = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionStorepromotionQueryStorePromotionDto) SetLink(v AlibabaAlscUnionElemePromotionStorepromotionQueryPromotionLink) *AlibabaAlscUnionElemePromotionStorepromotionQueryStorePromotionDto {
    s.Link = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionStorepromotionQueryStorePromotionDto) SetCategory1Id(v string) *AlibabaAlscUnionElemePromotionStorepromotionQueryStorePromotionDto {
    s.Category1Id = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionStorepromotionQueryStorePromotionDto) SetDeliveryDistance(v int64) *AlibabaAlscUnionElemePromotionStorepromotionQueryStorePromotionDto {
    s.DeliveryDistance = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionStorepromotionQueryStorePromotionDto) SetDeliveryTime(v int64) *AlibabaAlscUnionElemePromotionStorepromotionQueryStorePromotionDto {
    s.DeliveryTime = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionStorepromotionQueryStorePromotionDto) SetDeliveryPrice(v string) *AlibabaAlscUnionElemePromotionStorepromotionQueryStorePromotionDto {
    s.DeliveryPrice = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionStorepromotionQueryStorePromotionDto) SetRecommendReasons(v []string) *AlibabaAlscUnionElemePromotionStorepromotionQueryStorePromotionDto {
    s.RecommendReasons = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionStorepromotionQueryStorePromotionDto) SetServiceRating(v string) *AlibabaAlscUnionElemePromotionStorepromotionQueryStorePromotionDto {
    s.ServiceRating = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionStorepromotionQueryStorePromotionDto) SetItems(v []AlibabaAlscUnionElemePromotionStorepromotionQueryPromotionItem) *AlibabaAlscUnionElemePromotionStorepromotionQueryStorePromotionDto {
    s.Items = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionStorepromotionQueryStorePromotionDto) SetShopId(v string) *AlibabaAlscUnionElemePromotionStorepromotionQueryStorePromotionDto {
    s.ShopId = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionStorepromotionQueryStorePromotionDto) SetCategory1Name(v string) *AlibabaAlscUnionElemePromotionStorepromotionQueryStorePromotionDto {
    s.Category1Name = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionStorepromotionQueryStorePromotionDto) SetCommission(v int64) *AlibabaAlscUnionElemePromotionStorepromotionQueryStorePromotionDto {
    s.Commission = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionStorepromotionQueryStorePromotionDto) SetDiscountTags(v []string) *AlibabaAlscUnionElemePromotionStorepromotionQueryStorePromotionDto {
    s.DiscountTags = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionStorepromotionQueryStorePromotionDto) SetDeliveryFee(v string) *AlibabaAlscUnionElemePromotionStorepromotionQueryStorePromotionDto {
    s.DeliveryFee = &v
    return s
}
