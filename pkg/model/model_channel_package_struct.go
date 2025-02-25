/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 渠道包信息
type ChannelPackageStruct struct {
	AppAndroidChannelPackageId *string               `json:"app_android_channel_package_id,omitempty"`
	PackageName                *string               `json:"package_name,omitempty"`
	SystemStatus               UnionPackageSysStatus `json:"system_status,omitempty"`
	PackageOriginUrl           *string               `json:"package_origin_url,omitempty"`
}
