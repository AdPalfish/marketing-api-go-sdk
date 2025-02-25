/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type DynamicCreativesUpdateRequest struct {
	DynamicCreativeId         *int64                     `json:"dynamic_creative_id,omitempty"`
	DynamicCreativeTemplateId *int64                     `json:"dynamic_creative_template_id,omitempty"`
	DynamicCreativeElements   *DynamicCreativeElements   `json:"dynamic_creative_elements,omitempty"`
	DeepLinkUrl               *string                    `json:"deep_link_url,omitempty"`
	ImpressionTrackingUrl     *string                    `json:"impression_tracking_url,omitempty"`
	ClickTrackingUrl          *string                    `json:"click_tracking_url,omitempty"`
	FeedsVideoCommentSwitch   *bool                      `json:"feeds_video_comment_switch,omitempty"`
	UnionMarketSwitch         *bool                      `json:"union_market_switch,omitempty"`
	ProfileId                 *int64                     `json:"profile_id,omitempty"`
	VideoEndPage              *VideoEndPageSpec          `json:"video_end_page,omitempty"`
	BarrageList               *[]BarrageListCreateStruct `json:"barrage_list,omitempty"`
	AccountId                 *int64                     `json:"account_id,omitempty"`
}
