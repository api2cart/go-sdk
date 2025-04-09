# \BatchAPI

All URIs are relative to *https://api.api2cart.com/v1.1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BatchJobList**](BatchAPI.md#BatchJobList) | **Get** /batch.job.list.json | batch.job.list
[**BatchJobResult**](BatchAPI.md#BatchJobResult) | **Get** /batch.job.result.json | batch.job.result



## BatchJobList

> ModelResponseBatchJobList BatchJobList(ctx).Count(count).PageCursor(pageCursor).CreatedFrom(createdFrom).CreatedTo(createdTo).ProcessedFrom(processedFrom).ProcessedTo(processedTo).Ids(ids).ResponseFields(responseFields).Execute()

batch.job.list



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
	createdFrom := "2010-07-29 13:45:52" // string | Retrieve entities from their creation date (optional)
	createdTo := "2100-08-29 13:45:52" // string | Retrieve entities to their creation date (optional)
	processedFrom := "2100-08-29 13:45:52" // string | Retrieve entities according to their processing datetime (optional)
	processedTo := "2100-08-29 13:45:52" // string | Retrieve entities according to their processing datetime (optional)
	ids := "24,25" // string | Filter batch jobs by ids (optional)
	responseFields := "{result}" // string | Set this parameter in order to choose which entity fields you want to retrieve (optional) (default to "{return_code,return_message,pagination,result}")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BatchAPI.BatchJobList(context.Background()).Count(count).PageCursor(pageCursor).CreatedFrom(createdFrom).CreatedTo(createdTo).ProcessedFrom(processedFrom).ProcessedTo(processedTo).Ids(ids).ResponseFields(responseFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchAPI.BatchJobList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BatchJobList`: ModelResponseBatchJobList
	fmt.Fprintf(os.Stdout, "Response from `BatchAPI.BatchJobList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBatchJobListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **count** | **int32** | This parameter sets the entity amount that has to be retrieved. Max allowed count&#x3D;250 | [default to 10]
 **pageCursor** | **string** | Used to retrieve entities via cursor-based pagination (it can&#39;t be used with any other filtering parameter) | 
 **createdFrom** | **string** | Retrieve entities from their creation date | 
 **createdTo** | **string** | Retrieve entities to their creation date | 
 **processedFrom** | **string** | Retrieve entities according to their processing datetime | 
 **processedTo** | **string** | Retrieve entities according to their processing datetime | 
 **ids** | **string** | Filter batch jobs by ids | 
 **responseFields** | **string** | Set this parameter in order to choose which entity fields you want to retrieve | [default to &quot;{return_code,return_message,pagination,result}&quot;]

### Return type

[**ModelResponseBatchJobList**](ModelResponseBatchJobList.md)

### Authorization

[StoreKeyAuth](../README.md#StoreKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BatchJobResult

> ResponseBatchJobResult BatchJobResult(ctx).Id(id).Execute()

batch.job.result



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
	resp, r, err := apiClient.BatchAPI.BatchJobResult(context.Background()).Id(id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchAPI.BatchJobResult``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BatchJobResult`: ResponseBatchJobResult
	fmt.Fprintf(os.Stdout, "Response from `BatchAPI.BatchJobResult`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBatchJobResultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | Entity id | 

### Return type

[**ResponseBatchJobResult**](ResponseBatchJobResult.md)

### Authorization

[StoreKeyAuth](../README.md#StoreKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

