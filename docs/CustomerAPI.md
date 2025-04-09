# \CustomerAPI

All URIs are relative to *https://api.api2cart.com/v1.1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CustomerAdd**](CustomerAPI.md#CustomerAdd) | **Post** /customer.add.json | customer.add
[**CustomerAddressAdd**](CustomerAPI.md#CustomerAddressAdd) | **Post** /customer.address.add.json | customer.address.add
[**CustomerAttributeList**](CustomerAPI.md#CustomerAttributeList) | **Get** /customer.attribute.list.json | customer.attribute.list
[**CustomerCount**](CustomerAPI.md#CustomerCount) | **Get** /customer.count.json | customer.count
[**CustomerDelete**](CustomerAPI.md#CustomerDelete) | **Delete** /customer.delete.json | customer.delete
[**CustomerFind**](CustomerAPI.md#CustomerFind) | **Get** /customer.find.json | customer.find
[**CustomerGroupAdd**](CustomerAPI.md#CustomerGroupAdd) | **Post** /customer.group.add.json | customer.group.add
[**CustomerGroupList**](CustomerAPI.md#CustomerGroupList) | **Get** /customer.group.list.json | customer.group.list
[**CustomerInfo**](CustomerAPI.md#CustomerInfo) | **Get** /customer.info.json | customer.info
[**CustomerList**](CustomerAPI.md#CustomerList) | **Get** /customer.list.json | customer.list
[**CustomerUpdate**](CustomerAPI.md#CustomerUpdate) | **Put** /customer.update.json | customer.update
[**CustomerWishlistList**](CustomerAPI.md#CustomerWishlistList) | **Get** /customer.wishlist.list.json | customer.wishlist.list



## CustomerAdd

> CustomerAdd200Response CustomerAdd(ctx).CustomerAdd(customerAdd).Execute()

customer.add



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
	customerAdd := *openapiclient.NewCustomerAdd("mail@example.com", "John", "Smith") // CustomerAdd | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomerAPI.CustomerAdd(context.Background()).CustomerAdd(customerAdd).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomerAPI.CustomerAdd``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CustomerAdd`: CustomerAdd200Response
	fmt.Fprintf(os.Stdout, "Response from `CustomerAPI.CustomerAdd`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCustomerAddRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **customerAdd** | [**CustomerAdd**](CustomerAdd.md) |  | 

### Return type

[**CustomerAdd200Response**](CustomerAdd200Response.md)

### Authorization

[StoreKeyAuth](../README.md#StoreKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomerAddressAdd

> AttributeAdd200Response CustomerAddressAdd(ctx).CustomerAddressAdd(customerAddressAdd).Execute()

customer.address.add



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
	customerAddressAdd := *openapiclient.NewCustomerAddressAdd("5", "Green str. 35", "Chicago", "US", "12345") // CustomerAddressAdd | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomerAPI.CustomerAddressAdd(context.Background()).CustomerAddressAdd(customerAddressAdd).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomerAPI.CustomerAddressAdd``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CustomerAddressAdd`: AttributeAdd200Response
	fmt.Fprintf(os.Stdout, "Response from `CustomerAPI.CustomerAddressAdd`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCustomerAddressAddRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **customerAddressAdd** | [**CustomerAddressAdd**](CustomerAddressAdd.md) |  | 

### Return type

[**AttributeAdd200Response**](AttributeAdd200Response.md)

### Authorization

[StoreKeyAuth](../README.md#StoreKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomerAttributeList

> ModelResponseCustomerAttributeList CustomerAttributeList(ctx).CustomerId(customerId).Count(count).PageCursor(pageCursor).StoreId(storeId).LangId(langId).Params(params).Exclude(exclude).ResponseFields(responseFields).Execute()

customer.attribute.list



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
	count := int32(20) // int32 | This parameter sets the entity amount that has to be retrieved. Max allowed count=250 (optional) (default to 10)
	pageCursor := "pageCursor_example" // string | Used to retrieve entities via cursor-based pagination (it can't be used with any other filtering parameter) (optional)
	storeId := "1" // string | Store Id (optional)
	langId := "3" // string | Language id (optional)
	params := "id,model,price,images" // string | Set this parameter in order to choose which entity fields you want to retrieve (optional) (default to "force_all")
	exclude := "false" // string | Set this parameter in order to choose which entity fields you want to ignore. Works only if parameter `params` equal force_all (optional)
	responseFields := "{return_code,return_message,pagination,result}" // string | Set this parameter in order to choose which entity fields you want to retrieve (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomerAPI.CustomerAttributeList(context.Background()).CustomerId(customerId).Count(count).PageCursor(pageCursor).StoreId(storeId).LangId(langId).Params(params).Exclude(exclude).ResponseFields(responseFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomerAPI.CustomerAttributeList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CustomerAttributeList`: ModelResponseCustomerAttributeList
	fmt.Fprintf(os.Stdout, "Response from `CustomerAPI.CustomerAttributeList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCustomerAttributeListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **customerId** | **string** | Retrieves orders specified by customer id | 
 **count** | **int32** | This parameter sets the entity amount that has to be retrieved. Max allowed count&#x3D;250 | [default to 10]
 **pageCursor** | **string** | Used to retrieve entities via cursor-based pagination (it can&#39;t be used with any other filtering parameter) | 
 **storeId** | **string** | Store Id | 
 **langId** | **string** | Language id | 
 **params** | **string** | Set this parameter in order to choose which entity fields you want to retrieve | [default to &quot;force_all&quot;]
 **exclude** | **string** | Set this parameter in order to choose which entity fields you want to ignore. Works only if parameter &#x60;params&#x60; equal force_all | 
 **responseFields** | **string** | Set this parameter in order to choose which entity fields you want to retrieve | 

### Return type

[**ModelResponseCustomerAttributeList**](ModelResponseCustomerAttributeList.md)

### Authorization

[StoreKeyAuth](../README.md#StoreKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomerCount

> CustomerCount200Response CustomerCount(ctx).GroupId(groupId).CreatedFrom(createdFrom).CreatedTo(createdTo).ModifiedFrom(modifiedFrom).ModifiedTo(modifiedTo).StoreId(storeId).CustomerListId(customerListId).Avail(avail).FindValue(findValue).FindWhere(findWhere).Ids(ids).SinceId(sinceId).Execute()

customer.count



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
	groupId := "3" // string | Customer group_id (optional)
	createdFrom := "2010-07-29 13:45:52" // string | Retrieve entities from their creation date (optional)
	createdTo := "2100-08-29 13:45:52" // string | Retrieve entities to their creation date (optional)
	modifiedFrom := "2010-07-29 13:45:52" // string | Retrieve entities from their modification date (optional)
	modifiedTo := "2100-08-29 13:45:52" // string | Retrieve entities to their modification date (optional)
	storeId := "1" // string | Counts customer specified by store id (optional)
	customerListId := "exampleListId" // string | The numeric ID of the customer list in Demandware. (optional)
	avail := false // bool | Defines category's visibility status (optional) (default to true)
	findValue := "mail@gmail.com" // string | Entity search that is specified by some value (optional)
	findWhere := "email" // string | Counts customers that are searched specified by field (optional)
	ids := "24,25" // string | Counts customers specified by ids (optional)
	sinceId := "56" // string | Retrieve entities starting from the specified id. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomerAPI.CustomerCount(context.Background()).GroupId(groupId).CreatedFrom(createdFrom).CreatedTo(createdTo).ModifiedFrom(modifiedFrom).ModifiedTo(modifiedTo).StoreId(storeId).CustomerListId(customerListId).Avail(avail).FindValue(findValue).FindWhere(findWhere).Ids(ids).SinceId(sinceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomerAPI.CustomerCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CustomerCount`: CustomerCount200Response
	fmt.Fprintf(os.Stdout, "Response from `CustomerAPI.CustomerCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCustomerCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupId** | **string** | Customer group_id | 
 **createdFrom** | **string** | Retrieve entities from their creation date | 
 **createdTo** | **string** | Retrieve entities to their creation date | 
 **modifiedFrom** | **string** | Retrieve entities from their modification date | 
 **modifiedTo** | **string** | Retrieve entities to their modification date | 
 **storeId** | **string** | Counts customer specified by store id | 
 **customerListId** | **string** | The numeric ID of the customer list in Demandware. | 
 **avail** | **bool** | Defines category&#39;s visibility status | [default to true]
 **findValue** | **string** | Entity search that is specified by some value | 
 **findWhere** | **string** | Counts customers that are searched specified by field | 
 **ids** | **string** | Counts customers specified by ids | 
 **sinceId** | **string** | Retrieve entities starting from the specified id. | 

### Return type

[**CustomerCount200Response**](CustomerCount200Response.md)

### Authorization

[StoreKeyAuth](../README.md#StoreKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomerDelete

> CustomerDelete200Response CustomerDelete(ctx).Id(id).Execute()

customer.delete



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
	id := "10" // string | Identifies customer specified by the id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomerAPI.CustomerDelete(context.Background()).Id(id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomerAPI.CustomerDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CustomerDelete`: CustomerDelete200Response
	fmt.Fprintf(os.Stdout, "Response from `CustomerAPI.CustomerDelete`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCustomerDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | Identifies customer specified by the id | 

### Return type

[**CustomerDelete200Response**](CustomerDelete200Response.md)

### Authorization

[StoreKeyAuth](../README.md#StoreKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomerFind

> CustomerFind200Response CustomerFind(ctx).FindValue(findValue).FindWhere(findWhere).FindParams(findParams).StoreId(storeId).Execute()

customer.find



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
	findValue := "mail@gmail.com" // string | Entity search that is specified by some value
	findWhere := "email" // string | Entity search that is specified by the comma-separated unique fields (optional) (default to "email")
	findParams := "regex" // string | Entity search that is specified by comma-separated parameters (optional) (default to "whole_words")
	storeId := "1" // string | Store Id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomerAPI.CustomerFind(context.Background()).FindValue(findValue).FindWhere(findWhere).FindParams(findParams).StoreId(storeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomerAPI.CustomerFind``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CustomerFind`: CustomerFind200Response
	fmt.Fprintf(os.Stdout, "Response from `CustomerAPI.CustomerFind`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCustomerFindRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **findValue** | **string** | Entity search that is specified by some value | 
 **findWhere** | **string** | Entity search that is specified by the comma-separated unique fields | [default to &quot;email&quot;]
 **findParams** | **string** | Entity search that is specified by comma-separated parameters | [default to &quot;whole_words&quot;]
 **storeId** | **string** | Store Id | 

### Return type

[**CustomerFind200Response**](CustomerFind200Response.md)

### Authorization

[StoreKeyAuth](../README.md#StoreKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomerGroupAdd

> CustomerGroupAdd200Response CustomerGroupAdd(ctx).Name(name).StoreId(storeId).StoresIds(storesIds).Execute()

customer.group.add



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
	name := "new_group" // string | Customer group name
	storeId := "1" // string | Store Id (optional)
	storesIds := "1,2" // string | Assign customer group to the stores that is specified by comma-separated stores' id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomerAPI.CustomerGroupAdd(context.Background()).Name(name).StoreId(storeId).StoresIds(storesIds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomerAPI.CustomerGroupAdd``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CustomerGroupAdd`: CustomerGroupAdd200Response
	fmt.Fprintf(os.Stdout, "Response from `CustomerAPI.CustomerGroupAdd`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCustomerGroupAddRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | Customer group name | 
 **storeId** | **string** | Store Id | 
 **storesIds** | **string** | Assign customer group to the stores that is specified by comma-separated stores&#39; id | 

### Return type

[**CustomerGroupAdd200Response**](CustomerGroupAdd200Response.md)

### Authorization

[StoreKeyAuth](../README.md#StoreKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomerGroupList

> ModelResponseCustomerGroupList CustomerGroupList(ctx).DisableCache(disableCache).PageCursor(pageCursor).Start(start).Count(count).StoreId(storeId).LangId(langId).GroupIds(groupIds).Params(params).Exclude(exclude).ResponseFields(responseFields).Execute()

customer.group.list



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
	disableCache := false // bool | Disable cache for current request (optional) (default to false)
	pageCursor := "pageCursor_example" // string | Used to retrieve entities via cursor-based pagination (it can't be used with any other filtering parameter) (optional)
	start := int32(0) // int32 | This parameter sets the number from which you want to get entities (optional) (default to 0)
	count := int32(20) // int32 | This parameter sets the entity amount that has to be retrieved. Max allowed count=250 (optional) (default to 10)
	storeId := "1" // string | Store Id (optional)
	langId := "3" // string | Language id (optional)
	groupIds := "1,2,3" // string | Groups that will be assigned to a customer (optional)
	params := "id,model,price,images" // string | Set this parameter in order to choose which entity fields you want to retrieve (optional) (default to "id,name,additional_fields")
	exclude := "false" // string | Set this parameter in order to choose which entity fields you want to ignore. Works only if parameter `params` equal force_all (optional)
	responseFields := "{return_code,return_message,pagination,result}" // string | Set this parameter in order to choose which entity fields you want to retrieve (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomerAPI.CustomerGroupList(context.Background()).DisableCache(disableCache).PageCursor(pageCursor).Start(start).Count(count).StoreId(storeId).LangId(langId).GroupIds(groupIds).Params(params).Exclude(exclude).ResponseFields(responseFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomerAPI.CustomerGroupList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CustomerGroupList`: ModelResponseCustomerGroupList
	fmt.Fprintf(os.Stdout, "Response from `CustomerAPI.CustomerGroupList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCustomerGroupListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **disableCache** | **bool** | Disable cache for current request | [default to false]
 **pageCursor** | **string** | Used to retrieve entities via cursor-based pagination (it can&#39;t be used with any other filtering parameter) | 
 **start** | **int32** | This parameter sets the number from which you want to get entities | [default to 0]
 **count** | **int32** | This parameter sets the entity amount that has to be retrieved. Max allowed count&#x3D;250 | [default to 10]
 **storeId** | **string** | Store Id | 
 **langId** | **string** | Language id | 
 **groupIds** | **string** | Groups that will be assigned to a customer | 
 **params** | **string** | Set this parameter in order to choose which entity fields you want to retrieve | [default to &quot;id,name,additional_fields&quot;]
 **exclude** | **string** | Set this parameter in order to choose which entity fields you want to ignore. Works only if parameter &#x60;params&#x60; equal force_all | 
 **responseFields** | **string** | Set this parameter in order to choose which entity fields you want to retrieve | 

### Return type

[**ModelResponseCustomerGroupList**](ModelResponseCustomerGroupList.md)

### Authorization

[StoreKeyAuth](../README.md#StoreKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomerInfo

> CustomerInfo200Response CustomerInfo(ctx).Id(id).Params(params).ResponseFields(responseFields).Exclude(exclude).StoreId(storeId).Execute()

customer.info



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
	id := "10" // string | Retrieves customer's info specified by customer id
	params := "id,email" // string | Set this parameter in order to choose which entity fields you want to retrieve (optional) (default to "id,email,first_name,last_name")
	responseFields := "{result{id,parent_id,sku,upc,images,combination}}" // string | Set this parameter in order to choose which entity fields you want to retrieve (optional)
	exclude := "id,email" // string | Set this parameter in order to choose which entity fields you want to ignore. Works only if parameter `params` equal force_all (optional)
	storeId := "1" // string | Retrieves customer info specified by store id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomerAPI.CustomerInfo(context.Background()).Id(id).Params(params).ResponseFields(responseFields).Exclude(exclude).StoreId(storeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomerAPI.CustomerInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CustomerInfo`: CustomerInfo200Response
	fmt.Fprintf(os.Stdout, "Response from `CustomerAPI.CustomerInfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCustomerInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | Retrieves customer&#39;s info specified by customer id | 
 **params** | **string** | Set this parameter in order to choose which entity fields you want to retrieve | [default to &quot;id,email,first_name,last_name&quot;]
 **responseFields** | **string** | Set this parameter in order to choose which entity fields you want to retrieve | 
 **exclude** | **string** | Set this parameter in order to choose which entity fields you want to ignore. Works only if parameter &#x60;params&#x60; equal force_all | 
 **storeId** | **string** | Retrieves customer info specified by store id | 

### Return type

[**CustomerInfo200Response**](CustomerInfo200Response.md)

### Authorization

[StoreKeyAuth](../README.md#StoreKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomerList

> ModelResponseCustomerList CustomerList(ctx).PageCursor(pageCursor).Start(start).Count(count).CreatedFrom(createdFrom).CreatedTo(createdTo).ModifiedFrom(modifiedFrom).ModifiedTo(modifiedTo).Params(params).ResponseFields(responseFields).Exclude(exclude).GroupId(groupId).StoreId(storeId).CustomerListId(customerListId).Avail(avail).FindValue(findValue).FindWhere(findWhere).SortBy(sortBy).SortDirection(sortDirection).Ids(ids).SinceId(sinceId).Execute()

customer.list



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
	pageCursor := "pageCursor_example" // string | Used to retrieve entities via cursor-based pagination (it can't be used with any other filtering parameter) (optional)
	start := int32(0) // int32 | This parameter sets the number from which you want to get entities (optional) (default to 0)
	count := int32(20) // int32 | This parameter sets the entity amount that has to be retrieved. Max allowed count=250 (optional) (default to 10)
	createdFrom := "2010-07-29 13:45:52" // string | Retrieve entities from their creation date (optional)
	createdTo := "2100-08-29 13:45:52" // string | Retrieve entities to their creation date (optional)
	modifiedFrom := "2010-07-29 13:45:52" // string | Retrieve entities from their modification date (optional)
	modifiedTo := "2100-08-29 13:45:52" // string | Retrieve entities to their modification date (optional)
	params := "id,email" // string | Set this parameter in order to choose which entity fields you want to retrieve (optional) (default to "id,email,first_name,last_name")
	responseFields := "{result{customer}}" // string | Set this parameter in order to choose which entity fields you want to retrieve (optional)
	exclude := "id,email" // string | Set this parameter in order to choose which entity fields you want to ignore. Works only if parameter `params` equal force_all (optional)
	groupId := "3" // string | Customer group_id (optional)
	storeId := "1" // string | Retrieves customers specified by store id (optional)
	customerListId := "exampleListId" // string | The numeric ID of the customer list in Demandware. (optional)
	avail := false // bool | Defines category's visibility status (optional) (default to true)
	findValue := "mail@gmail.com" // string | Entity search that is specified by some value (optional)
	findWhere := "email" // string | Customer search that is specified by field (optional)
	sortBy := "value_id" // string | Set field to sort by (optional) (default to "created_time")
	sortDirection := "asc" // string | Set sorting direction (optional) (default to "asc")
	ids := "24,25" // string | Retrieves customers specified by ids (optional)
	sinceId := "56" // string | Retrieve entities starting from the specified id. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomerAPI.CustomerList(context.Background()).PageCursor(pageCursor).Start(start).Count(count).CreatedFrom(createdFrom).CreatedTo(createdTo).ModifiedFrom(modifiedFrom).ModifiedTo(modifiedTo).Params(params).ResponseFields(responseFields).Exclude(exclude).GroupId(groupId).StoreId(storeId).CustomerListId(customerListId).Avail(avail).FindValue(findValue).FindWhere(findWhere).SortBy(sortBy).SortDirection(sortDirection).Ids(ids).SinceId(sinceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomerAPI.CustomerList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CustomerList`: ModelResponseCustomerList
	fmt.Fprintf(os.Stdout, "Response from `CustomerAPI.CustomerList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCustomerListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageCursor** | **string** | Used to retrieve entities via cursor-based pagination (it can&#39;t be used with any other filtering parameter) | 
 **start** | **int32** | This parameter sets the number from which you want to get entities | [default to 0]
 **count** | **int32** | This parameter sets the entity amount that has to be retrieved. Max allowed count&#x3D;250 | [default to 10]
 **createdFrom** | **string** | Retrieve entities from their creation date | 
 **createdTo** | **string** | Retrieve entities to their creation date | 
 **modifiedFrom** | **string** | Retrieve entities from their modification date | 
 **modifiedTo** | **string** | Retrieve entities to their modification date | 
 **params** | **string** | Set this parameter in order to choose which entity fields you want to retrieve | [default to &quot;id,email,first_name,last_name&quot;]
 **responseFields** | **string** | Set this parameter in order to choose which entity fields you want to retrieve | 
 **exclude** | **string** | Set this parameter in order to choose which entity fields you want to ignore. Works only if parameter &#x60;params&#x60; equal force_all | 
 **groupId** | **string** | Customer group_id | 
 **storeId** | **string** | Retrieves customers specified by store id | 
 **customerListId** | **string** | The numeric ID of the customer list in Demandware. | 
 **avail** | **bool** | Defines category&#39;s visibility status | [default to true]
 **findValue** | **string** | Entity search that is specified by some value | 
 **findWhere** | **string** | Customer search that is specified by field | 
 **sortBy** | **string** | Set field to sort by | [default to &quot;created_time&quot;]
 **sortDirection** | **string** | Set sorting direction | [default to &quot;asc&quot;]
 **ids** | **string** | Retrieves customers specified by ids | 
 **sinceId** | **string** | Retrieve entities starting from the specified id. | 

### Return type

[**ModelResponseCustomerList**](ModelResponseCustomerList.md)

### Authorization

[StoreKeyAuth](../README.md#StoreKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomerUpdate

> AccountConfigUpdate200Response CustomerUpdate(ctx).CustomerUpdate(customerUpdate).Execute()

customer.update



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
	customerUpdate := *openapiclient.NewCustomerUpdate() // CustomerUpdate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomerAPI.CustomerUpdate(context.Background()).CustomerUpdate(customerUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomerAPI.CustomerUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CustomerUpdate`: AccountConfigUpdate200Response
	fmt.Fprintf(os.Stdout, "Response from `CustomerAPI.CustomerUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCustomerUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **customerUpdate** | [**CustomerUpdate**](CustomerUpdate.md) |  | 

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


## CustomerWishlistList

> ModelResponseCustomerWishlistList CustomerWishlistList(ctx).CustomerId(customerId).Id(id).StoreId(storeId).Start(start).Count(count).PageCursor(pageCursor).ResponseFields(responseFields).Execute()

customer.wishlist.list



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
	id := "10" // string | Entity id (optional)
	storeId := "1" // string | Store Id (optional)
	start := int32(0) // int32 | This parameter sets the number from which you want to get entities (optional) (default to 0)
	count := int32(20) // int32 | This parameter sets the entity amount that has to be retrieved. Max allowed count=250 (optional) (default to 10)
	pageCursor := "pageCursor_example" // string | Used to retrieve entities via cursor-based pagination (it can't be used with any other filtering parameter) (optional)
	responseFields := "{return_code,return_message,pagination,result}" // string | Set this parameter in order to choose which entity fields you want to retrieve (optional) (default to "{return_code,return_message,pagination,result}")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomerAPI.CustomerWishlistList(context.Background()).CustomerId(customerId).Id(id).StoreId(storeId).Start(start).Count(count).PageCursor(pageCursor).ResponseFields(responseFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomerAPI.CustomerWishlistList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CustomerWishlistList`: ModelResponseCustomerWishlistList
	fmt.Fprintf(os.Stdout, "Response from `CustomerAPI.CustomerWishlistList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCustomerWishlistListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **customerId** | **string** | Retrieves orders specified by customer id | 
 **id** | **string** | Entity id | 
 **storeId** | **string** | Store Id | 
 **start** | **int32** | This parameter sets the number from which you want to get entities | [default to 0]
 **count** | **int32** | This parameter sets the entity amount that has to be retrieved. Max allowed count&#x3D;250 | [default to 10]
 **pageCursor** | **string** | Used to retrieve entities via cursor-based pagination (it can&#39;t be used with any other filtering parameter) | 
 **responseFields** | **string** | Set this parameter in order to choose which entity fields you want to retrieve | [default to &quot;{return_code,return_message,pagination,result}&quot;]

### Return type

[**ModelResponseCustomerWishlistList**](ModelResponseCustomerWishlistList.md)

### Authorization

[StoreKeyAuth](../README.md#StoreKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

