# \BridgeAPI

All URIs are relative to *https://api.api2cart.local.com/v1.1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BridgeDelete**](BridgeAPI.md#BridgeDelete) | **Post** /bridge.delete.json | bridge.delete
[**BridgeDownload**](BridgeAPI.md#BridgeDownload) | **Get** /bridge.download.file | bridge.download
[**BridgeUpdate**](BridgeAPI.md#BridgeUpdate) | **Post** /bridge.update.json | bridge.update



## BridgeDelete

> AttributeValueDelete200Response BridgeDelete(ctx).Execute()

bridge.delete



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
	resp, r, err := apiClient.BridgeAPI.BridgeDelete(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BridgeAPI.BridgeDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BridgeDelete`: AttributeValueDelete200Response
	fmt.Fprintf(os.Stdout, "Response from `BridgeAPI.BridgeDelete`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiBridgeDeleteRequest struct via the builder pattern


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


## BridgeDownload

> *os.File BridgeDownload(ctx).Whitelabel(whitelabel).Execute()

bridge.download



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
	whitelabel := true // bool | Identifies if there is a necessity to download whitelabel bridge. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BridgeAPI.BridgeDownload(context.Background()).Whitelabel(whitelabel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BridgeAPI.BridgeDownload``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BridgeDownload`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `BridgeAPI.BridgeDownload`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBridgeDownloadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **whitelabel** | **bool** | Identifies if there is a necessity to download whitelabel bridge. | [default to false]

### Return type

[***os.File**](*os.File.md)

### Authorization

[StoreKeyAuth](../README.md#StoreKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/zip

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BridgeUpdate

> AttributeUpdate200Response BridgeUpdate(ctx).Execute()

bridge.update



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
	resp, r, err := apiClient.BridgeAPI.BridgeUpdate(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BridgeAPI.BridgeUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BridgeUpdate`: AttributeUpdate200Response
	fmt.Fprintf(os.Stdout, "Response from `BridgeAPI.BridgeUpdate`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiBridgeUpdateRequest struct via the builder pattern


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

