/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package api

import (
	"context"
	"github.com/AdPalfish/marketing-api-go-sdk/pkg/errors"
	. "github.com/AdPalfish/marketing-api-go-sdk/pkg/model"
	"github.com/antihax/optional"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// Linger please
var (
	_ context.Context
)

type XijingPageListApiService service

/*
XijingPageListApiService 蹊径-获取落地页列表
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param accountId
 * @param optional nil or *XijingPageListGetOpts - Optional Parameters:
     * @param "PageId" (optional.Int64) -
     * @param "PageServiceId" (optional.String) -
     * @param "PageName" (optional.String) -
     * @param "PageType" (optional.Interface of []string) -
     * @param "PageLastModifyStartTime" (optional.String) -
     * @param "PageLastModifyEndTime" (optional.String) -
     * @param "PageSize" (optional.Int64) -
     * @param "PageIndex" (optional.Int64) -
     * @param "PagePublishStatus" (optional.Interface of []string) -
     * @param "PageStatus" (optional.Interface of []string) -
     * @param "PageSource" (optional.String) -
     * @param "PageOwnerId" (optional.Int64) -
     * @param "AppId" (optional.Int64) -
     * @param "AppType" (optional.String) -
     * @param "QueryType" (optional.String) -
     * @param "Fields" (optional.Interface of []string) -  返回参数的字段列表

@return XijingPageListGetResponse
*/

type XijingPageListGetOpts struct {
	PageId                  optional.Int64
	PageServiceId           optional.String
	PageName                optional.String
	PageType                optional.Interface
	PageLastModifyStartTime optional.String
	PageLastModifyEndTime   optional.String
	PageSize                optional.Int64
	PageIndex               optional.Int64
	PagePublishStatus       optional.Interface
	PageStatus              optional.Interface
	PageSource              optional.String
	PageOwnerId             optional.Int64
	AppId                   optional.Int64
	AppType                 optional.String
	QueryType               optional.String
	Fields                  optional.Interface
}

func (a *XijingPageListApiService) Get(ctx context.Context, accountId int64, localVarOptionals *XijingPageListGetOpts) (XijingPageListGetResponseData, http.Header, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Get")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarFileKey     string
		localVarReturnValue XijingPageListGetResponseData
		localVarResponse    XijingPageListGetResponse
	)

	// create path and map variables
	localVarPath := a.client.Cfg.BasePath + "/xijing_page_list/get"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("account_id", parameterToString(accountId, ""))
	if localVarOptionals != nil && localVarOptionals.PageId.IsSet() {
		localVarQueryParams.Add("page_id", parameterToString(localVarOptionals.PageId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.PageServiceId.IsSet() {
		localVarQueryParams.Add("page_service_id", parameterToString(localVarOptionals.PageServiceId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.PageName.IsSet() {
		localVarQueryParams.Add("page_name", parameterToString(localVarOptionals.PageName.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.PageType.IsSet() {
		localVarQueryParams.Add("page_type", parameterToString(localVarOptionals.PageType.Value(), "multi"))
	}
	if localVarOptionals != nil && localVarOptionals.PageLastModifyStartTime.IsSet() {
		localVarQueryParams.Add("page_last_modify_start_time", parameterToString(localVarOptionals.PageLastModifyStartTime.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.PageLastModifyEndTime.IsSet() {
		localVarQueryParams.Add("page_last_modify_end_time", parameterToString(localVarOptionals.PageLastModifyEndTime.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.PageSize.IsSet() {
		localVarQueryParams.Add("page_size", parameterToString(localVarOptionals.PageSize.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.PageIndex.IsSet() {
		localVarQueryParams.Add("page_index", parameterToString(localVarOptionals.PageIndex.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.PagePublishStatus.IsSet() {
		localVarQueryParams.Add("page_publish_status", parameterToString(localVarOptionals.PagePublishStatus.Value(), "multi"))
	}
	if localVarOptionals != nil && localVarOptionals.PageStatus.IsSet() {
		localVarQueryParams.Add("page_status", parameterToString(localVarOptionals.PageStatus.Value(), "multi"))
	}
	if localVarOptionals != nil && localVarOptionals.PageSource.IsSet() {
		localVarQueryParams.Add("page_source", parameterToString(localVarOptionals.PageSource.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.PageOwnerId.IsSet() {
		localVarQueryParams.Add("page_owner_id", parameterToString(localVarOptionals.PageOwnerId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.AppId.IsSet() {
		localVarQueryParams.Add("app_id", parameterToString(localVarOptionals.AppId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.AppType.IsSet() {
		localVarQueryParams.Add("app_type", parameterToString(localVarOptionals.AppType.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.QueryType.IsSet() {
		localVarQueryParams.Add("query_type", parameterToString(localVarOptionals.QueryType.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Fields.IsSet() {
		localVarQueryParams.Add("fields", parameterToString(localVarOptionals.Fields.Value(), "multi"))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"text/plain"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes, localVarFileKey)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, nil, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	defer localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, nil, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarResponse, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err == nil {
			if *localVarResponse.Code != 0 {
				var localVarResponseErrors []ApiErrorStruct
				if localVarResponse.Errors != nil {
					localVarResponseErrors = *localVarResponse.Errors
				}
				err = errors.NewError(localVarResponse.Code, localVarResponse.Message, localVarResponse.MessageCn, localVarResponseErrors)
				return localVarReturnValue, localVarHttpResponse.Header, err
			}
			return *localVarResponse.Data, localVarHttpResponse.Header, err
		} else {
			return localVarReturnValue, localVarHttpResponse.Header, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}

		if localVarHttpResponse.StatusCode == 200 {
			var v XijingPageListGetResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse.Header, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse.Header, newErr
		}

		return localVarReturnValue, localVarHttpResponse.Header, newErr
	}

	return localVarReturnValue, localVarHttpResponse.Header, nil
}
