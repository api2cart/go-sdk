/*
API2Cart OpenAPI

API2Cart

API version: 1.1
Contact: contact@api2cart.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)


// ReturnAPIService ReturnAPI service
type ReturnAPIService service

type ApiReturnActionListRequest struct {
	ctx context.Context
	ApiService *ReturnAPIService
}

func (r ApiReturnActionListRequest) Execute() (*ReturnActionList200Response, *http.Response, error) {
	return r.ApiService.ReturnActionListExecute(r)
}

/*
ReturnActionList return.action.list

Retrieve list of return actions

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiReturnActionListRequest
*/
func (a *ReturnAPIService) ReturnActionList(ctx context.Context) ApiReturnActionListRequest {
	return ApiReturnActionListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ReturnActionList200Response
func (a *ReturnAPIService) ReturnActionListExecute(r ApiReturnActionListRequest) (*ReturnActionList200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ReturnActionList200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReturnAPIService.ReturnActionList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/return.action.list.json"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["StoreKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["x-store-key"] = key
			}
		}
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["x-api-key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiReturnCountRequest struct {
	ctx context.Context
	ApiService *ReturnAPIService
	orderIds *string
	customerId *string
	storeId *string
	status *string
	returnType *string
	createdFrom *string
	createdTo *string
	modifiedFrom *string
	modifiedTo *string
	reportRequestId *string
	disableReportCache *bool
}

// Counts return requests specified by order ids
func (r ApiReturnCountRequest) OrderIds(orderIds string) ApiReturnCountRequest {
	r.orderIds = &orderIds
	return r
}

// Counts return requests quantity specified by customer id
func (r ApiReturnCountRequest) CustomerId(customerId string) ApiReturnCountRequest {
	r.customerId = &customerId
	return r
}

// Store Id
func (r ApiReturnCountRequest) StoreId(storeId string) ApiReturnCountRequest {
	r.storeId = &storeId
	return r
}

// Defines status
func (r ApiReturnCountRequest) Status(status string) ApiReturnCountRequest {
	r.status = &status
	return r
}

// Retrieves returns specified by return type
func (r ApiReturnCountRequest) ReturnType(returnType string) ApiReturnCountRequest {
	r.returnType = &returnType
	return r
}

// Retrieve entities from their creation date
func (r ApiReturnCountRequest) CreatedFrom(createdFrom string) ApiReturnCountRequest {
	r.createdFrom = &createdFrom
	return r
}

// Retrieve entities to their creation date
func (r ApiReturnCountRequest) CreatedTo(createdTo string) ApiReturnCountRequest {
	r.createdTo = &createdTo
	return r
}

// Retrieve entities from their modification date
func (r ApiReturnCountRequest) ModifiedFrom(modifiedFrom string) ApiReturnCountRequest {
	r.modifiedFrom = &modifiedFrom
	return r
}

// Retrieve entities to their modification date
func (r ApiReturnCountRequest) ModifiedTo(modifiedTo string) ApiReturnCountRequest {
	r.modifiedTo = &modifiedTo
	return r
}

// Report request id
func (r ApiReturnCountRequest) ReportRequestId(reportRequestId string) ApiReturnCountRequest {
	r.reportRequestId = &reportRequestId
	return r
}

// Disable report cache for current request
func (r ApiReturnCountRequest) DisableReportCache(disableReportCache bool) ApiReturnCountRequest {
	r.disableReportCache = &disableReportCache
	return r
}

func (r ApiReturnCountRequest) Execute() (*ReturnCount200Response, *http.Response, error) {
	return r.ApiService.ReturnCountExecute(r)
}

/*
ReturnCount return.count

Count returns in store

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiReturnCountRequest
*/
func (a *ReturnAPIService) ReturnCount(ctx context.Context) ApiReturnCountRequest {
	return ApiReturnCountRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ReturnCount200Response
func (a *ReturnAPIService) ReturnCountExecute(r ApiReturnCountRequest) (*ReturnCount200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ReturnCount200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReturnAPIService.ReturnCount")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/return.count.json"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.orderIds != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "order_ids", r.orderIds, "form", "")
	}
	if r.customerId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "customer_id", r.customerId, "form", "")
	}
	if r.storeId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "store_id", r.storeId, "form", "")
	}
	if r.status != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "status", r.status, "form", "")
	}
	if r.returnType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "return_type", r.returnType, "form", "")
	}
	if r.createdFrom != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "created_from", r.createdFrom, "form", "")
	}
	if r.createdTo != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "created_to", r.createdTo, "form", "")
	}
	if r.modifiedFrom != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "modified_from", r.modifiedFrom, "form", "")
	}
	if r.modifiedTo != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "modified_to", r.modifiedTo, "form", "")
	}
	if r.reportRequestId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "report_request_id", r.reportRequestId, "form", "")
	}
	if r.disableReportCache != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "disable_report_cache", r.disableReportCache, "form", "")
	} else {
		var defaultValue bool = false
		r.disableReportCache = &defaultValue
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["StoreKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["x-store-key"] = key
			}
		}
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["x-api-key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiReturnInfoRequest struct {
	ctx context.Context
	ApiService *ReturnAPIService
	id *string
	orderId *string
	storeId *string
	responseFields *string
	params *string
	exclude *string
}

// Entity id
func (r ApiReturnInfoRequest) Id(id string) ApiReturnInfoRequest {
	r.id = &id
	return r
}

// Defines the order id
func (r ApiReturnInfoRequest) OrderId(orderId string) ApiReturnInfoRequest {
	r.orderId = &orderId
	return r
}

// Store Id
func (r ApiReturnInfoRequest) StoreId(storeId string) ApiReturnInfoRequest {
	r.storeId = &storeId
	return r
}

// Set this parameter in order to choose which entity fields you want to retrieve
func (r ApiReturnInfoRequest) ResponseFields(responseFields string) ApiReturnInfoRequest {
	r.responseFields = &responseFields
	return r
}

// Set this parameter in order to choose which entity fields you want to retrieve
func (r ApiReturnInfoRequest) Params(params string) ApiReturnInfoRequest {
	r.params = &params
	return r
}

// Set this parameter in order to choose which entity fields you want to ignore. Works only if parameter &#x60;params&#x60; equal force_all
func (r ApiReturnInfoRequest) Exclude(exclude string) ApiReturnInfoRequest {
	r.exclude = &exclude
	return r
}

func (r ApiReturnInfoRequest) Execute() (*ReturnInfo200Response, *http.Response, error) {
	return r.ApiService.ReturnInfoExecute(r)
}

/*
ReturnInfo return.info

Retrieve return information.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiReturnInfoRequest
*/
func (a *ReturnAPIService) ReturnInfo(ctx context.Context) ApiReturnInfoRequest {
	return ApiReturnInfoRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ReturnInfo200Response
func (a *ReturnAPIService) ReturnInfoExecute(r ApiReturnInfoRequest) (*ReturnInfo200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ReturnInfo200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReturnAPIService.ReturnInfo")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/return.info.json"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.id == nil {
		return localVarReturnValue, nil, reportError("id is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "id", r.id, "form", "")
	if r.orderId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "order_id", r.orderId, "form", "")
	}
	if r.storeId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "store_id", r.storeId, "form", "")
	}
	if r.responseFields != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "response_fields", r.responseFields, "form", "")
	}
	if r.params != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "params", r.params, "form", "")
	} else {
		var defaultValue string = "id,order_products"
		r.params = &defaultValue
	}
	if r.exclude != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "exclude", r.exclude, "form", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["StoreKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["x-store-key"] = key
			}
		}
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["x-api-key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiReturnListRequest struct {
	ctx context.Context
	ApiService *ReturnAPIService
	start *int32
	count *int32
	pageCursor *string
	orderId *string
	orderIds *string
	customerId *string
	storeId *string
	status *string
	returnType *string
	createdFrom *string
	createdTo *string
	modifiedFrom *string
	modifiedTo *string
	responseFields *string
	params *string
	exclude *string
	reportRequestId *string
	disableReportCache *bool
}

// This parameter sets the number from which you want to get entities
func (r ApiReturnListRequest) Start(start int32) ApiReturnListRequest {
	r.start = &start
	return r
}

// This parameter sets the entity amount that has to be retrieved. Max allowed count&#x3D;250
func (r ApiReturnListRequest) Count(count int32) ApiReturnListRequest {
	r.count = &count
	return r
}

// Used to retrieve entities via cursor-based pagination (it can&#39;t be used with any other filtering parameter)
func (r ApiReturnListRequest) PageCursor(pageCursor string) ApiReturnListRequest {
	r.pageCursor = &pageCursor
	return r
}

// Defines the order id
func (r ApiReturnListRequest) OrderId(orderId string) ApiReturnListRequest {
	r.orderId = &orderId
	return r
}

// Retrieves return requests specified by order ids
func (r ApiReturnListRequest) OrderIds(orderIds string) ApiReturnListRequest {
	r.orderIds = &orderIds
	return r
}

// Retrieves return requests specified by customer id
func (r ApiReturnListRequest) CustomerId(customerId string) ApiReturnListRequest {
	r.customerId = &customerId
	return r
}

// Store Id
func (r ApiReturnListRequest) StoreId(storeId string) ApiReturnListRequest {
	r.storeId = &storeId
	return r
}

// Defines status
func (r ApiReturnListRequest) Status(status string) ApiReturnListRequest {
	r.status = &status
	return r
}

// Retrieves returns specified by return type
func (r ApiReturnListRequest) ReturnType(returnType string) ApiReturnListRequest {
	r.returnType = &returnType
	return r
}

// Retrieve entities from their creation date
func (r ApiReturnListRequest) CreatedFrom(createdFrom string) ApiReturnListRequest {
	r.createdFrom = &createdFrom
	return r
}

// Retrieve entities to their creation date
func (r ApiReturnListRequest) CreatedTo(createdTo string) ApiReturnListRequest {
	r.createdTo = &createdTo
	return r
}

// Retrieve entities from their modification date
func (r ApiReturnListRequest) ModifiedFrom(modifiedFrom string) ApiReturnListRequest {
	r.modifiedFrom = &modifiedFrom
	return r
}

// Retrieve entities to their modification date
func (r ApiReturnListRequest) ModifiedTo(modifiedTo string) ApiReturnListRequest {
	r.modifiedTo = &modifiedTo
	return r
}

// Set this parameter in order to choose which entity fields you want to retrieve
func (r ApiReturnListRequest) ResponseFields(responseFields string) ApiReturnListRequest {
	r.responseFields = &responseFields
	return r
}

// Set this parameter in order to choose which entity fields you want to retrieve
func (r ApiReturnListRequest) Params(params string) ApiReturnListRequest {
	r.params = &params
	return r
}

// Set this parameter in order to choose which entity fields you want to ignore. Works only if parameter &#x60;params&#x60; equal force_all
func (r ApiReturnListRequest) Exclude(exclude string) ApiReturnListRequest {
	r.exclude = &exclude
	return r
}

// Report request id
func (r ApiReturnListRequest) ReportRequestId(reportRequestId string) ApiReturnListRequest {
	r.reportRequestId = &reportRequestId
	return r
}

// Disable report cache for current request
func (r ApiReturnListRequest) DisableReportCache(disableReportCache bool) ApiReturnListRequest {
	r.disableReportCache = &disableReportCache
	return r
}

func (r ApiReturnListRequest) Execute() (*ModelResponseReturnList, *http.Response, error) {
	return r.ApiService.ReturnListExecute(r)
}

/*
ReturnList return.list

Get list of return requests from store.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiReturnListRequest
*/
func (a *ReturnAPIService) ReturnList(ctx context.Context) ApiReturnListRequest {
	return ApiReturnListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ModelResponseReturnList
func (a *ReturnAPIService) ReturnListExecute(r ApiReturnListRequest) (*ModelResponseReturnList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ModelResponseReturnList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReturnAPIService.ReturnList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/return.list.json"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.start != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "start", r.start, "form", "")
	} else {
		var defaultValue int32 = 0
		r.start = &defaultValue
	}
	if r.count != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "count", r.count, "form", "")
	} else {
		var defaultValue int32 = 10
		r.count = &defaultValue
	}
	if r.pageCursor != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page_cursor", r.pageCursor, "form", "")
	}
	if r.orderId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "order_id", r.orderId, "form", "")
	}
	if r.orderIds != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "order_ids", r.orderIds, "form", "")
	}
	if r.customerId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "customer_id", r.customerId, "form", "")
	}
	if r.storeId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "store_id", r.storeId, "form", "")
	}
	if r.status != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "status", r.status, "form", "")
	}
	if r.returnType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "return_type", r.returnType, "form", "")
	}
	if r.createdFrom != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "created_from", r.createdFrom, "form", "")
	}
	if r.createdTo != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "created_to", r.createdTo, "form", "")
	}
	if r.modifiedFrom != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "modified_from", r.modifiedFrom, "form", "")
	}
	if r.modifiedTo != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "modified_to", r.modifiedTo, "form", "")
	}
	if r.responseFields != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "response_fields", r.responseFields, "form", "")
	}
	if r.params != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "params", r.params, "form", "")
	} else {
		var defaultValue string = "id,order_products"
		r.params = &defaultValue
	}
	if r.exclude != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "exclude", r.exclude, "form", "")
	}
	if r.reportRequestId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "report_request_id", r.reportRequestId, "form", "")
	}
	if r.disableReportCache != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "disable_report_cache", r.disableReportCache, "form", "")
	} else {
		var defaultValue bool = false
		r.disableReportCache = &defaultValue
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["StoreKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["x-store-key"] = key
			}
		}
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["x-api-key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiReturnReasonListRequest struct {
	ctx context.Context
	ApiService *ReturnAPIService
	storeId *string
}

// Store Id
func (r ApiReturnReasonListRequest) StoreId(storeId string) ApiReturnReasonListRequest {
	r.storeId = &storeId
	return r
}

func (r ApiReturnReasonListRequest) Execute() (*ReturnReasonList200Response, *http.Response, error) {
	return r.ApiService.ReturnReasonListExecute(r)
}

/*
ReturnReasonList return.reason.list

Retrieve list of return reasons

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiReturnReasonListRequest
*/
func (a *ReturnAPIService) ReturnReasonList(ctx context.Context) ApiReturnReasonListRequest {
	return ApiReturnReasonListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ReturnReasonList200Response
func (a *ReturnAPIService) ReturnReasonListExecute(r ApiReturnReasonListRequest) (*ReturnReasonList200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ReturnReasonList200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReturnAPIService.ReturnReasonList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/return.reason.list.json"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.storeId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "store_id", r.storeId, "form", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["StoreKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["x-store-key"] = key
			}
		}
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["x-api-key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiReturnStatusListRequest struct {
	ctx context.Context
	ApiService *ReturnAPIService
}

func (r ApiReturnStatusListRequest) Execute() (*ReturnStatusList200Response, *http.Response, error) {
	return r.ApiService.ReturnStatusListExecute(r)
}

/*
ReturnStatusList return.status.list

Retrieve list of statuses

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiReturnStatusListRequest
*/
func (a *ReturnAPIService) ReturnStatusList(ctx context.Context) ApiReturnStatusListRequest {
	return ApiReturnStatusListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ReturnStatusList200Response
func (a *ReturnAPIService) ReturnStatusListExecute(r ApiReturnStatusListRequest) (*ReturnStatusList200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ReturnStatusList200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReturnAPIService.ReturnStatusList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/return.status.list.json"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["StoreKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["x-store-key"] = key
			}
		}
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["x-api-key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
