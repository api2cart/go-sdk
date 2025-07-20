# \BasketAPI

All URIs are relative to *https://api.api2cart.local.com/v1.1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BasketInfo**](BasketAPI.md#BasketInfo) | **Get** /basket.info.json | basket.info
[**BasketItemAdd**](BasketAPI.md#BasketItemAdd) | **Post** /basket.item.add.json | basket.item.add
[**BasketLiveShippingServiceCreate**](BasketAPI.md#BasketLiveShippingServiceCreate) | **Post** /basket.live_shipping_service.create.json | basket.live_shipping_service.create
[**BasketLiveShippingServiceDelete**](BasketAPI.md#BasketLiveShippingServiceDelete) | **Delete** /basket.live_shipping_service.delete.json | basket.live_shipping_service.delete
[**BasketLiveShippingServiceList**](BasketAPI.md#BasketLiveShippingServiceList) | **Get** /basket.live_shipping_service.list.json | basket.live_shipping_service.list



## BasketInfo

> BasketInfo200Response BasketInfo(ctx).Id(id).StoreId(storeId).ResponseFields(responseFields).Params(params).Exclude(exclude).Execute()

basket.info



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
	storeId := "1" // string | Store Id (optional)
	responseFields := "{result}" // string | Set this parameter in order to choose which entity fields you want to retrieve (optional)
	params := "id,model,price,images" // string | Set this parameter in order to choose which entity fields you want to retrieve (optional) (default to "force_all")
	exclude := "false" // string | Set this parameter in order to choose which entity fields you want to ignore. Works only if parameter `params` equal force_all (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BasketAPI.BasketInfo(context.Background()).Id(id).StoreId(storeId).ResponseFields(responseFields).Params(params).Exclude(exclude).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasketAPI.BasketInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BasketInfo`: BasketInfo200Response
	fmt.Fprintf(os.Stdout, "Response from `BasketAPI.BasketInfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBasketInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | Entity id | 
 **storeId** | **string** | Store Id | 
 **responseFields** | **string** | Set this parameter in order to choose which entity fields you want to retrieve | 
 **params** | **string** | Set this parameter in order to choose which entity fields you want to retrieve | [default to &quot;force_all&quot;]
 **exclude** | **string** | Set this parameter in order to choose which entity fields you want to ignore. Works only if parameter &#x60;params&#x60; equal force_all | 

### Return type

[**BasketInfo200Response**](BasketInfo200Response.md)

### Authorization

[StoreKeyAuth](../README.md#StoreKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BasketItemAdd

> BasketItemAdd200Response BasketItemAdd(ctx).CustomerId(customerId).ProductId(productId).VariantId(variantId).Quantity(quantity).StoreId(storeId).Execute()

basket.item.add



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
	customerId := "5" // string | Retrieves orders specified by customer id
	productId := "10" // string | Defines id of the product which should be added to the basket
	variantId := "45" // string | Defines product's variants specified by variant id (optional)
	quantity := float32(6) // float32 | Defines new items quantity (optional) (default to 0)
	storeId := "1" // string | Store Id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BasketAPI.BasketItemAdd(context.Background()).CustomerId(customerId).ProductId(productId).VariantId(variantId).Quantity(quantity).StoreId(storeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasketAPI.BasketItemAdd``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BasketItemAdd`: BasketItemAdd200Response
	fmt.Fprintf(os.Stdout, "Response from `BasketAPI.BasketItemAdd`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBasketItemAddRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **customerId** | **string** | Retrieves orders specified by customer id | 
 **productId** | **string** | Defines id of the product which should be added to the basket | 
 **variantId** | **string** | Defines product&#39;s variants specified by variant id | 
 **quantity** | **float32** | Defines new items quantity | [default to 0]
 **storeId** | **string** | Store Id | 

### Return type

[**BasketItemAdd200Response**](BasketItemAdd200Response.md)

### Authorization

[StoreKeyAuth](../README.md#StoreKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BasketLiveShippingServiceCreate

> BasketLiveShippingServiceCreate200Response BasketLiveShippingServiceCreate(ctx).Name(name).Callback(callback).StoreId(storeId).Execute()

basket.live_shipping_service.create



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
	name := "BestDelivery" // string | Shipping Service Name
	callback := "https://example.com/callback" // string | Callback url that returns shipping rates. It should be able to accept POST requests with json data.
	storeId := "1" // string | Store Id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BasketAPI.BasketLiveShippingServiceCreate(context.Background()).Name(name).Callback(callback).StoreId(storeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasketAPI.BasketLiveShippingServiceCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BasketLiveShippingServiceCreate`: BasketLiveShippingServiceCreate200Response
	fmt.Fprintf(os.Stdout, "Response from `BasketAPI.BasketLiveShippingServiceCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBasketLiveShippingServiceCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | Shipping Service Name | 
 **callback** | **string** | Callback url that returns shipping rates. It should be able to accept POST requests with json data. | 
 **storeId** | **string** | Store Id | 

### Return type

[**BasketLiveShippingServiceCreate200Response**](BasketLiveShippingServiceCreate200Response.md)

### Authorization

[StoreKeyAuth](../README.md#StoreKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BasketLiveShippingServiceDelete

> BasketLiveShippingServiceDelete200Response BasketLiveShippingServiceDelete(ctx).Id(id).Execute()

basket.live_shipping_service.delete



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
	id := int32(5) // int32 | Entity id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BasketAPI.BasketLiveShippingServiceDelete(context.Background()).Id(id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasketAPI.BasketLiveShippingServiceDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BasketLiveShippingServiceDelete`: BasketLiveShippingServiceDelete200Response
	fmt.Fprintf(os.Stdout, "Response from `BasketAPI.BasketLiveShippingServiceDelete`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBasketLiveShippingServiceDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int32** | Entity id | 

### Return type

[**BasketLiveShippingServiceDelete200Response**](BasketLiveShippingServiceDelete200Response.md)

### Authorization

[StoreKeyAuth](../README.md#StoreKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BasketLiveShippingServiceList

> BasketLiveShippingServiceList200Response BasketLiveShippingServiceList(ctx).Start(start).Count(count).StoreId(storeId).Execute()

basket.live_shipping_service.list



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
	storeId := "1" // string | Store Id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BasketAPI.BasketLiveShippingServiceList(context.Background()).Start(start).Count(count).StoreId(storeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasketAPI.BasketLiveShippingServiceList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BasketLiveShippingServiceList`: BasketLiveShippingServiceList200Response
	fmt.Fprintf(os.Stdout, "Response from `BasketAPI.BasketLiveShippingServiceList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBasketLiveShippingServiceListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **start** | **int32** | This parameter sets the number from which you want to get entities | [default to 0]
 **count** | **int32** | This parameter sets the entity amount that has to be retrieved. Max allowed count&#x3D;250 | [default to 10]
 **storeId** | **string** | Store Id | 

### Return type

[**BasketLiveShippingServiceList200Response**](BasketLiveShippingServiceList200Response.md)

### Authorization

[StoreKeyAuth](../README.md#StoreKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

