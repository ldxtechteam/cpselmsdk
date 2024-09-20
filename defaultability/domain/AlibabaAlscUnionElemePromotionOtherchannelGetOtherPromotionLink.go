package domain


type AlibabaAlscUnionElemePromotionOtherchannelGetOtherPromotionLink struct {
    /*
        抖音内部跳转shcema地址，30天有效     */
    DySchema  *string `json:"dy_schema,omitempty" `

    /*
        抖音外部唤起链接，30天有效     */
    DyLink  *string `json:"dy_link,omitempty" `

}

func (s *AlibabaAlscUnionElemePromotionOtherchannelGetOtherPromotionLink) SetDySchema(v string) *AlibabaAlscUnionElemePromotionOtherchannelGetOtherPromotionLink {
    s.DySchema = &v
    return s
}
func (s *AlibabaAlscUnionElemePromotionOtherchannelGetOtherPromotionLink) SetDyLink(v string) *AlibabaAlscUnionElemePromotionOtherchannelGetOtherPromotionLink {
    s.DyLink = &v
    return s
}
