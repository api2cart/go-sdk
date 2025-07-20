# \SubscriberAPI

All URIs are relative to *https://api.api2cart.local.com/v1.1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SubscriberList**](SubscriberAPI.md#SubscriberList) | **Get** /subscriber.list.json | subscriber.list



## SubscriberList

> ModelResponseSubscriberList SubscriberList(ctx).Ids(ids).Start(start).Count(count).PageCursor(pageCursor).Subscribed(subscribed).StoreId(storeId).Email(email).CreatedFrom(createdFrom).CreatedTo(createdTo).ModifiedFrom(modifiedFrom).ModifiedTo(modifiedTo).ResponseFields(responseFields).Params(params).Exclude(exclude).Execute()

subscriber.list



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
	ids := "24,25" // string | Retrieves subscribers specified by ids (optional)
	start := int32(0) // int32 | This parameter sets the number from which you want to get entities (optional) (default to 0)
	count := int32(20) // int32 | This parameter sets the entity amount that has to be retrieved. Max allowed count=250 (optional) (default to 10)
	pageCursor := "pageCursor_example" // string | Used to retrieve entities via cursor-based pagination (it can't be used with any other filtering parameter) (optional)
	subscribed := false // bool | Filter by subscription status (optional)
	storeId := "1" // string | Store Id (optional)
	email := "mail@example.com" // string | Filter subscribers by email (optional)
	createdFrom := "2010-07-29 13:45:52" // string | Retrieve entities from their creation date (optional)
	createdTo := "2100-08-29 13:45:52" // string | Retrieve entities to their creation date (optional)
	modifiedFrom := "2010-07-29 13:45:52" // string | Retrieve entities from their modification date (optional)
	modifiedTo := "2100-08-29 13:45:52" // string | Retrieve entities to their modification date (optional)
	responseFields := "{return_code,return_message,pagination,result}" // string | Set this parameter in order to choose which entity fields you want to retrieve (optional)
	params := "id,model,price,images" // string | Set this parameter in order to choose which entity fields you want to retrieve (optional) (default to "force_all")
	exclude := "false" // string | Set this parameter in order to choose which entity fields you want to ignore. Works only if parameter `params` equal force_all (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriberAPI.SubscriberList(context.Background()).Ids(ids).Start(start).Count(count).PageCursor(pageCursor).Subscribed(subscribed).StoreId(storeId).Email(email).CreatedFrom(createdFrom).CreatedTo(createdTo).ModifiedFrom(modifiedFrom).ModifiedTo(modifiedTo).ResponseFields(responseFields).Params(params).Exclude(exclude).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriberAPI.SubscriberList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriberList`: ModelResponseSubscriberList
	fmt.Fprintf(os.Stdout, "Response from `SubscriberAPI.SubscriberList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubscriberListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ids** | **string** | Retrieves subscribers specified by ids | 
 **start** | **int32** | This parameter sets the number from which you want to get entities | [default to 0]
 **count** | **int32** | This parameter sets the entity amount that has to be retrieved. Max allowed count&#x3D;250 | [default to 10]
 **pageCursor** | **string** | Used to retrieve entities via cursor-based pagination (it can&#39;t be used with any other filtering parameter) | 
 **subscribed** | **bool** | Filter by subscription status | 
 **storeId** | **string** | Store Id | 
 **email** | **string** | Filter subscribers by email | 
 **createdFrom** | **string** | Retrieve entities from their creation date | 
 **createdTo** | **string** | Retrieve entities to their creation date | 
 **modifiedFrom** | **string** | Retrieve entities from their modification date | 
 **modifiedTo** | **string** | Retrieve entities to their modification date | 
 **responseFields** | **string** | Set this parameter in order to choose which entity fields you want to retrieve | 
 **params** | **string** | Set this parameter in order to choose which entity fields you want to retrieve | [default to &quot;force_all&quot;]
 **exclude** | **string** | Set this parameter in order to choose which entity fields you want to ignore. Works only if parameter &#x60;params&#x60; equal force_all | 

### Return type

[**ModelResponseSubscriberList**](ModelResponseSubscriberList.md)

### Authorization

[StoreKeyAuth](../README.md#StoreKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

