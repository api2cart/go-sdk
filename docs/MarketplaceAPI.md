# \MarketplaceAPI

All URIs are relative to *https://api.api2cart.local.com/v1.1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MarketplaceProductFind**](MarketplaceAPI.md#MarketplaceProductFind) | **Get** /marketplace.product.find.json | marketplace.product.find



## MarketplaceProductFind

> ModelResponseMarketplaceProductFind MarketplaceProductFind(ctx).Count(count).PageCursor(pageCursor).Keyword(keyword).CategoriesIds(categoriesIds).StoreId(storeId).Asin(asin).Ean(ean).Gtin(gtin).Upc(upc).Mpn(mpn).Isbn(isbn).ResponseFields(responseFields).Params(params).Exclude(exclude).Execute()

marketplace.product.find



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
	count := int32(20) // int32 | This parameter sets the entity amount that has to be retrieved. Max allowed count=250 (optional) (default to 10)
	pageCursor := "pageCursor_example" // string | Used to retrieve entities via cursor-based pagination (it can't be used with any other filtering parameter) (optional)
	keyword := "T-shirt" // string | Defines search keyword (optional)
	categoriesIds := "23,56" // string | Defines product add that is specified by comma-separated categories id (optional)
	storeId := "1" // string | Store Id (optional)
	asin := "97703178470" // string | Amazon Standard Identification Number. (optional)
	ean := "5901234123457" // string | European Article Number. An EAN is a unique 8 or 13-digit identifier that many industries (such as book publishers) use to identify products. (optional)
	gtin := "12345678912345" // string | Global Trade Item Number. An GTIN is an identifier for trade items. (optional)
	upc := "9770317847001" // string | Universal Product Code. A UPC (UPC-A) is a commonly used identifer for many different products. (optional)
	mpn := "9770317847001" // string | Manufacturer Part Number. A MPN is an identifier of a particular part design or material used. (optional)
	isbn := "9783161484100" // string | International Standard Book Number. An ISBN is a unique identifier for books. (optional)
	responseFields := "{result}" // string | Set this parameter in order to choose which entity fields you want to retrieve (optional)
	params := "id,model,price,images" // string | Set this parameter in order to choose which entity fields you want to retrieve (optional) (default to "force_all")
	exclude := "false" // string | Set this parameter in order to choose which entity fields you want to ignore. Works only if parameter `params` equal force_all (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketplaceAPI.MarketplaceProductFind(context.Background()).Count(count).PageCursor(pageCursor).Keyword(keyword).CategoriesIds(categoriesIds).StoreId(storeId).Asin(asin).Ean(ean).Gtin(gtin).Upc(upc).Mpn(mpn).Isbn(isbn).ResponseFields(responseFields).Params(params).Exclude(exclude).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketplaceAPI.MarketplaceProductFind``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarketplaceProductFind`: ModelResponseMarketplaceProductFind
	fmt.Fprintf(os.Stdout, "Response from `MarketplaceAPI.MarketplaceProductFind`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarketplaceProductFindRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **count** | **int32** | This parameter sets the entity amount that has to be retrieved. Max allowed count&#x3D;250 | [default to 10]
 **pageCursor** | **string** | Used to retrieve entities via cursor-based pagination (it can&#39;t be used with any other filtering parameter) | 
 **keyword** | **string** | Defines search keyword | 
 **categoriesIds** | **string** | Defines product add that is specified by comma-separated categories id | 
 **storeId** | **string** | Store Id | 
 **asin** | **string** | Amazon Standard Identification Number. | 
 **ean** | **string** | European Article Number. An EAN is a unique 8 or 13-digit identifier that many industries (such as book publishers) use to identify products. | 
 **gtin** | **string** | Global Trade Item Number. An GTIN is an identifier for trade items. | 
 **upc** | **string** | Universal Product Code. A UPC (UPC-A) is a commonly used identifer for many different products. | 
 **mpn** | **string** | Manufacturer Part Number. A MPN is an identifier of a particular part design or material used. | 
 **isbn** | **string** | International Standard Book Number. An ISBN is a unique identifier for books. | 
 **responseFields** | **string** | Set this parameter in order to choose which entity fields you want to retrieve | 
 **params** | **string** | Set this parameter in order to choose which entity fields you want to retrieve | [default to &quot;force_all&quot;]
 **exclude** | **string** | Set this parameter in order to choose which entity fields you want to ignore. Works only if parameter &#x60;params&#x60; equal force_all | 

### Return type

[**ModelResponseMarketplaceProductFind**](ModelResponseMarketplaceProductFind.md)

### Authorization

[StoreKeyAuth](../README.md#StoreKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

