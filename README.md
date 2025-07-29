# cribl-control-plane-sdk-go
<!-- Start Summary [summary] -->
## Summary

Cribl API Reference: This API Reference lists available REST endpoints, along with their supported operations for accessing, creating, updating, or deleting resources. See our complementary product documentation at [docs.cribl.io](http://docs.cribl.io).
<!-- End Summary [summary] -->

<!-- Start Table of Contents [toc] -->
## Table of Contents
<!-- $toc-max-depth=2 -->
* [cribl-control-plane-sdk-go](#cribl-control-plane-sdk-go)
  * [SDK Installation](#sdk-installation)
  * [SDK Example Usage](#sdk-example-usage)
  * [Authentication](#authentication)
  * [Available Resources and Operations](#available-resources-and-operations)
  * [Retries](#retries)
  * [Error Handling](#error-handling)
  * [Custom HTTP Client](#custom-http-client)

<!-- End Table of Contents [toc] -->

<!-- Start SDK Installation [installation] -->
## SDK Installation

To add the SDK as a dependency to your project:
```bash
go get github.com/criblio/cribl-control-plane-sdk-go
```
<!-- End SDK Installation [installation] -->

<!-- Start SDK Example Usage [usage] -->
## SDK Example Usage

### Example

```go
package main

import (
	"context"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	s := criblcontrolplanesdkgo.New(
		"https://api.example.com",
		criblcontrolplanesdkgo.WithSecurity(components.Security{
			BearerAuth: criblcontrolplanesdkgo.String(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
		}),
	)

	res, err := s.Sources.ListSource(ctx)
	if err != nil {
		log.Fatal(err)
	}
	if res.Object != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->

<!-- Start Authentication [security] -->
## Authentication

### Per-Client Security Schemes

This SDK supports the following security schemes globally:

| Name          | Type   | Scheme       | Environment Variable             |
| ------------- | ------ | ------------ | -------------------------------- |
| `BearerAuth`  | http   | HTTP Bearer  | `CRIBLCONTROLPLANE_BEARER_AUTH`  |
| `ClientOauth` | oauth2 | OAuth2 token | `CRIBLCONTROLPLANE_CLIENT_OAUTH` |

You can set the security parameters through the `WithSecurity` option when initializing the SDK client instance. The selected scheme will be used by default to authenticate with the API for all operations that support it. For example:
```go
package main

import (
	"context"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	s := criblcontrolplanesdkgo.New(
		"https://api.example.com",
		criblcontrolplanesdkgo.WithSecurity(components.Security{
			BearerAuth: criblcontrolplanesdkgo.String(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
		}),
	)

	res, err := s.Sources.ListSource(ctx)
	if err != nil {
		log.Fatal(err)
	}
	if res.Object != nil {
		// handle response
	}
}

```
<!-- End Authentication [security] -->

<!-- Start Available Resources and Operations [operations] -->
## Available Resources and Operations

<details open>
<summary>Available methods</summary>

### [Auth](docs/sdks/auth/README.md)

* [Login](docs/sdks/auth/README.md#login) - Log in and obtain Auth token


### [Destinations](docs/sdks/destinations/README.md)

* [ListDestination](docs/sdks/destinations/README.md#listdestination) - Get a list of Destination objects
* [CreateDestination](docs/sdks/destinations/README.md#createdestination) - Create Destination
* [GetDestinationByID](docs/sdks/destinations/README.md#getdestinationbyid) - Get Destination by ID
* [UpdateDestinationByID](docs/sdks/destinations/README.md#updatedestinationbyid) - Update Destination
* [DeleteDestinationByID](docs/sdks/destinations/README.md#deletedestinationbyid) - Delete Destination
* [DeleteDestinationPqByID](docs/sdks/destinations/README.md#deletedestinationpqbyid) - Clears destination persistent queue
* [GetDestinationPqByID](docs/sdks/destinations/README.md#getdestinationpqbyid) - Retrieves status of latest clear PQ job for a destination
* [GetDestinationSamplesByID](docs/sdks/destinations/README.md#getdestinationsamplesbyid) - Retrieve samples data for the specified destination. Used to get sample data for the test action.
* [CreateDestinationTestByID](docs/sdks/destinations/README.md#createdestinationtestbyid) - Send sample data to a destination to validate configuration or test connectivity

### [Health](docs/sdks/health/README.md)

* [GetHealthInfo](docs/sdks/health/README.md#gethealthinfo) - Provides health info for REST server

### [Pipelines](docs/sdks/pipelines/README.md)

* [ListPipeline](docs/sdks/pipelines/README.md#listpipeline) - Get a list of Pipeline objects
* [CreatePipeline](docs/sdks/pipelines/README.md#createpipeline) - Create Pipeline
* [GetPipelineByID](docs/sdks/pipelines/README.md#getpipelinebyid) - Get Pipeline by ID
* [UpdatePipelineByID](docs/sdks/pipelines/README.md#updatepipelinebyid) - Update Pipeline
* [DeletePipelineByID](docs/sdks/pipelines/README.md#deletepipelinebyid) - Delete Pipeline

### [Routes](docs/sdks/routes/README.md)

* [ListRoutes](docs/sdks/routes/README.md#listroutes) - Get a list of Routes objects
* [GetRoutesByID](docs/sdks/routes/README.md#getroutesbyid) - Get Routes by ID
* [UpdateRoutesByID](docs/sdks/routes/README.md#updateroutesbyid) - Update Routes
* [CreateRoutesAppendByID](docs/sdks/routes/README.md#createroutesappendbyid) - Appends routes to the end of the routing table

### [Sources](docs/sdks/sources/README.md)

* [ListSource](docs/sdks/sources/README.md#listsource) - Get a list of Source objects
* [CreateSource](docs/sdks/sources/README.md#createsource) - Create Source
* [GetSourceByID](docs/sdks/sources/README.md#getsourcebyid) - Get Source by ID
* [UpdateSourceByID](docs/sdks/sources/README.md#updatesourcebyid) - Update Source
* [DeleteSourceByID](docs/sdks/sources/README.md#deletesourcebyid) - Delete Source
* [CreateSourceHecTokenByID](docs/sdks/sources/README.md#createsourcehectokenbyid) - Add token and optional metadata to an existing HEC Source
* [UpdateSourceHecTokenByIDAndToken](docs/sdks/sources/README.md#updatesourcehectokenbyidandtoken) - Update token metadata on existing HEC Source

### [Versioning](docs/sdks/versioning/README.md)

* [GetVersionBranch](docs/sdks/versioning/README.md#getversionbranch) - get the list of branches
* [CreateVersionCommit](docs/sdks/versioning/README.md#createversioncommit) - create a new commit containing the current configs the given log message describing the changes.
* [GetVersionCount](docs/sdks/versioning/README.md#getversioncount) - get the count of files of changed
* [GetVersionCurrentBranch](docs/sdks/versioning/README.md#getversioncurrentbranch) - returns git branch that the config is checked out to, if any
* [GetVersionDiff](docs/sdks/versioning/README.md#getversiondiff) - get the textual diff for given commit
* [GetVersionFiles](docs/sdks/versioning/README.md#getversionfiles) - get the files changed
* [GetVersionInfo](docs/sdks/versioning/README.md#getversioninfo) - Get info about versioning availability
* [CreateVersionPush](docs/sdks/versioning/README.md#createversionpush) - push the current configs to the remote repository.
* [CreateVersionRevert](docs/sdks/versioning/README.md#createversionrevert) - revert a commit
* [GetVersionShow](docs/sdks/versioning/README.md#getversionshow) - get the log message and textual diff for given commit
* [GetVersionStatus](docs/sdks/versioning/README.md#getversionstatus) - get the the working tree status
* [CreateVersionSync](docs/sdks/versioning/README.md#createversionsync) - syncs with remote repo via POST requests
* [CreateVersionUndo](docs/sdks/versioning/README.md#createversionundo) - undo the last commit

</details>
<!-- End Available Resources and Operations [operations] -->

<!-- Start Retries [retries] -->
## Retries

Some of the endpoints in this SDK support retries. If you use the SDK without any configuration, it will fall back to the default retry strategy provided by the API. However, the default retry strategy can be overridden on a per-operation basis, or across the entire SDK.

To change the default retry strategy for a single API call, simply provide a `retry.Config` object to the call by using the `WithRetries` option:
```go
package main

import (
	"context"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	"github.com/criblio/cribl-control-plane-sdk-go/retry"
	"log"
	"models/operations"
	"os"
)

func main() {
	ctx := context.Background()

	s := criblcontrolplanesdkgo.New(
		"https://api.example.com",
		criblcontrolplanesdkgo.WithSecurity(components.Security{
			BearerAuth: criblcontrolplanesdkgo.String(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
		}),
	)

	res, err := s.Sources.ListSource(ctx, operations.WithRetries(
		retry.Config{
			Strategy: "backoff",
			Backoff: &retry.BackoffStrategy{
				InitialInterval: 1,
				MaxInterval:     50,
				Exponent:        1.1,
				MaxElapsedTime:  100,
			},
			RetryConnectionErrors: false,
		}))
	if err != nil {
		log.Fatal(err)
	}
	if res.Object != nil {
		// handle response
	}
}

```

If you'd like to override the default retry strategy for all operations that support retries, you can use the `WithRetryConfig` option at SDK initialization:
```go
package main

import (
	"context"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	"github.com/criblio/cribl-control-plane-sdk-go/retry"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	s := criblcontrolplanesdkgo.New(
		"https://api.example.com",
		criblcontrolplanesdkgo.WithRetryConfig(
			retry.Config{
				Strategy: "backoff",
				Backoff: &retry.BackoffStrategy{
					InitialInterval: 1,
					MaxInterval:     50,
					Exponent:        1.1,
					MaxElapsedTime:  100,
				},
				RetryConnectionErrors: false,
			}),
		criblcontrolplanesdkgo.WithSecurity(components.Security{
			BearerAuth: criblcontrolplanesdkgo.String(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
		}),
	)

	res, err := s.Sources.ListSource(ctx)
	if err != nil {
		log.Fatal(err)
	}
	if res.Object != nil {
		// handle response
	}
}

```
<!-- End Retries [retries] -->

<!-- Start Error Handling [errors] -->
## Error Handling

Handling errors in this SDK should largely match your expectations. All operations return a response object or an error, they will never return both.

By Default, an API error will return `apierrors.APIError`. When custom error responses are specified for an operation, the SDK may also return their associated error. You can refer to respective *Errors* tables in SDK docs for more details on possible error types for each operation.

For example, the `ListSource` function may return the following errors:

| Error Type         | Status Code | Content Type     |
| ------------------ | ----------- | ---------------- |
| apierrors.Error    | 500         | application/json |
| apierrors.APIError | 4XX, 5XX    | \*/\*            |

### Example

```go
package main

import (
	"context"
	"errors"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"github.com/criblio/cribl-control-plane-sdk-go/models/apierrors"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	s := criblcontrolplanesdkgo.New(
		"https://api.example.com",
		criblcontrolplanesdkgo.WithSecurity(components.Security{
			BearerAuth: criblcontrolplanesdkgo.String(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
		}),
	)

	res, err := s.Sources.ListSource(ctx)
	if err != nil {

		var e *apierrors.Error
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *apierrors.APIError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}
	}
}

```
<!-- End Error Handling [errors] -->

<!-- Start Custom HTTP Client [http-client] -->
## Custom HTTP Client

The Go SDK makes API calls that wrap an internal HTTP client. The requirements for the HTTP client are very simple. It must match this interface:

```go
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}
```

The built-in `net/http` client satisfies this interface and a default client based on the built-in is provided by default. To replace this default with a client of your own, you can implement this interface yourself or provide your own client configured as desired. Here's a simple example, which adds a client with a 30 second timeout.

```go
import (
	"net/http"
	"time"
	"github.com/myorg/your-go-sdk"
)

var (
	httpClient = &http.Client{Timeout: 30 * time.Second}
	sdkClient  = sdk.New(sdk.WithClient(httpClient))
)
```

This can be a convenient way to configure timeouts, cookies, proxies, custom headers, and other low-level configuration.
<!-- End Custom HTTP Client [http-client] -->

<!-- Placeholder for Future Speakeasy SDK Sections -->
