# \OrderAPI

All URIs are relative to *https://api.api2cart.com/v1.1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OrderAbandonedList**](OrderAPI.md#OrderAbandonedList) | **Get** /order.abandoned.list.json | order.abandoned.list
[**OrderAdd**](OrderAPI.md#OrderAdd) | **Post** /order.add.json | order.add
[**OrderCount**](OrderAPI.md#OrderCount) | **Get** /order.count.json | order.count
[**OrderFinancialStatusList**](OrderAPI.md#OrderFinancialStatusList) | **Get** /order.financial_status.list.json | order.financial_status.list
[**OrderFind**](OrderAPI.md#OrderFind) | **Get** /order.find.json | order.find
[**OrderFulfillmentStatusList**](OrderAPI.md#OrderFulfillmentStatusList) | **Get** /order.fulfillment_status.list.json | order.fulfillment_status.list
[**OrderInfo**](OrderAPI.md#OrderInfo) | **Get** /order.info.json | order.info
[**OrderList**](OrderAPI.md#OrderList) | **Get** /order.list.json | order.list
[**OrderPreestimateShippingList**](OrderAPI.md#OrderPreestimateShippingList) | **Post** /order.preestimate_shipping.list.json | order.preestimate_shipping.list
[**OrderRefundAdd**](OrderAPI.md#OrderRefundAdd) | **Post** /order.refund.add.json | order.refund.add
[**OrderReturnAdd**](OrderAPI.md#OrderReturnAdd) | **Post** /order.return.add.json | order.return.add
[**OrderReturnDelete**](OrderAPI.md#OrderReturnDelete) | **Delete** /order.return.delete.json | order.return.delete
[**OrderReturnUpdate**](OrderAPI.md#OrderReturnUpdate) | **Put** /order.return.update.json | order.return.update
[**OrderShipmentAdd**](OrderAPI.md#OrderShipmentAdd) | **Post** /order.shipment.add.json | order.shipment.add
[**OrderShipmentAddBatch**](OrderAPI.md#OrderShipmentAddBatch) | **Post** /order.shipment.add.batch.json | order.shipment.add.batch
[**OrderShipmentDelete**](OrderAPI.md#OrderShipmentDelete) | **Delete** /order.shipment.delete.json | order.shipment.delete
[**OrderShipmentInfo**](OrderAPI.md#OrderShipmentInfo) | **Get** /order.shipment.info.json | order.shipment.info
[**OrderShipmentList**](OrderAPI.md#OrderShipmentList) | **Get** /order.shipment.list.json | order.shipment.list
[**OrderShipmentTrackingAdd**](OrderAPI.md#OrderShipmentTrackingAdd) | **Post** /order.shipment.tracking.add.json | order.shipment.tracking.add
[**OrderShipmentUpdate**](OrderAPI.md#OrderShipmentUpdate) | **Put** /order.shipment.update.json | order.shipment.update
[**OrderStatusList**](OrderAPI.md#OrderStatusList) | **Get** /order.status.list.json | order.status.list
[**OrderTransactionList**](OrderAPI.md#OrderTransactionList) | **Get** /order.transaction.list.json | order.transaction.list
[**OrderUpdate**](OrderAPI.md#OrderUpdate) | **Put** /order.update.json | order.update



## OrderAbandonedList

> ModelResponseOrderAbandonedList OrderAbandonedList(ctx).CustomerId(customerId).CustomerEmail(customerEmail).CreatedTo(createdTo).CreatedFrom(createdFrom).ModifiedTo(modifiedTo).ModifiedFrom(modifiedFrom).SkipEmptyEmail(skipEmptyEmail).StoreId(storeId).PageCursor(pageCursor).Count(count).Start(start).Params(params).ResponseFields(responseFields).Exclude(exclude).Execute()

order.abandoned.list



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
	customerId := "5" // string | Retrieves orders specified by customer id (optional)
	customerEmail := "jubari@hannsgroup.com" // string | Retrieves orders specified by customer email (optional)
	createdTo := "2100-08-29 13:45:52" // string | Retrieve entities to their creation date (optional)
	createdFrom := "2010-07-29 13:45:52" // string | Retrieve entities from their creation date (optional)
	modifiedTo := "2100-08-29 13:45:52" // string | Retrieve entities to their modification date (optional)
	modifiedFrom := "2010-07-29 13:45:52" // string | Retrieve entities from their modification date (optional)
	skipEmptyEmail := true // bool | Filter empty emails (optional) (default to false)
	storeId := "1" // string | Store Id (optional)
	pageCursor := "pageCursor_example" // string | Used to retrieve entities via cursor-based pagination (it can't be used with any other filtering parameter) (optional)
	count := int32(20) // int32 | This parameter sets the entity amount that has to be retrieved. Max allowed count=250 (optional) (default to 10)
	start := int32(0) // int32 | This parameter sets the number from which you want to get entities (optional) (default to 0)
	params := "force_all" // string | Set this parameter in order to choose which entity fields you want to retrieve (optional) (default to "customer,totals,items")
	responseFields := "{return_code,pagination,result{order{id,customer{email},created_at,totals{total},order_products}}}" // string | Set this parameter in order to choose which entity fields you want to retrieve (optional)
	exclude := "customer" // string | Set this parameter in order to choose which entity fields you want to ignore. Works only if parameter `params` equal force_all (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrderAPI.OrderAbandonedList(context.Background()).CustomerId(customerId).CustomerEmail(customerEmail).CreatedTo(createdTo).CreatedFrom(createdFrom).ModifiedTo(modifiedTo).ModifiedFrom(modifiedFrom).SkipEmptyEmail(skipEmptyEmail).StoreId(storeId).PageCursor(pageCursor).Count(count).Start(start).Params(params).ResponseFields(responseFields).Exclude(exclude).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrderAPI.OrderAbandonedList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrderAbandonedList`: ModelResponseOrderAbandonedList
	fmt.Fprintf(os.Stdout, "Response from `OrderAPI.OrderAbandonedList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrderAbandonedListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **customerId** | **string** | Retrieves orders specified by customer id | 
 **customerEmail** | **string** | Retrieves orders specified by customer email | 
 **createdTo** | **string** | Retrieve entities to their creation date | 
 **createdFrom** | **string** | Retrieve entities from their creation date | 
 **modifiedTo** | **string** | Retrieve entities to their modification date | 
 **modifiedFrom** | **string** | Retrieve entities from their modification date | 
 **skipEmptyEmail** | **bool** | Filter empty emails | [default to false]
 **storeId** | **string** | Store Id | 
 **pageCursor** | **string** | Used to retrieve entities via cursor-based pagination (it can&#39;t be used with any other filtering parameter) | 
 **count** | **int32** | This parameter sets the entity amount that has to be retrieved. Max allowed count&#x3D;250 | [default to 10]
 **start** | **int32** | This parameter sets the number from which you want to get entities | [default to 0]
 **params** | **string** | Set this parameter in order to choose which entity fields you want to retrieve | [default to &quot;customer,totals,items&quot;]
 **responseFields** | **string** | Set this parameter in order to choose which entity fields you want to retrieve | 
 **exclude** | **string** | Set this parameter in order to choose which entity fields you want to ignore. Works only if parameter &#x60;params&#x60; equal force_all | 

### Return type

[**ModelResponseOrderAbandonedList**](ModelResponseOrderAbandonedList.md)

### Authorization

[StoreKeyAuth](../README.md#StoreKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrderAdd

> OrderAdd200Response OrderAdd(ctx).OrderAdd(orderAdd).Execute()

order.add



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
	orderAdd := *openapiclient.NewOrderAdd("Completed", "jubari@hannsgroup.com", "Adam", "Smith", "Green str. 35", "Chicago", "12345", "IL", "US", []openapiclient.OrderAddOrderItemInner{*openapiclient.NewOrderAddOrderItemInner("125, where {x} - 1,2,3,... etc", "Product 1, where {x} - 1,2,3,... etc", float32(1.32, where {x} - 1,2,3,... etc), int32(5, where {x} - 1,2,3,... etc))}) // OrderAdd | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrderAPI.OrderAdd(context.Background()).OrderAdd(orderAdd).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrderAPI.OrderAdd``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrderAdd`: OrderAdd200Response
	fmt.Fprintf(os.Stdout, "Response from `OrderAPI.OrderAdd`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrderAddRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orderAdd** | [**OrderAdd**](OrderAdd.md) |  | 

### Return type

[**OrderAdd200Response**](OrderAdd200Response.md)

### Authorization

[StoreKeyAuth](../README.md#StoreKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrderCount

> OrderCount200Response OrderCount(ctx).CustomerId(customerId).CustomerEmail(customerEmail).OrderStatus(orderStatus).OrderStatusIds(orderStatusIds).CreatedTo(createdTo).CreatedFrom(createdFrom).ModifiedTo(modifiedTo).ModifiedFrom(modifiedFrom).StoreId(storeId).Ids(ids).OrderIds(orderIds).EbayOrderStatus(ebayOrderStatus).FinancialStatus(financialStatus).FinancialStatusIds(financialStatusIds).FulfillmentChannel(fulfillmentChannel).FulfillmentStatus(fulfillmentStatus).ShippingMethod(shippingMethod).DeliveryMethod(deliveryMethod).Tags(tags).ShipNodeType(shipNodeType).Execute()

order.count



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
	customerId := "5" // string | Counts orders quantity specified by customer id (optional)
	customerEmail := "jubari@hannsgroup.com" // string | Counts orders quantity specified by customer email (optional)
	orderStatus := "Completed" // string | Counts orders quantity specified by order status (optional)
	orderStatusIds := []string{"Inner_example"} // []string | Retrieves orders specified by order statuses (optional)
	createdTo := "2100-08-29 13:45:52" // string | Retrieve entities to their creation date (optional)
	createdFrom := "2010-07-29 13:45:52" // string | Retrieve entities from their creation date (optional)
	modifiedTo := "2100-08-29 13:45:52" // string | Retrieve entities to their modification date (optional)
	modifiedFrom := "2010-07-29 13:45:52" // string | Retrieve entities from their modification date (optional)
	storeId := "1" // string | Counts orders quantity specified by store id (optional)
	ids := "24,25" // string | Counts orders specified by ids (optional)
	orderIds := "24,25" // string | Counts orders specified by order ids (optional)
	ebayOrderStatus := "Active" // string | Counts orders quantity specified by order status (optional)
	financialStatus := "paid" // string | Counts orders quantity specified by financial status (optional)
	financialStatusIds := []string{"Inner_example"} // []string | Retrieves orders count specified by financial status ids (optional)
	fulfillmentChannel := "local" // string | Retrieves order with a fulfillment channel (optional)
	fulfillmentStatus := "fulfilled" // string | Create order with fulfillment status (optional)
	shippingMethod := "flatrate_flatrate" // string | Retrieve entities according to shipping method (optional)
	deliveryMethod := "local" // string | Retrieves order with delivery method (optional)
	tags := "tag1,tag2" // string | Order tags (optional)
	shipNodeType := "SellerFulfilled" // string | Retrieves order with ship node type (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrderAPI.OrderCount(context.Background()).CustomerId(customerId).CustomerEmail(customerEmail).OrderStatus(orderStatus).OrderStatusIds(orderStatusIds).CreatedTo(createdTo).CreatedFrom(createdFrom).ModifiedTo(modifiedTo).ModifiedFrom(modifiedFrom).StoreId(storeId).Ids(ids).OrderIds(orderIds).EbayOrderStatus(ebayOrderStatus).FinancialStatus(financialStatus).FinancialStatusIds(financialStatusIds).FulfillmentChannel(fulfillmentChannel).FulfillmentStatus(fulfillmentStatus).ShippingMethod(shippingMethod).DeliveryMethod(deliveryMethod).Tags(tags).ShipNodeType(shipNodeType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrderAPI.OrderCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrderCount`: OrderCount200Response
	fmt.Fprintf(os.Stdout, "Response from `OrderAPI.OrderCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrderCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **customerId** | **string** | Counts orders quantity specified by customer id | 
 **customerEmail** | **string** | Counts orders quantity specified by customer email | 
 **orderStatus** | **string** | Counts orders quantity specified by order status | 
 **orderStatusIds** | **[]string** | Retrieves orders specified by order statuses | 
 **createdTo** | **string** | Retrieve entities to their creation date | 
 **createdFrom** | **string** | Retrieve entities from their creation date | 
 **modifiedTo** | **string** | Retrieve entities to their modification date | 
 **modifiedFrom** | **string** | Retrieve entities from their modification date | 
 **storeId** | **string** | Counts orders quantity specified by store id | 
 **ids** | **string** | Counts orders specified by ids | 
 **orderIds** | **string** | Counts orders specified by order ids | 
 **ebayOrderStatus** | **string** | Counts orders quantity specified by order status | 
 **financialStatus** | **string** | Counts orders quantity specified by financial status | 
 **financialStatusIds** | **[]string** | Retrieves orders count specified by financial status ids | 
 **fulfillmentChannel** | **string** | Retrieves order with a fulfillment channel | 
 **fulfillmentStatus** | **string** | Create order with fulfillment status | 
 **shippingMethod** | **string** | Retrieve entities according to shipping method | 
 **deliveryMethod** | **string** | Retrieves order with delivery method | 
 **tags** | **string** | Order tags | 
 **shipNodeType** | **string** | Retrieves order with ship node type | 

### Return type

[**OrderCount200Response**](OrderCount200Response.md)

### Authorization

[StoreKeyAuth](../README.md#StoreKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrderFinancialStatusList

> OrderFinancialStatusList200Response OrderFinancialStatusList(ctx).Execute()

order.financial_status.list



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
	resp, r, err := apiClient.OrderAPI.OrderFinancialStatusList(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrderAPI.OrderFinancialStatusList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrderFinancialStatusList`: OrderFinancialStatusList200Response
	fmt.Fprintf(os.Stdout, "Response from `OrderAPI.OrderFinancialStatusList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiOrderFinancialStatusListRequest struct via the builder pattern


### Return type

[**OrderFinancialStatusList200Response**](OrderFinancialStatusList200Response.md)

### Authorization

[StoreKeyAuth](../README.md#StoreKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrderFind

> OrderFind200Response OrderFind(ctx).CustomerId(customerId).CustomerEmail(customerEmail).OrderStatus(orderStatus).Start(start).Count(count).Params(params).Exclude(exclude).CreatedTo(createdTo).CreatedFrom(createdFrom).ModifiedTo(modifiedTo).ModifiedFrom(modifiedFrom).FinancialStatus(financialStatus).Execute()

order.find



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
	customerId := "5" // string | Retrieves orders specified by customer id (optional)
	customerEmail := "jubari@hannsgroup.com" // string | Retrieves orders specified by customer email (optional)
	orderStatus := "Completed" // string | Retrieves orders specified by order status (optional)
	start := int32(0) // int32 | This parameter sets the number from which you want to get entities (optional) (default to 0)
	count := int32(20) // int32 | This parameter sets the entity amount that has to be retrieved. Max allowed count=250 (optional) (default to 10)
	params := "order_id,totals,status" // string | Set this parameter in order to choose which entity fields you want to retrieve (optional) (default to "order_id,customer,totals,address,items,bundles,status")
	exclude := "order_id,totals,status" // string | Set this parameter in order to choose which entity fields you want to ignore. Works only if parameter `params` equal force_all (optional)
	createdTo := "2100-08-29 13:45:52" // string | Retrieve entities to their creation date (optional)
	createdFrom := "2010-07-29 13:45:52" // string | Retrieve entities from their creation date (optional)
	modifiedTo := "2100-08-29 13:45:52" // string | Retrieve entities to their modification date (optional)
	modifiedFrom := "2010-07-29 13:45:52" // string | Retrieve entities from their modification date (optional)
	financialStatus := "paid" // string | Retrieves orders specified by financial status (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrderAPI.OrderFind(context.Background()).CustomerId(customerId).CustomerEmail(customerEmail).OrderStatus(orderStatus).Start(start).Count(count).Params(params).Exclude(exclude).CreatedTo(createdTo).CreatedFrom(createdFrom).ModifiedTo(modifiedTo).ModifiedFrom(modifiedFrom).FinancialStatus(financialStatus).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrderAPI.OrderFind``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrderFind`: OrderFind200Response
	fmt.Fprintf(os.Stdout, "Response from `OrderAPI.OrderFind`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrderFindRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **customerId** | **string** | Retrieves orders specified by customer id | 
 **customerEmail** | **string** | Retrieves orders specified by customer email | 
 **orderStatus** | **string** | Retrieves orders specified by order status | 
 **start** | **int32** | This parameter sets the number from which you want to get entities | [default to 0]
 **count** | **int32** | This parameter sets the entity amount that has to be retrieved. Max allowed count&#x3D;250 | [default to 10]
 **params** | **string** | Set this parameter in order to choose which entity fields you want to retrieve | [default to &quot;order_id,customer,totals,address,items,bundles,status&quot;]
 **exclude** | **string** | Set this parameter in order to choose which entity fields you want to ignore. Works only if parameter &#x60;params&#x60; equal force_all | 
 **createdTo** | **string** | Retrieve entities to their creation date | 
 **createdFrom** | **string** | Retrieve entities from their creation date | 
 **modifiedTo** | **string** | Retrieve entities to their modification date | 
 **modifiedFrom** | **string** | Retrieve entities from their modification date | 
 **financialStatus** | **string** | Retrieves orders specified by financial status | 

### Return type

[**OrderFind200Response**](OrderFind200Response.md)

### Authorization

[StoreKeyAuth](../README.md#StoreKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrderFulfillmentStatusList

> OrderFulfillmentStatusList200Response OrderFulfillmentStatusList(ctx).Action(action).Execute()

order.fulfillment_status.list



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
	action := "add" // string | Available statuses for the specified action. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrderAPI.OrderFulfillmentStatusList(context.Background()).Action(action).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrderAPI.OrderFulfillmentStatusList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrderFulfillmentStatusList`: OrderFulfillmentStatusList200Response
	fmt.Fprintf(os.Stdout, "Response from `OrderAPI.OrderFulfillmentStatusList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrderFulfillmentStatusListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **action** | **string** | Available statuses for the specified action. | 

### Return type

[**OrderFulfillmentStatusList200Response**](OrderFulfillmentStatusList200Response.md)

### Authorization

[StoreKeyAuth](../README.md#StoreKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrderInfo

> OrderInfo200Response OrderInfo(ctx).OrderId(orderId).Id(id).Params(params).ResponseFields(responseFields).Exclude(exclude).StoreId(storeId).EnableCache(enableCache).UseLatestApiVersion(useLatestApiVersion).Execute()

order.info



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
	orderId := "25" // string | Retrieves order’s info specified by order id (optional)
	id := "10" // string | Retrieves order info specified by id (optional)
	params := "order_id,totals,status" // string | Set this parameter in order to choose which entity fields you want to retrieve (optional) (default to "order_id,customer,totals,address,items,bundles,status")
	responseFields := "{result{order_id,customer,totals,address,items,bundles,status}}" // string | Set this parameter in order to choose which entity fields you want to retrieve (optional)
	exclude := "order_id,totals,status" // string | Set this parameter in order to choose which entity fields you want to ignore. Works only if parameter `params` equal force_all (optional)
	storeId := "1" // string | Defines store id where the order should be found (optional)
	enableCache := true // bool | If the value is 'true' and order exist in our cache, we will return order.info response from cache (optional) (default to false)
	useLatestApiVersion := true // bool | Use the latest platform API version (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrderAPI.OrderInfo(context.Background()).OrderId(orderId).Id(id).Params(params).ResponseFields(responseFields).Exclude(exclude).StoreId(storeId).EnableCache(enableCache).UseLatestApiVersion(useLatestApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrderAPI.OrderInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrderInfo`: OrderInfo200Response
	fmt.Fprintf(os.Stdout, "Response from `OrderAPI.OrderInfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrderInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orderId** | **string** | Retrieves order’s info specified by order id | 
 **id** | **string** | Retrieves order info specified by id | 
 **params** | **string** | Set this parameter in order to choose which entity fields you want to retrieve | [default to &quot;order_id,customer,totals,address,items,bundles,status&quot;]
 **responseFields** | **string** | Set this parameter in order to choose which entity fields you want to retrieve | 
 **exclude** | **string** | Set this parameter in order to choose which entity fields you want to ignore. Works only if parameter &#x60;params&#x60; equal force_all | 
 **storeId** | **string** | Defines store id where the order should be found | 
 **enableCache** | **bool** | If the value is &#39;true&#39; and order exist in our cache, we will return order.info response from cache | [default to false]
 **useLatestApiVersion** | **bool** | Use the latest platform API version | [default to false]

### Return type

[**OrderInfo200Response**](OrderInfo200Response.md)

### Authorization

[StoreKeyAuth](../README.md#StoreKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrderList

> ModelResponseOrderList OrderList(ctx).CustomerId(customerId).CustomerEmail(customerEmail).Phone(phone).OrderStatus(orderStatus).OrderStatusIds(orderStatusIds).Start(start).Count(count).PageCursor(pageCursor).SortBy(sortBy).SortDirection(sortDirection).Params(params).ResponseFields(responseFields).Exclude(exclude).CreatedTo(createdTo).CreatedFrom(createdFrom).ModifiedTo(modifiedTo).ModifiedFrom(modifiedFrom).StoreId(storeId).Ids(ids).OrderIds(orderIds).EbayOrderStatus(ebayOrderStatus).BasketId(basketId).FinancialStatus(financialStatus).FinancialStatusIds(financialStatusIds).FulfillmentStatus(fulfillmentStatus).FulfillmentChannel(fulfillmentChannel).ShippingMethod(shippingMethod).SkipOrderIds(skipOrderIds).SinceId(sinceId).IsDeleted(isDeleted).ShippingCountryIso3(shippingCountryIso3).EnableCache(enableCache).DeliveryMethod(deliveryMethod).Tags(tags).ShipNodeType(shipNodeType).CurrencyId(currencyId).ReturnStatus(returnStatus).UseLatestApiVersion(useLatestApiVersion).Execute()

order.list



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
	customerId := "5" // string | Retrieves orders specified by customer id (optional)
	customerEmail := "jubari@hannsgroup.com" // string | Retrieves orders specified by customer email (optional)
	phone := "56686868654" // string | Filter orders by customer's phone number (optional)
	orderStatus := "Completed" // string | Retrieves orders specified by order status (optional)
	orderStatusIds := []string{"Inner_example"} // []string | Retrieves orders specified by order statuses (optional)
	start := int32(0) // int32 | This parameter sets the number from which you want to get entities (optional) (default to 0)
	count := int32(20) // int32 | This parameter sets the entity amount that has to be retrieved. Max allowed count=250 (optional) (default to 10)
	pageCursor := "pageCursor_example" // string | Used to retrieve orders via cursor-based pagination (it can't be used with any other filtering parameter) (optional)
	sortBy := "modified_at" // string | Set field to sort by (optional) (default to "order_id")
	sortDirection := "asc" // string | Set sorting direction (optional) (default to "asc")
	params := "order_id,totals,status" // string | Set this parameter in order to choose which entity fields you want to retrieve (optional) (default to "order_id,customer,totals,address,items,bundles,status")
	responseFields := "{return_code,pagination,result{order{order_id,customer,totals,address,items,bundles,status}}}" // string | Set this parameter in order to choose which entity fields you want to retrieve (optional)
	exclude := "order_id,totals,status" // string | Set this parameter in order to choose which entity fields you want to ignore. Works only if parameter `params` equal force_all (optional)
	createdTo := "2100-08-29 13:45:52" // string | Retrieve entities to their creation date (optional)
	createdFrom := "2010-07-29 13:45:52" // string | Retrieve entities from their creation date (optional)
	modifiedTo := "2100-08-29 13:45:52" // string | Retrieve entities to their modification date (optional)
	modifiedFrom := "2010-07-29 13:45:52" // string | Retrieve entities from their modification date (optional)
	storeId := "1" // string | Store Id (optional)
	ids := "24,25" // string | Retrieves orders specified by ids (optional)
	orderIds := "24,25" // string | Retrieves orders specified by order ids (optional)
	ebayOrderStatus := "Active" // string | Retrieves orders specified by order status (optional)
	basketId := "1" // string | Retrieves order’s info specified by basket id. (optional)
	financialStatus := "paid" // string | Retrieves orders specified by financial status (optional)
	financialStatusIds := []string{"Inner_example"} // []string | Retrieves orders specified by financial status ids (optional)
	fulfillmentStatus := "fulfilled" // string | Create order with fulfillment status (optional)
	fulfillmentChannel := "local" // string | Retrieves order with a fulfillment channel (optional)
	shippingMethod := "flatrate_flatrate" // string | Retrieve entities according to shipping method (optional)
	skipOrderIds := "24,25" // string | Skipped orders by ids (optional)
	sinceId := "56" // string | Retrieve entities starting from the specified id. (optional)
	isDeleted := true // bool | Filter deleted orders (optional)
	shippingCountryIso3 := "DEU" // string | Retrieve entities according to shipping country (optional)
	enableCache := true // bool | If the value is 'true', we will cache orders for a 15 minutes in order to increase speed and reduce requests throttling for some methods and shoping platforms (for example order.shipment.add) (optional) (default to false)
	deliveryMethod := "local" // string | Retrieves order with delivery method (optional)
	tags := "tag1,tag2" // string | Order tags (optional)
	shipNodeType := "SellerFulfilled" // string | Retrieves order with ship node type (optional)
	currencyId := "usd" // string | Currency Id (optional)
	returnStatus := "RETURNED" // string | Retrieves orders specified by return status (optional)
	useLatestApiVersion := true // bool | Use the latest platform API version (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrderAPI.OrderList(context.Background()).CustomerId(customerId).CustomerEmail(customerEmail).Phone(phone).OrderStatus(orderStatus).OrderStatusIds(orderStatusIds).Start(start).Count(count).PageCursor(pageCursor).SortBy(sortBy).SortDirection(sortDirection).Params(params).ResponseFields(responseFields).Exclude(exclude).CreatedTo(createdTo).CreatedFrom(createdFrom).ModifiedTo(modifiedTo).ModifiedFrom(modifiedFrom).StoreId(storeId).Ids(ids).OrderIds(orderIds).EbayOrderStatus(ebayOrderStatus).BasketId(basketId).FinancialStatus(financialStatus).FinancialStatusIds(financialStatusIds).FulfillmentStatus(fulfillmentStatus).FulfillmentChannel(fulfillmentChannel).ShippingMethod(shippingMethod).SkipOrderIds(skipOrderIds).SinceId(sinceId).IsDeleted(isDeleted).ShippingCountryIso3(shippingCountryIso3).EnableCache(enableCache).DeliveryMethod(deliveryMethod).Tags(tags).ShipNodeType(shipNodeType).CurrencyId(currencyId).ReturnStatus(returnStatus).UseLatestApiVersion(useLatestApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrderAPI.OrderList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrderList`: ModelResponseOrderList
	fmt.Fprintf(os.Stdout, "Response from `OrderAPI.OrderList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrderListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **customerId** | **string** | Retrieves orders specified by customer id | 
 **customerEmail** | **string** | Retrieves orders specified by customer email | 
 **phone** | **string** | Filter orders by customer&#39;s phone number | 
 **orderStatus** | **string** | Retrieves orders specified by order status | 
 **orderStatusIds** | **[]string** | Retrieves orders specified by order statuses | 
 **start** | **int32** | This parameter sets the number from which you want to get entities | [default to 0]
 **count** | **int32** | This parameter sets the entity amount that has to be retrieved. Max allowed count&#x3D;250 | [default to 10]
 **pageCursor** | **string** | Used to retrieve orders via cursor-based pagination (it can&#39;t be used with any other filtering parameter) | 
 **sortBy** | **string** | Set field to sort by | [default to &quot;order_id&quot;]
 **sortDirection** | **string** | Set sorting direction | [default to &quot;asc&quot;]
 **params** | **string** | Set this parameter in order to choose which entity fields you want to retrieve | [default to &quot;order_id,customer,totals,address,items,bundles,status&quot;]
 **responseFields** | **string** | Set this parameter in order to choose which entity fields you want to retrieve | 
 **exclude** | **string** | Set this parameter in order to choose which entity fields you want to ignore. Works only if parameter &#x60;params&#x60; equal force_all | 
 **createdTo** | **string** | Retrieve entities to their creation date | 
 **createdFrom** | **string** | Retrieve entities from their creation date | 
 **modifiedTo** | **string** | Retrieve entities to their modification date | 
 **modifiedFrom** | **string** | Retrieve entities from their modification date | 
 **storeId** | **string** | Store Id | 
 **ids** | **string** | Retrieves orders specified by ids | 
 **orderIds** | **string** | Retrieves orders specified by order ids | 
 **ebayOrderStatus** | **string** | Retrieves orders specified by order status | 
 **basketId** | **string** | Retrieves order’s info specified by basket id. | 
 **financialStatus** | **string** | Retrieves orders specified by financial status | 
 **financialStatusIds** | **[]string** | Retrieves orders specified by financial status ids | 
 **fulfillmentStatus** | **string** | Create order with fulfillment status | 
 **fulfillmentChannel** | **string** | Retrieves order with a fulfillment channel | 
 **shippingMethod** | **string** | Retrieve entities according to shipping method | 
 **skipOrderIds** | **string** | Skipped orders by ids | 
 **sinceId** | **string** | Retrieve entities starting from the specified id. | 
 **isDeleted** | **bool** | Filter deleted orders | 
 **shippingCountryIso3** | **string** | Retrieve entities according to shipping country | 
 **enableCache** | **bool** | If the value is &#39;true&#39;, we will cache orders for a 15 minutes in order to increase speed and reduce requests throttling for some methods and shoping platforms (for example order.shipment.add) | [default to false]
 **deliveryMethod** | **string** | Retrieves order with delivery method | 
 **tags** | **string** | Order tags | 
 **shipNodeType** | **string** | Retrieves order with ship node type | 
 **currencyId** | **string** | Currency Id | 
 **returnStatus** | **string** | Retrieves orders specified by return status | 
 **useLatestApiVersion** | **bool** | Use the latest platform API version | [default to false]

### Return type

[**ModelResponseOrderList**](ModelResponseOrderList.md)

### Authorization

[StoreKeyAuth](../README.md#StoreKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrderPreestimateShippingList

> ModelResponseOrderPreestimateShippingList OrderPreestimateShippingList(ctx).OrderPreestimateShippingList(orderPreestimateShippingList).Execute()

order.preestimate_shipping.list



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
	orderPreestimateShippingList := *openapiclient.NewOrderPreestimateShippingList("US", []openapiclient.OrderPreestimateShippingListOrderItemInner{*openapiclient.NewOrderPreestimateShippingListOrderItemInner("125, where {x} - 1,2,3,... etc", int32(5, where {x} - 1,2,3,... etc))}) // OrderPreestimateShippingList | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrderAPI.OrderPreestimateShippingList(context.Background()).OrderPreestimateShippingList(orderPreestimateShippingList).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrderAPI.OrderPreestimateShippingList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrderPreestimateShippingList`: ModelResponseOrderPreestimateShippingList
	fmt.Fprintf(os.Stdout, "Response from `OrderAPI.OrderPreestimateShippingList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrderPreestimateShippingListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orderPreestimateShippingList** | [**OrderPreestimateShippingList**](OrderPreestimateShippingList.md) |  | 

### Return type

[**ModelResponseOrderPreestimateShippingList**](ModelResponseOrderPreestimateShippingList.md)

### Authorization

[StoreKeyAuth](../README.md#StoreKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrderRefundAdd

> OrderRefundAdd200Response OrderRefundAdd(ctx).OrderRefundAdd(orderRefundAdd).Execute()

order.refund.add



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
	orderRefundAdd := *openapiclient.NewOrderRefundAdd() // OrderRefundAdd | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrderAPI.OrderRefundAdd(context.Background()).OrderRefundAdd(orderRefundAdd).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrderAPI.OrderRefundAdd``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrderRefundAdd`: OrderRefundAdd200Response
	fmt.Fprintf(os.Stdout, "Response from `OrderAPI.OrderRefundAdd`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrderRefundAddRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orderRefundAdd** | [**OrderRefundAdd**](OrderRefundAdd.md) |  | 

### Return type

[**OrderRefundAdd200Response**](OrderRefundAdd200Response.md)

### Authorization

[StoreKeyAuth](../README.md#StoreKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrderReturnAdd

> OrderReturnAdd200Response OrderReturnAdd(ctx).OrderReturnAdd(orderReturnAdd).Execute()

order.return.add



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
	orderReturnAdd := *openapiclient.NewOrderReturnAdd("RETURNED", "RETURNED", "broken", []openapiclient.OrderReturnAddOrderProductsInner{*openapiclient.NewOrderReturnAddOrderProductsInner("125, where {x} - 1,2,3,... etc", int32(1, where {x} - 1,2,3,... etc), "DEFECTIVE, where {x} - 1,2,3,... etc", "REFUND, where {x} - 1,2,3,... etc")}) // OrderReturnAdd | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrderAPI.OrderReturnAdd(context.Background()).OrderReturnAdd(orderReturnAdd).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrderAPI.OrderReturnAdd``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrderReturnAdd`: OrderReturnAdd200Response
	fmt.Fprintf(os.Stdout, "Response from `OrderAPI.OrderReturnAdd`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrderReturnAddRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orderReturnAdd** | [**OrderReturnAdd**](OrderReturnAdd.md) |  | 

### Return type

[**OrderReturnAdd200Response**](OrderReturnAdd200Response.md)

### Authorization

[StoreKeyAuth](../README.md#StoreKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrderReturnDelete

> AttributeValueDelete200Response OrderReturnDelete(ctx).ReturnId(returnId).OrderId(orderId).StoreId(storeId).Execute()

order.return.delete



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
	returnId := "200000002" // string | Return ID
	orderId := "25" // string | Defines the order id
	storeId := "1" // string | Store Id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrderAPI.OrderReturnDelete(context.Background()).ReturnId(returnId).OrderId(orderId).StoreId(storeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrderAPI.OrderReturnDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrderReturnDelete`: AttributeValueDelete200Response
	fmt.Fprintf(os.Stdout, "Response from `OrderAPI.OrderReturnDelete`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrderReturnDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **returnId** | **string** | Return ID | 
 **orderId** | **string** | Defines the order id | 
 **storeId** | **string** | Store Id | 

### Return type

[**AttributeValueDelete200Response**](AttributeValueDelete200Response.md)

### Authorization

[StoreKeyAuth](../README.md#StoreKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrderReturnUpdate

> AccountConfigUpdate200Response OrderReturnUpdate(ctx).OrderReturnUpdate(orderReturnUpdate).Execute()

order.return.update



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
	orderReturnUpdate := *openapiclient.NewOrderReturnUpdate("200000002", []openapiclient.OrderReturnUpdateOrderProductsInner{*openapiclient.NewOrderReturnUpdateOrderProductsInner("125, where {x} - 1,2,3,... etc", int32(1, where {x} - 1,2,3,... etc), "REFUND, where {x} - 1,2,3,... etc")}) // OrderReturnUpdate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrderAPI.OrderReturnUpdate(context.Background()).OrderReturnUpdate(orderReturnUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrderAPI.OrderReturnUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrderReturnUpdate`: AccountConfigUpdate200Response
	fmt.Fprintf(os.Stdout, "Response from `OrderAPI.OrderReturnUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrderReturnUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orderReturnUpdate** | [**OrderReturnUpdate**](OrderReturnUpdate.md) |  | 

### Return type

[**AccountConfigUpdate200Response**](AccountConfigUpdate200Response.md)

### Authorization

[StoreKeyAuth](../README.md#StoreKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrderShipmentAdd

> OrderShipmentAdd200Response OrderShipmentAdd(ctx).OrderShipmentAdd(orderShipmentAdd).Execute()

order.shipment.add



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
	orderShipmentAdd := *openapiclient.NewOrderShipmentAdd() // OrderShipmentAdd | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrderAPI.OrderShipmentAdd(context.Background()).OrderShipmentAdd(orderShipmentAdd).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrderAPI.OrderShipmentAdd``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrderShipmentAdd`: OrderShipmentAdd200Response
	fmt.Fprintf(os.Stdout, "Response from `OrderAPI.OrderShipmentAdd`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrderShipmentAddRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orderShipmentAdd** | [**OrderShipmentAdd**](OrderShipmentAdd.md) |  | 

### Return type

[**OrderShipmentAdd200Response**](OrderShipmentAdd200Response.md)

### Authorization

[StoreKeyAuth](../README.md#StoreKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrderShipmentAddBatch

> CategoryAddBatch200Response OrderShipmentAddBatch(ctx).OrderShipmentAddBatch(orderShipmentAddBatch).Execute()

order.shipment.add.batch



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
	orderShipmentAddBatch := *openapiclient.NewOrderShipmentAddBatch([]openapiclient.OrderShipmentAddBatchPayloadInner{*openapiclient.NewOrderShipmentAddBatchPayloadInner("OrderId_example", "TrackingNumber_example")}) // OrderShipmentAddBatch | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrderAPI.OrderShipmentAddBatch(context.Background()).OrderShipmentAddBatch(orderShipmentAddBatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrderAPI.OrderShipmentAddBatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrderShipmentAddBatch`: CategoryAddBatch200Response
	fmt.Fprintf(os.Stdout, "Response from `OrderAPI.OrderShipmentAddBatch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrderShipmentAddBatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orderShipmentAddBatch** | [**OrderShipmentAddBatch**](OrderShipmentAddBatch.md) |  | 

### Return type

[**CategoryAddBatch200Response**](CategoryAddBatch200Response.md)

### Authorization

[StoreKeyAuth](../README.md#StoreKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrderShipmentDelete

> OrderShipmentDelete200Response OrderShipmentDelete(ctx).ShipmentId(shipmentId).OrderId(orderId).StoreId(storeId).Execute()

order.shipment.delete



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
	shipmentId := "200000002" // string | Shipment id indicates the number of delivery
	orderId := "25" // string | Defines the order for which the shipment will be deleted
	storeId := "1" // string | Store Id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrderAPI.OrderShipmentDelete(context.Background()).ShipmentId(shipmentId).OrderId(orderId).StoreId(storeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrderAPI.OrderShipmentDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrderShipmentDelete`: OrderShipmentDelete200Response
	fmt.Fprintf(os.Stdout, "Response from `OrderAPI.OrderShipmentDelete`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrderShipmentDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **shipmentId** | **string** | Shipment id indicates the number of delivery | 
 **orderId** | **string** | Defines the order for which the shipment will be deleted | 
 **storeId** | **string** | Store Id | 

### Return type

[**OrderShipmentDelete200Response**](OrderShipmentDelete200Response.md)

### Authorization

[StoreKeyAuth](../README.md#StoreKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrderShipmentInfo

> OrderShipmentInfo200Response OrderShipmentInfo(ctx).Id(id).OrderId(orderId).Start(start).Params(params).ResponseFields(responseFields).Exclude(exclude).StoreId(storeId).Execute()

order.shipment.info



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
	orderId := "25" // string | Defines the order id
	start := int32(0) // int32 | This parameter sets the number from which you want to get entities (optional) (default to 0)
	params := "id,model,price,images" // string | Set this parameter in order to choose which entity fields you want to retrieve (optional) (default to "id,order_id,items,tracking_numbers")
	responseFields := "{result{id,order_id,shipment_provider,tracking_numbers{tracking_number},items{product_id,quantity}}}" // string | Set this parameter in order to choose which entity fields you want to retrieve (optional)
	exclude := "false" // string | Set this parameter in order to choose which entity fields you want to ignore. Works only if parameter `params` equal force_all (optional)
	storeId := "1" // string | Store Id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrderAPI.OrderShipmentInfo(context.Background()).Id(id).OrderId(orderId).Start(start).Params(params).ResponseFields(responseFields).Exclude(exclude).StoreId(storeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrderAPI.OrderShipmentInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrderShipmentInfo`: OrderShipmentInfo200Response
	fmt.Fprintf(os.Stdout, "Response from `OrderAPI.OrderShipmentInfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrderShipmentInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | Entity id | 
 **orderId** | **string** | Defines the order id | 
 **start** | **int32** | This parameter sets the number from which you want to get entities | [default to 0]
 **params** | **string** | Set this parameter in order to choose which entity fields you want to retrieve | [default to &quot;id,order_id,items,tracking_numbers&quot;]
 **responseFields** | **string** | Set this parameter in order to choose which entity fields you want to retrieve | 
 **exclude** | **string** | Set this parameter in order to choose which entity fields you want to ignore. Works only if parameter &#x60;params&#x60; equal force_all | 
 **storeId** | **string** | Store Id | 

### Return type

[**OrderShipmentInfo200Response**](OrderShipmentInfo200Response.md)

### Authorization

[StoreKeyAuth](../README.md#StoreKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrderShipmentList

> ModelResponseOrderShipmentList OrderShipmentList(ctx).OrderId(orderId).PageCursor(pageCursor).Start(start).Count(count).Params(params).ResponseFields(responseFields).Exclude(exclude).CreatedFrom(createdFrom).CreatedTo(createdTo).ModifiedFrom(modifiedFrom).ModifiedTo(modifiedTo).StoreId(storeId).Execute()

order.shipment.list



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
	orderId := "25" // string | Retrieves shipments specified by order id
	pageCursor := "pageCursor_example" // string | Used to retrieve entities via cursor-based pagination (it can't be used with any other filtering parameter) (optional)
	start := int32(0) // int32 | This parameter sets the number from which you want to get entities (optional) (default to 0)
	count := int32(20) // int32 | This parameter sets the entity amount that has to be retrieved. Max allowed count=250 (optional) (default to 10)
	params := "id,model,price,images" // string | Set this parameter in order to choose which entity fields you want to retrieve (optional) (default to "id,order_id,items,tracking_numbers")
	responseFields := "{status_code,pagination,result{shipment{id,order_id,shipment_provider,tracking_numbers{tracking_number},items{product_id,quantity}}}}" // string | Set this parameter in order to choose which entity fields you want to retrieve (optional)
	exclude := "false" // string | Set this parameter in order to choose which entity fields you want to ignore. Works only if parameter `params` equal force_all (optional)
	createdFrom := "2010-07-29 13:45:52" // string | Retrieve entities from their creation date (optional)
	createdTo := "2100-08-29 13:45:52" // string | Retrieve entities to their creation date (optional)
	modifiedFrom := "2010-07-29 13:45:52" // string | Retrieve entities from their modification date (optional)
	modifiedTo := "2100-08-29 13:45:52" // string | Retrieve entities to their modification date (optional)
	storeId := "1" // string | Store Id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrderAPI.OrderShipmentList(context.Background()).OrderId(orderId).PageCursor(pageCursor).Start(start).Count(count).Params(params).ResponseFields(responseFields).Exclude(exclude).CreatedFrom(createdFrom).CreatedTo(createdTo).ModifiedFrom(modifiedFrom).ModifiedTo(modifiedTo).StoreId(storeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrderAPI.OrderShipmentList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrderShipmentList`: ModelResponseOrderShipmentList
	fmt.Fprintf(os.Stdout, "Response from `OrderAPI.OrderShipmentList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrderShipmentListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orderId** | **string** | Retrieves shipments specified by order id | 
 **pageCursor** | **string** | Used to retrieve entities via cursor-based pagination (it can&#39;t be used with any other filtering parameter) | 
 **start** | **int32** | This parameter sets the number from which you want to get entities | [default to 0]
 **count** | **int32** | This parameter sets the entity amount that has to be retrieved. Max allowed count&#x3D;250 | [default to 10]
 **params** | **string** | Set this parameter in order to choose which entity fields you want to retrieve | [default to &quot;id,order_id,items,tracking_numbers&quot;]
 **responseFields** | **string** | Set this parameter in order to choose which entity fields you want to retrieve | 
 **exclude** | **string** | Set this parameter in order to choose which entity fields you want to ignore. Works only if parameter &#x60;params&#x60; equal force_all | 
 **createdFrom** | **string** | Retrieve entities from their creation date | 
 **createdTo** | **string** | Retrieve entities to their creation date | 
 **modifiedFrom** | **string** | Retrieve entities from their modification date | 
 **modifiedTo** | **string** | Retrieve entities to their modification date | 
 **storeId** | **string** | Store Id | 

### Return type

[**ModelResponseOrderShipmentList**](ModelResponseOrderShipmentList.md)

### Authorization

[StoreKeyAuth](../README.md#StoreKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrderShipmentTrackingAdd

> OrderShipmentTrackingAdd200Response OrderShipmentTrackingAdd(ctx).OrderShipmentTrackingAdd(orderShipmentTrackingAdd).Execute()

order.shipment.tracking.add



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
	orderShipmentTrackingAdd := *openapiclient.NewOrderShipmentTrackingAdd("200000002", "1А6745") // OrderShipmentTrackingAdd | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrderAPI.OrderShipmentTrackingAdd(context.Background()).OrderShipmentTrackingAdd(orderShipmentTrackingAdd).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrderAPI.OrderShipmentTrackingAdd``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrderShipmentTrackingAdd`: OrderShipmentTrackingAdd200Response
	fmt.Fprintf(os.Stdout, "Response from `OrderAPI.OrderShipmentTrackingAdd`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrderShipmentTrackingAddRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orderShipmentTrackingAdd** | [**OrderShipmentTrackingAdd**](OrderShipmentTrackingAdd.md) |  | 

### Return type

[**OrderShipmentTrackingAdd200Response**](OrderShipmentTrackingAdd200Response.md)

### Authorization

[StoreKeyAuth](../README.md#StoreKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrderShipmentUpdate

> AccountConfigUpdate200Response OrderShipmentUpdate(ctx).OrderShipmentUpdate(orderShipmentUpdate).Execute()

order.shipment.update



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
	orderShipmentUpdate := *openapiclient.NewOrderShipmentUpdate("200000002") // OrderShipmentUpdate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrderAPI.OrderShipmentUpdate(context.Background()).OrderShipmentUpdate(orderShipmentUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrderAPI.OrderShipmentUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrderShipmentUpdate`: AccountConfigUpdate200Response
	fmt.Fprintf(os.Stdout, "Response from `OrderAPI.OrderShipmentUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrderShipmentUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orderShipmentUpdate** | [**OrderShipmentUpdate**](OrderShipmentUpdate.md) |  | 

### Return type

[**AccountConfigUpdate200Response**](AccountConfigUpdate200Response.md)

### Authorization

[StoreKeyAuth](../README.md#StoreKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrderStatusList

> ModelResponseOrderStatusList OrderStatusList(ctx).StoreId(storeId).Action(action).ResponseFields(responseFields).Execute()

order.status.list



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
	action := "add" // string | Available statuses for the specified action. (optional)
	responseFields := "{return_code,return_message,result}" // string | Set this parameter in order to choose which entity fields you want to retrieve (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrderAPI.OrderStatusList(context.Background()).StoreId(storeId).Action(action).ResponseFields(responseFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrderAPI.OrderStatusList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrderStatusList`: ModelResponseOrderStatusList
	fmt.Fprintf(os.Stdout, "Response from `OrderAPI.OrderStatusList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrderStatusListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **storeId** | **string** | Store Id | 
 **action** | **string** | Available statuses for the specified action. | 
 **responseFields** | **string** | Set this parameter in order to choose which entity fields you want to retrieve | 

### Return type

[**ModelResponseOrderStatusList**](ModelResponseOrderStatusList.md)

### Authorization

[StoreKeyAuth](../README.md#StoreKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrderTransactionList

> ModelResponseOrderTransactionList OrderTransactionList(ctx).OrderIds(orderIds).Count(count).StoreId(storeId).Params(params).ResponseFields(responseFields).Exclude(exclude).PageCursor(pageCursor).Execute()

order.transaction.list



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
	orderIds := "24,25" // string | Retrieves order transactions specified by order ids
	count := int32(20) // int32 | This parameter sets the entity amount that has to be retrieved. Max allowed count=250 (optional) (default to 10)
	storeId := "1" // string | Store Id (optional)
	params := "id,model,price,images" // string | Set this parameter in order to choose which entity fields you want to retrieve (optional) (default to "id,order_id,amount,description")
	responseFields := "{return_code,pagination,result{transactions_count,transactions{id,transaction_id,status,description,settlement_amount,gateway,card_brand,card_last_four}}}" // string | Set this parameter in order to choose which entity fields you want to retrieve (optional)
	exclude := "false" // string | Set this parameter in order to choose which entity fields you want to ignore. Works only if parameter `params` equal force_all (optional)
	pageCursor := "pageCursor_example" // string | Used to retrieve entities via cursor-based pagination (it can't be used with any other filtering parameter) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrderAPI.OrderTransactionList(context.Background()).OrderIds(orderIds).Count(count).StoreId(storeId).Params(params).ResponseFields(responseFields).Exclude(exclude).PageCursor(pageCursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrderAPI.OrderTransactionList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrderTransactionList`: ModelResponseOrderTransactionList
	fmt.Fprintf(os.Stdout, "Response from `OrderAPI.OrderTransactionList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrderTransactionListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orderIds** | **string** | Retrieves order transactions specified by order ids | 
 **count** | **int32** | This parameter sets the entity amount that has to be retrieved. Max allowed count&#x3D;250 | [default to 10]
 **storeId** | **string** | Store Id | 
 **params** | **string** | Set this parameter in order to choose which entity fields you want to retrieve | [default to &quot;id,order_id,amount,description&quot;]
 **responseFields** | **string** | Set this parameter in order to choose which entity fields you want to retrieve | 
 **exclude** | **string** | Set this parameter in order to choose which entity fields you want to ignore. Works only if parameter &#x60;params&#x60; equal force_all | 
 **pageCursor** | **string** | Used to retrieve entities via cursor-based pagination (it can&#39;t be used with any other filtering parameter) | 

### Return type

[**ModelResponseOrderTransactionList**](ModelResponseOrderTransactionList.md)

### Authorization

[StoreKeyAuth](../README.md#StoreKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrderUpdate

> AccountConfigUpdate200Response OrderUpdate(ctx).OrderId(orderId).StoreId(storeId).OrderStatus(orderStatus).CancellationReason(cancellationReason).Comment(comment).AdminComment(adminComment).AdminPrivateComment(adminPrivateComment).DateModified(dateModified).DateFinished(dateFinished).FinancialStatus(financialStatus).FulfillmentStatus(fulfillmentStatus).OrderPaymentMethod(orderPaymentMethod).SendNotifications(sendNotifications).Origin(origin).CreateInvoice(createInvoice).InvoiceAdminComment(invoiceAdminComment).Execute()

order.update



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
	orderId := "25" // string | Defines the orders specified by order id
	storeId := "1" // string | Defines store id where the order should be found (optional)
	orderStatus := "Completed" // string | Defines new order's status (optional)
	cancellationReason := "ORDER_UNPAID" // string | Defines the cancellation reason when the order will be canceled (optional)
	comment := "This coole order" // string | Specifies order comment (optional)
	adminComment := "Test admin comment" // string | Specifies admin's order comment (optional)
	adminPrivateComment := "Test admin private comment" // string | Specifies private admin's order comment (optional)
	dateModified := "2014-05-05 05:05:00" // string | Specifies order's  modification date (optional)
	dateFinished := "2014-06-05 05:05:00" // string | Specifies order's  finished date (optional)
	financialStatus := "paid" // string | Update order financial status to specified (optional)
	fulfillmentStatus := "fulfilled" // string | Create order with fulfillment status (optional)
	orderPaymentMethod := "PayPal" // string | Defines order payment method.<br/>Setting order_payment_method on Shopify will also change financial_status field value to 'paid' (optional)
	sendNotifications := true // bool | Send notifications to customer after order was created (optional) (default to false)
	origin := "newsletter" // string | The source of the order (optional)
	createInvoice := true // bool | Determines whether an invoice should be created if it has not already been created (optional)
	invoiceAdminComment := "Test admin comment" // string | Specifies admin's order invoice comment (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrderAPI.OrderUpdate(context.Background()).OrderId(orderId).StoreId(storeId).OrderStatus(orderStatus).CancellationReason(cancellationReason).Comment(comment).AdminComment(adminComment).AdminPrivateComment(adminPrivateComment).DateModified(dateModified).DateFinished(dateFinished).FinancialStatus(financialStatus).FulfillmentStatus(fulfillmentStatus).OrderPaymentMethod(orderPaymentMethod).SendNotifications(sendNotifications).Origin(origin).CreateInvoice(createInvoice).InvoiceAdminComment(invoiceAdminComment).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrderAPI.OrderUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrderUpdate`: AccountConfigUpdate200Response
	fmt.Fprintf(os.Stdout, "Response from `OrderAPI.OrderUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrderUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orderId** | **string** | Defines the orders specified by order id | 
 **storeId** | **string** | Defines store id where the order should be found | 
 **orderStatus** | **string** | Defines new order&#39;s status | 
 **cancellationReason** | **string** | Defines the cancellation reason when the order will be canceled | 
 **comment** | **string** | Specifies order comment | 
 **adminComment** | **string** | Specifies admin&#39;s order comment | 
 **adminPrivateComment** | **string** | Specifies private admin&#39;s order comment | 
 **dateModified** | **string** | Specifies order&#39;s  modification date | 
 **dateFinished** | **string** | Specifies order&#39;s  finished date | 
 **financialStatus** | **string** | Update order financial status to specified | 
 **fulfillmentStatus** | **string** | Create order with fulfillment status | 
 **orderPaymentMethod** | **string** | Defines order payment method.&lt;br/&gt;Setting order_payment_method on Shopify will also change financial_status field value to &#39;paid&#39; | 
 **sendNotifications** | **bool** | Send notifications to customer after order was created | [default to false]
 **origin** | **string** | The source of the order | 
 **createInvoice** | **bool** | Determines whether an invoice should be created if it has not already been created | 
 **invoiceAdminComment** | **string** | Specifies admin&#39;s order invoice comment | 

### Return type

[**AccountConfigUpdate200Response**](AccountConfigUpdate200Response.md)

### Authorization

[StoreKeyAuth](../README.md#StoreKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

