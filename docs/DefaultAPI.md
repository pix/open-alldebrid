# \DefaultAPI

All URIs are relative to *https://api.alldebrid.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**HostsDomainsGet**](DefaultAPI.md#HostsDomainsGet) | **Get** /hosts/domains | Use this endpoint to only retrieve the list of supported hosts domains and redirectors as an array.
[**HostsGet**](DefaultAPI.md#HostsGet) | **Get** /hosts | Use this endpoint to retrieve informations about what hosts we support.
[**HostsPriorityGet**](DefaultAPI.md#HostsPriorityGet) | **Get** /hosts/priority | Not all hosts are created equal, so some hosts are more limited than other.
[**LinkDelayedGet**](DefaultAPI.md#LinkDelayedGet) | **Get** /link/delayed | This endpoint give the status of a delayed link.
[**LinkInfosGet**](DefaultAPI.md#LinkInfosGet) | **Get** /link/infos | Use this endpoint to retrieve informations about a link.
[**LinkRedirectorGet**](DefaultAPI.md#LinkRedirectorGet) | **Get** /link/redirector | Use this endpoint to retrieve links protected by a redirector or link protector.
[**LinkStreamingGet**](DefaultAPI.md#LinkStreamingGet) | **Get** /link/streaming | The unlocking flow for streaming link is a bit more complex.
[**LinkUnlockGet**](DefaultAPI.md#LinkUnlockGet) | **Get** /link/unlock | This endpoint unlocks a given link.
[**MagnetDeleteGet**](DefaultAPI.md#MagnetDeleteGet) | **Get** /magnet/delete | Delete a magnet.
[**MagnetInstantGet**](DefaultAPI.md#MagnetInstantGet) | **Get** /magnet/instant | Check if a magnet is available instantly.
[**MagnetRestartGet**](DefaultAPI.md#MagnetRestartGet) | **Get** /magnet/restart | Restart a failed magnet, or multiple failed magnets at once.
[**MagnetStatusGet**](DefaultAPI.md#MagnetStatusGet) | **Get** /magnet/status | Get the status of current magnets, or only one if you specify a magnet ID.
[**MagnetUploadFilePost**](DefaultAPI.md#MagnetUploadFilePost) | **Post** /magnet/upload/file | Upload torrent files.
[**MagnetUploadGet**](DefaultAPI.md#MagnetUploadGet) | **Get** /magnet/upload | Upload a magnet with its URI or hash.
[**UserGet**](DefaultAPI.md#UserGet) | **Get** /user | Use this endpoint to get user informations.
[**UserHistoryDeleteGet**](DefaultAPI.md#UserHistoryDeleteGet) | **Get** /user/history/delete | Use this endpoint to delete all links currently in your recent links history.
[**UserHistoryGet**](DefaultAPI.md#UserHistoryGet) | **Get** /user/history | Use this endpoint to get recent links.
[**UserHostsGet**](DefaultAPI.md#UserHostsGet) | **Get** /user/hosts | This endpoint retrieves a complete list of all available hosts for this user.
[**UserLinksDeleteGet**](DefaultAPI.md#UserLinksDeleteGet) | **Get** /user/links/delete | Delete a saved link.
[**UserLinksGet**](DefaultAPI.md#UserLinksGet) | **Get** /user/links | Use this endpoint to get links the user saved for later use.
[**UserLinksSaveGet**](DefaultAPI.md#UserLinksSaveGet) | **Get** /user/links/save | Save a link.
[**UserNotificationClearGet**](DefaultAPI.md#UserNotificationClearGet) | **Get** /user/notification/clear | This endpoint clears a user notification with its code.



## HostsDomainsGet

> HostsDomainsGet200Response HostsDomainsGet(ctx).Agent(agent).Execute()

Use this endpoint to only retrieve the list of supported hosts domains and redirectors as an array.



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
	agent := "agent_example" // string | Your software user-agent. (default to "open-alldebrid")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.HostsDomainsGet(context.Background()).Agent(agent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.HostsDomainsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HostsDomainsGet`: HostsDomainsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.HostsDomainsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiHostsDomainsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **agent** | **string** | Your software user-agent. | [default to &quot;open-alldebrid&quot;]

### Return type

[**HostsDomainsGet200Response**](HostsDomainsGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HostsGet

> HostsGet200Response HostsGet(ctx).Agent(agent).HostOnly(hostOnly).Execute()

Use this endpoint to retrieve informations about what hosts we support.



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
	agent := "agent_example" // string | Your software user-agent. (default to "open-alldebrid")
	hostOnly := "hostOnly_example" // string | Endpoint will only return \"hosts\" data (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.HostsGet(context.Background()).Agent(agent).HostOnly(hostOnly).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.HostsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HostsGet`: HostsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.HostsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiHostsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **agent** | **string** | Your software user-agent. | [default to &quot;open-alldebrid&quot;]
 **hostOnly** | **string** | Endpoint will only return \&quot;hosts\&quot; data | 

### Return type

[**HostsGet200Response**](HostsGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HostsPriorityGet

> HostsPriorityGet200Response HostsPriorityGet(ctx).Agent(agent).Execute()

Not all hosts are created equal, so some hosts are more limited than other.



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
	agent := "agent_example" // string | Your software user-agent. (default to "open-alldebrid")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.HostsPriorityGet(context.Background()).Agent(agent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.HostsPriorityGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HostsPriorityGet`: HostsPriorityGet200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.HostsPriorityGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiHostsPriorityGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **agent** | **string** | Your software user-agent. | [default to &quot;open-alldebrid&quot;]

### Return type

[**HostsPriorityGet200Response**](HostsPriorityGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LinkDelayedGet

> LinkDelayedGet200Response LinkDelayedGet(ctx).Agent(agent).Id(id).Apikey(apikey).Execute()

This endpoint give the status of a delayed link.



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
	agent := "agent_example" // string | Your software user-agent. (default to "open-alldebrid")
	id := "id_example" // string | Delayed ID received in /link/unlock.
	apikey := "apikey_example" // string | Deprecated User apikey (Use Bearer Auth in header). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.LinkDelayedGet(context.Background()).Agent(agent).Id(id).Apikey(apikey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.LinkDelayedGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LinkDelayedGet`: LinkDelayedGet200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.LinkDelayedGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLinkDelayedGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **agent** | **string** | Your software user-agent. | [default to &quot;open-alldebrid&quot;]
 **id** | **string** | Delayed ID received in /link/unlock. | 
 **apikey** | **string** | Deprecated User apikey (Use Bearer Auth in header). | 

### Return type

[**LinkDelayedGet200Response**](LinkDelayedGet200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LinkInfosGet

> LinkInfosGet200Response LinkInfosGet(ctx).Agent(agent).Link(link).Apikey(apikey).Password(password).Execute()

Use this endpoint to retrieve informations about a link.



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
	agent := "agent_example" // string | Your software user-agent. (default to "open-alldebrid")
	link := []string{"Inner_example"} // []string | The link or array of links you request informations about.
	apikey := "apikey_example" // string | Deprecated User apikey (Use Bearer Auth in header). (optional)
	password := "password_example" // string | Link password (supported on uptobox / 1fichier). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.LinkInfosGet(context.Background()).Agent(agent).Link(link).Apikey(apikey).Password(password).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.LinkInfosGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LinkInfosGet`: LinkInfosGet200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.LinkInfosGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLinkInfosGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **agent** | **string** | Your software user-agent. | [default to &quot;open-alldebrid&quot;]
 **link** | **[]string** | The link or array of links you request informations about. | 
 **apikey** | **string** | Deprecated User apikey (Use Bearer Auth in header). | 
 **password** | **string** | Link password (supported on uptobox / 1fichier). | 

### Return type

[**LinkInfosGet200Response**](LinkInfosGet200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LinkRedirectorGet

> LinkRedirectorGet200Response LinkRedirectorGet(ctx).Agent(agent).Link(link).Apikey(apikey).Execute()

Use this endpoint to retrieve links protected by a redirector or link protector.



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
	agent := "agent_example" // string | Your software user-agent. (default to "open-alldebrid")
	link := "link_example" // string | The redirector or protector link to extract links.
	apikey := "apikey_example" // string | Deprecated User apikey (Use Bearer Auth in header). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.LinkRedirectorGet(context.Background()).Agent(agent).Link(link).Apikey(apikey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.LinkRedirectorGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LinkRedirectorGet`: LinkRedirectorGet200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.LinkRedirectorGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLinkRedirectorGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **agent** | **string** | Your software user-agent. | [default to &quot;open-alldebrid&quot;]
 **link** | **string** | The redirector or protector link to extract links. | 
 **apikey** | **string** | Deprecated User apikey (Use Bearer Auth in header). | 

### Return type

[**LinkRedirectorGet200Response**](LinkRedirectorGet200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LinkStreamingGet

> LinkStreamingGet200Response LinkStreamingGet(ctx).Agent(agent).Id(id).Stream(stream).Apikey(apikey).Execute()

The unlocking flow for streaming link is a bit more complex.



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
	agent := "agent_example" // string | Your software user-agent. (default to "open-alldebrid")
	id := "id_example" // string | The link ID you received from the /link/unlock call.
	stream := "stream_example" // string | The stream ID you choosed from the stream qualities list returned by /link/unlock.
	apikey := "apikey_example" // string | Deprecated User apikey (Use Bearer Auth in header). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.LinkStreamingGet(context.Background()).Agent(agent).Id(id).Stream(stream).Apikey(apikey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.LinkStreamingGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LinkStreamingGet`: LinkStreamingGet200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.LinkStreamingGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLinkStreamingGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **agent** | **string** | Your software user-agent. | [default to &quot;open-alldebrid&quot;]
 **id** | **string** | The link ID you received from the /link/unlock call. | 
 **stream** | **string** | The stream ID you choosed from the stream qualities list returned by /link/unlock. | 
 **apikey** | **string** | Deprecated User apikey (Use Bearer Auth in header). | 

### Return type

[**LinkStreamingGet200Response**](LinkStreamingGet200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LinkUnlockGet

> LinkUnlockGet200Response LinkUnlockGet(ctx).Agent(agent).Link(link).Apikey(apikey).Password(password).Execute()

This endpoint unlocks a given link.



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
	agent := "agent_example" // string | Your software user-agent. (default to "open-alldebrid")
	link := "link_example" // string | The redirector or protector link to extract links.
	apikey := "apikey_example" // string | Deprecated User apikey (Use Bearer Auth in header). (optional)
	password := "password_example" // string | Link password (supported on uptobox / 1fichier). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.LinkUnlockGet(context.Background()).Agent(agent).Link(link).Apikey(apikey).Password(password).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.LinkUnlockGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LinkUnlockGet`: LinkUnlockGet200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.LinkUnlockGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLinkUnlockGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **agent** | **string** | Your software user-agent. | [default to &quot;open-alldebrid&quot;]
 **link** | **string** | The redirector or protector link to extract links. | 
 **apikey** | **string** | Deprecated User apikey (Use Bearer Auth in header). | 
 **password** | **string** | Link password (supported on uptobox / 1fichier). | 

### Return type

[**LinkUnlockGet200Response**](LinkUnlockGet200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MagnetDeleteGet

> MagnetDeleteGet200Response MagnetDeleteGet(ctx).Agent(agent).Id(id).Apikey(apikey).Execute()

Delete a magnet.



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
	agent := "agent_example" // string | Your software user-agent. (default to "open-alldebrid")
	id := "id_example" // string | Magnet ID.
	apikey := "apikey_example" // string | Deprecated User apikey (Use Bearer Auth in header). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.MagnetDeleteGet(context.Background()).Agent(agent).Id(id).Apikey(apikey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.MagnetDeleteGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MagnetDeleteGet`: MagnetDeleteGet200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.MagnetDeleteGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMagnetDeleteGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **agent** | **string** | Your software user-agent. | [default to &quot;open-alldebrid&quot;]
 **id** | **string** | Magnet ID. | 
 **apikey** | **string** | Deprecated User apikey (Use Bearer Auth in header). | 

### Return type

[**MagnetDeleteGet200Response**](MagnetDeleteGet200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MagnetInstantGet

> MagnetInstantGet200Response MagnetInstantGet(ctx).Agent(agent).Magnets(magnets).Apikey(apikey).Execute()

Check if a magnet is available instantly.



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
	agent := "agent_example" // string | Your software user-agent. (default to "open-alldebrid")
	magnets := []string{"Inner_example"} // []string | Magnets URI or hash you wish to check instant availability, can be one or many links
	apikey := "apikey_example" // string | Deprecated User apikey (Use Bearer Auth in header). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.MagnetInstantGet(context.Background()).Agent(agent).Magnets(magnets).Apikey(apikey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.MagnetInstantGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MagnetInstantGet`: MagnetInstantGet200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.MagnetInstantGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMagnetInstantGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **agent** | **string** | Your software user-agent. | [default to &quot;open-alldebrid&quot;]
 **magnets** | **[]string** | Magnets URI or hash you wish to check instant availability, can be one or many links | 
 **apikey** | **string** | Deprecated User apikey (Use Bearer Auth in header). | 

### Return type

[**MagnetInstantGet200Response**](MagnetInstantGet200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MagnetRestartGet

> MagnetRestartGet200Response MagnetRestartGet(ctx).Agent(agent).Id(id).Apikey(apikey).Execute()

Restart a failed magnet, or multiple failed magnets at once.



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
	agent := "agent_example" // string | Your software user-agent. (default to "open-alldebrid")
	id := "id_example" // string | Magnet ID
	apikey := "apikey_example" // string | Deprecated User apikey (Use Bearer Auth in header). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.MagnetRestartGet(context.Background()).Agent(agent).Id(id).Apikey(apikey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.MagnetRestartGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MagnetRestartGet`: MagnetRestartGet200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.MagnetRestartGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMagnetRestartGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **agent** | **string** | Your software user-agent. | [default to &quot;open-alldebrid&quot;]
 **id** | **string** | Magnet ID |
 **apikey** | **string** | Deprecated User apikey (Use Bearer Auth in header). | 

### Return type

[**MagnetRestartGet200Response**](MagnetRestartGet200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MagnetStatusGet

> MagnetStatusGet200Response MagnetStatusGet(ctx).Agent(agent).Apikey(apikey).Id(id).Status(status).Session(session).Counter(counter).Execute()

Get the status of current magnets, or only one if you specify a magnet ID.



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
	agent := "agent_example" // string | Your software user-agent. (default to "open-alldebrid")
	apikey := "apikey_example" // string | Deprecated User apikey (Use Bearer Auth in header). (optional)
	id := "id_example" // string | Magnet ID. (optional)
	status := "status_example" // string | Magnets status filter. Either active, ready, expired or error (optional)
	session := "session_example" // string | Session ID for Live mode (see Live Mode). (optional)
	counter := "counter_example" // string | Counter for Live mode (see Live Mode). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.MagnetStatusGet(context.Background()).Agent(agent).Apikey(apikey).Id(id).Status(status).Session(session).Counter(counter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.MagnetStatusGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MagnetStatusGet`: MagnetStatusGet200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.MagnetStatusGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMagnetStatusGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **agent** | **string** | Your software user-agent. | [default to &quot;open-alldebrid&quot;]
 **apikey** | **string** | Deprecated User apikey (Use Bearer Auth in header). | 
 **id** | **string** | Magnet ID. | 
 **status** | **string** | Magnets status filter. Either active, ready, expired or error | 
 **session** | **string** | Session ID for Live mode (see Live Mode). | 
 **counter** | **string** | Counter for Live mode (see Live Mode). | 

### Return type

[**MagnetStatusGet200Response**](MagnetStatusGet200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MagnetUploadFilePost

> MagnetUploadFilePost200Response MagnetUploadFilePost(ctx).Agent(agent).Apikey(apikey).Files(files).Execute()

Upload torrent files.



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
	agent := "agent_example" // string | Your software user-agent. (default to "open-alldebrid")
	apikey := "apikey_example" // string | Deprecated User apikey (Use Bearer Auth in header). (optional)
	files := os.NewFile(1234, "some_file") // *os.File |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.MagnetUploadFilePost(context.Background()).Agent(agent).Apikey(apikey).Files(files).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.MagnetUploadFilePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MagnetUploadFilePost`: MagnetUploadFilePost200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.MagnetUploadFilePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMagnetUploadFilePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **agent** | **string** | Your software user-agent. | [default to &quot;open-alldebrid&quot;]
 **apikey** | **string** | Deprecated User apikey (Use Bearer Auth in header). | 
 **files** | ***os.File** |  |

### Return type

[**MagnetUploadFilePost200Response**](MagnetUploadFilePost200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MagnetUploadGet

> MagnetUploadGet200Response MagnetUploadGet(ctx).Agent(agent).Magnets(magnets).Apikey(apikey).Execute()

Upload a magnet with its URI or hash.



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
	agent := "agent_example" // string | Your software user-agent. (default to "open-alldebrid")
	magnets := []string{"Inner_example"} // []string | Magnet(s) URI or hash. Must send magnet either in GET param or in POST data.
	apikey := "apikey_example" // string | Deprecated User apikey (Use Bearer Auth in header). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.MagnetUploadGet(context.Background()).Agent(agent).Magnets(magnets).Apikey(apikey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.MagnetUploadGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MagnetUploadGet`: MagnetUploadGet200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.MagnetUploadGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMagnetUploadGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **agent** | **string** | Your software user-agent. | [default to &quot;open-alldebrid&quot;]
 **magnets** | **[]string** | Magnet(s) URI or hash. Must send magnet either in GET param or in POST data. | 
 **apikey** | **string** | Deprecated User apikey (Use Bearer Auth in header). | 

### Return type

[**MagnetUploadGet200Response**](MagnetUploadGet200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserGet

> UserGet200Response UserGet(ctx).Agent(agent).Apikey(apikey).Execute()

Use this endpoint to get user informations.



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
	agent := "agent_example" // string | Your software user-agent. (default to "open-alldebrid")
	apikey := "apikey_example" // string | Deprecated User apikey (Use Bearer Auth in header). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.UserGet(context.Background()).Agent(agent).Apikey(apikey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UserGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserGet`: UserGet200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.UserGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUserGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **agent** | **string** | Your software user-agent. | [default to &quot;open-alldebrid&quot;]
 **apikey** | **string** | Deprecated User apikey (Use Bearer Auth in header). | 

### Return type

[**UserGet200Response**](UserGet200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserHistoryDeleteGet

> UserHistoryDeleteGet200Response UserHistoryDeleteGet(ctx).Agent(agent).Apikey(apikey).Execute()

Use this endpoint to delete all links currently in your recent links history.



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
	agent := "agent_example" // string | Your software user-agent. (default to "open-alldebrid")
	apikey := "apikey_example" // string | Deprecated User apikey (Use Bearer Auth in header). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.UserHistoryDeleteGet(context.Background()).Agent(agent).Apikey(apikey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UserHistoryDeleteGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserHistoryDeleteGet`: UserHistoryDeleteGet200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.UserHistoryDeleteGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUserHistoryDeleteGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **agent** | **string** | Your software user-agent. | [default to &quot;open-alldebrid&quot;]
 **apikey** | **string** | Deprecated User apikey (Use Bearer Auth in header). | 

### Return type

[**UserHistoryDeleteGet200Response**](UserHistoryDeleteGet200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserHistoryGet

> UserLinksGet200Response UserHistoryGet(ctx).Agent(agent).Apikey(apikey).Execute()

Use this endpoint to get recent links.



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
	agent := "agent_example" // string | Your software user-agent. (default to "open-alldebrid")
	apikey := "apikey_example" // string | Deprecated User apikey (Use Bearer Auth in header). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.UserHistoryGet(context.Background()).Agent(agent).Apikey(apikey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UserHistoryGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserHistoryGet`: UserLinksGet200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.UserHistoryGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUserHistoryGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **agent** | **string** | Your software user-agent. | [default to &quot;open-alldebrid&quot;]
 **apikey** | **string** | Deprecated User apikey (Use Bearer Auth in header). | 

### Return type

[**UserLinksGet200Response**](UserLinksGet200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserHostsGet

> HostsGet200Response UserHostsGet(ctx).Agent(agent).Apikey(apikey).HostOnly(hostOnly).Execute()

This endpoint retrieves a complete list of all available hosts for this user.



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
	agent := "agent_example" // string | Your software user-agent. (default to "open-alldebrid")
	apikey := "apikey_example" // string | Deprecated User apikey (Use Bearer Auth in header). (optional)
	hostOnly := "hostOnly_example" // string | Endpoint will only return \"hosts\" data (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.UserHostsGet(context.Background()).Agent(agent).Apikey(apikey).HostOnly(hostOnly).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UserHostsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserHostsGet`: HostsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.UserHostsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUserHostsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **agent** | **string** | Your software user-agent. | [default to &quot;open-alldebrid&quot;]
 **apikey** | **string** | Deprecated User apikey (Use Bearer Auth in header). | 
 **hostOnly** | **string** | Endpoint will only return \&quot;hosts\&quot; data | 

### Return type

[**HostsGet200Response**](HostsGet200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserLinksDeleteGet

> UserLinksDeleteGet200Response UserLinksDeleteGet(ctx).Agent(agent).Apikey(apikey).Link(link).Link2(link2).Execute()

Delete a saved link.



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
	agent := "agent_example" // string | Your software user-agent. (default to "open-alldebrid")
	apikey := "apikey_example" // string | Deprecated User apikey (Use Bearer Auth in header). (optional)
	link := "link_example" // string | Link to delete. (optional)
	link2 := []string{"Inner_example"} // []string | Links to delete. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.UserLinksDeleteGet(context.Background()).Agent(agent).Apikey(apikey).Link(link).Link2(link2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UserLinksDeleteGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserLinksDeleteGet`: UserLinksDeleteGet200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.UserLinksDeleteGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUserLinksDeleteGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **agent** | **string** | Your software user-agent. | [default to &quot;open-alldebrid&quot;]
 **apikey** | **string** | Deprecated User apikey (Use Bearer Auth in header). | 
 **link** | **string** | Link to delete. | 
 **link2** | **[]string** | Links to delete. | 

### Return type

[**UserLinksDeleteGet200Response**](UserLinksDeleteGet200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserLinksGet

> UserLinksGet200Response UserLinksGet(ctx).Agent(agent).Apikey(apikey).Execute()

Use this endpoint to get links the user saved for later use.



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
	agent := "agent_example" // string | Your software user-agent. (default to "open-alldebrid")
	apikey := "apikey_example" // string | Deprecated User apikey (Use Bearer Auth in header). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.UserLinksGet(context.Background()).Agent(agent).Apikey(apikey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UserLinksGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserLinksGet`: UserLinksGet200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.UserLinksGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUserLinksGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **agent** | **string** | Your software user-agent. | [default to &quot;open-alldebrid&quot;]
 **apikey** | **string** | Deprecated User apikey (Use Bearer Auth in header). | 

### Return type

[**UserLinksGet200Response**](UserLinksGet200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserLinksSaveGet

> UserLinksSaveGet200Response UserLinksSaveGet(ctx).Agent(agent).Link(link).Apikey(apikey).Execute()

Save a link.



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
	agent := "agent_example" // string | Your software user-agent. (default to "open-alldebrid")
	link := []string{"Inner_example"} // []string | Links to save.
	apikey := "apikey_example" // string | Deprecated User apikey (Use Bearer Auth in header). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.UserLinksSaveGet(context.Background()).Agent(agent).Link(link).Apikey(apikey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UserLinksSaveGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserLinksSaveGet`: UserLinksSaveGet200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.UserLinksSaveGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUserLinksSaveGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **agent** | **string** | Your software user-agent. | [default to &quot;open-alldebrid&quot;]
 **link** | **[]string** | Links to save. | 
 **apikey** | **string** | Deprecated User apikey (Use Bearer Auth in header). | 

### Return type

[**UserLinksSaveGet200Response**](UserLinksSaveGet200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserNotificationClearGet

> UserNotificationClearGet200Response UserNotificationClearGet(ctx).Agent(agent).Code(code).Apikey(apikey).Execute()

This endpoint clears a user notification with its code.



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
	agent := "agent_example" // string | Your software user-agent. (default to "open-alldebrid")
	code := "code_example" // string | Notification code to clear
	apikey := "apikey_example" // string | Deprecated User apikey (Use Bearer Auth in header). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.UserNotificationClearGet(context.Background()).Agent(agent).Code(code).Apikey(apikey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UserNotificationClearGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserNotificationClearGet`: UserNotificationClearGet200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.UserNotificationClearGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUserNotificationClearGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **agent** | **string** | Your software user-agent. | [default to &quot;open-alldebrid&quot;]
 **code** | **string** | Notification code to clear | 
 **apikey** | **string** | Deprecated User apikey (Use Bearer Auth in header). | 

### Return type

[**UserNotificationClearGet200Response**](UserNotificationClearGet200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

