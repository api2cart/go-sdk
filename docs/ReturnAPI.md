# \ReturnAPI

All URIs are relative to *https://api.api2cart.com/v1.1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReturnActionList**](ReturnAPI.md#ReturnActionList) | **Get** /return.action.list.json | return.action.list
[**ReturnCount**](ReturnAPI.md#ReturnCount) | **Get** /return.count.json | return.count
[**ReturnInfo**](ReturnAPI.md#ReturnInfo) | **Get** /return.info.json | return.info
[**ReturnList**](ReturnAPI.md#ReturnList) | **Get** /return.list.json | return.list
[**ReturnReasonList**](ReturnAPI.md#ReturnReasonList) | **Get** /return.reason.list.json | return.reason.list
[**ReturnStatusList**](ReturnAPI.md#ReturnStatusList) | **Get** /return.status.list.json | return.status.list



## ReturnActionList

> ReturnActionList200Response ReturnActionList(ctx).Execute()

return.action.list



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
	resp, r, err := apiClient.ReturnAPI.ReturnActionList(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReturnAPI.ReturnActionList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReturnActionList`: ReturnActionList200Response
	fmt.Fprintf(os.Stdout, "Response from `ReturnAPI.ReturnActionList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiReturnActionListRequest struct via the builder pattern


### Return type

[**ReturnActionList200Response**](ReturnActionList200Response.md)

### Authorization

[StoreKeyAuth](../README.md#StoreKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReturnCount

> ReturnCount200Response ReturnCount(ctx).OrderIds(orderIds).CustomerId(customerId).StoreId(storeId).Status(status).ReturnType(returnType).CreatedFrom(createdFrom).CreatedTo(createdTo).ModifiedFrom(modifiedFrom).ModifiedTo(modifiedTo).ReportRequestId(reportRequestId).DisableReportCache(disableReportCache).Execute()

return.count



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
	orderIds := "24,25" // string | Counts return requests specified by order ids (optional)
	customerId := "5" // string | Counts return requests quantity specified by customer id (optional)
	storeId := "1" // string | Store Id (optional)
	status := "disabled" // string | Defines status (optional)
	returnType := "FBA" // string | Retrieves returns specified by return type (optional)
	createdFrom := "2010-07-29 13:45:52" // string | Retrieve entities from their creation date (optional)
	createdTo := "2100-08-29 13:45:52" // string | Retrieve entities to their creation date (optional)
	modifiedFrom := "2010-07-29 13:45:52" // string | Retrieve entities from their modification date (optional)
	modifiedTo := "2100-08-29 13:45:52" // string | Retrieve entities to their modification date (optional)
	reportRequestId := "105245017661" // string | Report request id (optional)
	disableReportCache := false // bool | Disable report cache for current request (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReturnAPI.ReturnCount(context.Background()).OrderIds(orderIds).CustomerId(customerId).StoreId(storeId).Status(status).ReturnType(returnType).CreatedFrom(createdFrom).CreatedTo(createdTo).ModifiedFrom(modifiedFrom).ModifiedTo(modifiedTo).ReportRequestId(reportRequestId).DisableReportCache(disableReportCache).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReturnAPI.ReturnCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReturnCount`: ReturnCount200Response
	fmt.Fprintf(os.Stdout, "Response from `ReturnAPI.ReturnCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReturnCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orderIds** | **string** | Counts return requests specified by order ids | 
 **customerId** | **string** | Counts return requests quantity specified by customer id | 
 **storeId** | **string** | Store Id | 
 **status** | **string** | Defines status | 
 **returnType** | **string** | Retrieves returns specified by return type | 
 **createdFrom** | **string** | Retrieve entities from their creation date | 
 **createdTo** | **string** | Retrieve entities to their creation date | 
 **modifiedFrom** | **string** | Retrieve entities from their modification date | 
 **modifiedTo** | **string** | Retrieve entities to their modification date | 
 **reportRequestId** | **string** | Report request id | 
 **disableReportCache** | **bool** | Disable report cache for current request | [default to false]

### Return type

[**ReturnCount200Response**](ReturnCount200Response.md)

### Authorization

[StoreKeyAuth](../README.md#StoreKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReturnInfo

> ReturnInfo200Response ReturnInfo(ctx).Id(id).OrderId(orderId).StoreId(storeId).ResponseFields(responseFields).Params(params).Exclude(exclude).Execute()

return.info



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
	id := "10" // string | Entity id
	orderId := "25" // string | Defines the order id (optional)
	storeId := "1" // string | Store Id (optional)
	responseFields := "{return_code,return_message,result}" // string | Set this parameter in order to choose which entity fields you want to retrieve (optional)
	params := "id,order_products" // string | Set this parameter in order to choose which entity fields you want to retrieve (optional) (default to "id,order_products")
	exclude := "id,order_id" // string | Set this parameter in order to choose which entity fields you want to ignore. Works only if parameter `params` equal force_all (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReturnAPI.ReturnInfo(context.Background()).Id(id).OrderId(orderId).StoreId(storeId).ResponseFields(responseFields).Params(params).Exclude(exclude).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReturnAPI.ReturnInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReturnInfo`: ReturnInfo200Response
	fmt.Fprintf(os.Stdout, "Response from `ReturnAPI.ReturnInfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReturnInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | Entity id | 
 **orderId** | **string** | Defines the order id | 
 **storeId** | **string** | Store Id | 
 **responseFields** | **string** | Set this parameter in order to choose which entity fields you want to retrieve | 
 **params** | **string** | Set this parameter in order to choose which entity fields you want to retrieve | [default to &quot;id,order_products&quot;]
 **exclude** | **string** | Set this parameter in order to choose which entity fields you want to ignore. Works only if parameter &#x60;params&#x60; equal force_all | 

### Return type

[**ReturnInfo200Response**](ReturnInfo200Response.md)

### Authorization

[StoreKeyAuth](../README.md#StoreKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReturnList

> ModelResponseReturnList ReturnList(ctx).Start(start).Count(count).PageCursor(pageCursor).OrderId(orderId).OrderIds(orderIds).CustomerId(customerId).StoreId(storeId).Status(status).ReturnType(returnType).CreatedFrom(createdFrom).CreatedTo(createdTo).ModifiedFrom(modifiedFrom).ModifiedTo(modifiedTo).ResponseFields(responseFields).Params(params).Exclude(exclude).ReportRequestId(reportRequestId).DisableReportCache(disableReportCache).Execute()

return.list



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
	pageCursor := "pageCursor_example" // string | Used to retrieve entities via cursor-based pagination (it can't be used with any other filtering parameter) (optional)
	orderId := "25" // string | Defines the order id (optional)
	orderIds := "24,25" // string | Retrieves return requests specified by order ids (optional)
	customerId := "5" // string | Retrieves return requests specified by customer id (optional)
	storeId := "1" // string | Store Id (optional)
	status := "disabled" // string | Defines status (optional)
	returnType := "FBA" // string | Retrieves returns specified by return type (optional)
	createdFrom := "2010-07-29 13:45:52" // string | Retrieve entities from their creation date (optional)
	createdTo := "2100-08-29 13:45:52" // string | Retrieve entities to their creation date (optional)
	modifiedFrom := "2010-07-29 13:45:52" // string | Retrieve entities from their modification date (optional)
	modifiedTo := "2100-08-29 13:45:52" // string | Retrieve entities to their modification date (optional)
	responseFields := "{return_code,return_message,pagination,result}" // string | Set this parameter in order to choose which entity fields you want to retrieve (optional)
	params := "id,order_products" // string | Set this parameter in order to choose which entity fields you want to retrieve (optional) (default to "id,order_products")
	exclude := "id,order_id" // string | Set this parameter in order to choose which entity fields you want to ignore. Works only if parameter `params` equal force_all (optional)
	reportRequestId := "105245017661" // string | Report request id (optional)
	disableReportCache := false // bool | Disable report cache for current request (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReturnAPI.ReturnList(context.Background()).Start(start).Count(count).PageCursor(pageCursor).OrderId(orderId).OrderIds(orderIds).CustomerId(customerId).StoreId(storeId).Status(status).ReturnType(returnType).CreatedFrom(createdFrom).CreatedTo(createdTo).ModifiedFrom(modifiedFrom).ModifiedTo(modifiedTo).ResponseFields(responseFields).Params(params).Exclude(exclude).ReportRequestId(reportRequestId).DisableReportCache(disableReportCache).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReturnAPI.ReturnList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReturnList`: ModelResponseReturnList
	fmt.Fprintf(os.Stdout, "Response from `ReturnAPI.ReturnList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReturnListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **start** | **int32** | This parameter sets the number from which you want to get entities | [default to 0]
 **count** | **int32** | This parameter sets the entity amount that has to be retrieved. Max allowed count&#x3D;250 | [default to 10]
 **pageCursor** | **string** | Used to retrieve entities via cursor-based pagination (it can&#39;t be used with any other filtering parameter) | 
 **orderId** | **string** | Defines the order id | 
 **orderIds** | **string** | Retrieves return requests specified by order ids | 
 **customerId** | **string** | Retrieves return requests specified by customer id | 
 **storeId** | **string** | Store Id | 
 **status** | **string** | Defines status | 
 **returnType** | **string** | Retrieves returns specified by return type | 
 **createdFrom** | **string** | Retrieve entities from their creation date | 
 **createdTo** | **string** | Retrieve entities to their creation date | 
 **modifiedFrom** | **string** | Retrieve entities from their modification date | 
 **modifiedTo** | **string** | Retrieve entities to their modification date | 
 **responseFields** | **string** | Set this parameter in order to choose which entity fields you want to retrieve | 
 **params** | **string** | Set this parameter in order to choose which entity fields you want to retrieve | [default to &quot;id,order_products&quot;]
 **exclude** | **string** | Set this parameter in order to choose which entity fields you want to ignore. Works only if parameter &#x60;params&#x60; equal force_all | 
 **reportRequestId** | **string** | Report request id | 
 **disableReportCache** | **bool** | Disable report cache for current request | [default to false]

### Return type

[**ModelResponseReturnList**](ModelResponseReturnList.md)

### Authorization

[StoreKeyAuth](../README.md#StoreKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReturnReasonList

> ReturnReasonList200Response ReturnReasonList(ctx).StoreId(storeId).Execute()

return.reason.list



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
	storeId := "1" // string | Store Id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReturnAPI.ReturnReasonList(context.Background()).StoreId(storeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReturnAPI.ReturnReasonList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReturnReasonList`: ReturnReasonList200Response
	fmt.Fprintf(os.Stdout, "Response from `ReturnAPI.ReturnReasonList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReturnReasonListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **storeId** | **string** | Store Id | 

### Return type

[**ReturnReasonList200Response**](ReturnReasonList200Response.md)

### Authorization

[StoreKeyAuth](../README.md#StoreKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReturnStatusList

> ReturnStatusList200Response ReturnStatusList(ctx).Execute()

return.status.list



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
	resp, r, err := apiClient.ReturnAPI.ReturnStatusList(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReturnAPI.ReturnStatusList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReturnStatusList`: ReturnStatusList200Response
	fmt.Fprintf(os.Stdout, "Response from `ReturnAPI.ReturnStatusList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiReturnStatusListRequest struct via the builder pattern


### Return type

[**ReturnStatusList200Response**](ReturnStatusList200Response.md)

### Authorization

[StoreKeyAuth](../README.md#StoreKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

