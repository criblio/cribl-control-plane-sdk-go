# Mappings
(*Groups.Mappings*)

## Overview

### Available Operations

* [Activate](#activate) - Set a Mapping Ruleset as the active configuration for the specified Cribl product
* [Create](#create) - Create a new Mapping Ruleset for the specified Cribl product
* [List](#list) - List all Mapping Rulesets for the specified Cribl product
* [Delete](#delete) - Delete the specified Mapping Ruleset from the Worker Group or Edge Fleet
* [Get](#get) - Retrieve a Specific Mapping Ruleset
* [Update](#update) - Update an existing Mapping Ruleset for a Worker Group or Edge Fleet

## Activate

Set a specific Mapping Ruleset as the currently active configuration for the specified Cribl product.

### Example Usage

<!-- UsageSnippet language="go" operationID="createAdminProductsMappingsActivateByProduct" method="post" path="/admin/products/{product}/mappings/activate" -->
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

    res, err := s.Groups.Mappings.Activate(ctx, components.ProductsCoreStream, components.RulesetID{
        ID: "default",
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

| Parameter                                                          | Type                                                               | Required                                                           | Description                                                        |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `ctx`                                                              | [context.Context](https://pkg.go.dev/context#Context)              | :heavy_check_mark:                                                 | The context to use for the request.                                |
| `product`                                                          | [components.ProductsCore](../../models/components/productscore.md) | :heavy_check_mark:                                                 | Name of the Cribl product to activate the Mapping Ruleset for      |
| `rulesetID`                                                        | [components.RulesetID](../../models/components/rulesetid.md)       | :heavy_check_mark:                                                 | RulesetId object                                                   |
| `opts`                                                             | [][operations.Option](../../models/operations/option.md)           | :heavy_minus_sign:                                                 | The options for this request.                                      |

### Response

**[*operations.CreateAdminProductsMappingsActivateByProductResponse](../../models/operations/createadminproductsmappingsactivatebyproductresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Create

Create and save a new Mapping Ruleset for the specified Cribl product.

### Example Usage

<!-- UsageSnippet language="go" operationID="createAdminProductsMappingsByProduct" method="post" path="/admin/products/{product}/mappings" -->
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

    res, err := s.Groups.Mappings.Create(ctx, components.ProductsCoreEdge, components.MappingRuleset{
        ID: "os-based-mapping",
        Conf: &components.MappingRulesetConf{
            Functions: []components.EvalFunction{
                components.EvalFunction{
                    Filter: criblcontrolplanesdkgo.Pointer("platform === \"linux\""),
                    ID: components.IDEval,
                    Description: criblcontrolplanesdkgo.Pointer("Map Linux Edge Nodes"),
                    Disabled: criblcontrolplanesdkgo.Pointer(false),
                    Final: criblcontrolplanesdkgo.Pointer(true),
                    Conf: components.EvalSchema{
                        Add: []components.Add{
                            components.Add{
                                Name: criblcontrolplanesdkgo.Pointer("groupId"),
                                Value: "'linux_fleet'",
                            },
                        },
                        Keep: []string{
                            "<value 1>",
                        },
                        Remove: []string{
                            "<value 1>",
                        },
                    },
                    GroupID: criblcontrolplanesdkgo.Pointer("<id>"),
                },
                components.EvalFunction{
                    Filter: criblcontrolplanesdkgo.Pointer("platform === \"win32\""),
                    ID: components.IDEval,
                    Description: criblcontrolplanesdkgo.Pointer("Map Windows Edge Nodes"),
                    Disabled: criblcontrolplanesdkgo.Pointer(false),
                    Final: criblcontrolplanesdkgo.Pointer(true),
                    Conf: components.EvalSchema{
                        Add: []components.Add{
                            components.Add{
                                Name: criblcontrolplanesdkgo.Pointer("groupId"),
                                Value: "'windows_fleet'",
                            },
                        },
                        Keep: []string{
                            "<value 1>",
                            "<value 2>",
                        },
                        Remove: []string{
                            "<value 1>",
                            "<value 2>",
                            "<value 3>",
                        },
                    },
                    GroupID: criblcontrolplanesdkgo.Pointer("<id>"),
                },
                components.EvalFunction{
                    Filter: criblcontrolplanesdkgo.Pointer("platform === \"darwin\""),
                    ID: components.IDEval,
                    Description: criblcontrolplanesdkgo.Pointer("Map macOS Edge Nodes"),
                    Disabled: criblcontrolplanesdkgo.Pointer(false),
                    Final: criblcontrolplanesdkgo.Pointer(true),
                    Conf: components.EvalSchema{
                        Add: []components.Add{
                            components.Add{
                                Name: criblcontrolplanesdkgo.Pointer("groupId"),
                                Value: "'macos_fleet'",
                            },
                        },
                        Keep: []string{
                            "<value 1>",
                            "<value 2>",
                        },
                        Remove: []string{
                            "<value 1>",
                        },
                    },
                    GroupID: criblcontrolplanesdkgo.Pointer("<id>"),
                },
                components.EvalFunction{
                    Filter: criblcontrolplanesdkgo.Pointer("!cribl.group"),
                    ID: components.IDEval,
                    Description: criblcontrolplanesdkgo.Pointer("Default mapping for unmapped nodes"),
                    Disabled: criblcontrolplanesdkgo.Pointer(false),
                    Final: criblcontrolplanesdkgo.Pointer(true),
                    Conf: components.EvalSchema{
                        Add: []components.Add{
                            components.Add{
                                Name: criblcontrolplanesdkgo.Pointer("groupId"),
                                Value: "'default_fleet'",
                            },
                        },
                        Keep: []string{
                            "<value 1>",
                            "<value 2>",
                        },
                        Remove: []string{
                            "<value 1>",
                        },
                    },
                    GroupID: criblcontrolplanesdkgo.Pointer("<id>"),
                },
            },
        },
        Active: criblcontrolplanesdkgo.Pointer(false),
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

| Parameter                                                              | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `ctx`                                                                  | [context.Context](https://pkg.go.dev/context#Context)                  | :heavy_check_mark:                                                     | The context to use for the request.                                    |
| `product`                                                              | [components.ProductsCore](../../models/components/productscore.md)     | :heavy_check_mark:                                                     | Name of the Cribl product to create the Mapping Ruleset for            |
| `mappingRuleset`                                                       | [components.MappingRuleset](../../models/components/mappingruleset.md) | :heavy_check_mark:                                                     | MappingRuleset object                                                  |
| `opts`                                                                 | [][operations.Option](../../models/operations/option.md)               | :heavy_minus_sign:                                                     | The options for this request.                                          |

### Response

**[*operations.CreateAdminProductsMappingsByProductResponse](../../models/operations/createadminproductsmappingsbyproductresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## List

Retrieve a list of all existing Mapping Rulesets for the specified Cribl product.

### Example Usage

<!-- UsageSnippet language="go" operationID="getAdminProductsMappingsByProduct" method="get" path="/admin/products/{product}/mappings" -->
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

    res, err := s.Groups.Mappings.List(ctx, components.ProductsCoreEdge)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                          | Type                                                               | Required                                                           | Description                                                        |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `ctx`                                                              | [context.Context](https://pkg.go.dev/context#Context)              | :heavy_check_mark:                                                 | The context to use for the request.                                |
| `product`                                                          | [components.ProductsCore](../../models/components/productscore.md) | :heavy_check_mark:                                                 | Name of the Cribl product to list the Mapping Rulesets for         |
| `opts`                                                             | [][operations.Option](../../models/operations/option.md)           | :heavy_minus_sign:                                                 | The options for this request.                                      |

### Response

**[*operations.GetAdminProductsMappingsByProductResponse](../../models/operations/getadminproductsmappingsbyproductresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Delete

Permanently remove the specified Mapping Ruleset from the Worker Group or Edge Fleet.

### Example Usage

<!-- UsageSnippet language="go" operationID="deleteAdminProductsMappingsByProductAndId" method="delete" path="/admin/products/{product}/mappings/{id}" -->
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

    res, err := s.Groups.Mappings.Delete(ctx, components.ProductsCoreStream, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                          | Type                                                               | Required                                                           | Description                                                        |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `ctx`                                                              | [context.Context](https://pkg.go.dev/context#Context)              | :heavy_check_mark:                                                 | The context to use for the request.                                |
| `product`                                                          | [components.ProductsCore](../../models/components/productscore.md) | :heavy_check_mark:                                                 | Name of the Cribl product to delete the Mapping Ruleset for        |
| `id`                                                               | *string*                                                           | :heavy_check_mark:                                                 | The <code>id</code> of the Mapping Ruleset to delete.              |
| `opts`                                                             | [][operations.Option](../../models/operations/option.md)           | :heavy_minus_sign:                                                 | The options for this request.                                      |

### Response

**[*operations.DeleteAdminProductsMappingsByProductAndIDResponse](../../models/operations/deleteadminproductsmappingsbyproductandidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Get

Get the details for a single Mapping Ruleset, identified by its id, for a Worker Group or Edge Fleet.

### Example Usage

<!-- UsageSnippet language="go" operationID="getAdminProductsMappingsByProductAndId" method="get" path="/admin/products/{product}/mappings/{id}" -->
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

    res, err := s.Groups.Mappings.Get(ctx, components.ProductsCoreStream, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `product`                                                                    | [components.ProductsCore](../../models/components/productscore.md)           | :heavy_check_mark:                                                           | Name of the Cribl product to get the Mappings for                            |
| `id`                                                                         | *string*                                                                     | :heavy_check_mark:                                                           | The <code>id</code> of the Worker Group or Edge Fleet Mapping Ruleset to get |
| `opts`                                                                       | [][operations.Option](../../models/operations/option.md)                     | :heavy_minus_sign:                                                           | The options for this request.                                                |

### Response

**[*operations.GetAdminProductsMappingsByProductAndIDResponse](../../models/operations/getadminproductsmappingsbyproductandidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Update

Modify the configuration of the specified Mapping Ruleset for a Worker Group or Edge Fleet.

### Example Usage

<!-- UsageSnippet language="go" operationID="updateAdminProductsMappingsByProductAndId" method="patch" path="/admin/products/{product}/mappings/{id}" -->
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

    res, err := s.Groups.Mappings.Update(ctx, components.ProductsCoreStream, "<id>", components.MappingRuleset{
        ID: "complex-mapping",
        Conf: &components.MappingRulesetConf{
            Functions: []components.EvalFunction{
                components.EvalFunction{
                    Filter: criblcontrolplanesdkgo.Pointer("(conn_ip.startsWith(\"10.10.42.\") && cpus > 6) || env.CRIBL_HOME.match(\"DMZ\")"),
                    ID: components.IDEval,
                    Description: criblcontrolplanesdkgo.Pointer("Map high-performance nodes in specific network or DMZ"),
                    Disabled: criblcontrolplanesdkgo.Pointer(false),
                    Final: criblcontrolplanesdkgo.Pointer(true),
                    Conf: components.EvalSchema{
                        Add: []components.Add{
                            components.Add{
                                Name: criblcontrolplanesdkgo.Pointer("groupId"),
                                Value: "'high_perf_fleet'",
                            },
                        },
                        Keep: []string{
                            "<value 1>",
                        },
                        Remove: []string{
                            "<value 1>",
                            "<value 2>",
                        },
                    },
                    GroupID: criblcontrolplanesdkgo.Pointer("<id>"),
                },
                components.EvalFunction{
                    Filter: criblcontrolplanesdkgo.Pointer("!cribl.group"),
                    ID: components.IDEval,
                    Description: criblcontrolplanesdkgo.Pointer("Default mapping"),
                    Disabled: criblcontrolplanesdkgo.Pointer(false),
                    Final: criblcontrolplanesdkgo.Pointer(true),
                    Conf: components.EvalSchema{
                        Add: []components.Add{
                            components.Add{
                                Name: criblcontrolplanesdkgo.Pointer("groupId"),
                                Value: "'default_fleet'",
                            },
                        },
                        Keep: []string{
                            "<value 1>",
                        },
                        Remove: []string{
                            "<value 1>",
                        },
                    },
                    GroupID: criblcontrolplanesdkgo.Pointer("<id>"),
                },
            },
        },
        Active: criblcontrolplanesdkgo.Pointer(true),
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

| Parameter                                                              | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `ctx`                                                                  | [context.Context](https://pkg.go.dev/context#Context)                  | :heavy_check_mark:                                                     | The context to use for the request.                                    |
| `product`                                                              | [components.ProductsCore](../../models/components/productscore.md)     | :heavy_check_mark:                                                     | Name of the Cribl product to update the Mapping Ruleset for            |
| `id`                                                                   | *string*                                                               | :heavy_check_mark:                                                     | The <code>id</code> of the Mapping Ruleset to update.                  |
| `mappingRuleset`                                                       | [components.MappingRuleset](../../models/components/mappingruleset.md) | :heavy_check_mark:                                                     | MappingRuleset object                                                  |
| `opts`                                                                 | [][operations.Option](../../models/operations/option.md)               | :heavy_minus_sign:                                                     | The options for this request.                                          |

### Response

**[*operations.UpdateAdminProductsMappingsByProductAndIDResponse](../../models/operations/updateadminproductsmappingsbyproductandidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |