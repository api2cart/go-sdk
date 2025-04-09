# \CategoryAPI

All URIs are relative to *https://api.api2cart.com/v1.1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CategoryAdd**](CategoryAPI.md#CategoryAdd) | **Post** /category.add.json | category.add
[**CategoryAddBatch**](CategoryAPI.md#CategoryAddBatch) | **Post** /category.add.batch.json | category.add.batch
[**CategoryAssign**](CategoryAPI.md#CategoryAssign) | **Post** /category.assign.json | category.assign
[**CategoryCount**](CategoryAPI.md#CategoryCount) | **Get** /category.count.json | category.count
[**CategoryDelete**](CategoryAPI.md#CategoryDelete) | **Delete** /category.delete.json | category.delete
[**CategoryFind**](CategoryAPI.md#CategoryFind) | **Get** /category.find.json | category.find
[**CategoryImageAdd**](CategoryAPI.md#CategoryImageAdd) | **Post** /category.image.add.json | category.image.add
[**CategoryImageDelete**](CategoryAPI.md#CategoryImageDelete) | **Delete** /category.image.delete.json | category.image.delete
[**CategoryInfo**](CategoryAPI.md#CategoryInfo) | **Get** /category.info.json | category.info
[**CategoryList**](CategoryAPI.md#CategoryList) | **Get** /category.list.json | category.list
[**CategoryUnassign**](CategoryAPI.md#CategoryUnassign) | **Post** /category.unassign.json | category.unassign
[**CategoryUpdate**](CategoryAPI.md#CategoryUpdate) | **Put** /category.update.json | category.update



## CategoryAdd

> CategoryAdd200Response CategoryAdd(ctx).Name(name).ParentId(parentId).StoresIds(storesIds).StoreId(storeId).LangId(langId).Avail(avail).SortOrder(sortOrder).CreatedTime(createdTime).ModifiedTime(modifiedTime).Description(description).ShortDescription(shortDescription).MetaTitle(metaTitle).MetaDescription(metaDescription).MetaKeywords(metaKeywords).SeoUrl(seoUrl).Execute()

category.add



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
	name := "Shoes" // string | Defines category's name that has to be added
	parentId := "6" // string | Adds categories specified by parent id (optional)
	storesIds := "1,2" // string | Create category in the stores that is specified by comma-separated stores' id (optional)
	storeId := "1" // string | Store Id (optional)
	langId := "3" // string | Language id (optional)
	avail := false // bool | Defines category's visibility status (optional) (default to true)
	sortOrder := int32(2) // int32 | Sort number in the list (optional) (default to 0)
	createdTime := "2014-01-30 15:58:41" // string | Entity's date creation (optional)
	modifiedTime := "2014-07-30 15:58:41" // string | Entity's date modification (optional)
	description := "Test category" // string | Defines category's description (optional)
	shortDescription := "Short description. This is very short description" // string | Defines short description (optional)
	metaTitle := "category,test" // string | Defines unique meta title for each entity (optional)
	metaDescription := "category,test" // string | Defines unique meta description of a entity (optional)
	metaKeywords := "category,test" // string | Defines unique meta keywords for each entity (optional)
	seoUrl := "category,test" // string | Defines unique category's URL for SEO (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CategoryAPI.CategoryAdd(context.Background()).Name(name).ParentId(parentId).StoresIds(storesIds).StoreId(storeId).LangId(langId).Avail(avail).SortOrder(sortOrder).CreatedTime(createdTime).ModifiedTime(modifiedTime).Description(description).ShortDescription(shortDescription).MetaTitle(metaTitle).MetaDescription(metaDescription).MetaKeywords(metaKeywords).SeoUrl(seoUrl).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CategoryAPI.CategoryAdd``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CategoryAdd`: CategoryAdd200Response
	fmt.Fprintf(os.Stdout, "Response from `CategoryAPI.CategoryAdd`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCategoryAddRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | Defines category&#39;s name that has to be added | 
 **parentId** | **string** | Adds categories specified by parent id | 
 **storesIds** | **string** | Create category in the stores that is specified by comma-separated stores&#39; id | 
 **storeId** | **string** | Store Id | 
 **langId** | **string** | Language id | 
 **avail** | **bool** | Defines category&#39;s visibility status | [default to true]
 **sortOrder** | **int32** | Sort number in the list | [default to 0]
 **createdTime** | **string** | Entity&#39;s date creation | 
 **modifiedTime** | **string** | Entity&#39;s date modification | 
 **description** | **string** | Defines category&#39;s description | 
 **shortDescription** | **string** | Defines short description | 
 **metaTitle** | **string** | Defines unique meta title for each entity | 
 **metaDescription** | **string** | Defines unique meta description of a entity | 
 **metaKeywords** | **string** | Defines unique meta keywords for each entity | 
 **seoUrl** | **string** | Defines unique category&#39;s URL for SEO | 

### Return type

[**CategoryAdd200Response**](CategoryAdd200Response.md)

### Authorization

[StoreKeyAuth](../README.md#StoreKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CategoryAddBatch

> CategoryAddBatch200Response CategoryAddBatch(ctx).CategoryAddBatch(categoryAddBatch).Execute()

category.add.batch



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
	categoryAddBatch := *openapiclient.NewCategoryAddBatch([]openapiclient.CategoryAddBatchPayloadInner{*openapiclient.NewCategoryAddBatchPayloadInner("Name_example")}) // CategoryAddBatch | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CategoryAPI.CategoryAddBatch(context.Background()).CategoryAddBatch(categoryAddBatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CategoryAPI.CategoryAddBatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CategoryAddBatch`: CategoryAddBatch200Response
	fmt.Fprintf(os.Stdout, "Response from `CategoryAPI.CategoryAddBatch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCategoryAddBatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **categoryAddBatch** | [**CategoryAddBatch**](CategoryAddBatch.md) |  | 

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


## CategoryAssign

> CartConfigUpdate200Response CategoryAssign(ctx).ProductId(productId).CategoryId(categoryId).StoreId(storeId).Execute()

category.assign



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
	productId := "10" // string | Defines category assign to the product, specified by product id
	categoryId := "6" // string | Defines category assign, specified by category id
	storeId := "1" // string | Store Id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CategoryAPI.CategoryAssign(context.Background()).ProductId(productId).CategoryId(categoryId).StoreId(storeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CategoryAPI.CategoryAssign``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CategoryAssign`: CartConfigUpdate200Response
	fmt.Fprintf(os.Stdout, "Response from `CategoryAPI.CategoryAssign`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCategoryAssignRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **productId** | **string** | Defines category assign to the product, specified by product id | 
 **categoryId** | **string** | Defines category assign, specified by category id | 
 **storeId** | **string** | Store Id | 

### Return type

[**CartConfigUpdate200Response**](CartConfigUpdate200Response.md)

### Authorization

[StoreKeyAuth](../README.md#StoreKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CategoryCount

> CategoryCount200Response CategoryCount(ctx).ParentId(parentId).StoreId(storeId).LangId(langId).CreatedFrom(createdFrom).CreatedTo(createdTo).ModifiedFrom(modifiedFrom).ModifiedTo(modifiedTo).Avail(avail).ProductType(productType).FindValue(findValue).FindWhere(findWhere).ReportRequestId(reportRequestId).DisableReportCache(disableReportCache).Execute()

category.count



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
	parentId := "6" // string | Counts categories specified by parent id (optional)
	storeId := "1" // string | Counts category specified by store id (optional)
	langId := "3" // string | Counts category specified by language id (optional)
	createdFrom := "2010-07-29 13:45:52" // string | Retrieve entities from their creation date (optional)
	createdTo := "2100-08-29 13:45:52" // string | Retrieve entities to their creation date (optional)
	modifiedFrom := "2010-07-29 13:45:52" // string | Retrieve entities from their modification date (optional)
	modifiedTo := "2100-08-29 13:45:52" // string | Retrieve entities to their modification date (optional)
	avail := false // bool | Defines category's visibility status (optional) (default to true)
	productType := "BICYCLE" // string | A categorization for the product (optional)
	findValue := "Demo category 1" // string | Entity search that is specified by some value (optional)
	findWhere := "email" // string | Counts categories that are searched specified by field (optional)
	reportRequestId := "105245017661" // string | Report request id (optional)
	disableReportCache := false // bool | Disable report cache for current request (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CategoryAPI.CategoryCount(context.Background()).ParentId(parentId).StoreId(storeId).LangId(langId).CreatedFrom(createdFrom).CreatedTo(createdTo).ModifiedFrom(modifiedFrom).ModifiedTo(modifiedTo).Avail(avail).ProductType(productType).FindValue(findValue).FindWhere(findWhere).ReportRequestId(reportRequestId).DisableReportCache(disableReportCache).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CategoryAPI.CategoryCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CategoryCount`: CategoryCount200Response
	fmt.Fprintf(os.Stdout, "Response from `CategoryAPI.CategoryCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCategoryCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **parentId** | **string** | Counts categories specified by parent id | 
 **storeId** | **string** | Counts category specified by store id | 
 **langId** | **string** | Counts category specified by language id | 
 **createdFrom** | **string** | Retrieve entities from their creation date | 
 **createdTo** | **string** | Retrieve entities to their creation date | 
 **modifiedFrom** | **string** | Retrieve entities from their modification date | 
 **modifiedTo** | **string** | Retrieve entities to their modification date | 
 **avail** | **bool** | Defines category&#39;s visibility status | [default to true]
 **productType** | **string** | A categorization for the product | 
 **findValue** | **string** | Entity search that is specified by some value | 
 **findWhere** | **string** | Counts categories that are searched specified by field | 
 **reportRequestId** | **string** | Report request id | 
 **disableReportCache** | **bool** | Disable report cache for current request | [default to false]

### Return type

[**CategoryCount200Response**](CategoryCount200Response.md)

### Authorization

[StoreKeyAuth](../README.md#StoreKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CategoryDelete

> CategoryDelete200Response CategoryDelete(ctx).Id(id).StoreId(storeId).Execute()

category.delete



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
	id := "10" // string | Defines category removal, specified by category id
	storeId := "1" // string | Store Id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CategoryAPI.CategoryDelete(context.Background()).Id(id).StoreId(storeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CategoryAPI.CategoryDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CategoryDelete`: CategoryDelete200Response
	fmt.Fprintf(os.Stdout, "Response from `CategoryAPI.CategoryDelete`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCategoryDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | Defines category removal, specified by category id | 
 **storeId** | **string** | Store Id | 

### Return type

[**CategoryDelete200Response**](CategoryDelete200Response.md)

### Authorization

[StoreKeyAuth](../README.md#StoreKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CategoryFind

> CategoryFind200Response CategoryFind(ctx).FindValue(findValue).FindWhere(findWhere).FindParams(findParams).StoreId(storeId).LangId(langId).Execute()

category.find



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
	findValue := "Demo category 1" // string | Entity search that is specified by some value
	findWhere := "name" // string | Entity search that is specified by the comma-separated unique fields (optional) (default to "name")
	findParams := "regex" // string | Entity search that is specified by comma-separated parameters (optional) (default to "whole_words")
	storeId := "1" // string | Store Id (optional)
	langId := "3" // string | Language id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CategoryAPI.CategoryFind(context.Background()).FindValue(findValue).FindWhere(findWhere).FindParams(findParams).StoreId(storeId).LangId(langId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CategoryAPI.CategoryFind``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CategoryFind`: CategoryFind200Response
	fmt.Fprintf(os.Stdout, "Response from `CategoryAPI.CategoryFind`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCategoryFindRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **findValue** | **string** | Entity search that is specified by some value | 
 **findWhere** | **string** | Entity search that is specified by the comma-separated unique fields | [default to &quot;name&quot;]
 **findParams** | **string** | Entity search that is specified by comma-separated parameters | [default to &quot;whole_words&quot;]
 **storeId** | **string** | Store Id | 
 **langId** | **string** | Language id | 

### Return type

[**CategoryFind200Response**](CategoryFind200Response.md)

### Authorization

[StoreKeyAuth](../README.md#StoreKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CategoryImageAdd

> CategoryImageAdd200Response CategoryImageAdd(ctx).CategoryId(categoryId).ImageName(imageName).Url(url).Type_(type_).Label(label).Mime(mime).Position(position).StoreId(storeId).Execute()

category.image.add



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
	categoryId := "6" // string | Defines category id where the image should be added
	imageName := "bag-gray.png" // string | Defines image's name
	url := "http://docs.api2cart.com/img/logo.png" // string | Defines URL of the image that has to be added
	type_ := "base" // string | Defines image's types that are specified by comma-separated list
	label := "This cool image" // string | Defines alternative text that has to be attached to the picture (optional)
	mime := "image/jpeg" // string | Mime type of image http://en.wikipedia.org/wiki/Internet_media_type. (optional)
	position := int32(5) // int32 | Defines image’s position in the list (optional) (default to 0)
	storeId := "1" // string | Store Id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CategoryAPI.CategoryImageAdd(context.Background()).CategoryId(categoryId).ImageName(imageName).Url(url).Type_(type_).Label(label).Mime(mime).Position(position).StoreId(storeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CategoryAPI.CategoryImageAdd``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CategoryImageAdd`: CategoryImageAdd200Response
	fmt.Fprintf(os.Stdout, "Response from `CategoryAPI.CategoryImageAdd`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCategoryImageAddRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **categoryId** | **string** | Defines category id where the image should be added | 
 **imageName** | **string** | Defines image&#39;s name | 
 **url** | **string** | Defines URL of the image that has to be added | 
 **type_** | **string** | Defines image&#39;s types that are specified by comma-separated list | 
 **label** | **string** | Defines alternative text that has to be attached to the picture | 
 **mime** | **string** | Mime type of image http://en.wikipedia.org/wiki/Internet_media_type. | 
 **position** | **int32** | Defines image’s position in the list | [default to 0]
 **storeId** | **string** | Store Id | 

### Return type

[**CategoryImageAdd200Response**](CategoryImageAdd200Response.md)

### Authorization

[StoreKeyAuth](../README.md#StoreKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CategoryImageDelete

> AttributeDelete200Response CategoryImageDelete(ctx).CategoryId(categoryId).ImageId(imageId).StoreId(storeId).Execute()

category.image.delete



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
	categoryId := "6" // string | Defines category id where the image should be deleted
	imageId := "82950b84f468edff480680f99cedbe0d" // string | Define image id
	storeId := "1" // string | Store Id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CategoryAPI.CategoryImageDelete(context.Background()).CategoryId(categoryId).ImageId(imageId).StoreId(storeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CategoryAPI.CategoryImageDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CategoryImageDelete`: AttributeDelete200Response
	fmt.Fprintf(os.Stdout, "Response from `CategoryAPI.CategoryImageDelete`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCategoryImageDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **categoryId** | **string** | Defines category id where the image should be deleted | 
 **imageId** | **string** | Define image id | 
 **storeId** | **string** | Store Id | 

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


## CategoryInfo

> CategoryInfo200Response CategoryInfo(ctx).Id(id).Params(params).ResponseFields(responseFields).Exclude(exclude).StoreId(storeId).LangId(langId).SchemaType(schemaType).ReportRequestId(reportRequestId).DisableReportCache(disableReportCache).Execute()

category.info



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
	id := "10" // string | Retrieves category's info specified by category id
	params := "id,parent_id,name" // string | Set this parameter in order to choose which entity fields you want to retrieve (optional) (default to "id,parent_id,name,description")
	responseFields := "{result{id,name,parent_id,modified_at{value},images}}" // string | Set this parameter in order to choose which entity fields you want to retrieve (optional)
	exclude := "id,parent_id,name" // string | Set this parameter in order to choose which entity fields you want to ignore. Works only if parameter `params` equal force_all (optional)
	storeId := "1" // string | Retrieves category info  specified by store id (optional)
	langId := "3" // string | Retrieves category info  specified by language id (optional)
	schemaType := "LISTING" // string | The name of the requirements set for the provided schema. (optional)
	reportRequestId := "105245017661" // string | Report request id (optional)
	disableReportCache := false // bool | Disable report cache for current request (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CategoryAPI.CategoryInfo(context.Background()).Id(id).Params(params).ResponseFields(responseFields).Exclude(exclude).StoreId(storeId).LangId(langId).SchemaType(schemaType).ReportRequestId(reportRequestId).DisableReportCache(disableReportCache).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CategoryAPI.CategoryInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CategoryInfo`: CategoryInfo200Response
	fmt.Fprintf(os.Stdout, "Response from `CategoryAPI.CategoryInfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCategoryInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | Retrieves category&#39;s info specified by category id | 
 **params** | **string** | Set this parameter in order to choose which entity fields you want to retrieve | [default to &quot;id,parent_id,name,description&quot;]
 **responseFields** | **string** | Set this parameter in order to choose which entity fields you want to retrieve | 
 **exclude** | **string** | Set this parameter in order to choose which entity fields you want to ignore. Works only if parameter &#x60;params&#x60; equal force_all | 
 **storeId** | **string** | Retrieves category info  specified by store id | 
 **langId** | **string** | Retrieves category info  specified by language id | 
 **schemaType** | **string** | The name of the requirements set for the provided schema. | 
 **reportRequestId** | **string** | Report request id | 
 **disableReportCache** | **bool** | Disable report cache for current request | [default to false]

### Return type

[**CategoryInfo200Response**](CategoryInfo200Response.md)

### Authorization

[StoreKeyAuth](../README.md#StoreKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CategoryList

> ModelResponseCategoryList CategoryList(ctx).Start(start).Count(count).PageCursor(pageCursor).ParentId(parentId).Params(params).ResponseFields(responseFields).Exclude(exclude).StoreId(storeId).LangId(langId).CreatedFrom(createdFrom).CreatedTo(createdTo).ModifiedFrom(modifiedFrom).ModifiedTo(modifiedTo).Avail(avail).ProductType(productType).FindValue(findValue).FindWhere(findWhere).ReportRequestId(reportRequestId).DisableReportCache(disableReportCache).DisableCache(disableCache).Execute()

category.list



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
	parentId := "6" // string | Retrieves categories specified by parent id (optional)
	params := "id,parent_id,name" // string | Set this parameter in order to choose which entity fields you want to retrieve (optional) (default to "id,parent_id,name,description")
	responseFields := "{result{categories_count,category{id,parent_id,modified_at{value},images}}}" // string | Set this parameter in order to choose which entity fields you want to retrieve (optional)
	exclude := "id,parent_id,name" // string | Set this parameter in order to choose which entity fields you want to ignore. Works only if parameter `params` equal force_all (optional)
	storeId := "1" // string | Retrieves categories specified by store id (optional)
	langId := "3" // string | Retrieves categorys specified by language id (optional)
	createdFrom := "2010-07-29 13:45:52" // string | Retrieve entities from their creation date (optional)
	createdTo := "2100-08-29 13:45:52" // string | Retrieve entities to their creation date (optional)
	modifiedFrom := "2010-07-29 13:45:52" // string | Retrieve entities from their modification date (optional)
	modifiedTo := "2100-08-29 13:45:52" // string | Retrieve entities to their modification date (optional)
	avail := false // bool | Defines category's visibility status (optional) (default to true)
	productType := "BICYCLE" // string | A categorization for the product (optional)
	findValue := "Demo category 1" // string | Entity search that is specified by some value (optional)
	findWhere := "name" // string | Category search that is specified by field (optional)
	reportRequestId := "105245017661" // string | Report request id (optional)
	disableReportCache := false // bool | Disable report cache for current request (optional) (default to false)
	disableCache := false // bool | Disable cache for current request (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CategoryAPI.CategoryList(context.Background()).Start(start).Count(count).PageCursor(pageCursor).ParentId(parentId).Params(params).ResponseFields(responseFields).Exclude(exclude).StoreId(storeId).LangId(langId).CreatedFrom(createdFrom).CreatedTo(createdTo).ModifiedFrom(modifiedFrom).ModifiedTo(modifiedTo).Avail(avail).ProductType(productType).FindValue(findValue).FindWhere(findWhere).ReportRequestId(reportRequestId).DisableReportCache(disableReportCache).DisableCache(disableCache).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CategoryAPI.CategoryList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CategoryList`: ModelResponseCategoryList
	fmt.Fprintf(os.Stdout, "Response from `CategoryAPI.CategoryList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCategoryListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **start** | **int32** | This parameter sets the number from which you want to get entities | [default to 0]
 **count** | **int32** | This parameter sets the entity amount that has to be retrieved. Max allowed count&#x3D;250 | [default to 10]
 **pageCursor** | **string** | Used to retrieve entities via cursor-based pagination (it can&#39;t be used with any other filtering parameter) | 
 **parentId** | **string** | Retrieves categories specified by parent id | 
 **params** | **string** | Set this parameter in order to choose which entity fields you want to retrieve | [default to &quot;id,parent_id,name,description&quot;]
 **responseFields** | **string** | Set this parameter in order to choose which entity fields you want to retrieve | 
 **exclude** | **string** | Set this parameter in order to choose which entity fields you want to ignore. Works only if parameter &#x60;params&#x60; equal force_all | 
 **storeId** | **string** | Retrieves categories specified by store id | 
 **langId** | **string** | Retrieves categorys specified by language id | 
 **createdFrom** | **string** | Retrieve entities from their creation date | 
 **createdTo** | **string** | Retrieve entities to their creation date | 
 **modifiedFrom** | **string** | Retrieve entities from their modification date | 
 **modifiedTo** | **string** | Retrieve entities to their modification date | 
 **avail** | **bool** | Defines category&#39;s visibility status | [default to true]
 **productType** | **string** | A categorization for the product | 
 **findValue** | **string** | Entity search that is specified by some value | 
 **findWhere** | **string** | Category search that is specified by field | 
 **reportRequestId** | **string** | Report request id | 
 **disableReportCache** | **bool** | Disable report cache for current request | [default to false]
 **disableCache** | **bool** | Disable cache for current request | [default to false]

### Return type

[**ModelResponseCategoryList**](ModelResponseCategoryList.md)

### Authorization

[StoreKeyAuth](../README.md#StoreKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CategoryUnassign

> CartConfigUpdate200Response CategoryUnassign(ctx).CategoryId(categoryId).ProductId(productId).StoreId(storeId).Execute()

category.unassign



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
	categoryId := "6" // string | Defines category unassign, specified by category id
	productId := "10" // string | Defines category unassign to the product, specified by product id
	storeId := "1" // string | Store Id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CategoryAPI.CategoryUnassign(context.Background()).CategoryId(categoryId).ProductId(productId).StoreId(storeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CategoryAPI.CategoryUnassign``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CategoryUnassign`: CartConfigUpdate200Response
	fmt.Fprintf(os.Stdout, "Response from `CategoryAPI.CategoryUnassign`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCategoryUnassignRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **categoryId** | **string** | Defines category unassign, specified by category id | 
 **productId** | **string** | Defines category unassign to the product, specified by product id | 
 **storeId** | **string** | Store Id | 

### Return type

[**CartConfigUpdate200Response**](CartConfigUpdate200Response.md)

### Authorization

[StoreKeyAuth](../README.md#StoreKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CategoryUpdate

> AccountConfigUpdate200Response CategoryUpdate(ctx).Id(id).Name(name).ParentId(parentId).StoresIds(storesIds).Avail(avail).SortOrder(sortOrder).ModifiedTime(modifiedTime).Description(description).ShortDescription(shortDescription).MetaTitle(metaTitle).MetaDescription(metaDescription).MetaKeywords(metaKeywords).SeoUrl(seoUrl).LangId(langId).StoreId(storeId).Execute()

category.update



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
	id := "10" // string | Defines category update specified by category id
	name := "NEW Shoes" // string | Defines new category’s name (optional)
	parentId := "6" // string | Defines new parent category id (optional)
	storesIds := "1,2" // string | Update category in the stores that is specified by comma-separated stores' id (optional)
	avail := false // bool | Defines category's visibility status (optional)
	sortOrder := int32(2) // int32 | Sort number in the list (optional)
	modifiedTime := "2014-07-30 15:58:41" // string | Entity's date modification (optional)
	description := "New test category" // string | Defines new category's description (optional)
	shortDescription := "Short description. This is very short description" // string | Defines short description (optional)
	metaTitle := "category,test" // string | Defines unique meta title for each entity (optional)
	metaDescription := "category,test" // string | Defines unique meta description of a entity (optional)
	metaKeywords := "category,test" // string | Defines unique meta keywords for each entity (optional)
	seoUrl := "category,test" // string | Defines unique category's URL for SEO (optional)
	langId := "3" // string | Language id (optional)
	storeId := "1" // string | Store Id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CategoryAPI.CategoryUpdate(context.Background()).Id(id).Name(name).ParentId(parentId).StoresIds(storesIds).Avail(avail).SortOrder(sortOrder).ModifiedTime(modifiedTime).Description(description).ShortDescription(shortDescription).MetaTitle(metaTitle).MetaDescription(metaDescription).MetaKeywords(metaKeywords).SeoUrl(seoUrl).LangId(langId).StoreId(storeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CategoryAPI.CategoryUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CategoryUpdate`: AccountConfigUpdate200Response
	fmt.Fprintf(os.Stdout, "Response from `CategoryAPI.CategoryUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCategoryUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | Defines category update specified by category id | 
 **name** | **string** | Defines new category’s name | 
 **parentId** | **string** | Defines new parent category id | 
 **storesIds** | **string** | Update category in the stores that is specified by comma-separated stores&#39; id | 
 **avail** | **bool** | Defines category&#39;s visibility status | 
 **sortOrder** | **int32** | Sort number in the list | 
 **modifiedTime** | **string** | Entity&#39;s date modification | 
 **description** | **string** | Defines new category&#39;s description | 
 **shortDescription** | **string** | Defines short description | 
 **metaTitle** | **string** | Defines unique meta title for each entity | 
 **metaDescription** | **string** | Defines unique meta description of a entity | 
 **metaKeywords** | **string** | Defines unique meta keywords for each entity | 
 **seoUrl** | **string** | Defines unique category&#39;s URL for SEO | 
 **langId** | **string** | Language id | 
 **storeId** | **string** | Store Id | 

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

