/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// custom_audience返回结构
type CustomAudience struct {
	AudienceId         *int64        `json:"audience_id,omitempty"`
	AccountId          *int64        `json:"account_id,omitempty"`
	Name               *string       `json:"name,omitempty"`
	Description        *string       `json:"description,omitempty"`
	Type_              AudienceType  `json:"type,omitempty"`
	Status             ProcessStatus `json:"status,omitempty"`
	ErrorCode          *int64        `json:"error_code,omitempty"`
	UserCount          *int64        `json:"user_count,omitempty"`
	CreatedTime        *string       `json:"created_time,omitempty"`
	LastModifiedTime   *string       `json:"last_modified_time,omitempty"`
	AudienceSpec       *AudienceSpec `json:"audience_spec,omitempty"`
	ExternalAudienceId *string       `json:"external_audience_id,omitempty"`
}
