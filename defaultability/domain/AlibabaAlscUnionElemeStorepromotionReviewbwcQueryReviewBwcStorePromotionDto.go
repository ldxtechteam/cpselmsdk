package domain


type AlibabaAlscUnionElemeStorepromotionReviewbwcQueryReviewBwcStorePromotionDto struct {
    /*
        门店ID（加密）     */
    ShopId  *string `json:"shop_id,omitempty" `

    /*
        门店名称     */
    ShopName  *string `json:"shop_name,omitempty" `

    /*
        门店主图     */
    ShopPicture  *string `json:"shop_picture,omitempty" `

    /*
        月销量（模糊）     */
    IndistinctMonthlySales  *string `json:"indistinct_monthly_sales,omitempty" `

    /*
        一级类目ID     */
    Category1Id  *string `json:"category_1_id,omitempty" `

    /*
        一级类目名称     */
    Category1Name  *string `json:"category_1_name,omitempty" `

    /*
        配送距离（米）     */
    DeliveryDistance  *int64 `json:"delivery_distance,omitempty" `

    /*
        配送时间（分）     */
    DeliveryTime  *int64 `json:"delivery_time,omitempty" `

    /*
        起送价（分）     */
    DeliveryPriceCent  *int64 `json:"delivery_price_cent,omitempty" `

    /*
        服务评级     */
    ServiceRating  *string `json:"service_rating,omitempty" `

    /*
        推广链接，8.1号及以后该属性移除，默认不返回链接     */
    Link  *AlibabaAlscUnionElemeStorepromotionReviewbwcQueryPromotionItem `json:"link,omitempty" `

    /*
        活动     */
    Activities  *[]AlibabaAlscUnionElemeStorepromotionReviewbwcQueryReviewBwcActivity `json:"activities,omitempty" `

}

func (s *AlibabaAlscUnionElemeStorepromotionReviewbwcQueryReviewBwcStorePromotionDto) SetShopId(v string) *AlibabaAlscUnionElemeStorepromotionReviewbwcQueryReviewBwcStorePromotionDto {
    s.ShopId = &v
    return s
}
func (s *AlibabaAlscUnionElemeStorepromotionReviewbwcQueryReviewBwcStorePromotionDto) SetShopName(v string) *AlibabaAlscUnionElemeStorepromotionReviewbwcQueryReviewBwcStorePromotionDto {
    s.ShopName = &v
    return s
}
func (s *AlibabaAlscUnionElemeStorepromotionReviewbwcQueryReviewBwcStorePromotionDto) SetShopPicture(v string) *AlibabaAlscUnionElemeStorepromotionReviewbwcQueryReviewBwcStorePromotionDto {
    s.ShopPicture = &v
    return s
}
func (s *AlibabaAlscUnionElemeStorepromotionReviewbwcQueryReviewBwcStorePromotionDto) SetIndistinctMonthlySales(v string) *AlibabaAlscUnionElemeStorepromotionReviewbwcQueryReviewBwcStorePromotionDto {
    s.IndistinctMonthlySales = &v
    return s
}
func (s *AlibabaAlscUnionElemeStorepromotionReviewbwcQueryReviewBwcStorePromotionDto) SetCategory1Id(v string) *AlibabaAlscUnionElemeStorepromotionReviewbwcQueryReviewBwcStorePromotionDto {
    s.Category1Id = &v
    return s
}
func (s *AlibabaAlscUnionElemeStorepromotionReviewbwcQueryReviewBwcStorePromotionDto) SetCategory1Name(v string) *AlibabaAlscUnionElemeStorepromotionReviewbwcQueryReviewBwcStorePromotionDto {
    s.Category1Name = &v
    return s
}
func (s *AlibabaAlscUnionElemeStorepromotionReviewbwcQueryReviewBwcStorePromotionDto) SetDeliveryDistance(v int64) *AlibabaAlscUnionElemeStorepromotionReviewbwcQueryReviewBwcStorePromotionDto {
    s.DeliveryDistance = &v
    return s
}
func (s *AlibabaAlscUnionElemeStorepromotionReviewbwcQueryReviewBwcStorePromotionDto) SetDeliveryTime(v int64) *AlibabaAlscUnionElemeStorepromotionReviewbwcQueryReviewBwcStorePromotionDto {
    s.DeliveryTime = &v
    return s
}
func (s *AlibabaAlscUnionElemeStorepromotionReviewbwcQueryReviewBwcStorePromotionDto) SetDeliveryPriceCent(v int64) *AlibabaAlscUnionElemeStorepromotionReviewbwcQueryReviewBwcStorePromotionDto {
    s.DeliveryPriceCent = &v
    return s
}
func (s *AlibabaAlscUnionElemeStorepromotionReviewbwcQueryReviewBwcStorePromotionDto) SetServiceRating(v string) *AlibabaAlscUnionElemeStorepromotionReviewbwcQueryReviewBwcStorePromotionDto {
    s.ServiceRating = &v
    return s
}
func (s *AlibabaAlscUnionElemeStorepromotionReviewbwcQueryReviewBwcStorePromotionDto) SetLink(v AlibabaAlscUnionElemeStorepromotionReviewbwcQueryPromotionItem) *AlibabaAlscUnionElemeStorepromotionReviewbwcQueryReviewBwcStorePromotionDto {
    s.Link = &v
    return s
}
func (s *AlibabaAlscUnionElemeStorepromotionReviewbwcQueryReviewBwcStorePromotionDto) SetActivities(v []AlibabaAlscUnionElemeStorepromotionReviewbwcQueryReviewBwcActivity) *AlibabaAlscUnionElemeStorepromotionReviewbwcQueryReviewBwcStorePromotionDto {
    s.Activities = &v
    return s
}
