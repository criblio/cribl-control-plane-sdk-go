# Sources
(*Sources*)

## Overview

### Available Operations

* [List](#list) - List all Sources
* [Create](#create) - Create a Source
* [Get](#get) - Retrieve a Source
* [Update](#update) - Update a Source
* [Delete](#delete) - Delete a Source
* [CreateHecToken](#createhectoken) - Add an HEC token and optional metadata to a Splunk HEC Source
* [UpdateHecTokenMetadata](#updatehectokenmetadata) - Update metadata for an HEC token for a Splunk HEC Source

## List

Get a list of Source objects

### Example Usage

<!-- UsageSnippet language="go" operationID="listInput" method="get" path="/system/inputs" -->
```go
package main

import(
	"context"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.String(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Sources.List(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.ListInputResponse](../../models/operations/listinputresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Create

Create Source

### Example Usage

<!-- UsageSnippet language="go" operationID="createInput" method="post" path="/system/inputs" -->
```go
package main

import(
	"context"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.String(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Sources.Create(ctx, components.CreateInputInputTCP(
        components.InputTCP{
            ID: criblcontrolplanesdkgo.String("<id>"),
            Type: components.InputTCPTypeTCP.ToPointer(),
            Pipeline: criblcontrolplanesdkgo.String("<value>"),
            Environment: criblcontrolplanesdkgo.String("<value>"),
            Streamtags: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            Connections: []components.InputTCPConnection{
                components.InputTCPConnection{
                    Pipeline: criblcontrolplanesdkgo.String("<value>"),
                    Output: "<value>",
                },
            },
            Pq: &components.InputTCPPq{},
            Port: 301.76,
            TLS: &components.InputTCPTLSSettingsServerSide{
                CertificateName: criblcontrolplanesdkgo.String("<value>"),
                PrivKeyPath: criblcontrolplanesdkgo.String("<value>"),
                Passphrase: criblcontrolplanesdkgo.String("<value>"),
                CertPath: criblcontrolplanesdkgo.String("<value>"),
                CaPath: criblcontrolplanesdkgo.String("<value>"),
                RejectUnauthorized: "<value>",
                CommonNameRegex: "<value>",
                MinVersion: components.InputTCPMinimumTLSVersionTlSv1.ToPointer(),
                MaxVersion: components.InputTCPMaximumTLSVersionTlSv11.ToPointer(),
            },
            Metadata: []components.InputTCPMetadatum{
                components.InputTCPMetadatum{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            BreakerRulesets: []string{
                "<value 1>",
            },
            Preprocess: &components.InputTCPPreprocess{
                Command: criblcontrolplanesdkgo.String("<value>"),
                Args: []string{
                    "<value 1>",
                    "<value 2>",
                    "<value 3>",
                },
            },
            Description: criblcontrolplanesdkgo.String("classic pish supposing misguided carefully fen"),
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `request`                                                | [components.Input](../../models/components/input.md)     | :heavy_check_mark:                                       | The request object to use for the request.               |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.CreateInputResponse](../../models/operations/createinputresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Get

Get Source by ID

### Example Usage

<!-- UsageSnippet language="go" operationID="getInputById" method="get" path="/system/inputs/{id}" -->
```go
package main

import(
	"context"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.String(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Sources.Get(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | Unique ID to GET                                         |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetInputByIDResponse](../../models/operations/getinputbyidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Update

Update Source

### Example Usage

<!-- UsageSnippet language="go" operationID="updateInputById" method="patch" path="/system/inputs/{id}" -->
```go
package main

import(
	"context"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.String(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Sources.Update(ctx, "<id>", components.CreateInputInputKubeEvents(
        components.InputKubeEvents{
            ID: "<id>",
            Type: components.InputKubeEventsTypeKubeEvents,
            Pipeline: criblcontrolplanesdkgo.String("<value>"),
            Environment: criblcontrolplanesdkgo.String("<value>"),
            Streamtags: []string{
                "<value 1>",
                "<value 2>",
            },
            Connections: []components.InputKubeEventsConnection{
                components.InputKubeEventsConnection{
                    Pipeline: criblcontrolplanesdkgo.String("<value>"),
                    Output: "<value>",
                },
            },
            Pq: &components.InputKubeEventsPq{},
            Rules: []components.InputKubeEventsRule{
                components.InputKubeEventsRule{
                    Filter: "<value>",
                    Description: criblcontrolplanesdkgo.String("invite meh corny incidentally down"),
                },
            },
            Metadata: []components.InputKubeEventsMetadatum{
                components.InputKubeEventsMetadatum{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            Description: criblcontrolplanesdkgo.String("gown deployment portray gah mindless carp stabilise"),
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | Unique ID to PATCH                                       |
| `input`                                                  | [components.Input](../../models/components/input.md)     | :heavy_check_mark:                                       | Input object                                             |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.UpdateInputByIDResponse](../../models/operations/updateinputbyidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Delete

Delete Source

### Example Usage

<!-- UsageSnippet language="go" operationID="deleteInputById" method="delete" path="/system/inputs/{id}" -->
```go
package main

import(
	"context"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.String(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Sources.Delete(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | Unique ID to DELETE                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeleteInputByIDResponse](../../models/operations/deleteinputbyidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## CreateHecToken

Add token and optional metadata to an existing HEC Source

### Example Usage

<!-- UsageSnippet language="go" operationID="createInputHecTokenById" method="post" path="/system/inputs/{id}/hectoken" -->
```go
package main

import(
	"context"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.String(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Sources.CreateHecToken(ctx, "<id>", components.AddHecTokenRequest{
        Description: criblcontrolplanesdkgo.String("bah ick stingy"),
        Enabled: criblcontrolplanesdkgo.Bool(false),
        Metadata: []components.AddHecTokenRequestMetadatum{
            components.AddHecTokenRequestMetadatum{
                Name: "<value>",
                Value: "<value>",
            },
        },
        Token: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `id`                                                                           | *string*                                                                       | :heavy_check_mark:                                                             | HEC Source id                                                                  |
| `addHecTokenRequest`                                                           | [components.AddHecTokenRequest](../../models/components/addhectokenrequest.md) | :heavy_check_mark:                                                             | AddHecTokenRequest object                                                      |
| `opts`                                                                         | [][operations.Option](../../models/operations/option.md)                       | :heavy_minus_sign:                                                             | The options for this request.                                                  |

### Response

**[*operations.CreateInputHecTokenByIDResponse](../../models/operations/createinputhectokenbyidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## UpdateHecTokenMetadata

Update token metadata on existing HEC Source

### Example Usage

<!-- UsageSnippet language="go" operationID="updateInputHecTokenByIdAndToken" method="patch" path="/system/inputs/{id}/hectoken/{token}" -->
```go
package main

import(
	"context"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.String(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Sources.UpdateHecTokenMetadata(ctx, "<id>", "<value>", components.UpdateHecTokenRequest{
        Description: criblcontrolplanesdkgo.String("by bleakly fortunately phew barring"),
        Enabled: criblcontrolplanesdkgo.Bool(false),
        Metadata: []components.UpdateHecTokenRequestMetadatum{
            components.UpdateHecTokenRequestMetadatum{
                Name: "<value>",
                Value: "<value>",
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `id`                                                                                 | *string*                                                                             | :heavy_check_mark:                                                                   | HEC Source id                                                                        |
| `token`                                                                              | *string*                                                                             | :heavy_check_mark:                                                                   | token to update                                                                      |
| `updateHecTokenRequest`                                                              | [components.UpdateHecTokenRequest](../../models/components/updatehectokenrequest.md) | :heavy_check_mark:                                                                   | UpdateHecTokenRequest object                                                         |
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.UpdateInputHecTokenByIDAndTokenResponse](../../models/operations/updateinputhectokenbyidandtokenresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |