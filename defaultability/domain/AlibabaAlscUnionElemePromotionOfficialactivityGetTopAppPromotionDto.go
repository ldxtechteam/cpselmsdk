package domain


type AlibabaAlscUnionElemePromotionOfficialactivityGetTopAppPromotionDto struct {
    /*
        deeplink地址     */
    DeepLink  *string `json:"deep_link,omitempty" `

    /*
        饿口令     */
    ElemeWord  *string `json:"eleme_word,omitempty" `

    /*
        饿口令有效期     */
    WordValidDate  *string `json:"word_valid_date,omitempty" `

}

func (s *AlibabaAlscUnionElemePromotionOfficialactivityGetTopAppPromotionDto) SetDeepLink(v string) *AlibabaAlscUnionElemePromotionOfficialactivityGetTopAppPromotionDto {
    s.DeepLink = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionOfficialactivityGetTopAppPromotionDto) SetElemeWord(v string) *AlibabaAlscUnionElemePromotionOfficialactivityGetTopAppPromotionDto {
    s.ElemeWord = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionOfficialactivityGetTopAppPromotionDto) SetWordValidDate(v string) *AlibabaAlscUnionElemePromotionOfficialactivityGetTopAppPromotionDto {
    s.WordValidDate = &v
    return s
}
