/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 修改广告组日预算
type UpdateAdgroupDailyBudgetItem struct {
	AdgroupId   int64 `json:"adgroup_id,omitempty"`
	DailyBudget int64 `json:"daily_budget,omitempty"`
}
