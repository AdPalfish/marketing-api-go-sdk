/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type OuterCluesClaiminfoUpdateResponse struct {
	Code      *int64                                 `json:"code,omitempty"`
	Message   *string                                `json:"message,omitempty"`
	MessageCn *string                                `json:"message_cn,omitempty"`
	Errors    *[]ApiErrorStruct                      `json:"errors,omitempty"`
	Data      *OuterCluesClaiminfoUpdateResponseData `json:"data,omitempty"`
}
