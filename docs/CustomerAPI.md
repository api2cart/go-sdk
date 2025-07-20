# \CustomerAPI

All URIs are relative to *https://api.api2cart.local.com/v1.1*

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
	customerAdd := *openapiclient.NewCustomerAdd("mail@example.com") // CustomerAdd | 

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

> ModelResponseCustomerAttributeList CustomerAttributeList(ctx).CustomerId(customerId).Count(count).PageCursor(pageCursor).StoreId(storeId).LangId(langId).ResponseFields(responseFields).Params(params).Exclude(exclude).Execute()

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
	responseFields := "{return_code,return_message,pagination,result}" // string | Set this parameter in order to choose which entity fields you want to retrieve (optional)
	params := "id,model,price,images" // string | Set this parameter in order to choose which entity fields you want to retrieve (optional) (default to "force_all")
	exclude := "false" // string | Set this parameter in order to choose which entity fields you want to ignore. Works only if parameter `params` equal force_all (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomerAPI.CustomerAttributeList(context.Background()).CustomerId(customerId).Count(count).PageCursor(pageCursor).StoreId(storeId).LangId(langId).ResponseFields(responseFields).Params(params).Exclude(exclude).Execute()
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
 **responseFields** | **string** | Set this parameter in order to choose which entity fields you want to retrieve | 
 **params** | **string** | Set this parameter in order to choose which entity fields you want to retrieve | [default to &quot;force_all&quot;]
 **exclude** | **string** | Set this parameter in order to choose which entity fields you want to ignore. Works only if parameter &#x60;params&#x60; equal force_all | 

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

> CustomerCount200Response CustomerCount(ctx).Ids(ids).SinceId(sinceId).CustomerListId(customerListId).GroupId(groupId).StoreId(storeId).Avail(avail).IncludeGuests(includeGuests).FindValue(findValue).FindWhere(findWhere).CreatedFrom(createdFrom).CreatedTo(createdTo).ModifiedFrom(modifiedFrom).ModifiedTo(modifiedTo).Execute()

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
	ids := "24,25" // string | Counts customers specified by ids (optional)
	sinceId := "56" // string | Retrieve entities starting from the specified id. (optional)
	customerListId := "exampleListId" // string | The numeric ID of the customer list in Demandware. (optional)
	groupId := "3" // string | Customer group_id (optional)
	storeId := "1" // string | Counts customer specified by store id (optional)
	avail := false // bool | Defines category's visibility status (optional) (default to true)
	includeGuests := true // bool | Indicates whether to include guest customers in the total count. (optional) (default to false)
	findValue := "mail@gmail.com" // string | Entity search that is specified by some value (optional)
	findWhere := "email" // string | Counts customers that are searched specified by field (optional)
	createdFrom := "2010-07-29 13:45:52" // string | Retrieve entities from their creation date (optional)
	createdTo := "2100-08-29 13:45:52" // string | Retrieve entities to their creation date (optional)
	modifiedFrom := "2010-07-29 13:45:52" // string | Retrieve entities from their modification date (optional)
	modifiedTo := "2100-08-29 13:45:52" // string | Retrieve entities to their modification date (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomerAPI.CustomerCount(context.Background()).Ids(ids).SinceId(sinceId).CustomerListId(customerListId).GroupId(groupId).StoreId(storeId).Avail(avail).IncludeGuests(includeGuests).FindValue(findValue).FindWhere(findWhere).CreatedFrom(createdFrom).CreatedTo(createdTo).ModifiedFrom(modifiedFrom).ModifiedTo(modifiedTo).Execute()
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
 **ids** | **string** | Counts customers specified by ids | 
 **sinceId** | **string** | Retrieve entities starting from the specified id. | 
 **customerListId** | **string** | The numeric ID of the customer list in Demandware. | 
 **groupId** | **string** | Customer group_id | 
 **storeId** | **string** | Counts customer specified by store id | 
 **avail** | **bool** | Defines category&#39;s visibility status | [default to true]
 **includeGuests** | **bool** | Indicates whether to include guest customers in the total count. | [default to false]
 **findValue** | **string** | Entity search that is specified by some value | 
 **findWhere** | **string** | Counts customers that are searched specified by field | 
 **createdFrom** | **string** | Retrieve entities from their creation date | 
 **createdTo** | **string** | Retrieve entities to their creation date | 
 **modifiedFrom** | **string** | Retrieve entities from their modification date | 
 **modifiedTo** | **string** | Retrieve entities to their modification date | 

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

> CustomerFind200Response CustomerFind(ctx).FindValue(findValue).FindWhere(findWhere).FindParams(findParams).StoreId(storeId).IncludeGuests(includeGuests).Execute()

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
	includeGuests := true // bool | Indicates whether to search among guest customers when looking up a customer. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomerAPI.CustomerFind(context.Background()).FindValue(findValue).FindWhere(findWhere).FindParams(findParams).StoreId(storeId).IncludeGuests(includeGuests).Execute()
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
 **includeGuests** | **bool** | Indicates whether to search among guest customers when looking up a customer. | [default to false]

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

> ModelResponseCustomerGroupList CustomerGroupList(ctx).Start(start).Count(count).PageCursor(pageCursor).GroupIds(groupIds).StoreId(storeId).LangId(langId).ResponseFields(responseFields).Params(params).Exclude(exclude).DisableCache(disableCache).Execute()

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
	start := int32(0) // int32 | This parameter sets the number from which you want to get entities (optional) (default to 0)
	count := int32(20) // int32 | This parameter sets the entity amount that has to be retrieved. Max allowed count=250 (optional) (default to 10)
	pageCursor := "pageCursor_example" // string | Used to retrieve entities via cursor-based pagination (it can't be used with any other filtering parameter) (optional)
	groupIds := "1,2,3" // string | Groups that will be assigned to a customer (optional)
	storeId := "1" // string | Store Id (optional)
	langId := "3" // string | Language id (optional)
	responseFields := "{return_code,return_message,pagination,result}" // string | Set this parameter in order to choose which entity fields you want to retrieve (optional)
	params := "id,model,price,images" // string | Set this parameter in order to choose which entity fields you want to retrieve (optional) (default to "id,name,additional_fields")
	exclude := "false" // string | Set this parameter in order to choose which entity fields you want to ignore. Works only if parameter `params` equal force_all (optional)
	disableCache := false // bool | Disable cache for current request (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomerAPI.CustomerGroupList(context.Background()).Start(start).Count(count).PageCursor(pageCursor).GroupIds(groupIds).StoreId(storeId).LangId(langId).ResponseFields(responseFields).Params(params).Exclude(exclude).DisableCache(disableCache).Execute()
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
 **start** | **int32** | This parameter sets the number from which you want to get entities | [default to 0]
 **count** | **int32** | This parameter sets the entity amount that has to be retrieved. Max allowed count&#x3D;250 | [default to 10]
 **pageCursor** | **string** | Used to retrieve entities via cursor-based pagination (it can&#39;t be used with any other filtering parameter) | 
 **groupIds** | **string** | Groups that will be assigned to a customer | 
 **storeId** | **string** | Store Id | 
 **langId** | **string** | Language id | 
 **responseFields** | **string** | Set this parameter in order to choose which entity fields you want to retrieve | 
 **params** | **string** | Set this parameter in order to choose which entity fields you want to retrieve | [default to &quot;id,name,additional_fields&quot;]
 **exclude** | **string** | Set this parameter in order to choose which entity fields you want to ignore. Works only if parameter &#x60;params&#x60; equal force_all | 
 **disableCache** | **bool** | Disable cache for current request | [default to false]

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

> CustomerInfo200Response CustomerInfo(ctx).Id(id).StoreId(storeId).ResponseFields(responseFields).Params(params).Exclude(exclude).Execute()

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
	storeId := "1" // string | Retrieves customer info specified by store id (optional)
	responseFields := "{result{id,parent_id,sku,upc,images,combination}}" // string | Set this parameter in order to choose which entity fields you want to retrieve (optional)
	params := "id,email" // string | Set this parameter in order to choose which entity fields you want to retrieve (optional) (default to "id,email,first_name,last_name")
	exclude := "id,email" // string | Set this parameter in order to choose which entity fields you want to ignore. Works only if parameter `params` equal force_all (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomerAPI.CustomerInfo(context.Background()).Id(id).StoreId(storeId).ResponseFields(responseFields).Params(params).Exclude(exclude).Execute()
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
 **storeId** | **string** | Retrieves customer info specified by store id | 
 **responseFields** | **string** | Set this parameter in order to choose which entity fields you want to retrieve | 
 **params** | **string** | Set this parameter in order to choose which entity fields you want to retrieve | [default to &quot;id,email,first_name,last_name&quot;]
 **exclude** | **string** | Set this parameter in order to choose which entity fields you want to ignore. Works only if parameter &#x60;params&#x60; equal force_all | 

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

> ModelResponseCustomerList CustomerList(ctx).Start(start).Count(count).PageCursor(pageCursor).Ids(ids).SinceId(sinceId).CustomerListId(customerListId).GroupId(groupId).StoreId(storeId).Avail(avail).IncludeGuests(includeGuests).FindValue(findValue).FindWhere(findWhere).CreatedFrom(createdFrom).CreatedTo(createdTo).ModifiedFrom(modifiedFrom).ModifiedTo(modifiedTo).SortBy(sortBy).SortDirection(sortDirection).ResponseFields(responseFields).Params(params).Exclude(exclude).Execute()

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
	start := int32(0) // int32 | This parameter sets the number from which you want to get entities (optional) (default to 0)
	count := int32(20) // int32 | This parameter sets the entity amount that has to be retrieved. Max allowed count=250 (optional) (default to 10)
	pageCursor := "pageCursor_example" // string | Used to retrieve entities via cursor-based pagination (it can't be used with any other filtering parameter) (optional)
	ids := "24,25" // string | Retrieves customers specified by ids (optional)
	sinceId := "56" // string | Retrieve entities starting from the specified id. (optional)
	customerListId := "exampleListId" // string | The numeric ID of the customer list in Demandware. (optional)
	groupId := "3" // string | Customer group_id (optional)
	storeId := "1" // string | Retrieves customers specified by store id (optional)
	avail := false // bool | Defines category's visibility status (optional) (default to true)
	includeGuests := true // bool | Indicates whether to include guest customers in the list results. (optional) (default to false)
	findValue := "mail@gmail.com" // string | Entity search that is specified by some value (optional)
	findWhere := "email" // string | Customer search that is specified by field (optional)
	createdFrom := "2010-07-29 13:45:52" // string | Retrieve entities from their creation date (optional)
	createdTo := "2100-08-29 13:45:52" // string | Retrieve entities to their creation date (optional)
	modifiedFrom := "2010-07-29 13:45:52" // string | Retrieve entities from their modification date (optional)
	modifiedTo := "2100-08-29 13:45:52" // string | Retrieve entities to their modification date (optional)
	sortBy := "value_id" // string | Set field to sort by (optional) (default to "created_time")
	sortDirection := "asc" // string | Set sorting direction (optional) (default to "asc")
	responseFields := "{result{customer}}" // string | Set this parameter in order to choose which entity fields you want to retrieve (optional)
	params := "id,email" // string | Set this parameter in order to choose which entity fields you want to retrieve (optional) (default to "id,email,first_name,last_name")
	exclude := "id,email" // string | Set this parameter in order to choose which entity fields you want to ignore. Works only if parameter `params` equal force_all (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomerAPI.CustomerList(context.Background()).Start(start).Count(count).PageCursor(pageCursor).Ids(ids).SinceId(sinceId).CustomerListId(customerListId).GroupId(groupId).StoreId(storeId).Avail(avail).IncludeGuests(includeGuests).FindValue(findValue).FindWhere(findWhere).CreatedFrom(createdFrom).CreatedTo(createdTo).ModifiedFrom(modifiedFrom).ModifiedTo(modifiedTo).SortBy(sortBy).SortDirection(sortDirection).ResponseFields(responseFields).Params(params).Exclude(exclude).Execute()
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
 **start** | **int32** | This parameter sets the number from which you want to get entities | [default to 0]
 **count** | **int32** | This parameter sets the entity amount that has to be retrieved. Max allowed count&#x3D;250 | [default to 10]
 **pageCursor** | **string** | Used to retrieve entities via cursor-based pagination (it can&#39;t be used with any other filtering parameter) | 
 **ids** | **string** | Retrieves customers specified by ids | 
 **sinceId** | **string** | Retrieve entities starting from the specified id. | 
 **customerListId** | **string** | The numeric ID of the customer list in Demandware. | 
 **groupId** | **string** | Customer group_id | 
 **storeId** | **string** | Retrieves customers specified by store id | 
 **avail** | **bool** | Defines category&#39;s visibility status | [default to true]
 **includeGuests** | **bool** | Indicates whether to include guest customers in the list results. | [default to false]
 **findValue** | **string** | Entity search that is specified by some value | 
 **findWhere** | **string** | Customer search that is specified by field | 
 **createdFrom** | **string** | Retrieve entities from their creation date | 
 **createdTo** | **string** | Retrieve entities to their creation date | 
 **modifiedFrom** | **string** | Retrieve entities from their modification date | 
 **modifiedTo** | **string** | Retrieve entities to their modification date | 
 **sortBy** | **string** | Set field to sort by | [default to &quot;created_time&quot;]
 **sortDirection** | **string** | Set sorting direction | [default to &quot;asc&quot;]
 **responseFields** | **string** | Set this parameter in order to choose which entity fields you want to retrieve | 
 **params** | **string** | Set this parameter in order to choose which entity fields you want to retrieve | [default to &quot;id,email,first_name,last_name&quot;]
 **exclude** | **string** | Set this parameter in order to choose which entity fields you want to ignore. Works only if parameter &#x60;params&#x60; equal force_all | 

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

> ModelResponseCustomerWishlistList CustomerWishlistList(ctx).CustomerId(customerId).Start(start).Count(count).PageCursor(pageCursor).Id(id).StoreId(storeId).ResponseFields(responseFields).Execute()

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
	start := int32(0) // int32 | This parameter sets the number from which you want to get entities (optional) (default to 0)
	count := int32(20) // int32 | This parameter sets the entity amount that has to be retrieved. Max allowed count=250 (optional) (default to 10)
	pageCursor := "pageCursor_example" // string | Used to retrieve entities via cursor-based pagination (it can't be used with any other filtering parameter) (optional)
	id := "10" // string | Entity id (optional)
	storeId := "1" // string | Store Id (optional)
	responseFields := "{return_code,return_message,pagination,result}" // string | Set this parameter in order to choose which entity fields you want to retrieve (optional) (default to "{return_code,return_message,pagination,result}")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomerAPI.CustomerWishlistList(context.Background()).CustomerId(customerId).Start(start).Count(count).PageCursor(pageCursor).Id(id).StoreId(storeId).ResponseFields(responseFields).Execute()
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
 **start** | **int32** | This parameter sets the number from which you want to get entities | [default to 0]
 **count** | **int32** | This parameter sets the entity amount that has to be retrieved. Max allowed count&#x3D;250 | [default to 10]
 **pageCursor** | **string** | Used to retrieve entities via cursor-based pagination (it can&#39;t be used with any other filtering parameter) | 
 **id** | **string** | Entity id | 
 **storeId** | **string** | Store Id | 
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

