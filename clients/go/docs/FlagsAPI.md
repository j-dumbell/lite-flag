# \FlagsAPI

All URIs are relative to *https://api.example.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FlagsGet**](FlagsAPI.md#FlagsGet) | **Get** /flags | Retrieve all feature flags
[**FlagsKeyDelete**](FlagsAPI.md#FlagsKeyDelete) | **Delete** /flags/{key} | Delete a feature flag
[**FlagsKeyGet**](FlagsAPI.md#FlagsKeyGet) | **Get** /flags/{key} | Retrieve a single feature flag by key
[**FlagsKeyPut**](FlagsAPI.md#FlagsKeyPut) | **Put** /flags/{key} | Update an existing feature flag
[**FlagsPost**](FlagsAPI.md#FlagsPost) | **Post** /flags | Create a new feature flag



## FlagsGet

> []Flag FlagsGet(ctx).Execute()

Retrieve all feature flags

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
	resp, r, err := apiClient.FlagsAPI.FlagsGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlagsAPI.FlagsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FlagsGet`: []Flag
	fmt.Fprintf(os.Stdout, "Response from `FlagsAPI.FlagsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiFlagsGetRequest struct via the builder pattern


### Return type

[**[]Flag**](Flag.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FlagsKeyDelete

> FlagsKeyDelete(ctx, key).Execute()

Delete a feature flag

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
	key := "key_example" // string | Unique key of the feature flag

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FlagsAPI.FlagsKeyDelete(context.Background(), key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlagsAPI.FlagsKeyDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** | Unique key of the feature flag | 

### Other Parameters

Other parameters are passed through a pointer to a apiFlagsKeyDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FlagsKeyGet

> Flag FlagsKeyGet(ctx, key).Execute()

Retrieve a single feature flag by key

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
	key := "key_example" // string | Unique key of the feature flag

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlagsAPI.FlagsKeyGet(context.Background(), key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlagsAPI.FlagsKeyGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FlagsKeyGet`: Flag
	fmt.Fprintf(os.Stdout, "Response from `FlagsAPI.FlagsKeyGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** | Unique key of the feature flag | 

### Other Parameters

Other parameters are passed through a pointer to a apiFlagsKeyGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Flag**](Flag.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FlagsKeyPut

> Flag FlagsKeyPut(ctx, key).FlagInput(flagInput).Execute()

Update an existing feature flag

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
	key := "key_example" // string | Unique key of the feature flag
	flagInput := *openapiclient.NewFlagInput(false, "Type_example", openapiclient.Flag_value{Bool: new(bool)}) // FlagInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlagsAPI.FlagsKeyPut(context.Background(), key).FlagInput(flagInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlagsAPI.FlagsKeyPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FlagsKeyPut`: Flag
	fmt.Fprintf(os.Stdout, "Response from `FlagsAPI.FlagsKeyPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** | Unique key of the feature flag | 

### Other Parameters

Other parameters are passed through a pointer to a apiFlagsKeyPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **flagInput** | [**FlagInput**](FlagInput.md) |  | 

### Return type

[**Flag**](Flag.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FlagsPost

> Flag FlagsPost(ctx).Flag(flag).Execute()

Create a new feature flag

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
	flag := *openapiclient.NewFlag("Key_example", "Type_example", false, openapiclient.Flag_value{Bool: new(bool)}) // Flag | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlagsAPI.FlagsPost(context.Background()).Flag(flag).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlagsAPI.FlagsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FlagsPost`: Flag
	fmt.Fprintf(os.Stdout, "Response from `FlagsAPI.FlagsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFlagsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **flag** | [**Flag**](Flag.md) |  | 

### Return type

[**Flag**](Flag.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

