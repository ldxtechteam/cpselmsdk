package defaultability

import (
	"errors"
	"github.com/ldxtechteam/cpselmsdk"
	"github.com/ldxtechteam/cpselmsdk/defaultability/request"
	"github.com/ldxtechteam/cpselmsdk/defaultability/response"
	"github.com/ldxtechteam/cpselmsdk/util"
	"log"
)

type Defaultability struct {
	Client *topsdk.TopClient
}

func NewDefaultability(client *topsdk.TopClient) *Defaultability {
	return &Defaultability{client}
}

/*
本地生活媒体平台口碑选品筛选项集合
*/
func (ability *Defaultability) AlibabaAlscUnionKbItemPromotionFilterList(req *request.AlibabaAlscUnionKbItemPromotionFilterListRequest) (*response.AlibabaAlscUnionKbItemPromotionFilterListResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alsc.union.kb.item.promotion.filter.list", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlscUnionKbItemPromotionFilterListResponse{}
	if err != nil {
		log.Println("alibabaAlscUnionKbItemPromotionFilterList error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
本地生活媒体推广位查询
*/
func (ability *Defaultability) AlibabaAlscUnionMediaZoneGet(req *request.AlibabaAlscUnionMediaZoneGetRequest) (*response.AlibabaAlscUnionMediaZoneGetResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alsc.union.media.zone.get", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlscUnionMediaZoneGetResponse{}
	if err != nil {
		log.Println("alibabaAlscUnionMediaZoneGet error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
本地生活媒体平台口碑选品
*/
func (ability *Defaultability) AlibabaAlscUnionKbItemPromotion(req *request.AlibabaAlscUnionKbItemPromotionRequest) (*response.AlibabaAlscUnionKbItemPromotionResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alsc.union.kb.item.promotion", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlscUnionKbItemPromotionResponse{}
	if err != nil {
		log.Println("alibabaAlscUnionKbItemPromotion error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
本地联盟饿了么归因排查工具
*/
func (ability *Defaultability) AlibabaAlscUnionElemeToolOrderAttrbuteCheck(req *request.AlibabaAlscUnionElemeToolOrderAttrbuteCheckRequest) (*response.AlibabaAlscUnionElemeToolOrderAttrbuteCheckResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alsc.union.eleme.tool.order.attrbute.check", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlscUnionElemeToolOrderAttrbuteCheckResponse{}
	if err != nil {
		log.Println("alibabaAlscUnionElemeToolOrderAttrbuteCheck error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
本地生活媒体推广位创建
*/
func (ability *Defaultability) AlibabaAlscUnionMediaZoneAdd(req *request.AlibabaAlscUnionMediaZoneAddRequest) (*response.AlibabaAlscUnionMediaZoneAddResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alsc.union.media.zone.add", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlscUnionMediaZoneAddResponse{}
	if err != nil {
		log.Println("alibabaAlscUnionMediaZoneAdd error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
本地联盟饿了么单品推广门店列表（尚未开放）
*/
func (ability *Defaultability) AlibabaAlscUnionElemePromotionItempromotionStoreQuery(req *request.AlibabaAlscUnionElemePromotionItempromotionStoreQueryRequest) (*response.AlibabaAlscUnionElemePromotionItempromotionStoreQueryResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alsc.union.eleme.promotion.itempromotion.store.query", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlscUnionElemePromotionItempromotionStoreQueryResponse{}
	if err != nil {
		log.Println("alibabaAlscUnionElemePromotionItempromotionStoreQuery error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
本地生活媒体推广口碑CPA用户反作弊订单数据报表
*/
func (ability *Defaultability) AlibabaAlscUnionKbcpaPunishOrderGet(req *request.AlibabaAlscUnionKbcpaPunishOrderGetRequest) (*response.AlibabaAlscUnionKbcpaPunishOrderGetResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alsc.union.kbcpa.punish.order.get", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlscUnionKbcpaPunishOrderGetResponse{}
	if err != nil {
		log.Println("alibabaAlscUnionKbcpaPunishOrderGet error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
本地联盟媒体出资活动红包发放
*/
func (ability *Defaultability) AlibabaAlscUnionElemeMediaActivityCouponSend(req *request.AlibabaAlscUnionElemeMediaActivityCouponSendRequest) (*response.AlibabaAlscUnionElemeMediaActivityCouponSendResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alsc.union.eleme.media.activity.coupon.send", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlscUnionElemeMediaActivityCouponSendResponse{}
	if err != nil {
		log.Println("alibabaAlscUnionElemeMediaActivityCouponSend error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
获取商家所在分组及其已授权(广播)消息topics
*/
func (ability *Defaultability) TaobaoTmcUserGet(req *request.TaobaoTmcUserGetRequest) (*response.TaobaoTmcUserGetResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("taobao.tmc.user.get", req.ToMap(), req.ToFileMap())
	var respStruct = response.TaobaoTmcUserGetResponse{}
	if err != nil {
		log.Println("taobaoTmcUserGet error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
本地联盟饿了么推广其他推广链接获取，抖音等
*/
func (ability *Defaultability) AlibabaAlscUnionElemePromotionOtherchannelGet(req *request.AlibabaAlscUnionElemePromotionOtherchannelGetRequest) (*response.AlibabaAlscUnionElemePromotionOtherchannelGetResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alsc.union.eleme.promotion.otherchannel.get", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlscUnionElemePromotionOtherchannelGetResponse{}
	if err != nil {
		log.Println("alibabaAlscUnionElemePromotionOtherchannelGet error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
本地联盟饿了么单店推广店铺列表
*/
func (ability *Defaultability) AlibabaAlscUnionElemePromotionStorepromotionQuery(req *request.AlibabaAlscUnionElemePromotionStorepromotionQueryRequest) (*response.AlibabaAlscUnionElemePromotionStorepromotionQueryResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alsc.union.eleme.promotion.storepromotion.query", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlscUnionElemePromotionStorepromotionQueryResponse{}
	if err != nil {
		log.Println("alibabaAlscUnionElemePromotionStorepromotionQuery error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
本地联盟饿了么单店推广单店铺查询
*/
func (ability *Defaultability) AlibabaAlscUnionElemePromotionStorepromotionGet(req *request.AlibabaAlscUnionElemePromotionStorepromotionGetRequest) (*response.AlibabaAlscUnionElemePromotionStorepromotionGetResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alsc.union.eleme.promotion.storepromotion.get", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlscUnionElemePromotionStorepromotionGetResponse{}
	if err != nil {
		log.Println("alibabaAlscUnionElemePromotionStorepromotionGet error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
本地联盟饿了么推广官方活动查询
*/
func (ability *Defaultability) AlibabaAlscUnionElemePromotionOfficialactivityGet(req *request.AlibabaAlscUnionElemePromotionOfficialactivityGetRequest) (*response.AlibabaAlscUnionElemePromotionOfficialactivityGetResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alsc.union.eleme.promotion.officialactivity.get", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlscUnionElemePromotionOfficialactivityGetResponse{}
	if err != nil {
		log.Println("alibabaAlscUnionElemePromotionOfficialactivityGet error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
本地联盟饿了么单店推广批量店铺查询
*/
func (ability *Defaultability) AlibabaAlscUnionElemePromotionStorepromotionBatchGet(req *request.AlibabaAlscUnionElemePromotionStorepromotionBatchGetRequest) (*response.AlibabaAlscUnionElemePromotionStorepromotionBatchGetResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alsc.union.eleme.promotion.storepromotion.batch.get", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlscUnionElemePromotionStorepromotionBatchGetResponse{}
	if err != nil {
		log.Println("alibabaAlscUnionElemePromotionStorepromotionBatchGet error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
本地生活媒体推广口碑CPA用户维权订单数据报表
*/
func (ability *Defaultability) AlibabaAlscUnionKbcpaRefundOrderGet(req *request.AlibabaAlscUnionKbcpaRefundOrderGetRequest) (*response.AlibabaAlscUnionKbcpaRefundOrderGetResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alsc.union.kbcpa.refund.order.get", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlscUnionKbcpaRefundOrderGetResponse{}
	if err != nil {
		log.Println("alibabaAlscUnionKbcpaRefundOrderGet error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
本地联盟口碑商户列表
*/
func (ability *Defaultability) AlibabaAlscUnionKbStoreQuery(req *request.AlibabaAlscUnionKbStoreQueryRequest) (*response.AlibabaAlscUnionKbStoreQueryResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alsc.union.kb.store.query", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlscUnionKbStoreQueryResponse{}
	if err != nil {
		log.Println("alibabaAlscUnionKbStoreQuery error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
本地联盟口碑门店商品列表
*/
func (ability *Defaultability) AlibabaAlscUnionKbStoreItemQuery(req *request.AlibabaAlscUnionKbStoreItemQueryRequest) (*response.AlibabaAlscUnionKbStoreItemQueryResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alsc.union.kb.store.item.query", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlscUnionKbStoreItemQueryResponse{}
	if err != nil {
		log.Println("alibabaAlscUnionKbStoreItemQuery error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
本地联盟饿了么单品推广详情（餐饮）
*/
func (ability *Defaultability) AlibabaAlscUnionElemePromotionItempromotionGet(req *request.AlibabaAlscUnionElemePromotionItempromotionGetRequest) (*response.AlibabaAlscUnionElemePromotionItempromotionGetResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alsc.union.eleme.promotion.itempromotion.get", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlscUnionElemePromotionItempromotionGetResponse{}
	if err != nil {
		log.Println("alibabaAlscUnionElemePromotionItempromotionGet error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
本地联盟饿了么单品推广列表（餐饮）
*/
func (ability *Defaultability) AlibabaAlscUnionElemePromotionItempromotionQuery(req *request.AlibabaAlscUnionElemePromotionItempromotionQueryRequest) (*response.AlibabaAlscUnionElemePromotionItempromotionQueryResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alsc.union.eleme.promotion.itempromotion.query", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlscUnionElemePromotionItempromotionQueryResponse{}
	if err != nil {
		log.Println("alibabaAlscUnionElemePromotionItempromotionQuery error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
本地生活媒体推广订单明细查询
*/
func (ability *Defaultability) AlibabaAlscUnionKbcpaOrderDetailsGet(req *request.AlibabaAlscUnionKbcpaOrderDetailsGetRequest) (*response.AlibabaAlscUnionKbcpaOrderDetailsGetResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alsc.union.kbcpa.order.details.get", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlscUnionKbcpaOrderDetailsGetResponse{}
	if err != nil {
		log.Println("alibabaAlscUnionKbcpaOrderDetailsGet error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
本地联盟口碑商品详情
*/
func (ability *Defaultability) AlibabaAlscUnionKbItemDetailGet(req *request.AlibabaAlscUnionKbItemDetailGetRequest) (*response.AlibabaAlscUnionKbItemDetailGetResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alsc.union.kb.item.detail.get", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlscUnionKbItemDetailGetResponse{}
	if err != nil {
		log.Println("alibabaAlscUnionKbItemDetailGet error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
本地联盟口碑商品门店关系
*/
func (ability *Defaultability) AlibabaAlscUnionKbItemStoreRelationQuery(req *request.AlibabaAlscUnionKbItemStoreRelationQueryRequest) (*response.AlibabaAlscUnionKbItemStoreRelationQueryResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alsc.union.kb.item.store.relation.query", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlscUnionKbItemStoreRelationQueryResponse{}
	if err != nil {
		log.Println("alibabaAlscUnionKbItemStoreRelationQuery error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
本地生活媒体推广订单明细报表查询
*/
func (ability *Defaultability) AlibabaAlscUnionKbcpxPositiveOrderGet(req *request.AlibabaAlscUnionKbcpxPositiveOrderGetRequest) (*response.AlibabaAlscUnionKbcpxPositiveOrderGetResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alsc.union.kbcpx.positive.order.get", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlscUnionKbcpxPositiveOrderGetResponse{}
	if err != nil {
		log.Println("alibabaAlscUnionKbcpxPositiveOrderGet error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
本地联盟口碑门店详情
*/
func (ability *Defaultability) AlibabaAlscUnionKbItemStoreDetailGet(req *request.AlibabaAlscUnionKbItemStoreDetailGetRequest) (*response.AlibabaAlscUnionKbItemStoreDetailGetResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alsc.union.kb.item.store.detail.get", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlscUnionKbItemStoreDetailGetResponse{}
	if err != nil {
		log.Println("alibabaAlscUnionKbItemStoreDetailGet error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
本地生活媒体推广用户维权订单数据报表
*/
func (ability *Defaultability) AlibabaAlscUnionKbcpxRefundOrderGet(req *request.AlibabaAlscUnionKbcpxRefundOrderGetRequest) (*response.AlibabaAlscUnionKbcpxRefundOrderGetResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alsc.union.kbcpx.refund.order.get", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlscUnionKbcpxRefundOrderGetResponse{}
	if err != nil {
		log.Println("alibabaAlscUnionKbcpxRefundOrderGet error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
本地生活媒体推广用户反作弊订单数据报表
*/
func (ability *Defaultability) AlibabaAlscUnionKbcpxPunishOrderGet(req *request.AlibabaAlscUnionKbcpxPunishOrderGetRequest) (*response.AlibabaAlscUnionKbcpxPunishOrderGetResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alsc.union.kbcpx.punish.order.get", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlscUnionKbcpxPunishOrderGetResponse{}
	if err != nil {
		log.Println("alibabaAlscUnionKbcpxPunishOrderGet error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
本地联盟口碑商品列表
*/
func (ability *Defaultability) AlibabaAlscUnionKbItemQuery(req *request.AlibabaAlscUnionKbItemQueryRequest) (*response.AlibabaAlscUnionKbItemQueryResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alsc.union.kb.item.query", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlscUnionKbItemQueryResponse{}
	if err != nil {
		log.Println("alibabaAlscUnionKbItemQuery error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
本地联盟饿了么评价有礼库存锁定
*/
func (ability *Defaultability) AlibabaAlscUnionElemeStorepromotionReviewbwcStockLock(req *request.AlibabaAlscUnionElemeStorepromotionReviewbwcStockLockRequest) (*response.AlibabaAlscUnionElemeStorepromotionReviewbwcStockLockResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alsc.union.eleme.storepromotion.reviewbwc.stock.lock", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlscUnionElemeStorepromotionReviewbwcStockLockResponse{}
	if err != nil {
		log.Println("alibabaAlscUnionElemeStorepromotionReviewbwcStockLock error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
本地联盟饿了么评价有礼锁定库存释放
*/
func (ability *Defaultability) AlibabaAlscUnionElemeStorepromotionReviewbwcStockRelease(req *request.AlibabaAlscUnionElemeStorepromotionReviewbwcStockReleaseRequest) (*response.AlibabaAlscUnionElemeStorepromotionReviewbwcStockReleaseResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alsc.union.eleme.storepromotion.reviewbwc.stock.release", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlscUnionElemeStorepromotionReviewbwcStockReleaseResponse{}
	if err != nil {
		log.Println("alibabaAlscUnionElemeStorepromotionReviewbwcStockRelease error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
本地联盟饿了么评价有礼店铺列表
*/
func (ability *Defaultability) AlibabaAlscUnionElemeStorepromotionReviewbwcQuery(req *request.AlibabaAlscUnionElemeStorepromotionReviewbwcQueryRequest) (*response.AlibabaAlscUnionElemeStorepromotionReviewbwcQueryResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alsc.union.eleme.storepromotion.reviewbwc.query", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlscUnionElemeStorepromotionReviewbwcQueryResponse{}
	if err != nil {
		log.Println("alibabaAlscUnionElemeStorepromotionReviewbwcQuery error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
本地联盟饿了么评价有礼店铺详情
*/
func (ability *Defaultability) AlibabaAlscUnionElemeStorepromotionReviewbwcDetailGet(req *request.AlibabaAlscUnionElemeStorepromotionReviewbwcDetailGetRequest) (*response.AlibabaAlscUnionElemeStorepromotionReviewbwcDetailGetResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alsc.union.eleme.storepromotion.reviewbwc.detail.get", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlscUnionElemeStorepromotionReviewbwcDetailGetResponse{}
	if err != nil {
		log.Println("alibabaAlscUnionElemeStorepromotionReviewbwcDetailGet error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
发布单条消息
*/
func (ability *Defaultability) TaobaoTmcMessageProduce(req *request.TaobaoTmcMessageProduceRequest) (*response.TaobaoTmcMessageProduceResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("taobao.tmc.message.produce", req.ToMap(), req.ToFileMap())
	var respStruct = response.TaobaoTmcMessageProduceResponse{}
	if err != nil {
		log.Println("taobaoTmcMessageProduce error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
本地生活媒体创建商品推广链接
*/
func (ability *Defaultability) AlibabaAlscUnionKbItemPromotionShareCreate(req *request.AlibabaAlscUnionKbItemPromotionShareCreateRequest) (*response.AlibabaAlscUnionKbItemPromotionShareCreateResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alsc.union.kb.item.promotion.share.create", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlscUnionKbItemPromotionShareCreateResponse{}
	if err != nil {
		log.Println("alibabaAlscUnionKbItemPromotionShareCreate error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
取消用户的消息服务
*/
func (ability *Defaultability) TaobaoTmcUserCancel(req *request.TaobaoTmcUserCancelRequest) (*response.TaobaoTmcUserCancelResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("taobao.tmc.user.cancel", req.ToMap(), req.ToFileMap())
	var respStruct = response.TaobaoTmcUserCancelResponse{}
	if err != nil {
		log.Println("taobaoTmcUserCancel error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
为已授权的用户开通消息服务
*/
func (ability *Defaultability) TaobaoTmcUserPermit(req *request.TaobaoTmcUserPermitRequest, session string) (*response.TaobaoTmcUserPermitResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.tmc.user.permit", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoTmcUserPermitResponse{}
	if err != nil {
		log.Println("taobaoTmcUserPermit error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
本地联盟推广链接推广对象解析
*/
func (ability *Defaultability) AlibabaAlscUnionPromotionLinkAnalyze(req *request.AlibabaAlscUnionPromotionLinkAnalyzeRequest) (*response.AlibabaAlscUnionPromotionLinkAnalyzeResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alsc.union.promotion.link.analyze", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlscUnionPromotionLinkAnalyzeResponse{}
	if err != nil {
		log.Println("alibabaAlscUnionPromotionLinkAnalyze error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
本地联盟饿了么评价有礼诊断
*/
func (ability *Defaultability) AlibabaAlscUnionElemeStorepromotionReviewbwcDiagnose(req *request.AlibabaAlscUnionElemeStorepromotionReviewbwcDiagnoseRequest) (*response.AlibabaAlscUnionElemeStorepromotionReviewbwcDiagnoseResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alsc.union.eleme.storepromotion.reviewbwc.diagnose", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlscUnionElemeStorepromotionReviewbwcDiagnoseResponse{}
	if err != nil {
		log.Println("alibabaAlscUnionElemeStorepromotionReviewbwcDiagnose error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
本地联盟饿了么评价有身份绑定
*/
func (ability *Defaultability) AlibabaAlscUnionElemeStorepromotionReviewbwcBindLinkGet(req *request.AlibabaAlscUnionElemeStorepromotionReviewbwcBindLinkGetRequest) (*response.AlibabaAlscUnionElemeStorepromotionReviewbwcBindLinkGetResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alsc.union.eleme.storepromotion.reviewbwc.bind.link.get", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlscUnionElemeStorepromotionReviewbwcBindLinkGetResponse{}
	if err != nil {
		log.Println("alibabaAlscUnionElemeStorepromotionReviewbwcBindLinkGet error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}
