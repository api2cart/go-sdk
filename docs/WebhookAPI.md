# \WebhookAPI

All URIs are relative to *https://api.api2cart.local.com/v1.1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WebhookCount**](WebhookAPI.md#WebhookCount) | **Get** /webhook.count.json | webhook.count
[**WebhookCreate**](WebhookAPI.md#WebhookCreate) | **Post** /webhook.create.json | webhook.create
[**WebhookDelete**](WebhookAPI.md#WebhookDelete) | **Delete** /webhook.delete.json | webhook.delete
[**WebhookEvents**](WebhookAPI.md#WebhookEvents) | **Get** /webhook.events.json | webhook.events
[**WebhookList**](WebhookAPI.md#WebhookList) | **Get** /webhook.list.json | webhook.list
[**WebhookUpdate**](WebhookAPI.md#WebhookUpdate) | **Put** /webhook.update.json | webhook.update



## WebhookCount

> WebhookCount200Response WebhookCount(ctx).Entity(entity).Action(action).Active(active).Execute()

webhook.count



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	entity := "product" // string | The entity you want to filter webhooks by (e.g. order or product) (optional)
	action := "add" // string | The action you want to filter webhooks by (e.g. order or product) (optional)
	active := true // bool | The webhook status you want to filter webhooks by (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhookAPI.WebhookCount(context.Background()).Entity(entity).Action(action).Active(active).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookAPI.WebhookCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WebhookCount`: WebhookCount200Response
	fmt.Fprintf(os.Stdout, "Response from `WebhookAPI.WebhookCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWebhookCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **entity** | **string** | The entity you want to filter webhooks by (e.g. order or product) | 
 **action** | **string** | The action you want to filter webhooks by (e.g. order or product) | 
 **active** | **bool** | The webhook status you want to filter webhooks by | 

### Return type

[**WebhookCount200Response**](WebhookCount200Response.md)

### Authorization

[StoreKeyAuth](../README.md#StoreKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebhookCreate

> BasketLiveShippingServiceCreate200Response WebhookCreate(ctx).WebhookCreate(webhookCreate).Execute()

webhook.create



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	webhookCreate := *openapiclient.NewWebhookCreate("product", "add") // WebhookCreate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhookAPI.WebhookCreate(context.Background()).WebhookCreate(webhookCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookAPI.WebhookCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WebhookCreate`: BasketLiveShippingServiceCreate200Response
	fmt.Fprintf(os.Stdout, "Response from `WebhookAPI.WebhookCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWebhookCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookCreate** | [**WebhookCreate**](WebhookCreate.md) |  | 

### Return type

[**BasketLiveShippingServiceCreate200Response**](BasketLiveShippingServiceCreate200Response.md)

### Authorization

[StoreKeyAuth](../README.md#StoreKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebhookDelete

> AttributeDelete200Response WebhookDelete(ctx).Id(id).Execute()

webhook.delete



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	id := "25" // string | Webhook id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhookAPI.WebhookDelete(context.Background()).Id(id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookAPI.WebhookDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WebhookDelete`: AttributeDelete200Response
	fmt.Fprintf(os.Stdout, "Response from `WebhookAPI.WebhookDelete`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWebhookDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | Webhook id | 

### Return type

[**AttributeDelete200Response**](AttributeDelete200Response.md)

### Authorization

[StoreKeyAuth](../README.md#StoreKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebhookEvents

> WebhookEvents200Response WebhookEvents(ctx).Execute()

webhook.events



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhookAPI.WebhookEvents(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookAPI.WebhookEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WebhookEvents`: WebhookEvents200Response
	fmt.Fprintf(os.Stdout, "Response from `WebhookAPI.WebhookEvents`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiWebhookEventsRequest struct via the builder pattern


### Return type

[**WebhookEvents200Response**](WebhookEvents200Response.md)

### Authorization

[StoreKeyAuth](../README.md#StoreKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebhookList

> WebhookList200Response WebhookList(ctx).Start(start).Count(count).Entity(entity).Action(action).Active(active).Ids(ids).Params(params).Execute()

webhook.list



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	start := int32(0) // int32 | This parameter sets the number from which you want to get entities (optional) (default to 0)
	count := int32(20) // int32 | This parameter sets the entity amount that has to be retrieved. Max allowed count=250 (optional) (default to 10)
	entity := "product" // string | The entity you want to filter webhooks by (e.g. order or product) (optional)
	action := "add" // string | The action you want to filter webhooks by (e.g. add, update, or delete) (optional)
	active := true // bool | The webhook status you want to filter webhooks by (optional)
	ids := "3,14,25" // string | List of сomma-separated webhook ids (optional)
	params := "id,entity,callback,fields" // string | Set this parameter in order to choose which entity fields you want to retrieve (optional) (default to "id,entity,action,callback")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhookAPI.WebhookList(context.Background()).Start(start).Count(count).Entity(entity).Action(action).Active(active).Ids(ids).Params(params).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookAPI.WebhookList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WebhookList`: WebhookList200Response
	fmt.Fprintf(os.Stdout, "Response from `WebhookAPI.WebhookList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWebhookListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **start** | **int32** | This parameter sets the number from which you want to get entities | [default to 0]
 **count** | **int32** | This parameter sets the entity amount that has to be retrieved. Max allowed count&#x3D;250 | [default to 10]
 **entity** | **string** | The entity you want to filter webhooks by (e.g. order or product) | 
 **action** | **string** | The action you want to filter webhooks by (e.g. add, update, or delete) | 
 **active** | **bool** | The webhook status you want to filter webhooks by | 
 **ids** | **string** | List of сomma-separated webhook ids | 
 **params** | **string** | Set this parameter in order to choose which entity fields you want to retrieve | [default to &quot;id,entity,action,callback&quot;]

### Return type

[**WebhookList200Response**](WebhookList200Response.md)

### Authorization

[StoreKeyAuth](../README.md#StoreKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebhookUpdate

> ProductImageUpdate200Response WebhookUpdate(ctx).WebhookUpdate(webhookUpdate).Execute()

webhook.update



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	webhookUpdate := *openapiclient.NewWebhookUpdate("25") // WebhookUpdate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhookAPI.WebhookUpdate(context.Background()).WebhookUpdate(webhookUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookAPI.WebhookUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WebhookUpdate`: ProductImageUpdate200Response
	fmt.Fprintf(os.Stdout, "Response from `WebhookAPI.WebhookUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWebhookUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookUpdate** | [**WebhookUpdate**](WebhookUpdate.md) |  | 

### Return type

[**ProductImageUpdate200Response**](ProductImageUpdate200Response.md)

### Authorization

[StoreKeyAuth](../README.md#StoreKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

