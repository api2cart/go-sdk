# \CartAPI

All URIs are relative to *https://api.api2cart.local.com/v1.1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CartCatalogPriceRulesCount**](CartAPI.md#CartCatalogPriceRulesCount) | **Get** /cart.catalog_price_rules.count.json | cart.catalog_price_rules.count
[**CartCatalogPriceRulesList**](CartAPI.md#CartCatalogPriceRulesList) | **Get** /cart.catalog_price_rules.list.json | cart.catalog_price_rules.list
[**CartCouponAdd**](CartAPI.md#CartCouponAdd) | **Post** /cart.coupon.add.json | cart.coupon.add
[**CartCouponConditionAdd**](CartAPI.md#CartCouponConditionAdd) | **Post** /cart.coupon.condition.add.json | cart.coupon.condition.add
[**CartCouponCount**](CartAPI.md#CartCouponCount) | **Get** /cart.coupon.count.json | cart.coupon.count
[**CartCouponDelete**](CartAPI.md#CartCouponDelete) | **Delete** /cart.coupon.delete.json | cart.coupon.delete
[**CartCouponList**](CartAPI.md#CartCouponList) | **Get** /cart.coupon.list.json | cart.coupon.list
[**CartDelete**](CartAPI.md#CartDelete) | **Delete** /cart.delete.json | cart.delete
[**CartGiftcardAdd**](CartAPI.md#CartGiftcardAdd) | **Post** /cart.giftcard.add.json | cart.giftcard.add
[**CartGiftcardCount**](CartAPI.md#CartGiftcardCount) | **Get** /cart.giftcard.count.json | cart.giftcard.count
[**CartGiftcardDelete**](CartAPI.md#CartGiftcardDelete) | **Delete** /cart.giftcard.delete.json | cart.giftcard.delete
[**CartGiftcardList**](CartAPI.md#CartGiftcardList) | **Get** /cart.giftcard.list.json | cart.giftcard.list
[**CartInfo**](CartAPI.md#CartInfo) | **Get** /cart.info.json | cart.info
[**CartMetaDataList**](CartAPI.md#CartMetaDataList) | **Get** /cart.meta_data.list.json | cart.meta_data.list
[**CartMetaDataSet**](CartAPI.md#CartMetaDataSet) | **Post** /cart.meta_data.set.json | cart.meta_data.set
[**CartMetaDataUnset**](CartAPI.md#CartMetaDataUnset) | **Delete** /cart.meta_data.unset.json | cart.meta_data.unset
[**CartMethods**](CartAPI.md#CartMethods) | **Get** /cart.methods.json | cart.methods
[**CartPluginList**](CartAPI.md#CartPluginList) | **Get** /cart.plugin.list.json | cart.plugin.list
[**CartScriptAdd**](CartAPI.md#CartScriptAdd) | **Post** /cart.script.add.json | cart.script.add
[**CartScriptDelete**](CartAPI.md#CartScriptDelete) | **Delete** /cart.script.delete.json | cart.script.delete
[**CartScriptList**](CartAPI.md#CartScriptList) | **Get** /cart.script.list.json | cart.script.list
[**CartShippingZonesList**](CartAPI.md#CartShippingZonesList) | **Get** /cart.shipping_zones.list.json | cart.shipping_zones.list
[**CartValidate**](CartAPI.md#CartValidate) | **Get** /cart.validate.json | cart.validate



## CartCatalogPriceRulesCount

> CartCatalogPriceRulesCount200Response CartCatalogPriceRulesCount(ctx).Execute()

cart.catalog_price_rules.count



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
	resp, r, err := apiClient.CartAPI.CartCatalogPriceRulesCount(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CartAPI.CartCatalogPriceRulesCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CartCatalogPriceRulesCount`: CartCatalogPriceRulesCount200Response
	fmt.Fprintf(os.Stdout, "Response from `CartAPI.CartCatalogPriceRulesCount`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCartCatalogPriceRulesCountRequest struct via the builder pattern


### Return type

[**CartCatalogPriceRulesCount200Response**](CartCatalogPriceRulesCount200Response.md)

### Authorization

[StoreKeyAuth](../README.md#StoreKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CartCatalogPriceRulesList

> ModelResponseCartCatalogPriceRulesList CartCatalogPriceRulesList(ctx).Start(start).Count(count).PageCursor(pageCursor).Ids(ids).ResponseFields(responseFields).Params(params).Exclude(exclude).Execute()

cart.catalog_price_rules.list



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
	ids := "24,25" // string | Retrieves  catalog_price_rules by ids (optional)
	responseFields := "{result{catalog_price_rules_count,catalog_price_rules{id,type,name,avail,usage_count,actions,conditions}}}" // string | Set this parameter in order to choose which entity fields you want to retrieve (optional)
	params := "id,model,price,images" // string | Set this parameter in order to choose which entity fields you want to retrieve (optional) (default to "id,name,description")
	exclude := "false" // string | Set this parameter in order to choose which entity fields you want to ignore. Works only if parameter `params` equal force_all (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CartAPI.CartCatalogPriceRulesList(context.Background()).Start(start).Count(count).PageCursor(pageCursor).Ids(ids).ResponseFields(responseFields).Params(params).Exclude(exclude).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CartAPI.CartCatalogPriceRulesList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CartCatalogPriceRulesList`: ModelResponseCartCatalogPriceRulesList
	fmt.Fprintf(os.Stdout, "Response from `CartAPI.CartCatalogPriceRulesList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCartCatalogPriceRulesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **start** | **int32** | This parameter sets the number from which you want to get entities | [default to 0]
 **count** | **int32** | This parameter sets the entity amount that has to be retrieved. Max allowed count&#x3D;250 | [default to 10]
 **pageCursor** | **string** | Used to retrieve entities via cursor-based pagination (it can&#39;t be used with any other filtering parameter) | 
 **ids** | **string** | Retrieves  catalog_price_rules by ids | 
 **responseFields** | **string** | Set this parameter in order to choose which entity fields you want to retrieve | 
 **params** | **string** | Set this parameter in order to choose which entity fields you want to retrieve | [default to &quot;id,name,description&quot;]
 **exclude** | **string** | Set this parameter in order to choose which entity fields you want to ignore. Works only if parameter &#x60;params&#x60; equal force_all | 

### Return type

[**ModelResponseCartCatalogPriceRulesList**](ModelResponseCartCatalogPriceRulesList.md)

### Authorization

[StoreKeyAuth](../README.md#StoreKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CartCouponAdd

> CartCouponAdd200Response CartCouponAdd(ctx).CartCouponAdd(cartCouponAdd).Execute()

cart.coupon.add



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
	cartCouponAdd := *openapiclient.NewCartCouponAdd("000_BIG_SALE_000", "percent", "order_total", "matching_items", float32(15.5)) // CartCouponAdd | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CartAPI.CartCouponAdd(context.Background()).CartCouponAdd(cartCouponAdd).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CartAPI.CartCouponAdd``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CartCouponAdd`: CartCouponAdd200Response
	fmt.Fprintf(os.Stdout, "Response from `CartAPI.CartCouponAdd`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCartCouponAddRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cartCouponAdd** | [**CartCouponAdd**](CartCouponAdd.md) |  | 

### Return type

[**CartCouponAdd200Response**](CartCouponAdd200Response.md)

### Authorization

[StoreKeyAuth](../README.md#StoreKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CartCouponConditionAdd

> BasketLiveShippingServiceDelete200Response CartCouponConditionAdd(ctx).CouponId(couponId).Entity(entity).Key(key).Operator(operator).Value(value).Target(target).IncludeTax(includeTax).IncludeShipping(includeShipping).StoreId(storeId).Execute()

cart.coupon.condition.add



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
	couponId := "45845" // string | Coupon Id
	entity := "order" // string | Defines condition entity type
	key := "subtotal" // string | Defines condition entity attribute key
	operator := "==" // string | Defines condition operator
	value := "2" // string | Defines condition value, can be comma separated according to the operator.
	target := "coupon_action" // string | Defines condition operator (optional) (default to "coupon_prerequisite")
	includeTax := true // bool | Indicates whether to apply a discount for taxes. (optional) (default to false)
	includeShipping := true // bool | Indicates whether to apply a discount for shipping. (optional) (default to false)
	storeId := "1" // string | Store Id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CartAPI.CartCouponConditionAdd(context.Background()).CouponId(couponId).Entity(entity).Key(key).Operator(operator).Value(value).Target(target).IncludeTax(includeTax).IncludeShipping(includeShipping).StoreId(storeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CartAPI.CartCouponConditionAdd``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CartCouponConditionAdd`: BasketLiveShippingServiceDelete200Response
	fmt.Fprintf(os.Stdout, "Response from `CartAPI.CartCouponConditionAdd`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCartCouponConditionAddRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **couponId** | **string** | Coupon Id | 
 **entity** | **string** | Defines condition entity type | 
 **key** | **string** | Defines condition entity attribute key | 
 **operator** | **string** | Defines condition operator | 
 **value** | **string** | Defines condition value, can be comma separated according to the operator. | 
 **target** | **string** | Defines condition operator | [default to &quot;coupon_prerequisite&quot;]
 **includeTax** | **bool** | Indicates whether to apply a discount for taxes. | [default to false]
 **includeShipping** | **bool** | Indicates whether to apply a discount for shipping. | [default to false]
 **storeId** | **string** | Store Id | 

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


## CartCouponCount

> CartCouponCount200Response CartCouponCount(ctx).StoreId(storeId).Avail(avail).DateStartFrom(dateStartFrom).DateStartTo(dateStartTo).DateEndFrom(dateEndFrom).DateEndTo(dateEndTo).Execute()

cart.coupon.count



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
	avail := false // bool | Defines category's visibility status (optional) (default to true)
	dateStartFrom := "2016-12-29 16:44:30" // string | Filter entity by date_start (greater or equal) (optional)
	dateStartTo := "2016-12-29 16:44:30" // string | Filter entity by date_start (less or equal) (optional)
	dateEndFrom := "2016-12-29 16:44:30" // string | Filter entity by date_end (greater or equal) (optional)
	dateEndTo := "2016-12-29 16:44:30" // string | Filter entity by date_end (less or equal) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CartAPI.CartCouponCount(context.Background()).StoreId(storeId).Avail(avail).DateStartFrom(dateStartFrom).DateStartTo(dateStartTo).DateEndFrom(dateEndFrom).DateEndTo(dateEndTo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CartAPI.CartCouponCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CartCouponCount`: CartCouponCount200Response
	fmt.Fprintf(os.Stdout, "Response from `CartAPI.CartCouponCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCartCouponCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **storeId** | **string** | Store Id | 
 **avail** | **bool** | Defines category&#39;s visibility status | [default to true]
 **dateStartFrom** | **string** | Filter entity by date_start (greater or equal) | 
 **dateStartTo** | **string** | Filter entity by date_start (less or equal) | 
 **dateEndFrom** | **string** | Filter entity by date_end (greater or equal) | 
 **dateEndTo** | **string** | Filter entity by date_end (less or equal) | 

### Return type

[**CartCouponCount200Response**](CartCouponCount200Response.md)

### Authorization

[StoreKeyAuth](../README.md#StoreKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CartCouponDelete

> AttributeDelete200Response CartCouponDelete(ctx).Id(id).StoreId(storeId).Execute()

cart.coupon.delete



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
	resp, r, err := apiClient.CartAPI.CartCouponDelete(context.Background()).Id(id).StoreId(storeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CartAPI.CartCouponDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CartCouponDelete`: AttributeDelete200Response
	fmt.Fprintf(os.Stdout, "Response from `CartAPI.CartCouponDelete`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCartCouponDeleteRequest struct via the builder pattern


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


## CartCouponList

> ModelResponseCartCouponList CartCouponList(ctx).Start(start).Count(count).PageCursor(pageCursor).CouponsIds(couponsIds).StoreId(storeId).LangId(langId).Avail(avail).DateStartFrom(dateStartFrom).DateStartTo(dateStartTo).DateEndFrom(dateEndFrom).DateEndTo(dateEndTo).ResponseFields(responseFields).Params(params).Exclude(exclude).Execute()

cart.coupon.list



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
	couponsIds := "1,2,3" // string | Filter coupons by ids (optional)
	storeId := "1" // string | Filter coupons by store id (optional)
	langId := "3" // string | Language id (optional)
	avail := false // bool | Filter coupons by avail status (optional)
	dateStartFrom := "2016-12-29 16:44:30" // string | Filter entity by date_start (greater or equal) (optional)
	dateStartTo := "2016-12-29 16:44:30" // string | Filter entity by date_start (less or equal) (optional)
	dateEndFrom := "2016-12-29 16:44:30" // string | Filter entity by date_end (greater or equal) (optional)
	dateEndTo := "2016-12-29 16:44:30" // string | Filter entity by date_end (less or equal) (optional)
	responseFields := "{pagination,result{coupon_count,coupon{id,code,name,conditions,actions{scope,amount,conditions{id,value,sub-conditions}},date_start,avail}}}" // string | Set this parameter in order to choose which entity fields you want to retrieve (optional)
	params := "id,code,type,amount" // string | Set this parameter in order to choose which entity fields you want to retrieve (optional) (default to "id,code,name,description")
	exclude := "usage_history,type" // string | Set this parameter in order to choose which entity fields you want to ignore. Works only if parameter `params` equal force_all (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CartAPI.CartCouponList(context.Background()).Start(start).Count(count).PageCursor(pageCursor).CouponsIds(couponsIds).StoreId(storeId).LangId(langId).Avail(avail).DateStartFrom(dateStartFrom).DateStartTo(dateStartTo).DateEndFrom(dateEndFrom).DateEndTo(dateEndTo).ResponseFields(responseFields).Params(params).Exclude(exclude).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CartAPI.CartCouponList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CartCouponList`: ModelResponseCartCouponList
	fmt.Fprintf(os.Stdout, "Response from `CartAPI.CartCouponList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCartCouponListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **start** | **int32** | This parameter sets the number from which you want to get entities | [default to 0]
 **count** | **int32** | This parameter sets the entity amount that has to be retrieved. Max allowed count&#x3D;250 | [default to 10]
 **pageCursor** | **string** | Used to retrieve entities via cursor-based pagination (it can&#39;t be used with any other filtering parameter) | 
 **couponsIds** | **string** | Filter coupons by ids | 
 **storeId** | **string** | Filter coupons by store id | 
 **langId** | **string** | Language id | 
 **avail** | **bool** | Filter coupons by avail status | 
 **dateStartFrom** | **string** | Filter entity by date_start (greater or equal) | 
 **dateStartTo** | **string** | Filter entity by date_start (less or equal) | 
 **dateEndFrom** | **string** | Filter entity by date_end (greater or equal) | 
 **dateEndTo** | **string** | Filter entity by date_end (less or equal) | 
 **responseFields** | **string** | Set this parameter in order to choose which entity fields you want to retrieve | 
 **params** | **string** | Set this parameter in order to choose which entity fields you want to retrieve | [default to &quot;id,code,name,description&quot;]
 **exclude** | **string** | Set this parameter in order to choose which entity fields you want to ignore. Works only if parameter &#x60;params&#x60; equal force_all | 

### Return type

[**ModelResponseCartCouponList**](ModelResponseCartCouponList.md)

### Authorization

[StoreKeyAuth](../README.md#StoreKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CartDelete

> CartDelete200Response CartDelete(ctx).DeleteBridge(deleteBridge).Execute()

cart.delete



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
	deleteBridge := true // bool | Identifies if there is a necessity to delete bridge (optional) (default to true)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CartAPI.CartDelete(context.Background()).DeleteBridge(deleteBridge).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CartAPI.CartDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CartDelete`: CartDelete200Response
	fmt.Fprintf(os.Stdout, "Response from `CartAPI.CartDelete`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCartDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteBridge** | **bool** | Identifies if there is a necessity to delete bridge | [default to true]

### Return type

[**CartDelete200Response**](CartDelete200Response.md)

### Authorization

[StoreKeyAuth](../README.md#StoreKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CartGiftcardAdd

> CartGiftcardAdd200Response CartGiftcardAdd(ctx).Amount(amount).Code(code).OwnerEmail(ownerEmail).RecipientEmail(recipientEmail).RecipientName(recipientName).OwnerName(ownerName).Execute()

cart.giftcard.add



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
	amount := float32(15.5) // float32 | Defines the gift card amount value.
	code := "GFT1 A4S5 AA11 RD61" // string | Gift card code (optional)
	ownerEmail := "jubari@hannsgroup.com" // string | Gift card owner email (optional)
	recipientEmail := "jubari@hannsgroup.com" // string | Gift card recipient email (optional)
	recipientName := "John Doe" // string | Gift card recipient name (optional)
	ownerName := "John Doe" // string | Gift card owner name (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CartAPI.CartGiftcardAdd(context.Background()).Amount(amount).Code(code).OwnerEmail(ownerEmail).RecipientEmail(recipientEmail).RecipientName(recipientName).OwnerName(ownerName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CartAPI.CartGiftcardAdd``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CartGiftcardAdd`: CartGiftcardAdd200Response
	fmt.Fprintf(os.Stdout, "Response from `CartAPI.CartGiftcardAdd`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCartGiftcardAddRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **amount** | **float32** | Defines the gift card amount value. | 
 **code** | **string** | Gift card code | 
 **ownerEmail** | **string** | Gift card owner email | 
 **recipientEmail** | **string** | Gift card recipient email | 
 **recipientName** | **string** | Gift card recipient name | 
 **ownerName** | **string** | Gift card owner name | 

### Return type

[**CartGiftcardAdd200Response**](CartGiftcardAdd200Response.md)

### Authorization

[StoreKeyAuth](../README.md#StoreKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CartGiftcardCount

> CartGiftcardCount200Response CartGiftcardCount(ctx).StoreId(storeId).Execute()

cart.giftcard.count



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
	resp, r, err := apiClient.CartAPI.CartGiftcardCount(context.Background()).StoreId(storeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CartAPI.CartGiftcardCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CartGiftcardCount`: CartGiftcardCount200Response
	fmt.Fprintf(os.Stdout, "Response from `CartAPI.CartGiftcardCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCartGiftcardCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **storeId** | **string** | Store Id | 

### Return type

[**CartGiftcardCount200Response**](CartGiftcardCount200Response.md)

### Authorization

[StoreKeyAuth](../README.md#StoreKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CartGiftcardDelete

> AttributeDelete200Response CartGiftcardDelete(ctx).Id(id).Execute()

cart.giftcard.delete



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CartAPI.CartGiftcardDelete(context.Background()).Id(id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CartAPI.CartGiftcardDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CartGiftcardDelete`: AttributeDelete200Response
	fmt.Fprintf(os.Stdout, "Response from `CartAPI.CartGiftcardDelete`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCartGiftcardDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | Entity id | 

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


## CartGiftcardList

> ModelResponseCartGiftCardList CartGiftcardList(ctx).Start(start).Count(count).PageCursor(pageCursor).StoreId(storeId).ResponseFields(responseFields).Params(params).Exclude(exclude).Execute()

cart.giftcard.list



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
	storeId := "1" // string | Store Id (optional)
	responseFields := "{pagination,result{gift_card{id,code,amount,status}}}" // string | Set this parameter in order to choose which entity fields you want to retrieve (optional)
	params := "id,model,price,images" // string | Set this parameter in order to choose which entity fields you want to retrieve (optional) (default to "id,code,name")
	exclude := "false" // string | Set this parameter in order to choose which entity fields you want to ignore. Works only if parameter `params` equal force_all (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CartAPI.CartGiftcardList(context.Background()).Start(start).Count(count).PageCursor(pageCursor).StoreId(storeId).ResponseFields(responseFields).Params(params).Exclude(exclude).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CartAPI.CartGiftcardList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CartGiftcardList`: ModelResponseCartGiftCardList
	fmt.Fprintf(os.Stdout, "Response from `CartAPI.CartGiftcardList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCartGiftcardListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **start** | **int32** | This parameter sets the number from which you want to get entities | [default to 0]
 **count** | **int32** | This parameter sets the entity amount that has to be retrieved. Max allowed count&#x3D;250 | [default to 10]
 **pageCursor** | **string** | Used to retrieve entities via cursor-based pagination (it can&#39;t be used with any other filtering parameter) | 
 **storeId** | **string** | Store Id | 
 **responseFields** | **string** | Set this parameter in order to choose which entity fields you want to retrieve | 
 **params** | **string** | Set this parameter in order to choose which entity fields you want to retrieve | [default to &quot;id,code,name&quot;]
 **exclude** | **string** | Set this parameter in order to choose which entity fields you want to ignore. Works only if parameter &#x60;params&#x60; equal force_all | 

### Return type

[**ModelResponseCartGiftCardList**](ModelResponseCartGiftCardList.md)

### Authorization

[StoreKeyAuth](../README.md#StoreKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CartInfo

> CartInfo200Response CartInfo(ctx).StoreId(storeId).ResponseFields(responseFields).Params(params).Exclude(exclude).Execute()

cart.info



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
	responseFields := "{result{name,url,stores_info{store_id,name,currency{id,iso3},store_owner_info}}}" // string | Set this parameter in order to choose which entity fields you want to retrieve (optional)
	params := "name,url" // string | Set this parameter in order to choose which entity fields you want to retrieve (optional) (default to "store_name,store_url,db_prefix")
	exclude := "name,url" // string | Set this parameter in order to choose which entity fields you want to ignore. Works only if parameter `params` equal force_all (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CartAPI.CartInfo(context.Background()).StoreId(storeId).ResponseFields(responseFields).Params(params).Exclude(exclude).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CartAPI.CartInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CartInfo`: CartInfo200Response
	fmt.Fprintf(os.Stdout, "Response from `CartAPI.CartInfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCartInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **storeId** | **string** | Store Id | 
 **responseFields** | **string** | Set this parameter in order to choose which entity fields you want to retrieve | 
 **params** | **string** | Set this parameter in order to choose which entity fields you want to retrieve | [default to &quot;store_name,store_url,db_prefix&quot;]
 **exclude** | **string** | Set this parameter in order to choose which entity fields you want to ignore. Works only if parameter &#x60;params&#x60; equal force_all | 

### Return type

[**CartInfo200Response**](CartInfo200Response.md)

### Authorization

[StoreKeyAuth](../README.md#StoreKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CartMetaDataList

> ModelResponseCartMetaDataList CartMetaDataList(ctx).EntityId(entityId).Count(count).PageCursor(pageCursor).Entity(entity).StoreId(storeId).LangId(langId).Key(key).ResponseFields(responseFields).Params(params).Exclude(exclude).Execute()

cart.meta_data.list



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
	entityId := "1" // string | Entity Id
	count := int32(20) // int32 | This parameter sets the entity amount that has to be retrieved. Max allowed count=250 (optional) (default to 10)
	pageCursor := "pageCursor_example" // string | Used to retrieve entities via cursor-based pagination (it can't be used with any other filtering parameter) (optional)
	entity := "order" // string | Entity (optional) (default to "product")
	storeId := "1" // string | Store Id (optional)
	langId := "3" // string | Language id (optional)
	key := "subtotal" // string | Key (optional)
	responseFields := "{result{items{key,value}}}" // string | Set this parameter in order to choose which entity fields you want to retrieve (optional)
	params := "id,model,price,images" // string | Set this parameter in order to choose which entity fields you want to retrieve (optional) (default to "key,value")
	exclude := "false" // string | Set this parameter in order to choose which entity fields you want to ignore. Works only if parameter `params` equal force_all (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CartAPI.CartMetaDataList(context.Background()).EntityId(entityId).Count(count).PageCursor(pageCursor).Entity(entity).StoreId(storeId).LangId(langId).Key(key).ResponseFields(responseFields).Params(params).Exclude(exclude).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CartAPI.CartMetaDataList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CartMetaDataList`: ModelResponseCartMetaDataList
	fmt.Fprintf(os.Stdout, "Response from `CartAPI.CartMetaDataList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCartMetaDataListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **entityId** | **string** | Entity Id | 
 **count** | **int32** | This parameter sets the entity amount that has to be retrieved. Max allowed count&#x3D;250 | [default to 10]
 **pageCursor** | **string** | Used to retrieve entities via cursor-based pagination (it can&#39;t be used with any other filtering parameter) | 
 **entity** | **string** | Entity | [default to &quot;product&quot;]
 **storeId** | **string** | Store Id | 
 **langId** | **string** | Language id | 
 **key** | **string** | Key | 
 **responseFields** | **string** | Set this parameter in order to choose which entity fields you want to retrieve | 
 **params** | **string** | Set this parameter in order to choose which entity fields you want to retrieve | [default to &quot;key,value&quot;]
 **exclude** | **string** | Set this parameter in order to choose which entity fields you want to ignore. Works only if parameter &#x60;params&#x60; equal force_all | 

### Return type

[**ModelResponseCartMetaDataList**](ModelResponseCartMetaDataList.md)

### Authorization

[StoreKeyAuth](../README.md#StoreKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CartMetaDataSet

> AttributeAdd200Response CartMetaDataSet(ctx).EntityId(entityId).Key(key).Value(value).Namespace(namespace).Entity(entity).StoreId(storeId).LangId(langId).Execute()

cart.meta_data.set



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
	entityId := "1" // string | Entity Id
	key := "subtotal" // string | Key
	value := "2" // string | Value
	namespace := "order" // string | Metafield namespace
	entity := "order" // string | Entity (optional) (default to "product")
	storeId := "1" // string | Store Id (optional)
	langId := "3" // string | Language id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CartAPI.CartMetaDataSet(context.Background()).EntityId(entityId).Key(key).Value(value).Namespace(namespace).Entity(entity).StoreId(storeId).LangId(langId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CartAPI.CartMetaDataSet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CartMetaDataSet`: AttributeAdd200Response
	fmt.Fprintf(os.Stdout, "Response from `CartAPI.CartMetaDataSet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCartMetaDataSetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **entityId** | **string** | Entity Id | 
 **key** | **string** | Key | 
 **value** | **string** | Value | 
 **namespace** | **string** | Metafield namespace | 
 **entity** | **string** | Entity | [default to &quot;product&quot;]
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


## CartMetaDataUnset

> BasketLiveShippingServiceDelete200Response CartMetaDataUnset(ctx).EntityId(entityId).Key(key).Id(id).Entity(entity).StoreId(storeId).Execute()

cart.meta_data.unset



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
	entityId := "1" // string | Entity Id
	key := "subtotal" // string | Key
	id := "10" // string | Entity id
	entity := "order" // string | Entity (optional) (default to "product")
	storeId := "1" // string | Store Id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CartAPI.CartMetaDataUnset(context.Background()).EntityId(entityId).Key(key).Id(id).Entity(entity).StoreId(storeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CartAPI.CartMetaDataUnset``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CartMetaDataUnset`: BasketLiveShippingServiceDelete200Response
	fmt.Fprintf(os.Stdout, "Response from `CartAPI.CartMetaDataUnset`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCartMetaDataUnsetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **entityId** | **string** | Entity Id | 
 **key** | **string** | Key | 
 **id** | **string** | Entity id | 
 **entity** | **string** | Entity | [default to &quot;product&quot;]
 **storeId** | **string** | Store Id | 

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


## CartMethods

> CartMethods200Response CartMethods(ctx).Execute()

cart.methods



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
	resp, r, err := apiClient.CartAPI.CartMethods(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CartAPI.CartMethods``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CartMethods`: CartMethods200Response
	fmt.Fprintf(os.Stdout, "Response from `CartAPI.CartMethods`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCartMethodsRequest struct via the builder pattern


### Return type

[**CartMethods200Response**](CartMethods200Response.md)

### Authorization

[StoreKeyAuth](../README.md#StoreKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CartPluginList

> CartPluginList200Response CartPluginList(ctx).Start(start).Count(count).StoreId(storeId).Execute()

cart.plugin.list



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
	resp, r, err := apiClient.CartAPI.CartPluginList(context.Background()).Start(start).Count(count).StoreId(storeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CartAPI.CartPluginList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CartPluginList`: CartPluginList200Response
	fmt.Fprintf(os.Stdout, "Response from `CartAPI.CartPluginList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCartPluginListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **start** | **int32** | This parameter sets the number from which you want to get entities | [default to 0]
 **count** | **int32** | This parameter sets the entity amount that has to be retrieved. Max allowed count&#x3D;250 | [default to 10]
 **storeId** | **string** | Store Id | 

### Return type

[**CartPluginList200Response**](CartPluginList200Response.md)

### Authorization

[StoreKeyAuth](../README.md#StoreKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CartScriptAdd

> CartScriptAdd200Response CartScriptAdd(ctx).Name(name).Description(description).Html(html).Src(src).LoadMethod(loadMethod).Scope(scope).Events(events).StoreId(storeId).Execute()

cart.script.add



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
	name := "jQuery Minimized" // string | The user-friendly script name (optional)
	description := "The Write Less, Do More, JavaScript Library" // string | The user-friendly description (optional)
	html := "&#x3C;script&#x3E;alert(&#x27;foo&#x27;)&#x3C;/script&#x3E;" // string | An html string containing exactly one `script` tag. (optional)
	src := "https://js-aplenty.com/foo.js" // string | The URL of the remote script (optional)
	loadMethod := "async" // string | The load method to use for the script (optional)
	scope := "all" // string | The page or pages on the online store where the script should be included (optional) (default to "storefront")
	events := "purchase_event" // string | Event for run scripts (optional)
	storeId := "1" // string | Store Id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CartAPI.CartScriptAdd(context.Background()).Name(name).Description(description).Html(html).Src(src).LoadMethod(loadMethod).Scope(scope).Events(events).StoreId(storeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CartAPI.CartScriptAdd``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CartScriptAdd`: CartScriptAdd200Response
	fmt.Fprintf(os.Stdout, "Response from `CartAPI.CartScriptAdd`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCartScriptAddRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | The user-friendly script name | 
 **description** | **string** | The user-friendly description | 
 **html** | **string** | An html string containing exactly one &#x60;script&#x60; tag. | 
 **src** | **string** | The URL of the remote script | 
 **loadMethod** | **string** | The load method to use for the script | 
 **scope** | **string** | The page or pages on the online store where the script should be included | [default to &quot;storefront&quot;]
 **events** | **string** | Event for run scripts | 
 **storeId** | **string** | Store Id | 

### Return type

[**CartScriptAdd200Response**](CartScriptAdd200Response.md)

### Authorization

[StoreKeyAuth](../README.md#StoreKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CartScriptDelete

> AttributeDelete200Response CartScriptDelete(ctx).Id(id).StoreId(storeId).Execute()

cart.script.delete



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
	resp, r, err := apiClient.CartAPI.CartScriptDelete(context.Background()).Id(id).StoreId(storeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CartAPI.CartScriptDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CartScriptDelete`: AttributeDelete200Response
	fmt.Fprintf(os.Stdout, "Response from `CartAPI.CartScriptDelete`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCartScriptDeleteRequest struct via the builder pattern


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


## CartScriptList

> ModelResponseCartScriptList CartScriptList(ctx).Start(start).Count(count).PageCursor(pageCursor).ScriptIds(scriptIds).StoreId(storeId).CreatedFrom(createdFrom).CreatedTo(createdTo).ModifiedFrom(modifiedFrom).ModifiedTo(modifiedTo).ResponseFields(responseFields).Params(params).Exclude(exclude).Execute()

cart.script.list



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
	scriptIds := "34023324,34024032" // string | Retrieves only scripts with specific ids (optional)
	storeId := "1" // string | Store Id (optional)
	createdFrom := "2010-07-29 13:45:52" // string | Retrieve entities from their creation date (optional)
	createdTo := "2100-08-29 13:45:52" // string | Retrieve entities to their creation date (optional)
	modifiedFrom := "2010-07-29 13:45:52" // string | Retrieve entities from their modification date (optional)
	modifiedTo := "2100-08-29 13:45:52" // string | Retrieve entities to their modification date (optional)
	responseFields := "{pagination,result{total_count,scripts{id,name,src,created_time{value}}}}" // string | Set this parameter in order to choose which entity fields you want to retrieve (optional)
	params := "id,model,price,images" // string | Set this parameter in order to choose which entity fields you want to retrieve (optional) (default to "id,name,description")
	exclude := "false" // string | Set this parameter in order to choose which entity fields you want to ignore. Works only if parameter `params` equal force_all (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CartAPI.CartScriptList(context.Background()).Start(start).Count(count).PageCursor(pageCursor).ScriptIds(scriptIds).StoreId(storeId).CreatedFrom(createdFrom).CreatedTo(createdTo).ModifiedFrom(modifiedFrom).ModifiedTo(modifiedTo).ResponseFields(responseFields).Params(params).Exclude(exclude).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CartAPI.CartScriptList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CartScriptList`: ModelResponseCartScriptList
	fmt.Fprintf(os.Stdout, "Response from `CartAPI.CartScriptList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCartScriptListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **start** | **int32** | This parameter sets the number from which you want to get entities | [default to 0]
 **count** | **int32** | This parameter sets the entity amount that has to be retrieved. Max allowed count&#x3D;250 | [default to 10]
 **pageCursor** | **string** | Used to retrieve entities via cursor-based pagination (it can&#39;t be used with any other filtering parameter) | 
 **scriptIds** | **string** | Retrieves only scripts with specific ids | 
 **storeId** | **string** | Store Id | 
 **createdFrom** | **string** | Retrieve entities from their creation date | 
 **createdTo** | **string** | Retrieve entities to their creation date | 
 **modifiedFrom** | **string** | Retrieve entities from their modification date | 
 **modifiedTo** | **string** | Retrieve entities to their modification date | 
 **responseFields** | **string** | Set this parameter in order to choose which entity fields you want to retrieve | 
 **params** | **string** | Set this parameter in order to choose which entity fields you want to retrieve | [default to &quot;id,name,description&quot;]
 **exclude** | **string** | Set this parameter in order to choose which entity fields you want to ignore. Works only if parameter &#x60;params&#x60; equal force_all | 

### Return type

[**ModelResponseCartScriptList**](ModelResponseCartScriptList.md)

### Authorization

[StoreKeyAuth](../README.md#StoreKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CartShippingZonesList

> ModelResponseCartShippingZonesList CartShippingZonesList(ctx).Start(start).Count(count).StoreId(storeId).ResponseFields(responseFields).Params(params).Exclude(exclude).Execute()

cart.shipping_zones.list



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
	responseFields := "{result{id,name,enabled,countries,shipping_methods{name,rates}}}" // string | Set this parameter in order to choose which entity fields you want to retrieve (optional)
	params := "id,model,price,images" // string | Set this parameter in order to choose which entity fields you want to retrieve (optional) (default to "id,name,enabled")
	exclude := "false" // string | Set this parameter in order to choose which entity fields you want to ignore. Works only if parameter `params` equal force_all (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CartAPI.CartShippingZonesList(context.Background()).Start(start).Count(count).StoreId(storeId).ResponseFields(responseFields).Params(params).Exclude(exclude).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CartAPI.CartShippingZonesList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CartShippingZonesList`: ModelResponseCartShippingZonesList
	fmt.Fprintf(os.Stdout, "Response from `CartAPI.CartShippingZonesList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCartShippingZonesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **start** | **int32** | This parameter sets the number from which you want to get entities | [default to 0]
 **count** | **int32** | This parameter sets the entity amount that has to be retrieved. Max allowed count&#x3D;250 | [default to 10]
 **storeId** | **string** | Store Id | 
 **responseFields** | **string** | Set this parameter in order to choose which entity fields you want to retrieve | 
 **params** | **string** | Set this parameter in order to choose which entity fields you want to retrieve | [default to &quot;id,name,enabled&quot;]
 **exclude** | **string** | Set this parameter in order to choose which entity fields you want to ignore. Works only if parameter &#x60;params&#x60; equal force_all | 

### Return type

[**ModelResponseCartShippingZonesList**](ModelResponseCartShippingZonesList.md)

### Authorization

[StoreKeyAuth](../README.md#StoreKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CartValidate

> CartValidate200Response CartValidate(ctx).ValidateVersion(validateVersion).Execute()

cart.validate



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
	validateVersion := true // bool | Specify if api2cart should validate cart version (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CartAPI.CartValidate(context.Background()).ValidateVersion(validateVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CartAPI.CartValidate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CartValidate`: CartValidate200Response
	fmt.Fprintf(os.Stdout, "Response from `CartAPI.CartValidate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCartValidateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **validateVersion** | **bool** | Specify if api2cart should validate cart version | [default to false]

### Return type

[**CartValidate200Response**](CartValidate200Response.md)

### Authorization

[StoreKeyAuth](../README.md#StoreKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

