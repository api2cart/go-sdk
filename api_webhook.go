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


// WebhookAPIService WebhookAPI service
type WebhookAPIService service

type ApiWebhookCountRequest struct {
	ctx context.Context
	ApiService *WebhookAPIService
	entity *string
	action *string
	active *bool
}

// The entity you want to filter webhooks by (e.g. order or product)
func (r ApiWebhookCountRequest) Entity(entity string) ApiWebhookCountRequest {
	r.entity = &entity
	return r
}

// The action you want to filter webhooks by (e.g. order or product)
func (r ApiWebhookCountRequest) Action(action string) ApiWebhookCountRequest {
	r.action = &action
	return r
}

// The webhook status you want to filter webhooks by
func (r ApiWebhookCountRequest) Active(active bool) ApiWebhookCountRequest {
	r.active = &active
	return r
}

func (r ApiWebhookCountRequest) Execute() (*WebhookCount200Response, *http.Response, error) {
	return r.ApiService.WebhookCountExecute(r)
}

/*
WebhookCount webhook.count

Count registered webhooks on the store.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiWebhookCountRequest
*/
func (a *WebhookAPIService) WebhookCount(ctx context.Context) ApiWebhookCountRequest {
	return ApiWebhookCountRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return WebhookCount200Response
func (a *WebhookAPIService) WebhookCountExecute(r ApiWebhookCountRequest) (*WebhookCount200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *WebhookCount200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WebhookAPIService.WebhookCount")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/webhook.count.json"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.entity != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "entity", r.entity, "form", "")
	}
	if r.action != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "action", r.action, "form", "")
	}
	if r.active != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "active", r.active, "form", "")
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

type ApiWebhookCreateRequest struct {
	ctx context.Context
	ApiService *WebhookAPIService
	entity *string
	action *string
	callback *string
	label *string
	fields *string
	active *bool
	storeId *string
}

// Specify the entity that you want to enable webhooks for (e.g product, order, customer, category)
func (r ApiWebhookCreateRequest) Entity(entity string) ApiWebhookCreateRequest {
	r.entity = &entity
	return r
}

// Specify what action (event) will trigger the webhook (e.g add, delete, or update)
func (r ApiWebhookCreateRequest) Action(action string) ApiWebhookCreateRequest {
	r.action = &action
	return r
}

// Callback url that returns shipping rates. It should be able to accept POST requests with json data.
func (r ApiWebhookCreateRequest) Callback(callback string) ApiWebhookCreateRequest {
	r.callback = &callback
	return r
}

// The name you give to the webhook
func (r ApiWebhookCreateRequest) Label(label string) ApiWebhookCreateRequest {
	r.label = &label
	return r
}

// Fields the webhook should send
func (r ApiWebhookCreateRequest) Fields(fields string) ApiWebhookCreateRequest {
	r.fields = &fields
	return r
}

// Webhook status
func (r ApiWebhookCreateRequest) Active(active bool) ApiWebhookCreateRequest {
	r.active = &active
	return r
}

// Defines store id where the webhook should be assigned
func (r ApiWebhookCreateRequest) StoreId(storeId string) ApiWebhookCreateRequest {
	r.storeId = &storeId
	return r
}

func (r ApiWebhookCreateRequest) Execute() (*BasketLiveShippingServiceCreate200Response, *http.Response, error) {
	return r.ApiService.WebhookCreateExecute(r)
}

/*
WebhookCreate webhook.create

Create webhook on the store and subscribe to it.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiWebhookCreateRequest
*/
func (a *WebhookAPIService) WebhookCreate(ctx context.Context) ApiWebhookCreateRequest {
	return ApiWebhookCreateRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return BasketLiveShippingServiceCreate200Response
func (a *WebhookAPIService) WebhookCreateExecute(r ApiWebhookCreateRequest) (*BasketLiveShippingServiceCreate200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *BasketLiveShippingServiceCreate200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WebhookAPIService.WebhookCreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/webhook.create.json"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.entity == nil {
		return localVarReturnValue, nil, reportError("entity is required and must be specified")
	}
	if r.action == nil {
		return localVarReturnValue, nil, reportError("action is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "entity", r.entity, "form", "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "action", r.action, "form", "")
	if r.callback != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "callback", r.callback, "form", "")
	}
	if r.label != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "label", r.label, "form", "")
	}
	if r.fields != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields", r.fields, "form", "")
	} else {
		var defaultValue string = "force_all"
		r.fields = &defaultValue
	}
	if r.active != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "active", r.active, "form", "")
	} else {
		var defaultValue bool = true
		r.active = &defaultValue
	}
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

type ApiWebhookDeleteRequest struct {
	ctx context.Context
	ApiService *WebhookAPIService
	id *string
}

// Webhook id
func (r ApiWebhookDeleteRequest) Id(id string) ApiWebhookDeleteRequest {
	r.id = &id
	return r
}

func (r ApiWebhookDeleteRequest) Execute() (*AttributeDelete200Response, *http.Response, error) {
	return r.ApiService.WebhookDeleteExecute(r)
}

/*
WebhookDelete webhook.delete

Delete registered webhook on the store.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiWebhookDeleteRequest
*/
func (a *WebhookAPIService) WebhookDelete(ctx context.Context) ApiWebhookDeleteRequest {
	return ApiWebhookDeleteRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return AttributeDelete200Response
func (a *WebhookAPIService) WebhookDeleteExecute(r ApiWebhookDeleteRequest) (*AttributeDelete200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AttributeDelete200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WebhookAPIService.WebhookDelete")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/webhook.delete.json"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.id == nil {
		return localVarReturnValue, nil, reportError("id is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "id", r.id, "form", "")
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

type ApiWebhookEventsRequest struct {
	ctx context.Context
	ApiService *WebhookAPIService
}

func (r ApiWebhookEventsRequest) Execute() (*WebhookEvents200Response, *http.Response, error) {
	return r.ApiService.WebhookEventsExecute(r)
}

/*
WebhookEvents webhook.events

List all Webhooks that are available on this store.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiWebhookEventsRequest
*/
func (a *WebhookAPIService) WebhookEvents(ctx context.Context) ApiWebhookEventsRequest {
	return ApiWebhookEventsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return WebhookEvents200Response
func (a *WebhookAPIService) WebhookEventsExecute(r ApiWebhookEventsRequest) (*WebhookEvents200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *WebhookEvents200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WebhookAPIService.WebhookEvents")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/webhook.events.json"

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

type ApiWebhookListRequest struct {
	ctx context.Context
	ApiService *WebhookAPIService
	params *string
	start *int32
	count *int32
	entity *string
	action *string
	active *bool
	ids *string
}

// Set this parameter in order to choose which entity fields you want to retrieve
func (r ApiWebhookListRequest) Params(params string) ApiWebhookListRequest {
	r.params = &params
	return r
}

// This parameter sets the number from which you want to get entities
func (r ApiWebhookListRequest) Start(start int32) ApiWebhookListRequest {
	r.start = &start
	return r
}

// This parameter sets the entity amount that has to be retrieved. Max allowed count&#x3D;250
func (r ApiWebhookListRequest) Count(count int32) ApiWebhookListRequest {
	r.count = &count
	return r
}

// The entity you want to filter webhooks by (e.g. order or product)
func (r ApiWebhookListRequest) Entity(entity string) ApiWebhookListRequest {
	r.entity = &entity
	return r
}

// The action you want to filter webhooks by (e.g. add, update, or delete)
func (r ApiWebhookListRequest) Action(action string) ApiWebhookListRequest {
	r.action = &action
	return r
}

// The webhook status you want to filter webhooks by
func (r ApiWebhookListRequest) Active(active bool) ApiWebhookListRequest {
	r.active = &active
	return r
}

// List of сomma-separated webhook ids
func (r ApiWebhookListRequest) Ids(ids string) ApiWebhookListRequest {
	r.ids = &ids
	return r
}

func (r ApiWebhookListRequest) Execute() (*WebhookList200Response, *http.Response, error) {
	return r.ApiService.WebhookListExecute(r)
}

/*
WebhookList webhook.list

List registered webhook on the store.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiWebhookListRequest
*/
func (a *WebhookAPIService) WebhookList(ctx context.Context) ApiWebhookListRequest {
	return ApiWebhookListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return WebhookList200Response
func (a *WebhookAPIService) WebhookListExecute(r ApiWebhookListRequest) (*WebhookList200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *WebhookList200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WebhookAPIService.WebhookList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/webhook.list.json"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.params != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "params", r.params, "form", "")
	} else {
		var defaultValue string = "id,entity,action,callback"
		r.params = &defaultValue
	}
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
	if r.entity != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "entity", r.entity, "form", "")
	}
	if r.action != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "action", r.action, "form", "")
	}
	if r.active != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "active", r.active, "form", "")
	}
	if r.ids != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "ids", r.ids, "form", "")
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

type ApiWebhookUpdateRequest struct {
	ctx context.Context
	ApiService *WebhookAPIService
	id *string
	callback *string
	label *string
	fields *string
	active *bool
}

// Webhook id
func (r ApiWebhookUpdateRequest) Id(id string) ApiWebhookUpdateRequest {
	r.id = &id
	return r
}

// Callback url that returns shipping rates. It should be able to accept POST requests with json data.
func (r ApiWebhookUpdateRequest) Callback(callback string) ApiWebhookUpdateRequest {
	r.callback = &callback
	return r
}

// The name you give to the webhook
func (r ApiWebhookUpdateRequest) Label(label string) ApiWebhookUpdateRequest {
	r.label = &label
	return r
}

// Fields the webhook should send
func (r ApiWebhookUpdateRequest) Fields(fields string) ApiWebhookUpdateRequest {
	r.fields = &fields
	return r
}

// Webhook status
func (r ApiWebhookUpdateRequest) Active(active bool) ApiWebhookUpdateRequest {
	r.active = &active
	return r
}

func (r ApiWebhookUpdateRequest) Execute() (*ProductImageUpdate200Response, *http.Response, error) {
	return r.ApiService.WebhookUpdateExecute(r)
}

/*
WebhookUpdate webhook.update

Update Webhooks parameters.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiWebhookUpdateRequest
*/
func (a *WebhookAPIService) WebhookUpdate(ctx context.Context) ApiWebhookUpdateRequest {
	return ApiWebhookUpdateRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ProductImageUpdate200Response
func (a *WebhookAPIService) WebhookUpdateExecute(r ApiWebhookUpdateRequest) (*ProductImageUpdate200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ProductImageUpdate200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WebhookAPIService.WebhookUpdate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/webhook.update.json"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.id == nil {
		return localVarReturnValue, nil, reportError("id is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "id", r.id, "form", "")
	if r.callback != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "callback", r.callback, "form", "")
	}
	if r.label != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "label", r.label, "form", "")
	}
	if r.fields != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields", r.fields, "form", "")
	}
	if r.active != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "active", r.active, "form", "")
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
