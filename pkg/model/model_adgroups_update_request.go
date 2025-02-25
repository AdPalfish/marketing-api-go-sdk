/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type AdgroupsUpdateRequest struct {
	AdgroupId                  *int64                           `json:"adgroup_id,omitempty"`
	AdgroupName                *string                          `json:"adgroup_name,omitempty"`
	BeginDate                  *string                          `json:"begin_date,omitempty"`
	FirstDayBeginTime          *string                          `json:"first_day_begin_time,omitempty"`
	EndDate                    *string                          `json:"end_date,omitempty"`
	BidAmount                  *int64                           `json:"bid_amount,omitempty"`
	ConversionId               *int64                           `json:"conversion_id,omitempty"`
	OptimizationGoal           OptimizationGoal                 `json:"optimization_goal,omitempty"`
	TimeSeries                 *string                          `json:"time_series,omitempty"`
	DailyBudget                *int64                           `json:"daily_budget,omitempty"`
	AppAndroidChannelPackageId *string                          `json:"app_android_channel_package_id,omitempty"`
	TargetingId                *int64                           `json:"targeting_id,omitempty"`
	Targeting                  *WriteTargetingSettingForAdgroup `json:"targeting,omitempty"`
	SceneSpec                  *SceneTargetingForWrite          `json:"scene_spec,omitempty"`
	ConfiguredStatus           AdStatus                         `json:"configured_status,omitempty"`
	CustomizedCategory         *string                          `json:"customized_category,omitempty"`
	AdditionalUserActionSets   *[]UserActionSetStruct           `json:"additional_user_action_sets,omitempty"`
	BidStrategy                BidStrategy                      `json:"bid_strategy,omitempty"`
	ColdStartAudience          *[]int64                         `json:"cold_start_audience,omitempty"`
	AutoAudience               *bool                            `json:"auto_audience,omitempty"`
	ExpandEnabled              *bool                            `json:"expand_enabled,omitempty"`
	ExpandTargeting            *[]string                        `json:"expand_targeting,omitempty"`
	DeepConversionSpec         *DeepConversionSpec              `json:"deep_conversion_spec,omitempty"`
	PoiList                    *[]string                        `json:"poi_list,omitempty"`
	DeepConversionBehaviorBid  *int64                           `json:"deep_conversion_behavior_bid,omitempty"`
	DeepConversionWorthRate    *float64                         `json:"deep_conversion_worth_rate,omitempty"`
	BidMode                    BidMode                          `json:"bid_mode,omitempty"`
	BidAdjustment              *BidAdjustment                   `json:"bid_adjustment,omitempty"`
	AutoAcquisitionEnabled     *bool                            `json:"auto_acquisition_enabled,omitempty"`
	AutoAcquisitionBudget      *int64                           `json:"auto_acquisition_budget,omitempty"`
	AutoDerivedCreativeEnabled *bool                            `json:"auto_derived_creative_enabled,omitempty"`
	UserActionSets             *[]UserActionSetStruct           `json:"user_action_sets,omitempty"`
	DynamicAdSpec              *DynamicAdSpec                   `json:"dynamic_ad_spec,omitempty"`
	AccountId                  *int64                           `json:"account_id,omitempty"`
}
