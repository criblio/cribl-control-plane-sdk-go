# Sources
(*Sources*)

## Overview

Actions related to Sources

### Available Operations

* [List](#list) - List all Sources
* [Create](#create) - Create a Source
* [Get](#get) - Get a Source
* [Update](#update) - Update a Source
* [Delete](#delete) - Delete a Source

## List

Get a list of all Sources.

### Example Usage

<!-- UsageSnippet language="go" operationID="listInput" method="get" path="/system/inputs" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Sources.List(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInput != nil {
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

Create a new Source.

### Example Usage

<!-- UsageSnippet language="go" operationID="createInput" method="post" path="/system/inputs" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Sources.Create(ctx, components.CreateInputAppscope(
        components.InputAppscope{
            ID: criblcontrolplanesdkgo.Pointer("appscope-source"),
            Type: components.InputAppscopeTypeAppscope,
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            Connections: []components.InputAppscopeConnection{
                components.InputAppscopeConnection{
                    Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
                    Output: "<value>",
                },
            },
            Pq: &components.InputAppscopePq{
                PqControls: &components.InputAppscopePqControls{},
            },
            Metadata: []components.InputAppscopeMetadatum{
                components.InputAppscopeMetadatum{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            BreakerRulesets: []string{
                "<value 1>",
            },
            Filter: &components.InputAppscopeFilter{
                Allow: []components.Allow{
                    components.Allow{
                        Procname: "<value>",
                        Arg: criblcontrolplanesdkgo.Pointer("<value>"),
                        Config: "<value>",
                    },
                },
                TransportURL: criblcontrolplanesdkgo.Pointer("https://drab-scrap.info/"),
            },
            Persistence: &components.InputAppscopePersistence{},
            Description: criblcontrolplanesdkgo.Pointer("if deserted boohoo red chops excepting know stay bah"),
            Host: criblcontrolplanesdkgo.Pointer("0.0.0.0"),
            Port: criblcontrolplanesdkgo.Pointer[float64](9109),
            TLS: &components.InputAppscopeTLSSettingsServerSide{
                CertificateName: criblcontrolplanesdkgo.Pointer("<value>"),
                PrivKeyPath: criblcontrolplanesdkgo.Pointer("<value>"),
                Passphrase: criblcontrolplanesdkgo.Pointer("<value>"),
                CertPath: criblcontrolplanesdkgo.Pointer("<value>"),
                CaPath: criblcontrolplanesdkgo.Pointer("<value>"),
                MinVersion: components.InputAppscopeMinimumTLSVersionTlSv11.ToPointer(),
                MaxVersion: components.InputAppscopeMaximumTLSVersionTlSv1.ToPointer(),
            },
            UnixSocketPerms: criblcontrolplanesdkgo.Pointer("<value>"),
            TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInput != nil {
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

Get the specified Source.

### Example Usage

<!-- UsageSnippet language="go" operationID="getInputById" method="get" path="/system/inputs/{id}" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Sources.Get(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInput != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | The <code>id</code> of the Source to get.                |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetInputByIDResponse](../../models/operations/getinputbyidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Update

Update the specified Source.</br></br>Provide a complete representation of the Source that you want to update in the request body. This endpoint does not support partial updates. Cribl removes any omitted fields when updating the Source.</br></br>Confirm that the configuration in your request body is correct before sending the request. If the configuration is incorrect, the updated Source might not function as expected.

### Example Usage

<!-- UsageSnippet language="go" operationID="updateInputById" method="patch" path="/system/inputs/{id}" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Sources.Update(ctx, "<id>", components.CreateInputAppscope(
        components.InputAppscope{
            ID: criblcontrolplanesdkgo.Pointer("appscope-source"),
            Type: components.InputAppscopeTypeAppscope,
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
                "<value 2>",
            },
            Connections: []components.InputAppscopeConnection{
                components.InputAppscopeConnection{
                    Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
                    Output: "<value>",
                },
            },
            Pq: &components.InputAppscopePq{
                PqControls: &components.InputAppscopePqControls{},
            },
            Metadata: []components.InputAppscopeMetadatum{
                components.InputAppscopeMetadatum{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            BreakerRulesets: []string{
                "<value 1>",
            },
            Filter: &components.InputAppscopeFilter{
                Allow: []components.Allow{
                    components.Allow{
                        Procname: "<value>",
                        Arg: criblcontrolplanesdkgo.Pointer("<value>"),
                        Config: "<value>",
                    },
                },
                TransportURL: criblcontrolplanesdkgo.Pointer("https://youthful-hammock.net/"),
            },
            Persistence: &components.InputAppscopePersistence{},
            Description: criblcontrolplanesdkgo.Pointer("beyond hidden supposing ghost fictionalize disarm geez"),
            Host: criblcontrolplanesdkgo.Pointer("0.0.0.0"),
            Port: criblcontrolplanesdkgo.Pointer[float64](9109),
            TLS: &components.InputAppscopeTLSSettingsServerSide{
                CertificateName: criblcontrolplanesdkgo.Pointer("<value>"),
                PrivKeyPath: criblcontrolplanesdkgo.Pointer("<value>"),
                Passphrase: criblcontrolplanesdkgo.Pointer("<value>"),
                CertPath: criblcontrolplanesdkgo.Pointer("<value>"),
                CaPath: criblcontrolplanesdkgo.Pointer("<value>"),
                MinVersion: components.InputAppscopeMinimumTLSVersionTlSv11.ToPointer(),
                MaxVersion: components.InputAppscopeMaximumTLSVersionTlSv12.ToPointer(),
            },
            UnixSocketPerms: criblcontrolplanesdkgo.Pointer("<value>"),
            TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInput != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | The <code>id</code> of the Source to update.             |
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

Delete the specified Source.

### Example Usage

<!-- UsageSnippet language="go" operationID="deleteInputById" method="delete" path="/system/inputs/{id}" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Sources.Delete(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInput != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | The <code>id</code> of the Source to delete.             |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeleteInputByIDResponse](../../models/operations/deleteinputbyidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |