# \TaxAPI

All URIs are relative to *https://api.api2cart.com/v1.1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TaxClassInfo**](TaxAPI.md#TaxClassInfo) | **Get** /tax.class.info.json | tax.class.info
[**TaxClassList**](TaxAPI.md#TaxClassList) | **Get** /tax.class.list.json | tax.class.list



## TaxClassInfo

> ModelResponseTaxClassInfo TaxClassInfo(ctx).TaxClassId(taxClassId).StoreId(storeId).LangId(langId).Params(params).ResponseFields(responseFields).Exclude(exclude).Execute()

tax.class.info



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
	taxClassId := "9" // string | Retrieves taxes specified by class id
	storeId := "1" // string | Store Id (optional)
	langId := "3" // string | Language id (optional)
	params := "tax_class_id,tax" // string | Set this parameter in order to choose which entity fields you want to retrieve (optional) (default to "tax_class_id,name,avail")
	responseFields := "{result{id,name,tax,tax_rates{id,countries{id,name,states},cities,address,zip_codes{is_range,range,fields}}}}" // string | Set this parameter in order to choose which entity fields you want to retrieve (optional)
	exclude := "tax_class_id,tax" // string | Set this parameter in order to choose which entity fields you want to ignore. Works only if parameter `params` equal force_all (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaxAPI.TaxClassInfo(context.Background()).TaxClassId(taxClassId).StoreId(storeId).LangId(langId).Params(params).ResponseFields(responseFields).Exclude(exclude).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaxAPI.TaxClassInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TaxClassInfo`: ModelResponseTaxClassInfo
	fmt.Fprintf(os.Stdout, "Response from `TaxAPI.TaxClassInfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTaxClassInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **taxClassId** | **string** | Retrieves taxes specified by class id | 
 **storeId** | **string** | Store Id | 
 **langId** | **string** | Language id | 
 **params** | **string** | Set this parameter in order to choose which entity fields you want to retrieve | [default to &quot;tax_class_id,name,avail&quot;]
 **responseFields** | **string** | Set this parameter in order to choose which entity fields you want to retrieve | 
 **exclude** | **string** | Set this parameter in order to choose which entity fields you want to ignore. Works only if parameter &#x60;params&#x60; equal force_all | 

### Return type

[**ModelResponseTaxClassInfo**](ModelResponseTaxClassInfo.md)

### Authorization

[StoreKeyAuth](../README.md#StoreKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TaxClassList

> ModelResponseTaxClassList TaxClassList(ctx).CreatedTo(createdTo).CreatedFrom(createdFrom).ModifiedTo(modifiedTo).ModifiedFrom(modifiedFrom).FindValue(findValue).FindWhere(findWhere).StoreId(storeId).Count(count).PageCursor(pageCursor).ResponseFields(responseFields).Execute()

tax.class.list



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
	createdTo := "2100-08-29 13:45:52" // string | Retrieve entities to their creation date (optional)
	createdFrom := "2010-07-29 13:45:52" // string | Retrieve entities from their creation date (optional)
	modifiedTo := "2100-08-29 13:45:52" // string | Retrieve entities to their modification date (optional)
	modifiedFrom := "2010-07-29 13:45:52" // string | Retrieve entities from their modification date (optional)
	findValue := "tax" // string | Entity search that is specified by some value (optional)
	findWhere := "name" // string | Tax class search that is specified by field (optional)
	storeId := "1" // string | Store Id (optional)
	count := int32(20) // int32 | This parameter sets the entity amount that has to be retrieved. Max allowed count=250 (optional) (default to 10)
	pageCursor := "pageCursor_example" // string | Used to retrieve entities via cursor-based pagination (it can't be used with any other filtering parameter) (optional)
	responseFields := "{result}" // string | Set this parameter in order to choose which entity fields you want to retrieve (optional) (default to "{return_code,return_message,pagination,result}")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaxAPI.TaxClassList(context.Background()).CreatedTo(createdTo).CreatedFrom(createdFrom).ModifiedTo(modifiedTo).ModifiedFrom(modifiedFrom).FindValue(findValue).FindWhere(findWhere).StoreId(storeId).Count(count).PageCursor(pageCursor).ResponseFields(responseFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaxAPI.TaxClassList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TaxClassList`: ModelResponseTaxClassList
	fmt.Fprintf(os.Stdout, "Response from `TaxAPI.TaxClassList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTaxClassListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createdTo** | **string** | Retrieve entities to their creation date | 
 **createdFrom** | **string** | Retrieve entities from their creation date | 
 **modifiedTo** | **string** | Retrieve entities to their modification date | 
 **modifiedFrom** | **string** | Retrieve entities from their modification date | 
 **findValue** | **string** | Entity search that is specified by some value | 
 **findWhere** | **string** | Tax class search that is specified by field | 
 **storeId** | **string** | Store Id | 
 **count** | **int32** | This parameter sets the entity amount that has to be retrieved. Max allowed count&#x3D;250 | [default to 10]
 **pageCursor** | **string** | Used to retrieve entities via cursor-based pagination (it can&#39;t be used with any other filtering parameter) | 
 **responseFields** | **string** | Set this parameter in order to choose which entity fields you want to retrieve | [default to &quot;{return_code,return_message,pagination,result}&quot;]

### Return type

[**ModelResponseTaxClassList**](ModelResponseTaxClassList.md)

### Authorization

[StoreKeyAuth](../README.md#StoreKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

