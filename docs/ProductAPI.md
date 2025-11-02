# \ProductAPI

All URIs are relative to *https://api.api2cart.local.com/v1.1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ProductAdd**](ProductAPI.md#ProductAdd) | **Post** /product.add.json | product.add
[**ProductAddBatch**](ProductAPI.md#ProductAddBatch) | **Post** /product.add.batch.json | product.add.batch
[**ProductAttributeList**](ProductAPI.md#ProductAttributeList) | **Get** /product.attribute.list.json | product.attribute.list
[**ProductAttributeValueSet**](ProductAPI.md#ProductAttributeValueSet) | **Post** /product.attribute.value.set.json | product.attribute.value.set
[**ProductAttributeValueUnset**](ProductAPI.md#ProductAttributeValueUnset) | **Post** /product.attribute.value.unset.json | product.attribute.value.unset
[**ProductBrandList**](ProductAPI.md#ProductBrandList) | **Get** /product.brand.list.json | product.brand.list
[**ProductChildItemFind**](ProductAPI.md#ProductChildItemFind) | **Get** /product.child_item.find.json | product.child_item.find
[**ProductChildItemInfo**](ProductAPI.md#ProductChildItemInfo) | **Get** /product.child_item.info.json | product.child_item.info
[**ProductChildItemList**](ProductAPI.md#ProductChildItemList) | **Get** /product.child_item.list.json | product.child_item.list
[**ProductCount**](ProductAPI.md#ProductCount) | **Get** /product.count.json | product.count
[**ProductCurrencyAdd**](ProductAPI.md#ProductCurrencyAdd) | **Post** /product.currency.add.json | product.currency.add
[**ProductCurrencyList**](ProductAPI.md#ProductCurrencyList) | **Get** /product.currency.list.json | product.currency.list
[**ProductDelete**](ProductAPI.md#ProductDelete) | **Delete** /product.delete.json | product.delete
[**ProductDeleteBatch**](ProductAPI.md#ProductDeleteBatch) | **Post** /product.delete.batch.json | product.delete.batch
[**ProductFind**](ProductAPI.md#ProductFind) | **Get** /product.find.json | product.find
[**ProductImageAdd**](ProductAPI.md#ProductImageAdd) | **Post** /product.image.add.json | product.image.add
[**ProductImageDelete**](ProductAPI.md#ProductImageDelete) | **Delete** /product.image.delete.json | product.image.delete
[**ProductImageUpdate**](ProductAPI.md#ProductImageUpdate) | **Put** /product.image.update.json | product.image.update
[**ProductInfo**](ProductAPI.md#ProductInfo) | **Get** /product.info.json | product.info
[**ProductList**](ProductAPI.md#ProductList) | **Get** /product.list.json | product.list
[**ProductManufacturerAdd**](ProductAPI.md#ProductManufacturerAdd) | **Post** /product.manufacturer.add.json | product.manufacturer.add
[**ProductOptionAdd**](ProductAPI.md#ProductOptionAdd) | **Post** /product.option.add.json | product.option.add
[**ProductOptionAssign**](ProductAPI.md#ProductOptionAssign) | **Post** /product.option.assign.json | product.option.assign
[**ProductOptionDelete**](ProductAPI.md#ProductOptionDelete) | **Delete** /product.option.delete.json | product.option.delete
[**ProductOptionList**](ProductAPI.md#ProductOptionList) | **Get** /product.option.list.json | product.option.list
[**ProductOptionValueAdd**](ProductAPI.md#ProductOptionValueAdd) | **Post** /product.option.value.add.json | product.option.value.add
[**ProductOptionValueAssign**](ProductAPI.md#ProductOptionValueAssign) | **Post** /product.option.value.assign.json | product.option.value.assign
[**ProductOptionValueDelete**](ProductAPI.md#ProductOptionValueDelete) | **Delete** /product.option.value.delete.json | product.option.value.delete
[**ProductOptionValueUpdate**](ProductAPI.md#ProductOptionValueUpdate) | **Put** /product.option.value.update.json | product.option.value.update
[**ProductPriceAdd**](ProductAPI.md#ProductPriceAdd) | **Post** /product.price.add.json | product.price.add
[**ProductPriceDelete**](ProductAPI.md#ProductPriceDelete) | **Delete** /product.price.delete.json | product.price.delete
[**ProductPriceUpdate**](ProductAPI.md#ProductPriceUpdate) | **Put** /product.price.update.json | product.price.update
[**ProductReviewList**](ProductAPI.md#ProductReviewList) | **Get** /product.review.list.json | product.review.list
[**ProductStoreAssign**](ProductAPI.md#ProductStoreAssign) | **Post** /product.store.assign.json | product.store.assign
[**ProductTaxAdd**](ProductAPI.md#ProductTaxAdd) | **Post** /product.tax.add.json | product.tax.add
[**ProductUpdate**](ProductAPI.md#ProductUpdate) | **Put** /product.update.json | product.update
[**ProductUpdateBatch**](ProductAPI.md#ProductUpdateBatch) | **Post** /product.update.batch.json | product.update.batch
[**ProductVariantAdd**](ProductAPI.md#ProductVariantAdd) | **Post** /product.variant.add.json | product.variant.add
[**ProductVariantAddBatch**](ProductAPI.md#ProductVariantAddBatch) | **Post** /product.variant.add.batch.json | product.variant.add.batch
[**ProductVariantDelete**](ProductAPI.md#ProductVariantDelete) | **Delete** /product.variant.delete.json | product.variant.delete
[**ProductVariantDeleteBatch**](ProductAPI.md#ProductVariantDeleteBatch) | **Post** /product.variant.delete.batch.json | product.variant.delete.batch
[**ProductVariantImageAdd**](ProductAPI.md#ProductVariantImageAdd) | **Post** /product.variant.image.add.json | product.variant.image.add
[**ProductVariantImageDelete**](ProductAPI.md#ProductVariantImageDelete) | **Delete** /product.variant.image.delete.json | product.variant.image.delete
[**ProductVariantPriceAdd**](ProductAPI.md#ProductVariantPriceAdd) | **Post** /product.variant.price.add.json | product.variant.price.add
[**ProductVariantPriceDelete**](ProductAPI.md#ProductVariantPriceDelete) | **Delete** /product.variant.price.delete.json | product.variant.price.delete
[**ProductVariantPriceUpdate**](ProductAPI.md#ProductVariantPriceUpdate) | **Put** /product.variant.price.update.json | product.variant.price.update
[**ProductVariantUpdate**](ProductAPI.md#ProductVariantUpdate) | **Put** /product.variant.update.json | product.variant.update
[**ProductVariantUpdateBatch**](ProductAPI.md#ProductVariantUpdateBatch) | **Post** /product.variant.update.batch.json | product.variant.update.batch



## ProductAdd

> ProductAdd200Response ProductAdd(ctx).ProductAdd(productAdd).Execute()

product.add



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
	productAdd := *openapiclient.NewProductAdd("Bag", "bag_01", "Product description", float32(99.9)) // ProductAdd | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductAPI.ProductAdd(context.Background()).ProductAdd(productAdd).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductAPI.ProductAdd``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProductAdd`: ProductAdd200Response
	fmt.Fprintf(os.Stdout, "Response from `ProductAPI.ProductAdd`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProductAddRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **productAdd** | [**ProductAdd**](ProductAdd.md) |  | 

### Return type

[**ProductAdd200Response**](ProductAdd200Response.md)

### Authorization

[StoreKeyAuth](../README.md#StoreKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductAddBatch

> CategoryAddBatch200Response ProductAddBatch(ctx).ProductAddBatch(productAddBatch).Execute()

product.add.batch



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
	productAddBatch := *openapiclient.NewProductAddBatch([]openapiclient.ProductAddBatchPayloadInner{*openapiclient.NewProductAddBatchPayloadInner()}) // ProductAddBatch | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductAPI.ProductAddBatch(context.Background()).ProductAddBatch(productAddBatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductAPI.ProductAddBatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProductAddBatch`: CategoryAddBatch200Response
	fmt.Fprintf(os.Stdout, "Response from `ProductAPI.ProductAddBatch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProductAddBatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **productAddBatch** | [**ProductAddBatch**](ProductAddBatch.md) |  | 

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


## ProductAttributeList

> ModelResponseProductAttributeList ProductAttributeList(ctx).ProductId(productId).Start(start).Count(count).PageCursor(pageCursor).AttributeId(attributeId).VariantId(variantId).AttributeGroupId(attributeGroupId).LangId(langId).StoreId(storeId).SetName(setName).SortBy(sortBy).SortDirection(sortDirection).ResponseFields(responseFields).Params(params).Exclude(exclude).Execute()

product.attribute.list



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
	productId := "10" // string | Retrieves attributes specified by product id
	start := int32(0) // int32 | This parameter sets the number from which you want to get entities (optional) (default to 0)
	count := int32(20) // int32 | This parameter sets the entity amount that has to be retrieved. Max allowed count=250 (optional) (default to 10)
	pageCursor := "pageCursor_example" // string | Used to retrieve entities via cursor-based pagination (it can't be used with any other filtering parameter) (optional)
	attributeId := "156" // string | Retrieves info for specified attribute_id (optional)
	variantId := "45" // string | Defines product's variants specified by variant id (optional)
	attributeGroupId := "202" // string | Filter by attribute_group_id (optional)
	langId := "3" // string | Retrieves attributes specified by language id (optional)
	storeId := "1" // string | Retrieves attributes specified by store id (optional)
	setName := "Shoes" // string | Retrieves attributes specified by set_name in Magento (optional)
	sortBy := "value" // string | Set field to sort by (optional) (default to "attribute_id")
	sortDirection := "asc" // string | Set sorting direction (optional) (default to "asc")
	responseFields := "{pagination,result{attribute}}" // string | Set this parameter in order to choose which entity fields you want to retrieve (optional)
	params := "attribute_id,name" // string | Set this parameter in order to choose which entity fields you want to retrieve (optional) (default to "attribute_id,name")
	exclude := "attribute_id,name" // string | Set this parameter in order to choose which entity fields you want to ignore. Works only if parameter `params` equal force_all (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductAPI.ProductAttributeList(context.Background()).ProductId(productId).Start(start).Count(count).PageCursor(pageCursor).AttributeId(attributeId).VariantId(variantId).AttributeGroupId(attributeGroupId).LangId(langId).StoreId(storeId).SetName(setName).SortBy(sortBy).SortDirection(sortDirection).ResponseFields(responseFields).Params(params).Exclude(exclude).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductAPI.ProductAttributeList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProductAttributeList`: ModelResponseProductAttributeList
	fmt.Fprintf(os.Stdout, "Response from `ProductAPI.ProductAttributeList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProductAttributeListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **productId** | **string** | Retrieves attributes specified by product id | 
 **start** | **int32** | This parameter sets the number from which you want to get entities | [default to 0]
 **count** | **int32** | This parameter sets the entity amount that has to be retrieved. Max allowed count&#x3D;250 | [default to 10]
 **pageCursor** | **string** | Used to retrieve entities via cursor-based pagination (it can&#39;t be used with any other filtering parameter) | 
 **attributeId** | **string** | Retrieves info for specified attribute_id | 
 **variantId** | **string** | Defines product&#39;s variants specified by variant id | 
 **attributeGroupId** | **string** | Filter by attribute_group_id | 
 **langId** | **string** | Retrieves attributes specified by language id | 
 **storeId** | **string** | Retrieves attributes specified by store id | 
 **setName** | **string** | Retrieves attributes specified by set_name in Magento | 
 **sortBy** | **string** | Set field to sort by | [default to &quot;attribute_id&quot;]
 **sortDirection** | **string** | Set sorting direction | [default to &quot;asc&quot;]
 **responseFields** | **string** | Set this parameter in order to choose which entity fields you want to retrieve | 
 **params** | **string** | Set this parameter in order to choose which entity fields you want to retrieve | [default to &quot;attribute_id,name&quot;]
 **exclude** | **string** | Set this parameter in order to choose which entity fields you want to ignore. Works only if parameter &#x60;params&#x60; equal force_all | 

### Return type

[**ModelResponseProductAttributeList**](ModelResponseProductAttributeList.md)

### Authorization

[StoreKeyAuth](../README.md#StoreKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductAttributeValueSet

> ProductAttributeValueSet200Response ProductAttributeValueSet(ctx).ProductId(productId).AttributeId(attributeId).AttributeGroupId(attributeGroupId).AttributeName(attributeName).Value(value).ValueId(valueId).LangId(langId).StoreId(storeId).Execute()

product.attribute.value.set



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
	productId := "10" // string | Defines product id where the attribute should be added
	attributeId := "156" // string | Filter by attribute_id (optional)
	attributeGroupId := "202" // string | Filter by attribute_group_id (optional)
	attributeName := "Color" // string | Define attribute name (optional)
	value := "Red" // string | Define attribute value (optional)
	valueId := int32(22) // int32 | Define attribute value id (optional)
	langId := "3" // string | Language id (optional)
	storeId := "1" // string | Store Id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductAPI.ProductAttributeValueSet(context.Background()).ProductId(productId).AttributeId(attributeId).AttributeGroupId(attributeGroupId).AttributeName(attributeName).Value(value).ValueId(valueId).LangId(langId).StoreId(storeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductAPI.ProductAttributeValueSet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProductAttributeValueSet`: ProductAttributeValueSet200Response
	fmt.Fprintf(os.Stdout, "Response from `ProductAPI.ProductAttributeValueSet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProductAttributeValueSetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **productId** | **string** | Defines product id where the attribute should be added | 
 **attributeId** | **string** | Filter by attribute_id | 
 **attributeGroupId** | **string** | Filter by attribute_group_id | 
 **attributeName** | **string** | Define attribute name | 
 **value** | **string** | Define attribute value | 
 **valueId** | **int32** | Define attribute value id | 
 **langId** | **string** | Language id | 
 **storeId** | **string** | Store Id | 

### Return type

[**ProductAttributeValueSet200Response**](ProductAttributeValueSet200Response.md)

### Authorization

[StoreKeyAuth](../README.md#StoreKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductAttributeValueUnset

> ProductAttributeValueUnset200Response ProductAttributeValueUnset(ctx).ProductId(productId).AttributeId(attributeId).StoreId(storeId).IncludeDefault(includeDefault).Reindex(reindex).ClearCache(clearCache).Execute()

product.attribute.value.unset



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
	productId := "10" // string | Product id
	attributeId := "156" // string | Attribute Id
	storeId := "1" // string | Store Id (optional)
	includeDefault := true // bool | Boolean, whether or not to unset default value of the attribute, if applicable (optional) (default to false)
	reindex := false // bool | Is reindex required (optional) (default to true)
	clearCache := false // bool | Is cache clear required (optional) (default to true)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductAPI.ProductAttributeValueUnset(context.Background()).ProductId(productId).AttributeId(attributeId).StoreId(storeId).IncludeDefault(includeDefault).Reindex(reindex).ClearCache(clearCache).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductAPI.ProductAttributeValueUnset``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProductAttributeValueUnset`: ProductAttributeValueUnset200Response
	fmt.Fprintf(os.Stdout, "Response from `ProductAPI.ProductAttributeValueUnset`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProductAttributeValueUnsetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **productId** | **string** | Product id | 
 **attributeId** | **string** | Attribute Id | 
 **storeId** | **string** | Store Id | 
 **includeDefault** | **bool** | Boolean, whether or not to unset default value of the attribute, if applicable | [default to false]
 **reindex** | **bool** | Is reindex required | [default to true]
 **clearCache** | **bool** | Is cache clear required | [default to true]

### Return type

[**ProductAttributeValueUnset200Response**](ProductAttributeValueUnset200Response.md)

### Authorization

[StoreKeyAuth](../README.md#StoreKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductBrandList

> ModelResponseProductBrandList ProductBrandList(ctx).Start(start).Count(count).PageCursor(pageCursor).BrandIds(brandIds).CategoryId(categoryId).ParentId(parentId).StoreId(storeId).LangId(langId).FindWhere(findWhere).FindValue(findValue).CreatedFrom(createdFrom).CreatedTo(createdTo).ModifiedFrom(modifiedFrom).ModifiedTo(modifiedTo).ResponseFields(responseFields).Params(params).Exclude(exclude).Execute()

product.brand.list



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
	brandIds := "4,5" // string | Retrieves brands specified by brand ids (optional)
	categoryId := "6" // string | Retrieves product brands specified by category id (optional)
	parentId := "6" // string | Retrieves brands specified by parent id (optional)
	storeId := "1" // string | Store Id (optional)
	langId := "3" // string | Language id (optional)
	findWhere := "name" // string | Entity search that is specified by the comma-separated unique fields (optional)
	findValue := "Phone" // string | Entity search that is specified by some value (optional)
	createdFrom := "2010-07-29 13:45:52" // string | Retrieve entities from their creation date (optional)
	createdTo := "2100-08-29 13:45:52" // string | Retrieve entities to their creation date (optional)
	modifiedFrom := "2010-07-29 13:45:52" // string | Retrieve entities from their modification date (optional)
	modifiedTo := "2100-08-29 13:45:52" // string | Retrieve entities to their modification date (optional)
	responseFields := "{return_code,return_message,pagination,result}" // string | Set this parameter in order to choose which entity fields you want to retrieve (optional)
	params := "id,model,price,images" // string | Set this parameter in order to choose which entity fields you want to retrieve (optional) (default to "id,name,short_description,active,url")
	exclude := "false" // string | Set this parameter in order to choose which entity fields you want to ignore. Works only if parameter `params` equal force_all (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductAPI.ProductBrandList(context.Background()).Start(start).Count(count).PageCursor(pageCursor).BrandIds(brandIds).CategoryId(categoryId).ParentId(parentId).StoreId(storeId).LangId(langId).FindWhere(findWhere).FindValue(findValue).CreatedFrom(createdFrom).CreatedTo(createdTo).ModifiedFrom(modifiedFrom).ModifiedTo(modifiedTo).ResponseFields(responseFields).Params(params).Exclude(exclude).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductAPI.ProductBrandList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProductBrandList`: ModelResponseProductBrandList
	fmt.Fprintf(os.Stdout, "Response from `ProductAPI.ProductBrandList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProductBrandListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **start** | **int32** | This parameter sets the number from which you want to get entities | [default to 0]
 **count** | **int32** | This parameter sets the entity amount that has to be retrieved. Max allowed count&#x3D;250 | [default to 10]
 **pageCursor** | **string** | Used to retrieve entities via cursor-based pagination (it can&#39;t be used with any other filtering parameter) | 
 **brandIds** | **string** | Retrieves brands specified by brand ids | 
 **categoryId** | **string** | Retrieves product brands specified by category id | 
 **parentId** | **string** | Retrieves brands specified by parent id | 
 **storeId** | **string** | Store Id | 
 **langId** | **string** | Language id | 
 **findWhere** | **string** | Entity search that is specified by the comma-separated unique fields | 
 **findValue** | **string** | Entity search that is specified by some value | 
 **createdFrom** | **string** | Retrieve entities from their creation date | 
 **createdTo** | **string** | Retrieve entities to their creation date | 
 **modifiedFrom** | **string** | Retrieve entities from their modification date | 
 **modifiedTo** | **string** | Retrieve entities to their modification date | 
 **responseFields** | **string** | Set this parameter in order to choose which entity fields you want to retrieve | 
 **params** | **string** | Set this parameter in order to choose which entity fields you want to retrieve | [default to &quot;id,name,short_description,active,url&quot;]
 **exclude** | **string** | Set this parameter in order to choose which entity fields you want to ignore. Works only if parameter &#x60;params&#x60; equal force_all | 

### Return type

[**ModelResponseProductBrandList**](ModelResponseProductBrandList.md)

### Authorization

[StoreKeyAuth](../README.md#StoreKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductChildItemFind

> ProductChildItemFind200Response ProductChildItemFind(ctx).FindValue(findValue).FindWhere(findWhere).FindParams(findParams).StoreId(storeId).Execute()

product.child_item.find



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
	findValue := "bundled-item-123-" // string | Entity search that is specified by some value (optional)
	findWhere := "sku" // string | Entity search that is specified by the comma-separated unique fields (optional)
	findParams := "regex" // string | Entity search that is specified by comma-separated parameters (optional) (default to "whole_words")
	storeId := "1" // string | Store Id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductAPI.ProductChildItemFind(context.Background()).FindValue(findValue).FindWhere(findWhere).FindParams(findParams).StoreId(storeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductAPI.ProductChildItemFind``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProductChildItemFind`: ProductChildItemFind200Response
	fmt.Fprintf(os.Stdout, "Response from `ProductAPI.ProductChildItemFind`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProductChildItemFindRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **findValue** | **string** | Entity search that is specified by some value | 
 **findWhere** | **string** | Entity search that is specified by the comma-separated unique fields | 
 **findParams** | **string** | Entity search that is specified by comma-separated parameters | [default to &quot;whole_words&quot;]
 **storeId** | **string** | Store Id | 

### Return type

[**ProductChildItemFind200Response**](ProductChildItemFind200Response.md)

### Authorization

[StoreKeyAuth](../README.md#StoreKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductChildItemInfo

> ProductChildItemInfo200Response ProductChildItemInfo(ctx).ProductId(productId).Id(id).StoreId(storeId).LangId(langId).CurrencyId(currencyId).ResponseFields(responseFields).Params(params).Exclude(exclude).UseLatestApiVersion(useLatestApiVersion).Execute()

product.child_item.info



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
	productId := "10" // string | Filter by parent product id
	id := "10" // string | Entity id
	storeId := "1" // string | Store Id (optional)
	langId := "3" // string | Language id (optional)
	currencyId := "usd" // string | Currency Id (optional)
	responseFields := "{result{id,parent_id,sku,upc,images,combination}}" // string | Set this parameter in order to choose which entity fields you want to retrieve (optional)
	params := "id,model,price,images" // string | Set this parameter in order to choose which entity fields you want to retrieve (optional) (default to "force_all")
	exclude := "false" // string | Set this parameter in order to choose which entity fields you want to ignore. Works only if parameter `params` equal force_all (optional)
	useLatestApiVersion := true // bool | Use the latest platform API version (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductAPI.ProductChildItemInfo(context.Background()).ProductId(productId).Id(id).StoreId(storeId).LangId(langId).CurrencyId(currencyId).ResponseFields(responseFields).Params(params).Exclude(exclude).UseLatestApiVersion(useLatestApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductAPI.ProductChildItemInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProductChildItemInfo`: ProductChildItemInfo200Response
	fmt.Fprintf(os.Stdout, "Response from `ProductAPI.ProductChildItemInfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProductChildItemInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **productId** | **string** | Filter by parent product id | 
 **id** | **string** | Entity id | 
 **storeId** | **string** | Store Id | 
 **langId** | **string** | Language id | 
 **currencyId** | **string** | Currency Id | 
 **responseFields** | **string** | Set this parameter in order to choose which entity fields you want to retrieve | 
 **params** | **string** | Set this parameter in order to choose which entity fields you want to retrieve | [default to &quot;force_all&quot;]
 **exclude** | **string** | Set this parameter in order to choose which entity fields you want to ignore. Works only if parameter &#x60;params&#x60; equal force_all | 
 **useLatestApiVersion** | **bool** | Use the latest platform API version | [default to false]

### Return type

[**ProductChildItemInfo200Response**](ProductChildItemInfo200Response.md)

### Authorization

[StoreKeyAuth](../README.md#StoreKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductChildItemList

> ModelResponseProductChildItemList ProductChildItemList(ctx).Start(start).Count(count).PageCursor(pageCursor).ProductId(productId).ProductIds(productIds).Sku(sku).StoreId(storeId).LangId(langId).CurrencyId(currencyId).AvailSale(availSale).FindValue(findValue).FindWhere(findWhere).CreatedFrom(createdFrom).CreatedTo(createdTo).ModifiedFrom(modifiedFrom).ModifiedTo(modifiedTo).ReturnGlobal(returnGlobal).ResponseFields(responseFields).Params(params).Exclude(exclude).ReportRequestId(reportRequestId).DisableReportCache(disableReportCache).UseLatestApiVersion(useLatestApiVersion).Execute()

product.child_item.list



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
	pageCursor := "pageCursor_example" // string | Used to retrieve products child items via cursor-based pagination (it can't be used with any other filtering parameter) (optional)
	productId := "10" // string | Filter by parent product id (optional)
	productIds := "4,5" // string | Filter by parent product ids (optional)
	sku := "bag_01" // string | Filter by products variant's sku (optional)
	storeId := "1" // string | Store Id (optional)
	langId := "3" // string | Language id (optional)
	currencyId := "usd" // string | Currency Id (optional)
	availSale := false // bool | Specifies the set of available/not available products for sale (optional)
	findValue := "bundled-item-123-" // string | Entity search that is specified by some value (optional)
	findWhere := "sku" // string | Child products search that is specified by field (optional)
	createdFrom := "2010-07-29 13:45:52" // string | Retrieve entities from their creation date (optional)
	createdTo := "2100-08-29 13:45:52" // string | Retrieve entities to their creation date (optional)
	modifiedFrom := "2010-07-29 13:45:52" // string | Retrieve entities from their modification date (optional)
	modifiedTo := "2100-08-29 13:45:52" // string | Retrieve entities to their modification date (optional)
	returnGlobal := false // bool | Determines the type of products to be returned. If set to 'true', only global products will be returned; if set to 'false', only local products will be returned. (optional) (default to false)
	responseFields := "{result{children{id,parent_id,sku,upc,images,combination}}}" // string | Set this parameter in order to choose which entity fields you want to retrieve (optional)
	params := "id,model,price,images" // string | Set this parameter in order to choose which entity fields you want to retrieve (optional) (default to "force_all")
	exclude := "false" // string | Set this parameter in order to choose which entity fields you want to ignore. Works only if parameter `params` equal force_all (optional)
	reportRequestId := "105245017661" // string | Report request id (optional)
	disableReportCache := false // bool | Disable report cache for current request (optional) (default to false)
	useLatestApiVersion := true // bool | Use the latest platform API version (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductAPI.ProductChildItemList(context.Background()).Start(start).Count(count).PageCursor(pageCursor).ProductId(productId).ProductIds(productIds).Sku(sku).StoreId(storeId).LangId(langId).CurrencyId(currencyId).AvailSale(availSale).FindValue(findValue).FindWhere(findWhere).CreatedFrom(createdFrom).CreatedTo(createdTo).ModifiedFrom(modifiedFrom).ModifiedTo(modifiedTo).ReturnGlobal(returnGlobal).ResponseFields(responseFields).Params(params).Exclude(exclude).ReportRequestId(reportRequestId).DisableReportCache(disableReportCache).UseLatestApiVersion(useLatestApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductAPI.ProductChildItemList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProductChildItemList`: ModelResponseProductChildItemList
	fmt.Fprintf(os.Stdout, "Response from `ProductAPI.ProductChildItemList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProductChildItemListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **start** | **int32** | This parameter sets the number from which you want to get entities | [default to 0]
 **count** | **int32** | This parameter sets the entity amount that has to be retrieved. Max allowed count&#x3D;250 | [default to 10]
 **pageCursor** | **string** | Used to retrieve products child items via cursor-based pagination (it can&#39;t be used with any other filtering parameter) | 
 **productId** | **string** | Filter by parent product id | 
 **productIds** | **string** | Filter by parent product ids | 
 **sku** | **string** | Filter by products variant&#39;s sku | 
 **storeId** | **string** | Store Id | 
 **langId** | **string** | Language id | 
 **currencyId** | **string** | Currency Id | 
 **availSale** | **bool** | Specifies the set of available/not available products for sale | 
 **findValue** | **string** | Entity search that is specified by some value | 
 **findWhere** | **string** | Child products search that is specified by field | 
 **createdFrom** | **string** | Retrieve entities from their creation date | 
 **createdTo** | **string** | Retrieve entities to their creation date | 
 **modifiedFrom** | **string** | Retrieve entities from their modification date | 
 **modifiedTo** | **string** | Retrieve entities to their modification date | 
 **returnGlobal** | **bool** | Determines the type of products to be returned. If set to &#39;true&#39;, only global products will be returned; if set to &#39;false&#39;, only local products will be returned. | [default to false]
 **responseFields** | **string** | Set this parameter in order to choose which entity fields you want to retrieve | 
 **params** | **string** | Set this parameter in order to choose which entity fields you want to retrieve | [default to &quot;force_all&quot;]
 **exclude** | **string** | Set this parameter in order to choose which entity fields you want to ignore. Works only if parameter &#x60;params&#x60; equal force_all | 
 **reportRequestId** | **string** | Report request id | 
 **disableReportCache** | **bool** | Disable report cache for current request | [default to false]
 **useLatestApiVersion** | **bool** | Use the latest platform API version | [default to false]

### Return type

[**ModelResponseProductChildItemList**](ModelResponseProductChildItemList.md)

### Authorization

[StoreKeyAuth](../README.md#StoreKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductCount

> ProductCount200Response ProductCount(ctx).Sku(sku).ProductIds(productIds).SinceId(sinceId).CategoriesIds(categoriesIds).CategoryId(categoryId).StoreId(storeId).LangId(langId).AvailView(availView).AvailSale(availSale).CreatedFrom(createdFrom).CreatedTo(createdTo).ModifiedFrom(modifiedFrom).ModifiedTo(modifiedTo).BrandName(brandName).ManufacturerId(manufacturerId).ProductAttributes(productAttributes).Status(status).Type_(type_).Visible(visible).FindValue(findValue).FindWhere(findWhere).ReportRequestId(reportRequestId).ReturnGlobal(returnGlobal).DisableReportCache(disableReportCache).UseLatestApiVersion(useLatestApiVersion).Execute()

product.count



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
	sku := "bag_01" // string | Filter by product's sku (optional)
	productIds := "4,5" // string | Counts products specified by product ids (optional)
	sinceId := "56" // string | Retrieve entities starting from the specified id. (optional)
	categoriesIds := "23,56" // string | Defines product add that is specified by comma-separated categories id (optional)
	categoryId := "6" // string | Counts products specified by category id (optional)
	storeId := "1" // string | Counts products specified by store id (optional)
	langId := "3" // string | Counts products specified by language id (optional)
	availView := true // bool | Specifies the set of visible/invisible products (optional)
	availSale := false // bool | Specifies the set of available/not available products for sale (optional)
	createdFrom := "2010-07-29 13:45:52" // string | Retrieve entities from their creation date (optional)
	createdTo := "2100-08-29 13:45:52" // string | Retrieve entities to their creation date (optional)
	modifiedFrom := "2010-07-29 13:45:52" // string | Retrieve entities from their modification date (optional)
	modifiedTo := "2100-08-29 13:45:52" // string | Retrieve entities to their modification date (optional)
	brandName := "Abidas" // string | Retrieves brands specified by brand name (optional)
	manufacturerId := "1" // string | Defines product's manufacturer by manufacturer_id (optional)
	productAttributes := []string{"Inner_example"} // []string | Defines product attributes (optional)
	status := "disabled" // string | Defines product's status (optional)
	type_ := "simple" // string | Defines products's type (optional)
	visible := "everywhere" // string | Filter items by visibility status (optional) (default to "everywhere")
	findValue := "Phone" // string | Entity search that is specified by some value (optional)
	findWhere := "name" // string | Counts products that are searched specified by field (optional)
	reportRequestId := "105245017661" // string | Report request id (optional)
	returnGlobal := false // bool | Determines the type of products to be returned. If set to 'true', only global products will be returned; if set to 'false', only local products will be returned. (optional) (default to false)
	disableReportCache := false // bool | Disable report cache for current request (optional) (default to false)
	useLatestApiVersion := true // bool | Use the latest platform API version (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductAPI.ProductCount(context.Background()).Sku(sku).ProductIds(productIds).SinceId(sinceId).CategoriesIds(categoriesIds).CategoryId(categoryId).StoreId(storeId).LangId(langId).AvailView(availView).AvailSale(availSale).CreatedFrom(createdFrom).CreatedTo(createdTo).ModifiedFrom(modifiedFrom).ModifiedTo(modifiedTo).BrandName(brandName).ManufacturerId(manufacturerId).ProductAttributes(productAttributes).Status(status).Type_(type_).Visible(visible).FindValue(findValue).FindWhere(findWhere).ReportRequestId(reportRequestId).ReturnGlobal(returnGlobal).DisableReportCache(disableReportCache).UseLatestApiVersion(useLatestApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductAPI.ProductCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProductCount`: ProductCount200Response
	fmt.Fprintf(os.Stdout, "Response from `ProductAPI.ProductCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProductCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sku** | **string** | Filter by product&#39;s sku | 
 **productIds** | **string** | Counts products specified by product ids | 
 **sinceId** | **string** | Retrieve entities starting from the specified id. | 
 **categoriesIds** | **string** | Defines product add that is specified by comma-separated categories id | 
 **categoryId** | **string** | Counts products specified by category id | 
 **storeId** | **string** | Counts products specified by store id | 
 **langId** | **string** | Counts products specified by language id | 
 **availView** | **bool** | Specifies the set of visible/invisible products | 
 **availSale** | **bool** | Specifies the set of available/not available products for sale | 
 **createdFrom** | **string** | Retrieve entities from their creation date | 
 **createdTo** | **string** | Retrieve entities to their creation date | 
 **modifiedFrom** | **string** | Retrieve entities from their modification date | 
 **modifiedTo** | **string** | Retrieve entities to their modification date | 
 **brandName** | **string** | Retrieves brands specified by brand name | 
 **manufacturerId** | **string** | Defines product&#39;s manufacturer by manufacturer_id | 
 **productAttributes** | **[]string** | Defines product attributes | 
 **status** | **string** | Defines product&#39;s status | 
 **type_** | **string** | Defines products&#39;s type | 
 **visible** | **string** | Filter items by visibility status | [default to &quot;everywhere&quot;]
 **findValue** | **string** | Entity search that is specified by some value | 
 **findWhere** | **string** | Counts products that are searched specified by field | 
 **reportRequestId** | **string** | Report request id | 
 **returnGlobal** | **bool** | Determines the type of products to be returned. If set to &#39;true&#39;, only global products will be returned; if set to &#39;false&#39;, only local products will be returned. | [default to false]
 **disableReportCache** | **bool** | Disable report cache for current request | [default to false]
 **useLatestApiVersion** | **bool** | Use the latest platform API version | [default to false]

### Return type

[**ProductCount200Response**](ProductCount200Response.md)

### Authorization

[StoreKeyAuth](../README.md#StoreKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductCurrencyAdd

> ProductCurrencyAdd200Response ProductCurrencyAdd(ctx).Iso3(iso3).Rate(rate).Name(name).Avail(avail).SymbolLeft(symbolLeft).SymbolRight(symbolRight).Default_(default_).Execute()

product.currency.add



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
	iso3 := "USD" // string | Specifies standardized currency code
	rate := float32(1) // float32 | Defines the numerical identifier against to the major currency
	name := "US Dollar" // string | Defines currency's name (optional)
	avail := false // bool | Specifies whether the currency is available (optional) (default to true)
	symbolLeft := "$" // string | Defines the symbol that is located before the currency (optional)
	symbolRight := "грн" // string | Defines the symbol that is located after the currency (optional)
	default_ := true // bool | Specifies currency's default meaning (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductAPI.ProductCurrencyAdd(context.Background()).Iso3(iso3).Rate(rate).Name(name).Avail(avail).SymbolLeft(symbolLeft).SymbolRight(symbolRight).Default_(default_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductAPI.ProductCurrencyAdd``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProductCurrencyAdd`: ProductCurrencyAdd200Response
	fmt.Fprintf(os.Stdout, "Response from `ProductAPI.ProductCurrencyAdd`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProductCurrencyAddRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iso3** | **string** | Specifies standardized currency code | 
 **rate** | **float32** | Defines the numerical identifier against to the major currency | 
 **name** | **string** | Defines currency&#39;s name | 
 **avail** | **bool** | Specifies whether the currency is available | [default to true]
 **symbolLeft** | **string** | Defines the symbol that is located before the currency | 
 **symbolRight** | **string** | Defines the symbol that is located after the currency | 
 **default_** | **bool** | Specifies currency&#39;s default meaning | [default to false]

### Return type

[**ProductCurrencyAdd200Response**](ProductCurrencyAdd200Response.md)

### Authorization

[StoreKeyAuth](../README.md#StoreKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductCurrencyList

> ModelResponseProductCurrencyList ProductCurrencyList(ctx).Start(start).Count(count).PageCursor(pageCursor).Default_(default_).Avail(avail).ResponseFields(responseFields).Params(params).Exclude(exclude).Execute()

product.currency.list



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
	default_ := true // bool | Specifies the set of default/not default currencies (optional)
	avail := false // bool | Specifies the set of available/not available currencies (optional)
	responseFields := "{return_message,pagination,result{currency}}" // string | Set this parameter in order to choose which entity fields you want to retrieve (optional)
	params := "name,iso3,default,avail" // string | Set this parameter in order to choose which entity fields you want to retrieve (optional) (default to "name,iso3,default,avail")
	exclude := "name,iso3,default,avail" // string | Set this parameter in order to choose which entity fields you want to ignore. Works only if parameter `params` equal force_all (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductAPI.ProductCurrencyList(context.Background()).Start(start).Count(count).PageCursor(pageCursor).Default_(default_).Avail(avail).ResponseFields(responseFields).Params(params).Exclude(exclude).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductAPI.ProductCurrencyList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProductCurrencyList`: ModelResponseProductCurrencyList
	fmt.Fprintf(os.Stdout, "Response from `ProductAPI.ProductCurrencyList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProductCurrencyListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **start** | **int32** | This parameter sets the number from which you want to get entities | [default to 0]
 **count** | **int32** | This parameter sets the entity amount that has to be retrieved. Max allowed count&#x3D;250 | [default to 10]
 **pageCursor** | **string** | Used to retrieve entities via cursor-based pagination (it can&#39;t be used with any other filtering parameter) | 
 **default_** | **bool** | Specifies the set of default/not default currencies | 
 **avail** | **bool** | Specifies the set of available/not available currencies | 
 **responseFields** | **string** | Set this parameter in order to choose which entity fields you want to retrieve | 
 **params** | **string** | Set this parameter in order to choose which entity fields you want to retrieve | [default to &quot;name,iso3,default,avail&quot;]
 **exclude** | **string** | Set this parameter in order to choose which entity fields you want to ignore. Works only if parameter &#x60;params&#x60; equal force_all | 

### Return type

[**ModelResponseProductCurrencyList**](ModelResponseProductCurrencyList.md)

### Authorization

[StoreKeyAuth](../README.md#StoreKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductDelete

> CustomerDelete200Response ProductDelete(ctx).Id(id).StoreId(storeId).Execute()

product.delete



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
	id := "10" // string | Product id that will be removed
	storeId := "1" // string | Store Id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductAPI.ProductDelete(context.Background()).Id(id).StoreId(storeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductAPI.ProductDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProductDelete`: CustomerDelete200Response
	fmt.Fprintf(os.Stdout, "Response from `ProductAPI.ProductDelete`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProductDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | Product id that will be removed | 
 **storeId** | **string** | Store Id | 

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


## ProductDeleteBatch

> CategoryAddBatch200Response ProductDeleteBatch(ctx).ProductDeleteBatch(productDeleteBatch).Execute()

product.delete.batch



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
	productDeleteBatch := *openapiclient.NewProductDeleteBatch([]openapiclient.ProductDeleteBatchPayloadInner{*openapiclient.NewProductDeleteBatchPayloadInner()}) // ProductDeleteBatch | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductAPI.ProductDeleteBatch(context.Background()).ProductDeleteBatch(productDeleteBatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductAPI.ProductDeleteBatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProductDeleteBatch`: CategoryAddBatch200Response
	fmt.Fprintf(os.Stdout, "Response from `ProductAPI.ProductDeleteBatch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProductDeleteBatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **productDeleteBatch** | [**ProductDeleteBatch**](ProductDeleteBatch.md) |  | 

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


## ProductFind

> ProductFind200Response ProductFind(ctx).FindValue(findValue).FindWhere(findWhere).FindParams(findParams).FindWhat(findWhat).LangId(langId).StoreId(storeId).Execute()

product.find



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
	findValue := "Simple" // string | Entity search that is specified by some value
	findWhere := "name" // string | Entity search that is specified by the comma-separated unique fields (optional) (default to "name")
	findParams := "regex" // string | Entity search that is specified by comma-separated parameters (optional) (default to "whole_words")
	findWhat := "each" // string | Parameter's value specifies the entity that has to be found (optional) (default to "product")
	langId := "3" // string | Search products specified by language id (optional)
	storeId := "1" // string | Store Id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductAPI.ProductFind(context.Background()).FindValue(findValue).FindWhere(findWhere).FindParams(findParams).FindWhat(findWhat).LangId(langId).StoreId(storeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductAPI.ProductFind``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProductFind`: ProductFind200Response
	fmt.Fprintf(os.Stdout, "Response from `ProductAPI.ProductFind`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProductFindRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **findValue** | **string** | Entity search that is specified by some value | 
 **findWhere** | **string** | Entity search that is specified by the comma-separated unique fields | [default to &quot;name&quot;]
 **findParams** | **string** | Entity search that is specified by comma-separated parameters | [default to &quot;whole_words&quot;]
 **findWhat** | **string** | Parameter&#39;s value specifies the entity that has to be found | [default to &quot;product&quot;]
 **langId** | **string** | Search products specified by language id | 
 **storeId** | **string** | Store Id | 

### Return type

[**ProductFind200Response**](ProductFind200Response.md)

### Authorization

[StoreKeyAuth](../README.md#StoreKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductImageAdd

> ProductImageAdd200Response ProductImageAdd(ctx).ProductImageAdd(productImageAdd).Execute()

product.image.add



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
	productImageAdd := *openapiclient.NewProductImageAdd("base,small", "bag-gray.png") // ProductImageAdd | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductAPI.ProductImageAdd(context.Background()).ProductImageAdd(productImageAdd).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductAPI.ProductImageAdd``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProductImageAdd`: ProductImageAdd200Response
	fmt.Fprintf(os.Stdout, "Response from `ProductAPI.ProductImageAdd`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProductImageAddRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **productImageAdd** | [**ProductImageAdd**](ProductImageAdd.md) |  | 

### Return type

[**ProductImageAdd200Response**](ProductImageAdd200Response.md)

### Authorization

[StoreKeyAuth](../README.md#StoreKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductImageDelete

> AttributeDelete200Response ProductImageDelete(ctx).ProductId(productId).Id(id).StoreId(storeId).Execute()

product.image.delete



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
	productId := "10" // string | Defines product id where the image should be deleted
	id := "10" // string | Entity id
	storeId := "1" // string | Store Id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductAPI.ProductImageDelete(context.Background()).ProductId(productId).Id(id).StoreId(storeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductAPI.ProductImageDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProductImageDelete`: AttributeDelete200Response
	fmt.Fprintf(os.Stdout, "Response from `ProductAPI.ProductImageDelete`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProductImageDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **productId** | **string** | Defines product id where the image should be deleted | 
 **id** | **string** | Entity id | 
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


## ProductImageUpdate

> ProductImageUpdate200Response ProductImageUpdate(ctx).ProductId(productId).Id(id).VariantIds(variantIds).StoreId(storeId).LangId(langId).ImageName(imageName).Type_(type_).Label(label).Position(position).Hidden(hidden).Execute()

product.image.update



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
	productId := "10" // string | Defines product id where the image should be updated
	id := "10" // string | Defines image update specified by image id
	variantIds := "1,2,3,4,5" // string | Defines product's variants ids (optional)
	storeId := "1" // string | Store Id (optional)
	langId := "3" // string | Language id (optional)
	imageName := "data/product/main/product_69_bag-gray.png" // string | Defines image's name (optional)
	type_ := "thumbnail" // string | Defines image's types that are specified by comma-separated list (optional) (default to "additional")
	label := "This cool image" // string | Defines alternative text that has to be attached to the picture (optional)
	position := int32(5) // int32 | Defines image’s position in the list (optional)
	hidden := true // bool | Define is hide image (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductAPI.ProductImageUpdate(context.Background()).ProductId(productId).Id(id).VariantIds(variantIds).StoreId(storeId).LangId(langId).ImageName(imageName).Type_(type_).Label(label).Position(position).Hidden(hidden).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductAPI.ProductImageUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProductImageUpdate`: ProductImageUpdate200Response
	fmt.Fprintf(os.Stdout, "Response from `ProductAPI.ProductImageUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProductImageUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **productId** | **string** | Defines product id where the image should be updated | 
 **id** | **string** | Defines image update specified by image id | 
 **variantIds** | **string** | Defines product&#39;s variants ids | 
 **storeId** | **string** | Store Id | 
 **langId** | **string** | Language id | 
 **imageName** | **string** | Defines image&#39;s name | 
 **type_** | **string** | Defines image&#39;s types that are specified by comma-separated list | [default to &quot;additional&quot;]
 **label** | **string** | Defines alternative text that has to be attached to the picture | 
 **position** | **int32** | Defines image’s position in the list | 
 **hidden** | **bool** | Define is hide image | 

### Return type

[**ProductImageUpdate200Response**](ProductImageUpdate200Response.md)

### Authorization

[StoreKeyAuth](../README.md#StoreKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductInfo

> ProductInfo200Response ProductInfo(ctx).Id(id).StoreId(storeId).LangId(langId).CurrencyId(currencyId).ResponseFields(responseFields).Params(params).Exclude(exclude).ReportRequestId(reportRequestId).DisableReportCache(disableReportCache).UseLatestApiVersion(useLatestApiVersion).Execute()

product.info



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
	id := "10" // string | Retrieves product's info specified by product id
	storeId := "1" // string | Retrieves product info specified by store id (optional)
	langId := "3" // string | Retrieves product info specified by language id (optional)
	currencyId := "usd" // string | Currency Id (optional)
	responseFields := "{result{id,name,price,images}}" // string | Set this parameter in order to choose which entity fields you want to retrieve (optional)
	params := "id,model,price,images" // string | Set this parameter in order to choose which entity fields you want to retrieve (optional) (default to "id,name,description,price,categories_ids")
	exclude := "false" // string | Set this parameter in order to choose which entity fields you want to ignore. Works only if parameter `params` equal force_all (optional)
	reportRequestId := "105245017661" // string | Report request id (optional)
	disableReportCache := false // bool | Disable report cache for current request (optional) (default to false)
	useLatestApiVersion := true // bool | Use the latest platform API version (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductAPI.ProductInfo(context.Background()).Id(id).StoreId(storeId).LangId(langId).CurrencyId(currencyId).ResponseFields(responseFields).Params(params).Exclude(exclude).ReportRequestId(reportRequestId).DisableReportCache(disableReportCache).UseLatestApiVersion(useLatestApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductAPI.ProductInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProductInfo`: ProductInfo200Response
	fmt.Fprintf(os.Stdout, "Response from `ProductAPI.ProductInfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProductInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | Retrieves product&#39;s info specified by product id | 
 **storeId** | **string** | Retrieves product info specified by store id | 
 **langId** | **string** | Retrieves product info specified by language id | 
 **currencyId** | **string** | Currency Id | 
 **responseFields** | **string** | Set this parameter in order to choose which entity fields you want to retrieve | 
 **params** | **string** | Set this parameter in order to choose which entity fields you want to retrieve | [default to &quot;id,name,description,price,categories_ids&quot;]
 **exclude** | **string** | Set this parameter in order to choose which entity fields you want to ignore. Works only if parameter &#x60;params&#x60; equal force_all | 
 **reportRequestId** | **string** | Report request id | 
 **disableReportCache** | **bool** | Disable report cache for current request | [default to false]
 **useLatestApiVersion** | **bool** | Use the latest platform API version | [default to false]

### Return type

[**ProductInfo200Response**](ProductInfo200Response.md)

### Authorization

[StoreKeyAuth](../README.md#StoreKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductList

> ModelResponseProductList ProductList(ctx).Start(start).Count(count).PageCursor(pageCursor).ProductIds(productIds).SinceId(sinceId).CategoriesIds(categoriesIds).CategoryId(categoryId).StoreId(storeId).LangId(langId).CurrencyId(currencyId).AvailView(availView).AvailSale(availSale).CreatedFrom(createdFrom).CreatedTo(createdTo).ModifiedFrom(modifiedFrom).ModifiedTo(modifiedTo).Sku(sku).BrandName(brandName).ProductAttributes(productAttributes).ManufacturerId(manufacturerId).Status(status).Type_(type_).Visible(visible).FindValue(findValue).FindWhere(findWhere).ReturnGlobal(returnGlobal).Params(params).ResponseFields(responseFields).Exclude(exclude).SortBy(sortBy).SortDirection(sortDirection).ReportRequestId(reportRequestId).DisableCache(disableCache).DisableReportCache(disableReportCache).UseLatestApiVersion(useLatestApiVersion).ProductType(productType).Execute()

product.list



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
	pageCursor := "pageCursor_example" // string | Used to retrieve products via cursor-based pagination (it can't be used with any other filtering parameter) (optional)
	productIds := "4,5" // string | Retrieves products specified by product ids (optional)
	sinceId := "56" // string | Retrieve entities starting from the specified id. (optional)
	categoriesIds := "23,56" // string | Retrieves products specified by categories ids (optional)
	categoryId := "6" // string | Retrieves products specified by category id (optional)
	storeId := "1" // string | Retrieves products specified by store id (optional)
	langId := "3" // string | Retrieves products specified by language id (optional)
	currencyId := "usd" // string | Currency Id (optional)
	availView := true // bool | Specifies the set of visible/invisible products (optional)
	availSale := false // bool | Specifies the set of available/not available products for sale (optional)
	createdFrom := "2010-07-29 13:45:52" // string | Retrieve entities from their creation date (optional)
	createdTo := "2100-08-29 13:45:52" // string | Retrieve entities to their creation date (optional)
	modifiedFrom := "2010-07-29 13:45:52" // string | Retrieve entities from their modification date (optional)
	modifiedTo := "2100-08-29 13:45:52" // string | Retrieve entities to their modification date (optional)
	sku := "bag_01" // string | Filter by product's sku (optional)
	brandName := "Abidas" // string | Retrieves brands specified by brand name (optional)
	productAttributes := []string{"Inner_example"} // []string | Defines product attributes (optional)
	manufacturerId := "1" // string | Defines product's manufacturer by manufacturer_id (optional)
	status := "disabled" // string | Defines product's status (optional)
	type_ := "simple" // string | Defines products's type (optional)
	visible := "everywhere" // string | Filter items by visibility status (optional) (default to "everywhere")
	findValue := "Phone" // string | Entity search that is specified by some value (optional)
	findWhere := "name" // string | Product search that is specified by field (optional)
	returnGlobal := false // bool | Determines the type of products to be returned. If set to 'true', only global products will be returned; if set to 'false', only local products will be returned. (optional) (default to false)
	params := "id,model,price,images" // string | Set this parameter in order to choose which entity fields you want to retrieve (optional) (default to "id,name,description,price,categories_ids")
	responseFields := "{return_code,pagination,result{product{id,name,price,images}}}" // string | Set this parameter in order to choose which entity fields you want to retrieve (optional)
	exclude := "false" // string | Set this parameter in order to choose which entity fields you want to ignore. Works only if parameter `params` equal force_all (optional)
	sortBy := "value_id" // string | Set field to sort by (optional) (default to "id")
	sortDirection := "asc" // string | Set sorting direction (optional) (default to "asc")
	reportRequestId := "105245017661" // string | Report request id (optional)
	disableCache := false // bool | Disable cache for current request (optional) (default to false)
	disableReportCache := false // bool | Disable report cache for current request (optional) (default to false)
	useLatestApiVersion := true // bool | Use the latest platform API version (optional) (default to false)
	productType := "BICYCLE" // string | A categorization for the product (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductAPI.ProductList(context.Background()).Start(start).Count(count).PageCursor(pageCursor).ProductIds(productIds).SinceId(sinceId).CategoriesIds(categoriesIds).CategoryId(categoryId).StoreId(storeId).LangId(langId).CurrencyId(currencyId).AvailView(availView).AvailSale(availSale).CreatedFrom(createdFrom).CreatedTo(createdTo).ModifiedFrom(modifiedFrom).ModifiedTo(modifiedTo).Sku(sku).BrandName(brandName).ProductAttributes(productAttributes).ManufacturerId(manufacturerId).Status(status).Type_(type_).Visible(visible).FindValue(findValue).FindWhere(findWhere).ReturnGlobal(returnGlobal).Params(params).ResponseFields(responseFields).Exclude(exclude).SortBy(sortBy).SortDirection(sortDirection).ReportRequestId(reportRequestId).DisableCache(disableCache).DisableReportCache(disableReportCache).UseLatestApiVersion(useLatestApiVersion).ProductType(productType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductAPI.ProductList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProductList`: ModelResponseProductList
	fmt.Fprintf(os.Stdout, "Response from `ProductAPI.ProductList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProductListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **start** | **int32** | This parameter sets the number from which you want to get entities | [default to 0]
 **count** | **int32** | This parameter sets the entity amount that has to be retrieved. Max allowed count&#x3D;250 | [default to 10]
 **pageCursor** | **string** | Used to retrieve products via cursor-based pagination (it can&#39;t be used with any other filtering parameter) | 
 **productIds** | **string** | Retrieves products specified by product ids | 
 **sinceId** | **string** | Retrieve entities starting from the specified id. | 
 **categoriesIds** | **string** | Retrieves products specified by categories ids | 
 **categoryId** | **string** | Retrieves products specified by category id | 
 **storeId** | **string** | Retrieves products specified by store id | 
 **langId** | **string** | Retrieves products specified by language id | 
 **currencyId** | **string** | Currency Id | 
 **availView** | **bool** | Specifies the set of visible/invisible products | 
 **availSale** | **bool** | Specifies the set of available/not available products for sale | 
 **createdFrom** | **string** | Retrieve entities from their creation date | 
 **createdTo** | **string** | Retrieve entities to their creation date | 
 **modifiedFrom** | **string** | Retrieve entities from their modification date | 
 **modifiedTo** | **string** | Retrieve entities to their modification date | 
 **sku** | **string** | Filter by product&#39;s sku | 
 **brandName** | **string** | Retrieves brands specified by brand name | 
 **productAttributes** | **[]string** | Defines product attributes | 
 **manufacturerId** | **string** | Defines product&#39;s manufacturer by manufacturer_id | 
 **status** | **string** | Defines product&#39;s status | 
 **type_** | **string** | Defines products&#39;s type | 
 **visible** | **string** | Filter items by visibility status | [default to &quot;everywhere&quot;]
 **findValue** | **string** | Entity search that is specified by some value | 
 **findWhere** | **string** | Product search that is specified by field | 
 **returnGlobal** | **bool** | Determines the type of products to be returned. If set to &#39;true&#39;, only global products will be returned; if set to &#39;false&#39;, only local products will be returned. | [default to false]
 **params** | **string** | Set this parameter in order to choose which entity fields you want to retrieve | [default to &quot;id,name,description,price,categories_ids&quot;]
 **responseFields** | **string** | Set this parameter in order to choose which entity fields you want to retrieve | 
 **exclude** | **string** | Set this parameter in order to choose which entity fields you want to ignore. Works only if parameter &#x60;params&#x60; equal force_all | 
 **sortBy** | **string** | Set field to sort by | [default to &quot;id&quot;]
 **sortDirection** | **string** | Set sorting direction | [default to &quot;asc&quot;]
 **reportRequestId** | **string** | Report request id | 
 **disableCache** | **bool** | Disable cache for current request | [default to false]
 **disableReportCache** | **bool** | Disable report cache for current request | [default to false]
 **useLatestApiVersion** | **bool** | Use the latest platform API version | [default to false]
 **productType** | **string** | A categorization for the product | 

### Return type

[**ModelResponseProductList**](ModelResponseProductList.md)

### Authorization

[StoreKeyAuth](../README.md#StoreKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductManufacturerAdd

> ProductManufacturerAdd200Response ProductManufacturerAdd(ctx).ProductId(productId).Manufacturer(manufacturer).StoreId(storeId).MetaTitle(metaTitle).MetaKeywords(metaKeywords).MetaDescription(metaDescription).SearchKeywords(searchKeywords).ImageUrl(imageUrl).SeoUrl(seoUrl).Execute()

product.manufacturer.add



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
	productId := "10" // string | Defines products specified by product id
	manufacturer := "Samsung" // string | Defines product’s manufacturer's name
	storeId := "1" // string | Store Id (optional)
	metaTitle := "category,test" // string | Defines unique meta title for each entity (optional)
	metaKeywords := "category,test" // string | Defines unique meta keywords for each entity (optional)
	metaDescription := "category,test" // string | Defines unique meta description of a entity (optional)
	searchKeywords := "key1,key2,key3" // string | Defines unique search keywords (optional)
	imageUrl := "https://docs.api2cart.com/img/logo.png" // string | Image Url (optional)
	seoUrl := "some seo url" // string | Defines unique URL for SEO (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductAPI.ProductManufacturerAdd(context.Background()).ProductId(productId).Manufacturer(manufacturer).StoreId(storeId).MetaTitle(metaTitle).MetaKeywords(metaKeywords).MetaDescription(metaDescription).SearchKeywords(searchKeywords).ImageUrl(imageUrl).SeoUrl(seoUrl).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductAPI.ProductManufacturerAdd``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProductManufacturerAdd`: ProductManufacturerAdd200Response
	fmt.Fprintf(os.Stdout, "Response from `ProductAPI.ProductManufacturerAdd`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProductManufacturerAddRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **productId** | **string** | Defines products specified by product id | 
 **manufacturer** | **string** | Defines product’s manufacturer&#39;s name | 
 **storeId** | **string** | Store Id | 
 **metaTitle** | **string** | Defines unique meta title for each entity | 
 **metaKeywords** | **string** | Defines unique meta keywords for each entity | 
 **metaDescription** | **string** | Defines unique meta description of a entity | 
 **searchKeywords** | **string** | Defines unique search keywords | 
 **imageUrl** | **string** | Image Url | 
 **seoUrl** | **string** | Defines unique URL for SEO | 

### Return type

[**ProductManufacturerAdd200Response**](ProductManufacturerAdd200Response.md)

### Authorization

[StoreKeyAuth](../README.md#StoreKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductOptionAdd

> ProductOptionAdd200Response ProductOptionAdd(ctx).ProductOptionAdd(productOptionAdd).Execute()

product.option.add



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
	productOptionAdd := *openapiclient.NewProductOptionAdd("Color", "option_type_select") // ProductOptionAdd | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductAPI.ProductOptionAdd(context.Background()).ProductOptionAdd(productOptionAdd).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductAPI.ProductOptionAdd``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProductOptionAdd`: ProductOptionAdd200Response
	fmt.Fprintf(os.Stdout, "Response from `ProductAPI.ProductOptionAdd`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProductOptionAddRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **productOptionAdd** | [**ProductOptionAdd**](ProductOptionAdd.md) |  | 

### Return type

[**ProductOptionAdd200Response**](ProductOptionAdd200Response.md)

### Authorization

[StoreKeyAuth](../README.md#StoreKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductOptionAssign

> ProductOptionAssign200Response ProductOptionAssign(ctx).ProductId(productId).OptionId(optionId).Required(required).SortOrder(sortOrder).OptionValues(optionValues).ClearCache(clearCache).Execute()

product.option.assign



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
	productId := "10" // string | Defines product id where the option should be assigned
	optionId := "5" // string | Defines option id which has to be assigned
	required := true // bool | Defines if the option is required (optional) (default to false)
	sortOrder := int32(2) // int32 | Sort number in the list (optional) (default to 0)
	optionValues := "green,black,yellow" // string | Defines option values that has to be assigned (optional)
	clearCache := false // bool | Is cache clear required (optional) (default to true)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductAPI.ProductOptionAssign(context.Background()).ProductId(productId).OptionId(optionId).Required(required).SortOrder(sortOrder).OptionValues(optionValues).ClearCache(clearCache).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductAPI.ProductOptionAssign``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProductOptionAssign`: ProductOptionAssign200Response
	fmt.Fprintf(os.Stdout, "Response from `ProductAPI.ProductOptionAssign`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProductOptionAssignRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **productId** | **string** | Defines product id where the option should be assigned | 
 **optionId** | **string** | Defines option id which has to be assigned | 
 **required** | **bool** | Defines if the option is required | [default to false]
 **sortOrder** | **int32** | Sort number in the list | [default to 0]
 **optionValues** | **string** | Defines option values that has to be assigned | 
 **clearCache** | **bool** | Is cache clear required | [default to true]

### Return type

[**ProductOptionAssign200Response**](ProductOptionAssign200Response.md)

### Authorization

[StoreKeyAuth](../README.md#StoreKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductOptionDelete

> AttributeDelete200Response ProductOptionDelete(ctx).OptionId(optionId).ProductId(productId).StoreId(storeId).Execute()

product.option.delete



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
	optionId := "5" // string | Defines option id that should be deleted
	productId := "10" // string | Defines product id where the option should be deleted
	storeId := "1" // string | Store Id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductAPI.ProductOptionDelete(context.Background()).OptionId(optionId).ProductId(productId).StoreId(storeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductAPI.ProductOptionDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProductOptionDelete`: AttributeDelete200Response
	fmt.Fprintf(os.Stdout, "Response from `ProductAPI.ProductOptionDelete`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProductOptionDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **optionId** | **string** | Defines option id that should be deleted | 
 **productId** | **string** | Defines product id where the option should be deleted | 
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


## ProductOptionList

> ModelResponseProductOptionList ProductOptionList(ctx).Start(start).Count(count).ProductId(productId).LangId(langId).StoreId(storeId).ResponseFields(responseFields).Params(params).Exclude(exclude).Execute()

product.option.list



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
	productId := "10" // string | Retrieves products' options specified by product id (optional)
	langId := "3" // string | Language id (optional)
	storeId := "1" // string | Store Id (optional)
	responseFields := "{return_code,return_message,pagination,result}" // string | Set this parameter in order to choose which entity fields you want to retrieve (optional)
	params := "id,name,sort_order" // string | Set this parameter in order to choose which entity fields you want to retrieve (optional) (default to "id,name,description")
	exclude := "id,name,sort_order" // string | Set this parameter in order to choose which entity fields you want to ignore. Works only if parameter `params` equal force_all (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductAPI.ProductOptionList(context.Background()).Start(start).Count(count).ProductId(productId).LangId(langId).StoreId(storeId).ResponseFields(responseFields).Params(params).Exclude(exclude).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductAPI.ProductOptionList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProductOptionList`: ModelResponseProductOptionList
	fmt.Fprintf(os.Stdout, "Response from `ProductAPI.ProductOptionList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProductOptionListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **start** | **int32** | This parameter sets the number from which you want to get entities | [default to 0]
 **count** | **int32** | This parameter sets the entity amount that has to be retrieved. Max allowed count&#x3D;250 | [default to 10]
 **productId** | **string** | Retrieves products&#39; options specified by product id | 
 **langId** | **string** | Language id | 
 **storeId** | **string** | Store Id | 
 **responseFields** | **string** | Set this parameter in order to choose which entity fields you want to retrieve | 
 **params** | **string** | Set this parameter in order to choose which entity fields you want to retrieve | [default to &quot;id,name,description&quot;]
 **exclude** | **string** | Set this parameter in order to choose which entity fields you want to ignore. Works only if parameter &#x60;params&#x60; equal force_all | 

### Return type

[**ModelResponseProductOptionList**](ModelResponseProductOptionList.md)

### Authorization

[StoreKeyAuth](../README.md#StoreKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductOptionValueAdd

> ProductOptionValueAdd200Response ProductOptionValueAdd(ctx).ProductId(productId).OptionId(optionId).OptionValue(optionValue).SortOrder(sortOrder).DisplayValue(displayValue).IsDefault(isDefault).ClearCache(clearCache).Execute()

product.option.value.add



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
	productId := "10" // string | Defines product id where the option value should be added
	optionId := "5" // string | Defines option id where the value has to be added
	optionValue := "green" // string | Defines option value that has to be added (optional)
	sortOrder := int32(2) // int32 | Sort number in the list (optional) (default to 0)
	displayValue := "value" // string | Defines the value that will be displayed for the option value (optional)
	isDefault := true // bool | Defines as a default (optional)
	clearCache := false // bool | Is cache clear required (optional) (default to true)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductAPI.ProductOptionValueAdd(context.Background()).ProductId(productId).OptionId(optionId).OptionValue(optionValue).SortOrder(sortOrder).DisplayValue(displayValue).IsDefault(isDefault).ClearCache(clearCache).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductAPI.ProductOptionValueAdd``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProductOptionValueAdd`: ProductOptionValueAdd200Response
	fmt.Fprintf(os.Stdout, "Response from `ProductAPI.ProductOptionValueAdd`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProductOptionValueAddRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **productId** | **string** | Defines product id where the option value should be added | 
 **optionId** | **string** | Defines option id where the value has to be added | 
 **optionValue** | **string** | Defines option value that has to be added | 
 **sortOrder** | **int32** | Sort number in the list | [default to 0]
 **displayValue** | **string** | Defines the value that will be displayed for the option value | 
 **isDefault** | **bool** | Defines as a default | 
 **clearCache** | **bool** | Is cache clear required | [default to true]

### Return type

[**ProductOptionValueAdd200Response**](ProductOptionValueAdd200Response.md)

### Authorization

[StoreKeyAuth](../README.md#StoreKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductOptionValueAssign

> ProductOptionValueAssign200Response ProductOptionValueAssign(ctx).ProductOptionId(productOptionId).OptionValueId(optionValueId).ClearCache(clearCache).Execute()

product.option.value.assign



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
	productOptionId := int32(5) // int32 | Defines product's option id where the value has to be assigned
	optionValueId := "45" // string | Defines value id that has to be assigned
	clearCache := false // bool | Is cache clear required (optional) (default to true)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductAPI.ProductOptionValueAssign(context.Background()).ProductOptionId(productOptionId).OptionValueId(optionValueId).ClearCache(clearCache).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductAPI.ProductOptionValueAssign``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProductOptionValueAssign`: ProductOptionValueAssign200Response
	fmt.Fprintf(os.Stdout, "Response from `ProductAPI.ProductOptionValueAssign`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProductOptionValueAssignRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **productOptionId** | **int32** | Defines product&#39;s option id where the value has to be assigned | 
 **optionValueId** | **string** | Defines value id that has to be assigned | 
 **clearCache** | **bool** | Is cache clear required | [default to true]

### Return type

[**ProductOptionValueAssign200Response**](ProductOptionValueAssign200Response.md)

### Authorization

[StoreKeyAuth](../README.md#StoreKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductOptionValueDelete

> AttributeDelete200Response ProductOptionValueDelete(ctx).OptionId(optionId).OptionValueId(optionValueId).ProductId(productId).StoreId(storeId).Execute()

product.option.value.delete



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
	optionId := "5" // string | Defines option id where the value should be deleted
	optionValueId := "45" // string | Defines option value id that should be deleted
	productId := "10" // string | Defines product id where the option value should be deleted
	storeId := "1" // string | Store Id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductAPI.ProductOptionValueDelete(context.Background()).OptionId(optionId).OptionValueId(optionValueId).ProductId(productId).StoreId(storeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductAPI.ProductOptionValueDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProductOptionValueDelete`: AttributeDelete200Response
	fmt.Fprintf(os.Stdout, "Response from `ProductAPI.ProductOptionValueDelete`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProductOptionValueDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **optionId** | **string** | Defines option id where the value should be deleted | 
 **optionValueId** | **string** | Defines option value id that should be deleted | 
 **productId** | **string** | Defines product id where the option value should be deleted | 
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


## ProductOptionValueUpdate

> AccountConfigUpdate200Response ProductOptionValueUpdate(ctx).ProductId(productId).OptionId(optionId).OptionValueId(optionValueId).OptionValue(optionValue).Price(price).Quantity(quantity).DisplayValue(displayValue).ClearCache(clearCache).Execute()

product.option.value.update



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
	productId := "10" // string | Defines product id where the option value should be updated
	optionId := "5" // string | Defines option id where the value has to be updated
	optionValueId := "45" // string | Defines value id that has to be assigned
	optionValue := "green" // string | Defines option value that has to be added (optional)
	price := float32(99.9) // float32 | Defines new product option price (optional)
	quantity := float32(6) // float32 | Defines new products' options quantity (optional)
	displayValue := "value" // string | Defines the value that will be displayed for the option value (optional)
	clearCache := false // bool | Is cache clear required (optional) (default to true)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductAPI.ProductOptionValueUpdate(context.Background()).ProductId(productId).OptionId(optionId).OptionValueId(optionValueId).OptionValue(optionValue).Price(price).Quantity(quantity).DisplayValue(displayValue).ClearCache(clearCache).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductAPI.ProductOptionValueUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProductOptionValueUpdate`: AccountConfigUpdate200Response
	fmt.Fprintf(os.Stdout, "Response from `ProductAPI.ProductOptionValueUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProductOptionValueUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **productId** | **string** | Defines product id where the option value should be updated | 
 **optionId** | **string** | Defines option id where the value has to be updated | 
 **optionValueId** | **string** | Defines value id that has to be assigned | 
 **optionValue** | **string** | Defines option value that has to be added | 
 **price** | **float32** | Defines new product option price | 
 **quantity** | **float32** | Defines new products&#39; options quantity | 
 **displayValue** | **string** | Defines the value that will be displayed for the option value | 
 **clearCache** | **bool** | Is cache clear required | [default to true]

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


## ProductPriceAdd

> CartValidate200Response ProductPriceAdd(ctx).ProductPriceAdd(productPriceAdd).Execute()

product.price.add



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
	productPriceAdd := *openapiclient.NewProductPriceAdd() // ProductPriceAdd | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductAPI.ProductPriceAdd(context.Background()).ProductPriceAdd(productPriceAdd).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductAPI.ProductPriceAdd``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProductPriceAdd`: CartValidate200Response
	fmt.Fprintf(os.Stdout, "Response from `ProductAPI.ProductPriceAdd`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProductPriceAddRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **productPriceAdd** | [**ProductPriceAdd**](ProductPriceAdd.md) |  | 

### Return type

[**CartValidate200Response**](CartValidate200Response.md)

### Authorization

[StoreKeyAuth](../README.md#StoreKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductPriceDelete

> AttributeDelete200Response ProductPriceDelete(ctx).ProductId(productId).GroupPrices(groupPrices).StoreId(storeId).Execute()

product.price.delete



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
	productId := "10" // string | Defines the product where the price has to be deleted
	groupPrices := "group_prices=5,8,9" // string | Defines product's group prices (optional)
	storeId := "1" // string | Store Id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductAPI.ProductPriceDelete(context.Background()).ProductId(productId).GroupPrices(groupPrices).StoreId(storeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductAPI.ProductPriceDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProductPriceDelete`: AttributeDelete200Response
	fmt.Fprintf(os.Stdout, "Response from `ProductAPI.ProductPriceDelete`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProductPriceDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **productId** | **string** | Defines the product where the price has to be deleted | 
 **groupPrices** | **string** | Defines product&#39;s group prices | 
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


## ProductPriceUpdate

> AccountConfigUpdate200Response ProductPriceUpdate(ctx).ProductPriceUpdate(productPriceUpdate).Execute()

product.price.update



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
	productPriceUpdate := *openapiclient.NewProductPriceUpdate() // ProductPriceUpdate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductAPI.ProductPriceUpdate(context.Background()).ProductPriceUpdate(productPriceUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductAPI.ProductPriceUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProductPriceUpdate`: AccountConfigUpdate200Response
	fmt.Fprintf(os.Stdout, "Response from `ProductAPI.ProductPriceUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProductPriceUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **productPriceUpdate** | [**ProductPriceUpdate**](ProductPriceUpdate.md) |  | 

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


## ProductReviewList

> ModelResponseProductReviewList ProductReviewList(ctx).ProductId(productId).Start(start).Count(count).PageCursor(pageCursor).Ids(ids).StoreId(storeId).LangId(langId).Status(status).CreatedFrom(createdFrom).CreatedTo(createdTo).CustomerId(customerId).SortBy(sortBy).SortDirection(sortDirection).ResponseFields(responseFields).Params(params).Exclude(exclude).Execute()

product.review.list



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
	productId := "10" // string | Product id
	start := int32(0) // int32 | This parameter sets the number from which you want to get entities (optional) (default to 0)
	count := int32(20) // int32 | This parameter sets the entity amount that has to be retrieved. Max allowed count=250 (optional) (default to 10)
	pageCursor := "pageCursor_example" // string | Used to retrieve entities via cursor-based pagination (it can't be used with any other filtering parameter) (optional)
	ids := "24,25" // string | Retrieves reviews specified by ids (optional)
	storeId := "1" // string | Store Id (optional)
	langId := "3" // string | Language id (optional)
	status := "disabled" // string | Defines status (optional)
	createdFrom := "2010-07-29 13:45:52" // string | Retrieve entities from their creation date (optional)
	createdTo := "2100-08-29 13:45:52" // string | Retrieve entities to their creation date (optional)
	customerId := "5" // string | Retrieves orders specified by customer id (optional)
	sortBy := "value_id" // string | Set field to sort by (optional) (default to "id")
	sortDirection := "asc" // string | Set sorting direction (optional) (default to "asc")
	responseFields := "{return_code,return_message,pagination,result}" // string | Set this parameter in order to choose which entity fields you want to retrieve (optional)
	params := "id,model,price,images" // string | Set this parameter in order to choose which entity fields you want to retrieve (optional) (default to "id,customer_id,email,message,status,product_id,nick_name,summary,rating,ratings,status,created_time")
	exclude := "false" // string | Set this parameter in order to choose which entity fields you want to ignore. Works only if parameter `params` equal force_all (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductAPI.ProductReviewList(context.Background()).ProductId(productId).Start(start).Count(count).PageCursor(pageCursor).Ids(ids).StoreId(storeId).LangId(langId).Status(status).CreatedFrom(createdFrom).CreatedTo(createdTo).CustomerId(customerId).SortBy(sortBy).SortDirection(sortDirection).ResponseFields(responseFields).Params(params).Exclude(exclude).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductAPI.ProductReviewList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProductReviewList`: ModelResponseProductReviewList
	fmt.Fprintf(os.Stdout, "Response from `ProductAPI.ProductReviewList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProductReviewListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **productId** | **string** | Product id | 
 **start** | **int32** | This parameter sets the number from which you want to get entities | [default to 0]
 **count** | **int32** | This parameter sets the entity amount that has to be retrieved. Max allowed count&#x3D;250 | [default to 10]
 **pageCursor** | **string** | Used to retrieve entities via cursor-based pagination (it can&#39;t be used with any other filtering parameter) | 
 **ids** | **string** | Retrieves reviews specified by ids | 
 **storeId** | **string** | Store Id | 
 **langId** | **string** | Language id | 
 **status** | **string** | Defines status | 
 **createdFrom** | **string** | Retrieve entities from their creation date | 
 **createdTo** | **string** | Retrieve entities to their creation date | 
 **customerId** | **string** | Retrieves orders specified by customer id | 
 **sortBy** | **string** | Set field to sort by | [default to &quot;id&quot;]
 **sortDirection** | **string** | Set sorting direction | [default to &quot;asc&quot;]
 **responseFields** | **string** | Set this parameter in order to choose which entity fields you want to retrieve | 
 **params** | **string** | Set this parameter in order to choose which entity fields you want to retrieve | [default to &quot;id,customer_id,email,message,status,product_id,nick_name,summary,rating,ratings,status,created_time&quot;]
 **exclude** | **string** | Set this parameter in order to choose which entity fields you want to ignore. Works only if parameter &#x60;params&#x60; equal force_all | 

### Return type

[**ModelResponseProductReviewList**](ModelResponseProductReviewList.md)

### Authorization

[StoreKeyAuth](../README.md#StoreKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductStoreAssign

> AccountConfigUpdate200Response ProductStoreAssign(ctx).ProductId(productId).StoreId(storeId).Execute()

product.store.assign



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
	productId := "10" // string | Defines id of the product which should be assigned to a store
	storeId := "1" // string | Defines id of the store product should be assigned to

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductAPI.ProductStoreAssign(context.Background()).ProductId(productId).StoreId(storeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductAPI.ProductStoreAssign``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProductStoreAssign`: AccountConfigUpdate200Response
	fmt.Fprintf(os.Stdout, "Response from `ProductAPI.ProductStoreAssign`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProductStoreAssignRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **productId** | **string** | Defines id of the product which should be assigned to a store | 
 **storeId** | **string** | Defines id of the store product should be assigned to | 

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


## ProductTaxAdd

> ProductTaxAdd200Response ProductTaxAdd(ctx).ProductTaxAdd(productTaxAdd).Execute()

product.tax.add



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
	productTaxAdd := *openapiclient.NewProductTaxAdd("ECO-Tax", []openapiclient.ProductTaxAddTaxRatesInner{*openapiclient.NewProductTaxAddTaxRatesInner()}) // ProductTaxAdd | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductAPI.ProductTaxAdd(context.Background()).ProductTaxAdd(productTaxAdd).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductAPI.ProductTaxAdd``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProductTaxAdd`: ProductTaxAdd200Response
	fmt.Fprintf(os.Stdout, "Response from `ProductAPI.ProductTaxAdd`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProductTaxAddRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **productTaxAdd** | [**ProductTaxAdd**](ProductTaxAdd.md) |  | 

### Return type

[**ProductTaxAdd200Response**](ProductTaxAdd200Response.md)

### Authorization

[StoreKeyAuth](../README.md#StoreKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductUpdate

> AccountConfigUpdate200Response ProductUpdate(ctx).ProductUpdate(productUpdate).Execute()

product.update



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
	productUpdate := *openapiclient.NewProductUpdate() // ProductUpdate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductAPI.ProductUpdate(context.Background()).ProductUpdate(productUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductAPI.ProductUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProductUpdate`: AccountConfigUpdate200Response
	fmt.Fprintf(os.Stdout, "Response from `ProductAPI.ProductUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProductUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **productUpdate** | [**ProductUpdate**](ProductUpdate.md) |  | 

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


## ProductUpdateBatch

> CategoryAddBatch200Response ProductUpdateBatch(ctx).ProductUpdateBatch(productUpdateBatch).Execute()

product.update.batch



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
	productUpdateBatch := *openapiclient.NewProductUpdateBatch([]openapiclient.ProductUpdateBatchPayloadInner{*openapiclient.NewProductUpdateBatchPayloadInner("Id_example")}) // ProductUpdateBatch | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductAPI.ProductUpdateBatch(context.Background()).ProductUpdateBatch(productUpdateBatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductAPI.ProductUpdateBatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProductUpdateBatch`: CategoryAddBatch200Response
	fmt.Fprintf(os.Stdout, "Response from `ProductAPI.ProductUpdateBatch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProductUpdateBatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **productUpdateBatch** | [**ProductUpdateBatch**](ProductUpdateBatch.md) |  | 

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


## ProductVariantAdd

> ProductVariantAdd200Response ProductVariantAdd(ctx).ProductVariantAdd(productVariantAdd).Execute()

product.variant.add



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
	productVariantAdd := *openapiclient.NewProductVariantAdd("bag_01") // ProductVariantAdd | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductAPI.ProductVariantAdd(context.Background()).ProductVariantAdd(productVariantAdd).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductAPI.ProductVariantAdd``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProductVariantAdd`: ProductVariantAdd200Response
	fmt.Fprintf(os.Stdout, "Response from `ProductAPI.ProductVariantAdd`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProductVariantAddRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **productVariantAdd** | [**ProductVariantAdd**](ProductVariantAdd.md) |  | 

### Return type

[**ProductVariantAdd200Response**](ProductVariantAdd200Response.md)

### Authorization

[StoreKeyAuth](../README.md#StoreKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductVariantAddBatch

> CategoryAddBatch200Response ProductVariantAddBatch(ctx).ProductVariantAddBatch(productVariantAddBatch).Execute()

product.variant.add.batch



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
	productVariantAddBatch := *openapiclient.NewProductVariantAddBatch([]openapiclient.ProductVariantAddBatchPayloadInner{*openapiclient.NewProductVariantAddBatchPayloadInner("ProductId_example", []openapiclient.ProductVariantAddBatchPayloadInnerCombinationInner{*openapiclient.NewProductVariantAddBatchPayloadInnerCombinationInner("OptionName_example", "OptionValueName_example")}, "Sku_example")}) // ProductVariantAddBatch | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductAPI.ProductVariantAddBatch(context.Background()).ProductVariantAddBatch(productVariantAddBatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductAPI.ProductVariantAddBatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProductVariantAddBatch`: CategoryAddBatch200Response
	fmt.Fprintf(os.Stdout, "Response from `ProductAPI.ProductVariantAddBatch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProductVariantAddBatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **productVariantAddBatch** | [**ProductVariantAddBatch**](ProductVariantAddBatch.md) |  | 

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


## ProductVariantDelete

> AttributeValueDelete200Response ProductVariantDelete(ctx).Id(id).ProductId(productId).StoreId(storeId).Execute()

product.variant.delete



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
	id := "10" // string | Defines variant removal, specified by variant id
	productId := "10" // string | Defines product's id where the variant has to be deleted
	storeId := "1" // string | Store Id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductAPI.ProductVariantDelete(context.Background()).Id(id).ProductId(productId).StoreId(storeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductAPI.ProductVariantDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProductVariantDelete`: AttributeValueDelete200Response
	fmt.Fprintf(os.Stdout, "Response from `ProductAPI.ProductVariantDelete`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProductVariantDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | Defines variant removal, specified by variant id | 
 **productId** | **string** | Defines product&#39;s id where the variant has to be deleted | 
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


## ProductVariantDeleteBatch

> CategoryAddBatch200Response ProductVariantDeleteBatch(ctx).ProductVariantDeleteBatch(productVariantDeleteBatch).Execute()

product.variant.delete.batch



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
	productVariantDeleteBatch := *openapiclient.NewProductVariantDeleteBatch([]openapiclient.ProductVariantDeleteBatchPayloadInner{*openapiclient.NewProductVariantDeleteBatchPayloadInner("ProductId_example", "Id_example")}) // ProductVariantDeleteBatch | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductAPI.ProductVariantDeleteBatch(context.Background()).ProductVariantDeleteBatch(productVariantDeleteBatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductAPI.ProductVariantDeleteBatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProductVariantDeleteBatch`: CategoryAddBatch200Response
	fmt.Fprintf(os.Stdout, "Response from `ProductAPI.ProductVariantDeleteBatch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProductVariantDeleteBatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **productVariantDeleteBatch** | [**ProductVariantDeleteBatch**](ProductVariantDeleteBatch.md) |  | 

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


## ProductVariantImageAdd

> ProductVariantImageAdd200Response ProductVariantImageAdd(ctx).ProductVariantImageAdd(productVariantImageAdd).Execute()

product.variant.image.add



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
	productVariantImageAdd := *openapiclient.NewProductVariantImageAdd("45", "abibas.png", "base") // ProductVariantImageAdd | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductAPI.ProductVariantImageAdd(context.Background()).ProductVariantImageAdd(productVariantImageAdd).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductAPI.ProductVariantImageAdd``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProductVariantImageAdd`: ProductVariantImageAdd200Response
	fmt.Fprintf(os.Stdout, "Response from `ProductAPI.ProductVariantImageAdd`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProductVariantImageAddRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **productVariantImageAdd** | [**ProductVariantImageAdd**](ProductVariantImageAdd.md) |  | 

### Return type

[**ProductVariantImageAdd200Response**](ProductVariantImageAdd200Response.md)

### Authorization

[StoreKeyAuth](../README.md#StoreKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductVariantImageDelete

> AttributeDelete200Response ProductVariantImageDelete(ctx).ProductId(productId).ProductVariantId(productVariantId).Id(id).StoreId(storeId).Execute()

product.variant.image.delete



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
	productId := "10" // string | Defines product id where the variant image should be deleted
	productVariantId := "45" // string | Defines product's variants specified by variant id
	id := "10" // string | Entity id
	storeId := "1" // string | Store Id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductAPI.ProductVariantImageDelete(context.Background()).ProductId(productId).ProductVariantId(productVariantId).Id(id).StoreId(storeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductAPI.ProductVariantImageDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProductVariantImageDelete`: AttributeDelete200Response
	fmt.Fprintf(os.Stdout, "Response from `ProductAPI.ProductVariantImageDelete`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProductVariantImageDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **productId** | **string** | Defines product id where the variant image should be deleted | 
 **productVariantId** | **string** | Defines product&#39;s variants specified by variant id | 
 **id** | **string** | Entity id | 
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


## ProductVariantPriceAdd

> CartValidate200Response ProductVariantPriceAdd(ctx).ProductVariantPriceAdd(productVariantPriceAdd).Execute()

product.variant.price.add



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
	productVariantPriceAdd := *openapiclient.NewProductVariantPriceAdd([]openapiclient.ProductAddGroupPricesInner{*openapiclient.NewProductAddGroupPricesInner()}) // ProductVariantPriceAdd | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductAPI.ProductVariantPriceAdd(context.Background()).ProductVariantPriceAdd(productVariantPriceAdd).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductAPI.ProductVariantPriceAdd``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProductVariantPriceAdd`: CartValidate200Response
	fmt.Fprintf(os.Stdout, "Response from `ProductAPI.ProductVariantPriceAdd`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProductVariantPriceAddRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **productVariantPriceAdd** | [**ProductVariantPriceAdd**](ProductVariantPriceAdd.md) |  | 

### Return type

[**CartValidate200Response**](CartValidate200Response.md)

### Authorization

[StoreKeyAuth](../README.md#StoreKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductVariantPriceDelete

> AttributeDelete200Response ProductVariantPriceDelete(ctx).Id(id).ProductId(productId).GroupPrices(groupPrices).StoreId(storeId).Execute()

product.variant.price.delete



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
	id := "10" // string | Defines the variant where the price has to be deleted
	productId := "10" // string | Product id
	groupPrices := "group_prices=6,8,9" // string | Defines variants's group prices
	storeId := "1" // string | Store Id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductAPI.ProductVariantPriceDelete(context.Background()).Id(id).ProductId(productId).GroupPrices(groupPrices).StoreId(storeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductAPI.ProductVariantPriceDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProductVariantPriceDelete`: AttributeDelete200Response
	fmt.Fprintf(os.Stdout, "Response from `ProductAPI.ProductVariantPriceDelete`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProductVariantPriceDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | Defines the variant where the price has to be deleted | 
 **productId** | **string** | Product id | 
 **groupPrices** | **string** | Defines variants&#39;s group prices | 
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


## ProductVariantPriceUpdate

> AccountConfigUpdate200Response ProductVariantPriceUpdate(ctx).ProductVariantPriceUpdate(productVariantPriceUpdate).Execute()

product.variant.price.update



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
	productVariantPriceUpdate := *openapiclient.NewProductVariantPriceUpdate([]openapiclient.ProductPriceUpdateGroupPricesInner{*openapiclient.NewProductPriceUpdateGroupPricesInner()}) // ProductVariantPriceUpdate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductAPI.ProductVariantPriceUpdate(context.Background()).ProductVariantPriceUpdate(productVariantPriceUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductAPI.ProductVariantPriceUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProductVariantPriceUpdate`: AccountConfigUpdate200Response
	fmt.Fprintf(os.Stdout, "Response from `ProductAPI.ProductVariantPriceUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProductVariantPriceUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **productVariantPriceUpdate** | [**ProductVariantPriceUpdate**](ProductVariantPriceUpdate.md) |  | 

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


## ProductVariantUpdate

> AccountConfigUpdate200Response ProductVariantUpdate(ctx).ProductVariantUpdate(productVariantUpdate).Execute()

product.variant.update



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
	productVariantUpdate := *openapiclient.NewProductVariantUpdate() // ProductVariantUpdate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductAPI.ProductVariantUpdate(context.Background()).ProductVariantUpdate(productVariantUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductAPI.ProductVariantUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProductVariantUpdate`: AccountConfigUpdate200Response
	fmt.Fprintf(os.Stdout, "Response from `ProductAPI.ProductVariantUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProductVariantUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **productVariantUpdate** | [**ProductVariantUpdate**](ProductVariantUpdate.md) |  | 

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


## ProductVariantUpdateBatch

> CategoryAddBatch200Response ProductVariantUpdateBatch(ctx).ProductVariantUpdateBatch(productVariantUpdateBatch).Execute()

product.variant.update.batch



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
	productVariantUpdateBatch := *openapiclient.NewProductVariantUpdateBatch([]openapiclient.ProductVariantUpdateBatchPayloadInner{*openapiclient.NewProductVariantUpdateBatchPayloadInner("Id_example", "ProductId_example")}) // ProductVariantUpdateBatch | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductAPI.ProductVariantUpdateBatch(context.Background()).ProductVariantUpdateBatch(productVariantUpdateBatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductAPI.ProductVariantUpdateBatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProductVariantUpdateBatch`: CategoryAddBatch200Response
	fmt.Fprintf(os.Stdout, "Response from `ProductAPI.ProductVariantUpdateBatch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProductVariantUpdateBatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **productVariantUpdateBatch** | [**ProductVariantUpdateBatch**](ProductVariantUpdateBatch.md) |  | 

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

