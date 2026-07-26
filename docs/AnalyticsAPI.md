# \AnalyticsAPI

All URIs are relative to *https://api.api2cart.local.com/v1.1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AnalyticsCustomerReport**](AnalyticsAPI.md#AnalyticsCustomerReport) | **Get** /analytics.customer_report.json | analytics.customer_report
[**AnalyticsProductReport**](AnalyticsAPI.md#AnalyticsProductReport) | **Get** /analytics.product_report.json | analytics.product_report
[**AnalyticsReport**](AnalyticsAPI.md#AnalyticsReport) | **Get** /analytics.report.json | analytics.report



## AnalyticsCustomerReport

> ResponseAnalyticsCustomerReportResult AnalyticsCustomerReport(ctx).DateFrom(dateFrom).DateTo(dateTo).Count(count).CurrencyId(currencyId).StoreId(storeId).CustomerType(customerType).Email(email).SortBy(sortBy).SortDirection(sortDirection).PageCursor(pageCursor).ResponseFields(responseFields).Execute()

analytics.customer_report



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
	dateFrom := "2026-01-01" // string | Start date for the analytics period (Y-m-d or Y-m-d H:i:s) (optional)
	dateTo := "2026-01-31" // string | End date for the analytics period (Y-m-d or Y-m-d H:i:s). Defaults to the current date. (optional)
	count := int32(20) // int32 | This parameter sets the entity amount that has to be retrieved. Max allowed count=250 (optional) (default to 10)
	currencyId := "usd" // string | Currency Id (optional)
	storeId := "1" // string | Store Id (optional)
	customerType := "registered" // string | Filter analytics customers by customer type (optional)
	email := "mail@example.com" // string | Filter analytics customers by email (optional)
	sortBy := "total_spend" // string | Set field to sort by (optional) (default to "total_spend")
	sortDirection := "asc" // string | Set sorting direction (optional) (default to "desc")
	pageCursor := "pageCursor_example" // string | Used to retrieve entities via cursor-based pagination (it can't be used with any other filtering parameter) (optional)
	responseFields := "{result}" // string | Set this parameter to choose which entity fields to retrieve. Use comma-separated field names in curly braces, nested to match the response structure, e.g. {result{product{id,name}}}. The wildcard * returns every field at a level: {*} gives the whole response, {result{product{*}}} all product fields. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalyticsAPI.AnalyticsCustomerReport(context.Background()).DateFrom(dateFrom).DateTo(dateTo).Count(count).CurrencyId(currencyId).StoreId(storeId).CustomerType(customerType).Email(email).SortBy(sortBy).SortDirection(sortDirection).PageCursor(pageCursor).ResponseFields(responseFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsAPI.AnalyticsCustomerReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AnalyticsCustomerReport`: ResponseAnalyticsCustomerReportResult
	fmt.Fprintf(os.Stdout, "Response from `AnalyticsAPI.AnalyticsCustomerReport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAnalyticsCustomerReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dateFrom** | **string** | Start date for the analytics period (Y-m-d or Y-m-d H:i:s) | 
 **dateTo** | **string** | End date for the analytics period (Y-m-d or Y-m-d H:i:s). Defaults to the current date. | 
 **count** | **int32** | This parameter sets the entity amount that has to be retrieved. Max allowed count&#x3D;250 | [default to 10]
 **currencyId** | **string** | Currency Id | 
 **storeId** | **string** | Store Id | 
 **customerType** | **string** | Filter analytics customers by customer type | 
 **email** | **string** | Filter analytics customers by email | 
 **sortBy** | **string** | Set field to sort by | [default to &quot;total_spend&quot;]
 **sortDirection** | **string** | Set sorting direction | [default to &quot;desc&quot;]
 **pageCursor** | **string** | Used to retrieve entities via cursor-based pagination (it can&#39;t be used with any other filtering parameter) | 
 **responseFields** | **string** | Set this parameter to choose which entity fields to retrieve. Use comma-separated field names in curly braces, nested to match the response structure, e.g. {result{product{id,name}}}. The wildcard * returns every field at a level: {*} gives the whole response, {result{product{*}}} all product fields. | 

### Return type

[**ResponseAnalyticsCustomerReportResult**](ResponseAnalyticsCustomerReportResult.md)

### Authorization

[StoreKeyAuth](../README.md#StoreKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AnalyticsProductReport

> ResponseAnalyticsProductReportResult AnalyticsProductReport(ctx).DateFrom(dateFrom).DateTo(dateTo).Count(count).ProductIds(productIds).CurrencyId(currencyId).StoreId(storeId).CategoriesIds(categoriesIds).SortBy(sortBy).SortDirection(sortDirection).PageCursor(pageCursor).ResponseFields(responseFields).Execute()

analytics.product_report



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
	dateFrom := "2026-01-01" // string | Start date for the analytics period (Y-m-d or Y-m-d H:i:s) (optional)
	dateTo := "2026-01-31" // string | End date for the analytics period (Y-m-d or Y-m-d H:i:s). Defaults to the current date. (optional)
	count := int32(20) // int32 | This parameter sets the entity amount that has to be retrieved. Max allowed count=250 (optional) (default to 10)
	productIds := "4,5" // string | Filter analytics by product ids (optional)
	currencyId := "usd" // string | Currency Id (optional)
	storeId := "1" // string | Store Id (optional)
	categoriesIds := "23,56" // string | Defines product add that is specified by comma-separated categories id (optional)
	sortBy := "items_sold" // string | Set field to sort by (optional) (default to "items_sold")
	sortDirection := "asc" // string | Set sorting direction (optional) (default to "desc")
	pageCursor := "pageCursor_example" // string | Used to retrieve entities via cursor-based pagination (it can't be used with any other filtering parameter) (optional)
	responseFields := "{result}" // string | Set this parameter to choose which entity fields to retrieve. Use comma-separated field names in curly braces, nested to match the response structure, e.g. {result{product{id,name}}}. The wildcard * returns every field at a level: {*} gives the whole response, {result{product{*}}} all product fields. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalyticsAPI.AnalyticsProductReport(context.Background()).DateFrom(dateFrom).DateTo(dateTo).Count(count).ProductIds(productIds).CurrencyId(currencyId).StoreId(storeId).CategoriesIds(categoriesIds).SortBy(sortBy).SortDirection(sortDirection).PageCursor(pageCursor).ResponseFields(responseFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsAPI.AnalyticsProductReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AnalyticsProductReport`: ResponseAnalyticsProductReportResult
	fmt.Fprintf(os.Stdout, "Response from `AnalyticsAPI.AnalyticsProductReport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAnalyticsProductReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dateFrom** | **string** | Start date for the analytics period (Y-m-d or Y-m-d H:i:s) | 
 **dateTo** | **string** | End date for the analytics period (Y-m-d or Y-m-d H:i:s). Defaults to the current date. | 
 **count** | **int32** | This parameter sets the entity amount that has to be retrieved. Max allowed count&#x3D;250 | [default to 10]
 **productIds** | **string** | Filter analytics by product ids | 
 **currencyId** | **string** | Currency Id | 
 **storeId** | **string** | Store Id | 
 **categoriesIds** | **string** | Defines product add that is specified by comma-separated categories id | 
 **sortBy** | **string** | Set field to sort by | [default to &quot;items_sold&quot;]
 **sortDirection** | **string** | Set sorting direction | [default to &quot;desc&quot;]
 **pageCursor** | **string** | Used to retrieve entities via cursor-based pagination (it can&#39;t be used with any other filtering parameter) | 
 **responseFields** | **string** | Set this parameter to choose which entity fields to retrieve. Use comma-separated field names in curly braces, nested to match the response structure, e.g. {result{product{id,name}}}. The wildcard * returns every field at a level: {*} gives the whole response, {result{product{*}}} all product fields. | 

### Return type

[**ResponseAnalyticsProductReportResult**](ResponseAnalyticsProductReportResult.md)

### Authorization

[StoreKeyAuth](../README.md#StoreKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AnalyticsReport

> ResponseAnalyticsReportResult AnalyticsReport(ctx).DateFrom(dateFrom).DateTo(dateTo).Interval(interval).OrderStatus(orderStatus).FinancialStatus(financialStatus).CurrencyId(currencyId).StoreId(storeId).SortBy(sortBy).SortDirection(sortDirection).ResponseFields(responseFields).Execute()

analytics.report



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
	dateFrom := "2026-01-01" // string | Start date for the analytics period (Y-m-d or Y-m-d H:i:s)
	dateTo := "2026-01-31" // string | End date for the analytics period (Y-m-d or Y-m-d H:i:s). Defaults to the current date. (optional)
	interval := "day" // string | Interval for analytics report breakdown (optional)
	orderStatus := "Completed" // string | Retrieves orders specified by order status (optional)
	financialStatus := "paid" // string | Retrieves orders specified by financial status (optional)
	currencyId := "usd" // string | Currency Id (optional)
	storeId := "1" // string | Store Id (optional)
	sortBy := "date" // string | Set field to sort by (optional) (default to "date")
	sortDirection := "asc" // string | Set sorting direction (optional) (default to "asc")
	responseFields := "{result}" // string | Set this parameter to choose which entity fields to retrieve. Use comma-separated field names in curly braces, nested to match the response structure, e.g. {result{product{id,name}}}. The wildcard * returns every field at a level: {*} gives the whole response, {result{product{*}}} all product fields. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalyticsAPI.AnalyticsReport(context.Background()).DateFrom(dateFrom).DateTo(dateTo).Interval(interval).OrderStatus(orderStatus).FinancialStatus(financialStatus).CurrencyId(currencyId).StoreId(storeId).SortBy(sortBy).SortDirection(sortDirection).ResponseFields(responseFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsAPI.AnalyticsReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AnalyticsReport`: ResponseAnalyticsReportResult
	fmt.Fprintf(os.Stdout, "Response from `AnalyticsAPI.AnalyticsReport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAnalyticsReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dateFrom** | **string** | Start date for the analytics period (Y-m-d or Y-m-d H:i:s) | 
 **dateTo** | **string** | End date for the analytics period (Y-m-d or Y-m-d H:i:s). Defaults to the current date. | 
 **interval** | **string** | Interval for analytics report breakdown | 
 **orderStatus** | **string** | Retrieves orders specified by order status | 
 **financialStatus** | **string** | Retrieves orders specified by financial status | 
 **currencyId** | **string** | Currency Id | 
 **storeId** | **string** | Store Id | 
 **sortBy** | **string** | Set field to sort by | [default to &quot;date&quot;]
 **sortDirection** | **string** | Set sorting direction | [default to &quot;asc&quot;]
 **responseFields** | **string** | Set this parameter to choose which entity fields to retrieve. Use comma-separated field names in curly braces, nested to match the response structure, e.g. {result{product{id,name}}}. The wildcard * returns every field at a level: {*} gives the whole response, {result{product{*}}} all product fields. | 

### Return type

[**ResponseAnalyticsReportResult**](ResponseAnalyticsReportResult.md)

### Authorization

[StoreKeyAuth](../README.md#StoreKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

