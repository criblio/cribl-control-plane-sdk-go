# cribl-control-plane-sdk-go

The Cribl Go SDK for the control plane provides operational control over Cribl resources and helps streamline the process of integrating with Cribl.

In addition to the usage examples in this repository, you can adapt the [code examples for common use cases](https://docs.cribl.io/cribl-as-code/code-examples) in the Cribl documentation to use Go instead of Python.

Complementary API reference documentation is available at https://docs.cribl.io/cribl-as-code/api-reference. Product documentation is available at https://docs.cribl.io.

> [!IMPORTANT]
> **Preview Feature**
> The Cribl SDKs are Preview features that are still being developed. We do not recommend using them in a production environment, because the features might not be fully tested or optimized for performance, and related documentation could be incomplete.
>
> Please continue to submit feedback through normal Cribl support channels, but assistance might be limited while the features remain in Preview.

<!-- No Summary [summary] -->

<!-- Start Table of Contents [toc] -->
## Table of Contents
<!-- $toc-max-depth=2 -->
* [cribl-control-plane-sdk-go](#cribl-control-plane-sdk-go)
  * [SDK Installation](#sdk-installation)
  * [SDK Example Usage](#sdk-example-usage)
  * [Authentication](#authentication)
  * [Available Resources and Operations](#available-resources-and-operations)
  * [Json Streaming](#json-streaming)
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
			BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
		}),
	)

	res, err := s.LakeDatasets.Create(ctx, "<id>", components.CriblLakeDataset{
		AcceleratedFields: []string{
			"<value 1>",
			"<value 2>",
		},
		BucketName: criblcontrolplanesdkgo.Pointer("<value>"),
		CacheConnection: &components.CacheConnection{
			AcceleratedFields: []string{
				"<value 1>",
				"<value 2>",
			},
			BackfillStatus:          components.CacheConnectionBackfillStatusPending.ToPointer(),
			CacheRef:                "<value>",
			CreatedAt:               7795.06,
			LakehouseConnectionType: components.LakehouseConnectionTypeCache.ToPointer(),
			MigrationQueryID:        criblcontrolplanesdkgo.Pointer("<id>"),
			RetentionInDays:         1466.58,
		},
		DeletionStartedAt: criblcontrolplanesdkgo.Pointer[float64](8310.58),
		Description:       criblcontrolplanesdkgo.Pointer("pleased toothbrush long brush smooth swiftly rightfully phooey chapel"),
		Format:            components.FormatOptionsCriblLakeDatasetDdss.ToPointer(),
		HTTPDAUsed:        criblcontrolplanesdkgo.Pointer(true),
		ID:                "<id>",
		Metrics: &components.LakeDatasetMetrics{
			CurrentSizeBytes: 6170.04,
			MetricsDate:      "<value>",
		},
		RetentionPeriodInDays: criblcontrolplanesdkgo.Pointer[float64](456.37),
		SearchConfig: &components.LakeDatasetSearchConfig{
			Datatypes: []string{
				"<value 1>",
			},
			Metadata: &components.DatasetMetadata{
				Earliest:           "<value>",
				EnableAcceleration: true,
				FieldList: []string{
					"<value 1>",
					"<value 2>",
				},
				LatestRunInfo: &components.DatasetMetadataRunInfo{
					EarliestScannedTime: criblcontrolplanesdkgo.Pointer[float64](4334.7),
					FinishedAt:          criblcontrolplanesdkgo.Pointer[float64](6811.22),
					LatestScannedTime:   criblcontrolplanesdkgo.Pointer[float64](5303.3),
					ObjectCount:         criblcontrolplanesdkgo.Pointer[float64](9489.04),
				},
				ScanMode: components.ScanModeDetailed,
			},
		},
		StorageLocationID: criblcontrolplanesdkgo.Pointer("<id>"),
		ViewName:          criblcontrolplanesdkgo.Pointer("<value>"),
	})
	if err != nil {
		log.Fatal(err)
	}
	if res.CountedCriblLakeDataset != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->

## Authentication

Except for the `health.get` and `auth.tokens.get` methods, all Cribl SDK requests require you to authenticate with a Bearer token. You must include a valid Bearer token in the configuration when initializing your SDK client. The Bearer token verifies your identity and ensures secure access to the requested resources. The SDK automatically manages the `Authorization` header for subsequent requests once properly authenticated.

For information about Bearer token expiration, see [Token Management](https://docs.cribl.io/cribl-as-code/sdks-auth/#sdks-token-mgmt) in the Cribl as Code documentation.

**Authentication happens once during SDK initialization**. After you initialize the SDK client with authentication as shown in the [authentication examples](#authentication-examples), the SDK automatically handles authentication for all subsequent API calls. You do not need to include authentication parameters in individual API requests. The [SDK Example Usage](#sdk-example-usage) section shows how to initialize the SDK and make API calls, but if you've properly initialized your client as shown in the authentication examples, you only need to make the API method calls themselves without re-initializing.

### Per-Client Security Schemes

This SDK supports the following security schemes globally:

| Name          | Type   | Scheme       | Environment Variable             |
| ------------- | ------ | ------------ | -------------------------------- |
| `BearerAuth`  | http   | HTTP Bearer  | `CRIBLCONTROLPLANE_BEARER_AUTH`  |
| `ClientOauth` | oauth2 | OAuth2 token | `CRIBLCONTROLPLANE_CLIENT_OAUTH` |

To configure authentication on Cribl.Cloud and in hybrid deployments, use the `ClientOauth` security scheme. The SDK uses the OAuth credentials that you provide to obtain a Bearer token and refresh the token within its expiration window using the standard OAuth2 flow.

In on-prem deployments, use the `BearerAuth` security scheme. The SDK uses the username/password credentials that you provide to obtain a Bearer token. Automatically refreshing the Bearer token within its expiration window requires a callback function as shown in the [On-Prem Authentication Example](https://github.com/criblio/cribl-control-plane-sdk-go/blob/main/examples/example-onprem-auth.go).

Set the security scheme when initializing the SDK client instance using one of these optional parameters:
- `WithSecurity`: Use for static security values (OAuth2 credentials that the SDK will manage automatically)
- `WithSecuritySource`: Use for dynamic security with a callback function (automatic Bearer token refresh in on-prem deployments)

The SDK uses the selected scheme by default to authenticate with the API for all operations that support it.

### Authentication Examples

The [Cribl.Cloud and Hybrid Authentication Example](https://github.com/criblio/cribl-control-plane-sdk-go/blob/main/examples/example-cloud-auth.go) demonstrates how to configure authentication on Cribl.Cloud and in hybrid deployments. To obtain the Client ID and Client Secret you'll need to initialize using the `ClientOauth` security schema, follow the [instructions for creating an API Credential](https://docs.cribl.io/cribl-as-code/sdks-auth/#sdks-auth-cloud) in the Cribl as Code documentation.

The [On-Prem Authentication Example](https://github.com/criblio/cribl-control-plane-sdk-go/blob/main/examples/example-onprem-auth.go) demonstrates how to configure authentication in on-prem deployments using your username and password.

<!-- No Authentication [security] -->

<!-- Start Available Resources and Operations [operations] -->
## Available Resources and Operations

<details open>
<summary>Available methods</summary>

### [Auth.Tokens](docs/sdks/tokens/README.md)

* [Get](docs/sdks/tokens/README.md#get) - Log in and fetch an authentication token

### [Collectors](docs/sdks/collectors/README.md)

* [Create](docs/sdks/collectors/README.md#create) - Create a Collector
* [List](docs/sdks/collectors/README.md#list) - List all Collectors
* [Delete](docs/sdks/collectors/README.md#delete) - Delete a Collector
* [Get](docs/sdks/collectors/README.md#get) - Get a Collector
* [Update](docs/sdks/collectors/README.md#update) - Update a Collector

### [DatabaseConnections](docs/sdks/databaseconnections/README.md)

* [Create](docs/sdks/databaseconnections/README.md#create) - Create Database Connection
* [Delete](docs/sdks/databaseconnections/README.md#delete) - Delete a Database Connection
* [Get](docs/sdks/databaseconnections/README.md#get) - Get a Database Connection
* [Update](docs/sdks/databaseconnections/README.md#update) - Update a Database Connection

### [Destinations](docs/sdks/destinations/README.md)

* [List](docs/sdks/destinations/README.md#list) - List all Destinations
* [Create](docs/sdks/destinations/README.md#create) - Create a Destination
* [Get](docs/sdks/destinations/README.md#get) - Get a Destination
* [Update](docs/sdks/destinations/README.md#update) - Update a Destination
* [Delete](docs/sdks/destinations/README.md#delete) - Delete a Destination

#### [Destinations.Pq](docs/sdks/pq/README.md)

* [Clear](docs/sdks/pq/README.md#clear) - Clear the persistent queue for a Destination
* [Get](docs/sdks/pq/README.md#get) - Get information about the latest job to clear the persistent queue for a Destination

#### [Destinations.Samples](docs/sdks/samples/README.md)

* [Get](docs/sdks/samples/README.md#get) - Get sample event data for a Destination
* [Create](docs/sdks/samples/README.md#create) - Send sample event data to a Destination

### [Functions](docs/sdks/functions/README.md)

* [Get](docs/sdks/functions/README.md#get) - Get a Function
* [List](docs/sdks/functions/README.md#list) - List all Functions

### [Groups](docs/sdks/groups/README.md)

* [List](docs/sdks/groups/README.md#list) - List all Worker Groups, Outpost Groups, or Edge Fleets for the specified Cribl product
* [Create](docs/sdks/groups/README.md#create) - Create a Worker Group, Outpost Group, or Edge Fleet for the specified Cribl product
* [Get](docs/sdks/groups/README.md#get) - Get a Worker Group, Outpost Group, or Edge Fleet
* [Update](docs/sdks/groups/README.md#update) - Update a Worker Group, Outpost Group, or Edge Fleet
* [Delete](docs/sdks/groups/README.md#delete) - Delete a Worker Group, Outpost Group, or Edge Fleet
* [Deploy](docs/sdks/groups/README.md#deploy) - Deploy commits to a Worker Group, Outpost Group, or Edge Fleet

#### [Groups.Acl](docs/sdks/acl/README.md)

* [Get](docs/sdks/acl/README.md#get) - Get the Access Control List for a Worker Group, Outpost Group, or Edge Fleet

##### [Groups.Acl.Teams](docs/sdks/teams/README.md)

* [Get](docs/sdks/teams/README.md#get) - Get the Access Control List for teams with permissions on a Worker Group, Outpost Group, or Edge Fleet for the specified Cribl product

#### [Groups.Configs.Versions](docs/sdks/configsversions/README.md)

* [Get](docs/sdks/configsversions/README.md#get) - Get the configuration version for a Worker Group, Outpost Group, or Edge Fleet

### [Health](docs/sdks/health/README.md)

* [Get](docs/sdks/health/README.md#get) - Retrieve health status of the server

### [LakeDatasets](docs/sdks/lakedatasets/README.md)

* [Create](docs/sdks/lakedatasets/README.md#create) - Create a Lake Dataset
* [List](docs/sdks/lakedatasets/README.md#list) - List all Lake Datasets
* [Delete](docs/sdks/lakedatasets/README.md#delete) - Delete a Lake Dataset
* [Get](docs/sdks/lakedatasets/README.md#get) - Get a Lake Dataset
* [Update](docs/sdks/lakedatasets/README.md#update) - Update a Lake Dataset

### [Nodes](docs/sdks/nodes/README.md)

* [~~Count~~](docs/sdks/nodes/README.md#count) - Get a count of Worker and Edge Nodes :warning: **Deprecated**
* [~~List~~](docs/sdks/nodes/README.md#list) - Get detailed metadata for Worker and Edge Nodes :warning: **Deprecated**

#### [Nodes.Summaries](docs/sdks/summaries/README.md)

* [Get](docs/sdks/summaries/README.md#get) - Get a summary of the Distributed deployment

### [Packs](docs/sdks/packs/README.md)

* [Install](docs/sdks/packs/README.md#install) - Install a Pack
* [List](docs/sdks/packs/README.md#list) - List all Packs
* [Upload](docs/sdks/packs/README.md#upload) - Upload a Pack file
* [Delete](docs/sdks/packs/README.md#delete) - Uninstall a Pack
* [Get](docs/sdks/packs/README.md#get) - Get a Pack
* [Update](docs/sdks/packs/README.md#update) - Upgrade a Pack

### [Pipelines](docs/sdks/pipelines/README.md)

* [Create](docs/sdks/pipelines/README.md#create) - Create a Pipeline
* [List](docs/sdks/pipelines/README.md#list) - List all Pipelines
* [Delete](docs/sdks/pipelines/README.md#delete) - Delete a Pipeline
* [Get](docs/sdks/pipelines/README.md#get) - Get a Pipeline
* [Update](docs/sdks/pipelines/README.md#update) - Update a Pipeline

### [Routes](docs/sdks/routes/README.md)

* [Get](docs/sdks/routes/README.md#get) - Get a Routing table
* [Update](docs/sdks/routes/README.md#update) - Update a Route
* [List](docs/sdks/routes/README.md#list) - List all Routes
* [Append](docs/sdks/routes/README.md#append) - Add a Route to the end of the Routing table

### [Sources](docs/sdks/sources/README.md)

* [List](docs/sdks/sources/README.md#list) - List all Sources
* [Create](docs/sdks/sources/README.md#create) - Create a Source
* [Get](docs/sdks/sources/README.md#get) - Get a Source
* [Update](docs/sdks/sources/README.md#update) - Update a Source
* [Delete](docs/sdks/sources/README.md#delete) - Delete a Source

#### [Sources.HecTokens](docs/sdks/hectokens/README.md)

* [Create](docs/sdks/hectokens/README.md#create) - Add an HEC token and optional metadata to a Splunk HEC Source
* [Update](docs/sdks/hectokens/README.md#update) - Update metadata for an HEC token for a Splunk HEC Source

### [System.Captures](docs/sdks/captures/README.md)

* [Create](docs/sdks/captures/README.md#create) - Capture live incoming data

### [System.Settings.Cribl](docs/sdks/cribl/README.md)

* [List](docs/sdks/cribl/README.md#list) - Get Cribl system settings
* [Update](docs/sdks/cribl/README.md#update) - Update Cribl system settings

### [Versions.Branches](docs/sdks/branches/README.md)

* [List](docs/sdks/branches/README.md#list) - List all branches in the Git repository used for Cribl configuration
* [Get](docs/sdks/branches/README.md#get) - Get the name of the Git branch that the Cribl configuration is checked out to

### [Versions.Commits](docs/sdks/commits/README.md)

* [Create](docs/sdks/commits/README.md#create) - Create a new commit for pending changes to the Cribl configuration
* [Diff](docs/sdks/commits/README.md#diff) - Get the diff for a commit
* [List](docs/sdks/commits/README.md#list) - List the commit history
* [Push](docs/sdks/commits/README.md#push) - Push local commits to the remote repository
* [Revert](docs/sdks/commits/README.md#revert) - Revert a commit in the local repository
* [Get](docs/sdks/commits/README.md#get) - Get the diff and log message for a commit
* [Undo](docs/sdks/commits/README.md#undo) - Discard uncommitted (staged) changes

#### [Versions.Commits.Files](docs/sdks/files/README.md)

* [Count](docs/sdks/files/README.md#count) - Get a count of files that changed since a commit
* [List](docs/sdks/files/README.md#list) - Get the names and statuses of files that changed since a commit

### [Versions.Configs](docs/sdks/versionsconfigs/README.md)

* [Get](docs/sdks/versionsconfigs/README.md#get) - Get the configuration and status for the Git integration

### [Versions.Statuses](docs/sdks/statuses/README.md)

* [Get](docs/sdks/statuses/README.md#get) - Get the status of the current working tree

</details>
<!-- End Available Resources and Operations [operations] -->

<!-- Start Json Streaming [jsonl] -->
## Json Streaming

Json Streaming ([jsonl][jsonl-format] / [x-ndjson][x-ndjson]) content type can be used to stream content from certain operations. These operations expose the stream that can be consumed using a `for` loop in Go. The loop will terminate when the server no longer has any events to send and closes the underlying connection.

Here's an example of consuming a JSONL stream:

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
			BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
		}),
	)

	res, err := s.System.Captures.Create(ctx, components.CaptureParams{
		Duration:        5,
		Filter:          "sourcetype===\"pan:traffic\"",
		Level:           components.CaptureLevelBeforePreProcessingPipeline,
		MaxEvents:       100,
		StepDuration:    criblcontrolplanesdkgo.Pointer[int64](994184),
		WorkerID:        criblcontrolplanesdkgo.Pointer("<id>"),
		WorkerThreshold: criblcontrolplanesdkgo.Pointer[int64](771620),
	})
	if err != nil {
		log.Fatal(err)
	}
	if res.CapturedEvent != nil {
		for res.CapturedEvent.Next() {
			event, _ := res.CapturedEvent.Value()
			log.Print(event)
			// Handle the event
		}
	}
}

```

[jsonl-format]: https://jsonlines.org/
[x-ndjson]: https://github.com/ndjson/ndjson-spec
<!-- End Json Streaming [jsonl] -->

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
			BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
		}),
	)

	res, err := s.LakeDatasets.Create(ctx, "<id>", components.CriblLakeDataset{
		AcceleratedFields: []string{
			"<value 1>",
			"<value 2>",
		},
		BucketName: criblcontrolplanesdkgo.Pointer("<value>"),
		CacheConnection: &components.CacheConnection{
			AcceleratedFields: []string{
				"<value 1>",
				"<value 2>",
			},
			BackfillStatus:          components.CacheConnectionBackfillStatusPending.ToPointer(),
			CacheRef:                "<value>",
			CreatedAt:               7795.06,
			LakehouseConnectionType: components.LakehouseConnectionTypeCache.ToPointer(),
			MigrationQueryID:        criblcontrolplanesdkgo.Pointer("<id>"),
			RetentionInDays:         1466.58,
		},
		DeletionStartedAt: criblcontrolplanesdkgo.Pointer[float64](8310.58),
		Description:       criblcontrolplanesdkgo.Pointer("pleased toothbrush long brush smooth swiftly rightfully phooey chapel"),
		Format:            components.FormatOptionsCriblLakeDatasetDdss.ToPointer(),
		HTTPDAUsed:        criblcontrolplanesdkgo.Pointer(true),
		ID:                "<id>",
		Metrics: &components.LakeDatasetMetrics{
			CurrentSizeBytes: 6170.04,
			MetricsDate:      "<value>",
		},
		RetentionPeriodInDays: criblcontrolplanesdkgo.Pointer[float64](456.37),
		SearchConfig: &components.LakeDatasetSearchConfig{
			Datatypes: []string{
				"<value 1>",
			},
			Metadata: &components.DatasetMetadata{
				Earliest:           "<value>",
				EnableAcceleration: true,
				FieldList: []string{
					"<value 1>",
					"<value 2>",
				},
				LatestRunInfo: &components.DatasetMetadataRunInfo{
					EarliestScannedTime: criblcontrolplanesdkgo.Pointer[float64](4334.7),
					FinishedAt:          criblcontrolplanesdkgo.Pointer[float64](6811.22),
					LatestScannedTime:   criblcontrolplanesdkgo.Pointer[float64](5303.3),
					ObjectCount:         criblcontrolplanesdkgo.Pointer[float64](9489.04),
				},
				ScanMode: components.ScanModeDetailed,
			},
		},
		StorageLocationID: criblcontrolplanesdkgo.Pointer("<id>"),
		ViewName:          criblcontrolplanesdkgo.Pointer("<value>"),
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
	if res.CountedCriblLakeDataset != nil {
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
			BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
		}),
	)

	res, err := s.LakeDatasets.Create(ctx, "<id>", components.CriblLakeDataset{
		AcceleratedFields: []string{
			"<value 1>",
			"<value 2>",
		},
		BucketName: criblcontrolplanesdkgo.Pointer("<value>"),
		CacheConnection: &components.CacheConnection{
			AcceleratedFields: []string{
				"<value 1>",
				"<value 2>",
			},
			BackfillStatus:          components.CacheConnectionBackfillStatusPending.ToPointer(),
			CacheRef:                "<value>",
			CreatedAt:               7795.06,
			LakehouseConnectionType: components.LakehouseConnectionTypeCache.ToPointer(),
			MigrationQueryID:        criblcontrolplanesdkgo.Pointer("<id>"),
			RetentionInDays:         1466.58,
		},
		DeletionStartedAt: criblcontrolplanesdkgo.Pointer[float64](8310.58),
		Description:       criblcontrolplanesdkgo.Pointer("pleased toothbrush long brush smooth swiftly rightfully phooey chapel"),
		Format:            components.FormatOptionsCriblLakeDatasetDdss.ToPointer(),
		HTTPDAUsed:        criblcontrolplanesdkgo.Pointer(true),
		ID:                "<id>",
		Metrics: &components.LakeDatasetMetrics{
			CurrentSizeBytes: 6170.04,
			MetricsDate:      "<value>",
		},
		RetentionPeriodInDays: criblcontrolplanesdkgo.Pointer[float64](456.37),
		SearchConfig: &components.LakeDatasetSearchConfig{
			Datatypes: []string{
				"<value 1>",
			},
			Metadata: &components.DatasetMetadata{
				Earliest:           "<value>",
				EnableAcceleration: true,
				FieldList: []string{
					"<value 1>",
					"<value 2>",
				},
				LatestRunInfo: &components.DatasetMetadataRunInfo{
					EarliestScannedTime: criblcontrolplanesdkgo.Pointer[float64](4334.7),
					FinishedAt:          criblcontrolplanesdkgo.Pointer[float64](6811.22),
					LatestScannedTime:   criblcontrolplanesdkgo.Pointer[float64](5303.3),
					ObjectCount:         criblcontrolplanesdkgo.Pointer[float64](9489.04),
				},
				ScanMode: components.ScanModeDetailed,
			},
		},
		StorageLocationID: criblcontrolplanesdkgo.Pointer("<id>"),
		ViewName:          criblcontrolplanesdkgo.Pointer("<value>"),
	})
	if err != nil {
		log.Fatal(err)
	}
	if res.CountedCriblLakeDataset != nil {
		// handle response
	}
}

```
<!-- End Retries [retries] -->

<!-- Start Error Handling [errors] -->
## Error Handling

Handling errors in this SDK should largely match your expectations. All operations return a response object or an error, they will never return both.

By Default, an API error will return `apierrors.APIError`. When custom error responses are specified for an operation, the SDK may also return their associated error. You can refer to respective *Errors* tables in SDK docs for more details on possible error types for each operation.

For example, the `Create` function may return the following errors:

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
			BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
		}),
	)

	res, err := s.LakeDatasets.Create(ctx, "<id>", components.CriblLakeDataset{
		AcceleratedFields: []string{
			"<value 1>",
			"<value 2>",
		},
		BucketName: criblcontrolplanesdkgo.Pointer("<value>"),
		CacheConnection: &components.CacheConnection{
			AcceleratedFields: []string{
				"<value 1>",
				"<value 2>",
			},
			BackfillStatus:          components.CacheConnectionBackfillStatusPending.ToPointer(),
			CacheRef:                "<value>",
			CreatedAt:               7795.06,
			LakehouseConnectionType: components.LakehouseConnectionTypeCache.ToPointer(),
			MigrationQueryID:        criblcontrolplanesdkgo.Pointer("<id>"),
			RetentionInDays:         1466.58,
		},
		DeletionStartedAt: criblcontrolplanesdkgo.Pointer[float64](8310.58),
		Description:       criblcontrolplanesdkgo.Pointer("pleased toothbrush long brush smooth swiftly rightfully phooey chapel"),
		Format:            components.FormatOptionsCriblLakeDatasetDdss.ToPointer(),
		HTTPDAUsed:        criblcontrolplanesdkgo.Pointer(true),
		ID:                "<id>",
		Metrics: &components.LakeDatasetMetrics{
			CurrentSizeBytes: 6170.04,
			MetricsDate:      "<value>",
		},
		RetentionPeriodInDays: criblcontrolplanesdkgo.Pointer[float64](456.37),
		SearchConfig: &components.LakeDatasetSearchConfig{
			Datatypes: []string{
				"<value 1>",
			},
			Metadata: &components.DatasetMetadata{
				Earliest:           "<value>",
				EnableAcceleration: true,
				FieldList: []string{
					"<value 1>",
					"<value 2>",
				},
				LatestRunInfo: &components.DatasetMetadataRunInfo{
					EarliestScannedTime: criblcontrolplanesdkgo.Pointer[float64](4334.7),
					FinishedAt:          criblcontrolplanesdkgo.Pointer[float64](6811.22),
					LatestScannedTime:   criblcontrolplanesdkgo.Pointer[float64](5303.3),
					ObjectCount:         criblcontrolplanesdkgo.Pointer[float64](9489.04),
				},
				ScanMode: components.ScanModeDetailed,
			},
		},
		StorageLocationID: criblcontrolplanesdkgo.Pointer("<id>"),
		ViewName:          criblcontrolplanesdkgo.Pointer("<value>"),
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

	"github.com/criblio/cribl-control-plane-sdk-go"
)

var (
	httpClient = &http.Client{Timeout: 30 * time.Second}
	sdkClient  = criblcontrolplanesdkgo.New(criblcontrolplanesdkgo.WithClient(httpClient))
)
```

This can be a convenient way to configure timeouts, cookies, proxies, custom headers, and other low-level configuration.
<!-- End Custom HTTP Client [http-client] -->

<!-- Placeholder for Future Speakeasy SDK Sections -->
