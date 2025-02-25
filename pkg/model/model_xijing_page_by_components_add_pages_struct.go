/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 落地页配置结构
type XijingPageByComponentsAddPagesStruct struct {
	PageType          ComponentsPageType `json:"page_type,omitempty"`
	PageName          *string            `json:"page_name,omitempty"`
	PageTitle         *string            `json:"page_title,omitempty"`
	Clipboard         *string            `json:"clipboard,omitempty"`
	MobileAppId       *string            `json:"mobile_app_id,omitempty"`
	BgColor           *string            `json:"bg_color,omitempty"`
	BgImageId         *string            `json:"bg_image_id,omitempty"`
	ComponentSpecList *[]string          `json:"component_spec_list,omitempty"`
}
