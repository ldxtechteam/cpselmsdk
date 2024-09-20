package domain


type AlibabaAlscUnionElemeStorepromotionReviewbwcDetailGetReviewBwcStorePromotionDto struct {
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
        推广链接     */
    Link  *AlibabaAlscUnionElemeStorepromotionReviewbwcDetailGetPromotionLink `json:"link,omitempty" `

    /*
        活动     */
    Activities  *[]AlibabaAlscUnionElemeStorepromotionReviewbwcDetailGetReviewBwcActivity `json:"activities,omitempty" `

}

func (s *AlibabaAlscUnionElemeStorepromotionReviewbwcDetailGetReviewBwcStorePromotionDto) SetShopId(v string) *AlibabaAlscUnionElemeStorepromotionReviewbwcDetailGetReviewBwcStorePromotionDto {
    s.ShopId = &v
    return s
}
func (s *AlibabaAlscUnionElemeStorepromotionReviewbwcDetailGetReviewBwcStorePromotionDto) SetShopName(v string) *AlibabaAlscUnionElemeStorepromotionReviewbwcDetailGetReviewBwcStorePromotionDto {
    s.ShopName = &v
    return s
}
func (s *AlibabaAlscUnionElemeStorepromotionReviewbwcDetailGetReviewBwcStorePromotionDto) SetShopPicture(v string) *AlibabaAlscUnionElemeStorepromotionReviewbwcDetailGetReviewBwcStorePromotionDto {
    s.ShopPicture = &v
    return s
}
func (s *AlibabaAlscUnionElemeStorepromotionReviewbwcDetailGetReviewBwcStorePromotionDto) SetIndistinctMonthlySales(v string) *AlibabaAlscUnionElemeStorepromotionReviewbwcDetailGetReviewBwcStorePromotionDto {
    s.IndistinctMonthlySales = &v
    return s
}
func (s *AlibabaAlscUnionElemeStorepromotionReviewbwcDetailGetReviewBwcStorePromotionDto) SetCategory1Id(v string) *AlibabaAlscUnionElemeStorepromotionReviewbwcDetailGetReviewBwcStorePromotionDto {
    s.Category1Id = &v
    return s
}
func (s *AlibabaAlscUnionElemeStorepromotionReviewbwcDetailGetReviewBwcStorePromotionDto) SetCategory1Name(v string) *AlibabaAlscUnionElemeStorepromotionReviewbwcDetailGetReviewBwcStorePromotionDto {
    s.Category1Name = &v
    return s
}
func (s *AlibabaAlscUnionElemeStorepromotionReviewbwcDetailGetReviewBwcStorePromotionDto) SetDeliveryDistance(v int64) *AlibabaAlscUnionElemeStorepromotionReviewbwcDetailGetReviewBwcStorePromotionDto {
    s.DeliveryDistance = &v
    return s
}
func (s *AlibabaAlscUnionElemeStorepromotionReviewbwcDetailGetReviewBwcStorePromotionDto) SetDeliveryTime(v int64) *AlibabaAlscUnionElemeStorepromotionReviewbwcDetailGetReviewBwcStorePromotionDto {
    s.DeliveryTime = &v
    return s
}
func (s *AlibabaAlscUnionElemeStorepromotionReviewbwcDetailGetReviewBwcStorePromotionDto) SetDeliveryPriceCent(v int64) *AlibabaAlscUnionElemeStorepromotionReviewbwcDetailGetReviewBwcStorePromotionDto {
    s.DeliveryPriceCent = &v
    return s
}
func (s *AlibabaAlscUnionElemeStorepromotionReviewbwcDetailGetReviewBwcStorePromotionDto) SetServiceRating(v string) *AlibabaAlscUnionElemeStorepromotionReviewbwcDetailGetReviewBwcStorePromotionDto {
    s.ServiceRating = &v
    return s
}
func (s *AlibabaAlscUnionElemeStorepromotionReviewbwcDetailGetReviewBwcStorePromotionDto) SetLink(v AlibabaAlscUnionElemeStorepromotionReviewbwcDetailGetPromotionLink) *AlibabaAlscUnionElemeStorepromotionReviewbwcDetailGetReviewBwcStorePromotionDto {
    s.Link = &v
    return s
}
func (s *AlibabaAlscUnionElemeStorepromotionReviewbwcDetailGetReviewBwcStorePromotionDto) SetActivities(v []AlibabaAlscUnionElemeStorepromotionReviewbwcDetailGetReviewBwcActivity) *AlibabaAlscUnionElemeStorepromotionReviewbwcDetailGetReviewBwcStorePromotionDto {
    s.Activities = &v
    return s
}
