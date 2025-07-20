# \AttributeAPI

All URIs are relative to *https://api.api2cart.local.com/v1.1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AttributeAdd**](AttributeAPI.md#AttributeAdd) | **Post** /attribute.add.json | attribute.add
[**AttributeAssignGroup**](AttributeAPI.md#AttributeAssignGroup) | **Post** /attribute.assign.group.json | attribute.assign.group
[**AttributeAssignSet**](AttributeAPI.md#AttributeAssignSet) | **Post** /attribute.assign.set.json | attribute.assign.set
[**AttributeAttributesetList**](AttributeAPI.md#AttributeAttributesetList) | **Get** /attribute.attributeset.list.json | attribute.attributeset.list
[**AttributeCount**](AttributeAPI.md#AttributeCount) | **Get** /attribute.count.json | attribute.count
[**AttributeDelete**](AttributeAPI.md#AttributeDelete) | **Delete** /attribute.delete.json | attribute.delete
[**AttributeGroupList**](AttributeAPI.md#AttributeGroupList) | **Get** /attribute.group.list.json | attribute.group.list
[**AttributeInfo**](AttributeAPI.md#AttributeInfo) | **Get** /attribute.info.json | attribute.info
[**AttributeList**](AttributeAPI.md#AttributeList) | **Get** /attribute.list.json | attribute.list
[**AttributeTypeList**](AttributeAPI.md#AttributeTypeList) | **Get** /attribute.type.list.json | attribute.type.list
[**AttributeUnassignGroup**](AttributeAPI.md#AttributeUnassignGroup) | **Post** /attribute.unassign.group.json | attribute.unassign.group
[**AttributeUnassignSet**](AttributeAPI.md#AttributeUnassignSet) | **Post** /attribute.unassign.set.json | attribute.unassign.set
[**AttributeUpdate**](AttributeAPI.md#AttributeUpdate) | **Put** /attribute.update.json | attribute.update
[**AttributeValueAdd**](AttributeAPI.md#AttributeValueAdd) | **Post** /attribute.value.add.json | attribute.value.add
[**AttributeValueDelete**](AttributeAPI.md#AttributeValueDelete) | **Delete** /attribute.value.delete.json | attribute.value.delete
[**AttributeValueUpdate**](AttributeAPI.md#AttributeValueUpdate) | **Put** /attribute.value.update.json | attribute.value.update



## AttributeAdd

> AttributeAdd200Response AttributeAdd(ctx).Type_(type_).Name(name).Code(code).StoreId(storeId).LangId(langId).Visible(visible).Required(required).Position(position).AttributeGroupId(attributeGroupId).IsGlobal(isGlobal).IsSearchable(isSearchable).IsFilterable(isFilterable).IsComparable(isComparable).IsHtmlAllowedOnFront(isHtmlAllowedOnFront).IsFilterableInSearch(isFilterableInSearch).IsConfigurable(isConfigurable).IsVisibleInAdvancedSearch(isVisibleInAdvancedSearch).IsUsedForPromoRules(isUsedForPromoRules).UsedInProductListing(usedInProductListing).UsedForSortBy(usedForSortBy).ApplyTo(applyTo).Execute()

attribute.add



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
	type_ := "text" // string | Defines attribute's type
	name := "Specification" // string | Defines attributes's name
	code := "code" // string | Entity code (optional)
	storeId := "1" // string | Store Id (optional)
	langId := "3" // string | Language id (optional)
	visible := true // bool | Set visibility status (optional) (default to false)
	required := true // bool | Defines if the option is required (optional) (default to false)
	position := int32(5) // int32 | Attribute`s position (optional) (default to 0)
	attributeGroupId := "202" // string | Filter by attribute_group_id (optional)
	isGlobal := "Global" // string | Attribute saving scope (optional) (default to "Store")
	isSearchable := false // bool | Use attribute in Quick Search (optional) (default to false)
	isFilterable := "No" // string | Use In Layered Navigation (optional) (default to "No")
	isComparable := true // bool | Comparable on Front-end (optional) (default to false)
	isHtmlAllowedOnFront := true // bool | Allow HTML Tags on Frontend (optional) (default to false)
	isFilterableInSearch := true // bool | Use In Search Results Layered Navigation (optional) (default to false)
	isConfigurable := true // bool | Use To Create Configurable Product (optional) (default to false)
	isVisibleInAdvancedSearch := true // bool | Use in Advanced Search (optional) (default to false)
	isUsedForPromoRules := true // bool | Use for Promo Rule Conditions (optional) (default to false)
	usedInProductListing := true // bool | Used in Product Listing (optional) (default to false)
	usedForSortBy := true // bool | Used for Sorting in Product Listing (optional) (default to false)
	applyTo := "Global" // string | Types of products which can have this attribute (optional) (default to "all_types")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AttributeAPI.AttributeAdd(context.Background()).Type_(type_).Name(name).Code(code).StoreId(storeId).LangId(langId).Visible(visible).Required(required).Position(position).AttributeGroupId(attributeGroupId).IsGlobal(isGlobal).IsSearchable(isSearchable).IsFilterable(isFilterable).IsComparable(isComparable).IsHtmlAllowedOnFront(isHtmlAllowedOnFront).IsFilterableInSearch(isFilterableInSearch).IsConfigurable(isConfigurable).IsVisibleInAdvancedSearch(isVisibleInAdvancedSearch).IsUsedForPromoRules(isUsedForPromoRules).UsedInProductListing(usedInProductListing).UsedForSortBy(usedForSortBy).ApplyTo(applyTo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AttributeAPI.AttributeAdd``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AttributeAdd`: AttributeAdd200Response
	fmt.Fprintf(os.Stdout, "Response from `AttributeAPI.AttributeAdd`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAttributeAddRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | **string** | Defines attribute&#39;s type | 
 **name** | **string** | Defines attributes&#39;s name | 
 **code** | **string** | Entity code | 
 **storeId** | **string** | Store Id | 
 **langId** | **string** | Language id | 
 **visible** | **bool** | Set visibility status | [default to false]
 **required** | **bool** | Defines if the option is required | [default to false]
 **position** | **int32** | Attribute&#x60;s position | [default to 0]
 **attributeGroupId** | **string** | Filter by attribute_group_id | 
 **isGlobal** | **string** | Attribute saving scope | [default to &quot;Store&quot;]
 **isSearchable** | **bool** | Use attribute in Quick Search | [default to false]
 **isFilterable** | **string** | Use In Layered Navigation | [default to &quot;No&quot;]
 **isComparable** | **bool** | Comparable on Front-end | [default to false]
 **isHtmlAllowedOnFront** | **bool** | Allow HTML Tags on Frontend | [default to false]
 **isFilterableInSearch** | **bool** | Use In Search Results Layered Navigation | [default to false]
 **isConfigurable** | **bool** | Use To Create Configurable Product | [default to false]
 **isVisibleInAdvancedSearch** | **bool** | Use in Advanced Search | [default to false]
 **isUsedForPromoRules** | **bool** | Use for Promo Rule Conditions | [default to false]
 **usedInProductListing** | **bool** | Used in Product Listing | [default to false]
 **usedForSortBy** | **bool** | Used for Sorting in Product Listing | [default to false]
 **applyTo** | **string** | Types of products which can have this attribute | [default to &quot;all_types&quot;]

### Return type

[**AttributeAdd200Response**](AttributeAdd200Response.md)

### Authorization

[StoreKeyAuth](../README.md#StoreKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AttributeAssignGroup

> AttributeAssignGroup200Response AttributeAssignGroup(ctx).Id(id).GroupId(groupId).AttributeSetId(attributeSetId).Execute()

attribute.assign.group



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
	groupId := "3" // string | Attribute group_id
	attributeSetId := "4" // string | Attribute set id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AttributeAPI.AttributeAssignGroup(context.Background()).Id(id).GroupId(groupId).AttributeSetId(attributeSetId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AttributeAPI.AttributeAssignGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AttributeAssignGroup`: AttributeAssignGroup200Response
	fmt.Fprintf(os.Stdout, "Response from `AttributeAPI.AttributeAssignGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAttributeAssignGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | Entity id | 
 **groupId** | **string** | Attribute group_id | 
 **attributeSetId** | **string** | Attribute set id | 

### Return type

[**AttributeAssignGroup200Response**](AttributeAssignGroup200Response.md)

### Authorization

[StoreKeyAuth](../README.md#StoreKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AttributeAssignSet

> AttributeAssignGroup200Response AttributeAssignSet(ctx).Id(id).AttributeSetId(attributeSetId).GroupId(groupId).Execute()

attribute.assign.set



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
	attributeSetId := "4" // string | Attribute set id
	groupId := "3" // string | Attribute group_id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AttributeAPI.AttributeAssignSet(context.Background()).Id(id).AttributeSetId(attributeSetId).GroupId(groupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AttributeAPI.AttributeAssignSet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AttributeAssignSet`: AttributeAssignGroup200Response
	fmt.Fprintf(os.Stdout, "Response from `AttributeAPI.AttributeAssignSet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAttributeAssignSetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | Entity id | 
 **attributeSetId** | **string** | Attribute set id | 
 **groupId** | **string** | Attribute group_id | 

### Return type

[**AttributeAssignGroup200Response**](AttributeAssignGroup200Response.md)

### Authorization

[StoreKeyAuth](../README.md#StoreKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AttributeAttributesetList

> ModelResponseAttributeAttributesetList AttributeAttributesetList(ctx).Start(start).Count(count).ResponseFields(responseFields).Params(params).Exclude(exclude).Execute()

attribute.attributeset.list



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
	responseFields := "{result}" // string | Set this parameter in order to choose which entity fields you want to retrieve (optional)
	params := "id,name" // string | Set this parameter in order to choose which entity fields you want to retrieve (optional) (default to "id,name")
	exclude := "id,name" // string | Set this parameter in order to choose which entity fields you want to ignore. Works only if parameter `params` equal force_all (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AttributeAPI.AttributeAttributesetList(context.Background()).Start(start).Count(count).ResponseFields(responseFields).Params(params).Exclude(exclude).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AttributeAPI.AttributeAttributesetList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AttributeAttributesetList`: ModelResponseAttributeAttributesetList
	fmt.Fprintf(os.Stdout, "Response from `AttributeAPI.AttributeAttributesetList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAttributeAttributesetListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **start** | **int32** | This parameter sets the number from which you want to get entities | [default to 0]
 **count** | **int32** | This parameter sets the entity amount that has to be retrieved. Max allowed count&#x3D;250 | [default to 10]
 **responseFields** | **string** | Set this parameter in order to choose which entity fields you want to retrieve | 
 **params** | **string** | Set this parameter in order to choose which entity fields you want to retrieve | [default to &quot;id,name&quot;]
 **exclude** | **string** | Set this parameter in order to choose which entity fields you want to ignore. Works only if parameter &#x60;params&#x60; equal force_all | 

### Return type

[**ModelResponseAttributeAttributesetList**](ModelResponseAttributeAttributesetList.md)

### Authorization

[StoreKeyAuth](../README.md#StoreKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AttributeCount

> AttributeCount200Response AttributeCount(ctx).Type_(type_).AttributeSetId(attributeSetId).StoreId(storeId).LangId(langId).Visible(visible).Required(required).System(system).Execute()

attribute.count



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
	type_ := "text" // string | Defines attribute's type (optional)
	attributeSetId := "4" // string | Filter items by attribute set id (optional)
	storeId := "1" // string | Store Id (optional)
	langId := "3" // string | Language id (optional)
	visible := true // bool | Filter items by visibility status (optional)
	required := true // bool | Defines if the option is required (optional)
	system := false // bool | True if attribute is system (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AttributeAPI.AttributeCount(context.Background()).Type_(type_).AttributeSetId(attributeSetId).StoreId(storeId).LangId(langId).Visible(visible).Required(required).System(system).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AttributeAPI.AttributeCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AttributeCount`: AttributeCount200Response
	fmt.Fprintf(os.Stdout, "Response from `AttributeAPI.AttributeCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAttributeCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | **string** | Defines attribute&#39;s type | 
 **attributeSetId** | **string** | Filter items by attribute set id | 
 **storeId** | **string** | Store Id | 
 **langId** | **string** | Language id | 
 **visible** | **bool** | Filter items by visibility status | 
 **required** | **bool** | Defines if the option is required | 
 **system** | **bool** | True if attribute is system | 

### Return type

[**AttributeCount200Response**](AttributeCount200Response.md)

### Authorization

[StoreKeyAuth](../README.md#StoreKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AttributeDelete

> AttributeDelete200Response AttributeDelete(ctx).Id(id).StoreId(storeId).Execute()

attribute.delete



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AttributeAPI.AttributeDelete(context.Background()).Id(id).StoreId(storeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AttributeAPI.AttributeDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AttributeDelete`: AttributeDelete200Response
	fmt.Fprintf(os.Stdout, "Response from `AttributeAPI.AttributeDelete`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAttributeDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
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


## AttributeGroupList

> ModelResponseAttributeGroupList AttributeGroupList(ctx).Start(start).Count(count).AttributeSetId(attributeSetId).LangId(langId).ResponseFields(responseFields).Params(params).Exclude(exclude).Execute()

attribute.group.list



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
	attributeSetId := "4" // string | Attribute set id (optional)
	langId := "3" // string | Language id (optional)
	responseFields := "{result}" // string | Set this parameter in order to choose which entity fields you want to retrieve (optional)
	params := "id,name" // string | Set this parameter in order to choose which entity fields you want to retrieve (optional) (default to "id,name")
	exclude := "id,name" // string | Set this parameter in order to choose which entity fields you want to ignore. Works only if parameter `params` equal force_all (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AttributeAPI.AttributeGroupList(context.Background()).Start(start).Count(count).AttributeSetId(attributeSetId).LangId(langId).ResponseFields(responseFields).Params(params).Exclude(exclude).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AttributeAPI.AttributeGroupList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AttributeGroupList`: ModelResponseAttributeGroupList
	fmt.Fprintf(os.Stdout, "Response from `AttributeAPI.AttributeGroupList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAttributeGroupListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **start** | **int32** | This parameter sets the number from which you want to get entities | [default to 0]
 **count** | **int32** | This parameter sets the entity amount that has to be retrieved. Max allowed count&#x3D;250 | [default to 10]
 **attributeSetId** | **string** | Attribute set id | 
 **langId** | **string** | Language id | 
 **responseFields** | **string** | Set this parameter in order to choose which entity fields you want to retrieve | 
 **params** | **string** | Set this parameter in order to choose which entity fields you want to retrieve | [default to &quot;id,name&quot;]
 **exclude** | **string** | Set this parameter in order to choose which entity fields you want to ignore. Works only if parameter &#x60;params&#x60; equal force_all | 

### Return type

[**ModelResponseAttributeGroupList**](ModelResponseAttributeGroupList.md)

### Authorization

[StoreKeyAuth](../README.md#StoreKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AttributeInfo

> AttributeInfo200Response AttributeInfo(ctx).Id(id).AttributeSetId(attributeSetId).StoreId(storeId).LangId(langId).ResponseFields(responseFields).Params(params).Exclude(exclude).Execute()

attribute.info



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
	attributeSetId := "4" // string | Attribute set id (optional)
	storeId := "1" // string | Store Id (optional)
	langId := "3" // string | Language id (optional)
	responseFields := "{result}" // string | Set this parameter in order to choose which entity fields you want to retrieve (optional)
	params := "force_all" // string | Set this parameter in order to choose which entity fields you want to retrieve (optional) (default to "force_all")
	exclude := "name" // string | Set this parameter in order to choose which entity fields you want to ignore. Works only if parameter `params` equal force_all (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AttributeAPI.AttributeInfo(context.Background()).Id(id).AttributeSetId(attributeSetId).StoreId(storeId).LangId(langId).ResponseFields(responseFields).Params(params).Exclude(exclude).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AttributeAPI.AttributeInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AttributeInfo`: AttributeInfo200Response
	fmt.Fprintf(os.Stdout, "Response from `AttributeAPI.AttributeInfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAttributeInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | Entity id | 
 **attributeSetId** | **string** | Attribute set id | 
 **storeId** | **string** | Store Id | 
 **langId** | **string** | Language id | 
 **responseFields** | **string** | Set this parameter in order to choose which entity fields you want to retrieve | 
 **params** | **string** | Set this parameter in order to choose which entity fields you want to retrieve | [default to &quot;force_all&quot;]
 **exclude** | **string** | Set this parameter in order to choose which entity fields you want to ignore. Works only if parameter &#x60;params&#x60; equal force_all | 

### Return type

[**AttributeInfo200Response**](AttributeInfo200Response.md)

### Authorization

[StoreKeyAuth](../README.md#StoreKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AttributeList

> ModelResponseAttributeList AttributeList(ctx).Start(start).Count(count).AttributeIds(attributeIds).AttributeSetId(attributeSetId).StoreId(storeId).LangId(langId).Type_(type_).Visible(visible).Required(required).System(system).ResponseFields(responseFields).Params(params).Exclude(exclude).Execute()

attribute.list



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
	attributeIds := "1,2,3" // string | Filter attributes by ids (optional)
	attributeSetId := "4" // string | Filter items by attribute set id (optional)
	storeId := "1" // string | Store Id (optional)
	langId := "3" // string | Retrieves attributes on specified language id (optional)
	type_ := "text" // string | Defines attribute's type (optional)
	visible := true // bool | Filter items by visibility status (optional)
	required := true // bool | Defines if the option is required (optional)
	system := false // bool | True if attribute is system (optional)
	responseFields := "{return_code,return_message,pagination,result}" // string | Set this parameter in order to choose which entity fields you want to retrieve (optional)
	params := "id,name" // string | Set this parameter in order to choose which entity fields you want to retrieve (optional) (default to "id,name,code,type")
	exclude := "id,name" // string | Set this parameter in order to choose which entity fields you want to ignore. Works only if parameter `params` equal force_all (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AttributeAPI.AttributeList(context.Background()).Start(start).Count(count).AttributeIds(attributeIds).AttributeSetId(attributeSetId).StoreId(storeId).LangId(langId).Type_(type_).Visible(visible).Required(required).System(system).ResponseFields(responseFields).Params(params).Exclude(exclude).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AttributeAPI.AttributeList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AttributeList`: ModelResponseAttributeList
	fmt.Fprintf(os.Stdout, "Response from `AttributeAPI.AttributeList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAttributeListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **start** | **int32** | This parameter sets the number from which you want to get entities | [default to 0]
 **count** | **int32** | This parameter sets the entity amount that has to be retrieved. Max allowed count&#x3D;250 | [default to 10]
 **attributeIds** | **string** | Filter attributes by ids | 
 **attributeSetId** | **string** | Filter items by attribute set id | 
 **storeId** | **string** | Store Id | 
 **langId** | **string** | Retrieves attributes on specified language id | 
 **type_** | **string** | Defines attribute&#39;s type | 
 **visible** | **bool** | Filter items by visibility status | 
 **required** | **bool** | Defines if the option is required | 
 **system** | **bool** | True if attribute is system | 
 **responseFields** | **string** | Set this parameter in order to choose which entity fields you want to retrieve | 
 **params** | **string** | Set this parameter in order to choose which entity fields you want to retrieve | [default to &quot;id,name,code,type&quot;]
 **exclude** | **string** | Set this parameter in order to choose which entity fields you want to ignore. Works only if parameter &#x60;params&#x60; equal force_all | 

### Return type

[**ModelResponseAttributeList**](ModelResponseAttributeList.md)

### Authorization

[StoreKeyAuth](../README.md#StoreKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AttributeTypeList

> AttributeTypeList200Response AttributeTypeList(ctx).Execute()

attribute.type.list



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
	resp, r, err := apiClient.AttributeAPI.AttributeTypeList(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AttributeAPI.AttributeTypeList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AttributeTypeList`: AttributeTypeList200Response
	fmt.Fprintf(os.Stdout, "Response from `AttributeAPI.AttributeTypeList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAttributeTypeListRequest struct via the builder pattern


### Return type

[**AttributeTypeList200Response**](AttributeTypeList200Response.md)

### Authorization

[StoreKeyAuth](../README.md#StoreKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AttributeUnassignGroup

> AttributeUnassignGroup200Response AttributeUnassignGroup(ctx).Id(id).GroupId(groupId).Execute()

attribute.unassign.group



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
	groupId := "3" // string | Customer group_id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AttributeAPI.AttributeUnassignGroup(context.Background()).Id(id).GroupId(groupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AttributeAPI.AttributeUnassignGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AttributeUnassignGroup`: AttributeUnassignGroup200Response
	fmt.Fprintf(os.Stdout, "Response from `AttributeAPI.AttributeUnassignGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAttributeUnassignGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | Entity id | 
 **groupId** | **string** | Customer group_id | 

### Return type

[**AttributeUnassignGroup200Response**](AttributeUnassignGroup200Response.md)

### Authorization

[StoreKeyAuth](../README.md#StoreKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AttributeUnassignSet

> AttributeUnassignGroup200Response AttributeUnassignSet(ctx).Id(id).AttributeSetId(attributeSetId).Execute()

attribute.unassign.set



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
	attributeSetId := "4" // string | Attribute set id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AttributeAPI.AttributeUnassignSet(context.Background()).Id(id).AttributeSetId(attributeSetId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AttributeAPI.AttributeUnassignSet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AttributeUnassignSet`: AttributeUnassignGroup200Response
	fmt.Fprintf(os.Stdout, "Response from `AttributeAPI.AttributeUnassignSet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAttributeUnassignSetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | Entity id | 
 **attributeSetId** | **string** | Attribute set id | 

### Return type

[**AttributeUnassignGroup200Response**](AttributeUnassignGroup200Response.md)

### Authorization

[StoreKeyAuth](../README.md#StoreKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AttributeUpdate

> AttributeUpdate200Response AttributeUpdate(ctx).Id(id).Name(name).StoreId(storeId).LangId(langId).Execute()

attribute.update



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
	name := "Test name" // string | Defines new attributes's name
	storeId := "1" // string | Store Id (optional)
	langId := "3" // string | Language id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AttributeAPI.AttributeUpdate(context.Background()).Id(id).Name(name).StoreId(storeId).LangId(langId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AttributeAPI.AttributeUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AttributeUpdate`: AttributeUpdate200Response
	fmt.Fprintf(os.Stdout, "Response from `AttributeAPI.AttributeUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAttributeUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | Entity id | 
 **name** | **string** | Defines new attributes&#39;s name | 
 **storeId** | **string** | Store Id | 
 **langId** | **string** | Language id | 

### Return type

[**AttributeUpdate200Response**](AttributeUpdate200Response.md)

### Authorization

[StoreKeyAuth](../README.md#StoreKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AttributeValueAdd

> AttributeAdd200Response AttributeValueAdd(ctx).AttributeId(attributeId).Name(name).Code(code).Description(description).StoreId(storeId).LangId(langId).Execute()

attribute.value.add



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
	attributeId := "156" // string | Attribute Id
	name := "Test name" // string | Defines attribute value's name
	code := "code" // string | Entity code (optional)
	description := "Test value" // string | Defines attribute value's description (optional)
	storeId := "1" // string | Store Id (optional)
	langId := "3" // string | Language id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AttributeAPI.AttributeValueAdd(context.Background()).AttributeId(attributeId).Name(name).Code(code).Description(description).StoreId(storeId).LangId(langId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AttributeAPI.AttributeValueAdd``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AttributeValueAdd`: AttributeAdd200Response
	fmt.Fprintf(os.Stdout, "Response from `AttributeAPI.AttributeValueAdd`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAttributeValueAddRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **attributeId** | **string** | Attribute Id | 
 **name** | **string** | Defines attribute value&#39;s name | 
 **code** | **string** | Entity code | 
 **description** | **string** | Defines attribute value&#39;s description | 
 **storeId** | **string** | Store Id | 
 **langId** | **string** | Language id | 

### Return type

[**AttributeAdd200Response**](AttributeAdd200Response.md)

### Authorization

[StoreKeyAuth](../README.md#StoreKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AttributeValueDelete

> AttributeValueDelete200Response AttributeValueDelete(ctx).Id(id).AttributeId(attributeId).StoreId(storeId).Execute()

attribute.value.delete



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
	attributeId := "156" // string | Attribute Id
	storeId := "1" // string | Store Id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AttributeAPI.AttributeValueDelete(context.Background()).Id(id).AttributeId(attributeId).StoreId(storeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AttributeAPI.AttributeValueDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AttributeValueDelete`: AttributeValueDelete200Response
	fmt.Fprintf(os.Stdout, "Response from `AttributeAPI.AttributeValueDelete`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAttributeValueDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | Entity id | 
 **attributeId** | **string** | Attribute Id | 
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


## AttributeValueUpdate

> AttributeUpdate200Response AttributeValueUpdate(ctx).Id(id).AttributeId(attributeId).Name(name).Description(description).Code(code).StoreId(storeId).LangId(langId).Execute()

attribute.value.update



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
	id := "10" // string | Defines attribute value's id
	attributeId := "156" // string | Attribute Id
	name := "Test name" // string | Defines attribute value's name (optional)
	description := "Test value" // string | Defines new attribute value's description (optional)
	code := "code" // string | Entity code (optional)
	storeId := "1" // string | Store Id (optional)
	langId := "3" // string | Language id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AttributeAPI.AttributeValueUpdate(context.Background()).Id(id).AttributeId(attributeId).Name(name).Description(description).Code(code).StoreId(storeId).LangId(langId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AttributeAPI.AttributeValueUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AttributeValueUpdate`: AttributeUpdate200Response
	fmt.Fprintf(os.Stdout, "Response from `AttributeAPI.AttributeValueUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAttributeValueUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | Defines attribute value&#39;s id | 
 **attributeId** | **string** | Attribute Id | 
 **name** | **string** | Defines attribute value&#39;s name | 
 **description** | **string** | Defines new attribute value&#39;s description | 
 **code** | **string** | Entity code | 
 **storeId** | **string** | Store Id | 
 **langId** | **string** | Language id | 

### Return type

[**AttributeUpdate200Response**](AttributeUpdate200Response.md)

### Authorization

[StoreKeyAuth](../README.md#StoreKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

