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
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// Linger please
var (
	_ context.Context
)

type LocalApiService service

/*
LocalApiService 创建广告
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param data

@return LocalAddResponse
*/
func (a *LocalApiService) Add(ctx context.Context, data LocalAddRequest) (LocalAddResponseData, http.Header, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Post")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarFileKey     string
		localVarReturnValue LocalAddResponseData
		localVarResponse    LocalAddResponse
	)

	// create path and map variables
	localVarPath := a.client.Cfg.BasePath + "/local/add"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json", "application/xml"}

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
	// body params
	localVarPostBody = &data
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
			var v LocalAddResponse
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

/*
LocalApiService 拉取广告列表
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param data

@return LocalGetResponse
*/
func (a *LocalApiService) Get(ctx context.Context, data LocalGetRequest) (LocalGetResponseData, http.Header, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Post")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarFileKey     string
		localVarReturnValue LocalGetResponseData
		localVarResponse    LocalGetResponse
	)

	// create path and map variables
	localVarPath := a.client.Cfg.BasePath + "/local/get"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json", "application/xml"}

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
	// body params
	localVarPostBody = &data
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
			var v LocalGetResponse
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

/*
LocalApiService 更新广告
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param data

@return LocalUpdateResponse
*/
func (a *LocalApiService) Update(ctx context.Context, data LocalUpdateRequest) (LocalUpdateResponseData, http.Header, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Post")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarFileKey     string
		localVarReturnValue LocalUpdateResponseData
		localVarResponse    LocalUpdateResponse
	)

	// create path and map variables
	localVarPath := a.client.Cfg.BasePath + "/local/update"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json", "application/xml"}

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
	// body params
	localVarPostBody = &data
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
			var v LocalUpdateResponse
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
