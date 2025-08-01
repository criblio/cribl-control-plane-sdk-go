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

	res, err := s.Lake.CreateCriblLakeDatasetByLakeID(ctx, "<id>", components.CriblLakeDataset{
		ID: "<id>",
	})
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

	res, err := s.Lake.CreateCriblLakeDatasetByLakeID(ctx, "<id>", components.CriblLakeDataset{
		ID: "<id>",
	})
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

### [Distributed](docs/sdks/distributed/README.md)

* [GetSummary](docs/sdks/distributed/README.md#getsummary) - Get summary of Distributed deployment

### [Groups](docs/sdks/groups/README.md)

* [GetGroupsConfigVersionByID](docs/sdks/groups/README.md#getgroupsconfigversionbyid) - Get effective bundle version for given Group
* [CreateProductsGroupsByProduct](docs/sdks/groups/README.md#createproductsgroupsbyproduct) - Create a Fleet or Worker Group
* [GetProductsGroupsByProduct](docs/sdks/groups/README.md#getproductsgroupsbyproduct) - Get a list of ConfigGroup objects
* [DeleteGroupsByID](docs/sdks/groups/README.md#deletegroupsbyid) - Delete a Fleet or Worker Group
* [GetGroupsByID](docs/sdks/groups/README.md#getgroupsbyid) - Get a specific ConfigGroup object
* [UpdateGroupsByID](docs/sdks/groups/README.md#updategroupsbyid) - Update a Fleet or Worker Group
* [UpdateGroupsDeployByID](docs/sdks/groups/README.md#updategroupsdeploybyid) - Deploy commits for a Fleet or Worker Group
* [GetGroupsACLByID](docs/sdks/groups/README.md#getgroupsaclbyid) - ACL of members with permissions for resources in this Group

### [Health](docs/sdks/health/README.md)

* [GetHealthInfo](docs/sdks/health/README.md#gethealthinfo) - Provides health info for REST server

### [Lake](docs/sdks/lake/README.md)

* [CreateCriblLakeDatasetByLakeID](docs/sdks/lake/README.md#createcribllakedatasetbylakeid) - Create a Dataset in the specified Lake
* [GetCriblLakeDatasetByLakeID](docs/sdks/lake/README.md#getcribllakedatasetbylakeid) - Get the list of Dataset contained in the specified Lake
* [DeleteCriblLakeDatasetByLakeIDAndID](docs/sdks/lake/README.md#deletecribllakedatasetbylakeidandid) - Delete a Dataset in the specified Lake
* [GetCriblLakeDatasetByLakeIDAndID](docs/sdks/lake/README.md#getcribllakedatasetbylakeidandid) - Get a Dataset in the specified Lake
* [UpdateCriblLakeDatasetByLakeIDAndID](docs/sdks/lake/README.md#updatecribllakedatasetbylakeidandid) - Update a Dataset in the specified Lake

### [Packs](docs/sdks/packs/README.md)

* [CreatePacks](docs/sdks/packs/README.md#createpacks) - Install Pack
* [GetPacks](docs/sdks/packs/README.md#getpacks) - Get info on packs
* [UpdatePacks](docs/sdks/packs/README.md#updatepacks) - Upload Pack
* [DeletePacksByID](docs/sdks/packs/README.md#deletepacksbyid) - Uninstall Pack from the system
* [UpdatePacksByID](docs/sdks/packs/README.md#updatepacksbyid) - Upgrade Pack

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

### [Teams](docs/sdks/teams/README.md)

* [GetProductsGroupsACLTeamsByProductAndID](docs/sdks/teams/README.md#getproductsgroupsaclteamsbyproductandid) - ACL of team with permissions for resources in this Group

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

### [Workers](docs/sdks/workers/README.md)

* [GetSummaryWorkers](docs/sdks/workers/README.md#getsummaryworkers) - get worker and edge nodes count
* [GetWorkers](docs/sdks/workers/README.md#getworkers) - get worker and edge nodes
* [UpdateWorkersRestart](docs/sdks/workers/README.md#updateworkersrestart) - restarts worker nodes

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

	res, err := s.Lake.CreateCriblLakeDatasetByLakeID(ctx, "<id>", components.CriblLakeDataset{
		ID: "<id>",
	}, operations.WithRetries(
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

	res, err := s.Lake.CreateCriblLakeDatasetByLakeID(ctx, "<id>", components.CriblLakeDataset{
		ID: "<id>",
	})
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

For example, the `CreateCriblLakeDatasetByLakeID` function may return the following errors:

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

	res, err := s.Lake.CreateCriblLakeDatasetByLakeID(ctx, "<id>", components.CriblLakeDataset{
		ID: "<id>",
	})
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
