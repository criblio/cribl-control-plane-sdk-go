# Groups.Mappings

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

<!-- UsageSnippet language="go" operationID="createAdminProductsMappingsActivateByProduct" method="post" path="/admin/products/{product}/mappings/activate" example="MappingRulesetActivateExamplesActivate" -->
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

    res, err := s.Groups.Mappings.Activate(ctx, components.ProductsBaseStream, components.RulesetID{
        ID: "default",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedRulesetID != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                          | Type                                                               | Required                                                           | Description                                                        |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `ctx`                                                              | [context.Context](https://pkg.go.dev/context#Context)              | :heavy_check_mark:                                                 | The context to use for the request.                                |
| `product`                                                          | [components.ProductsBase](../../models/components/productsbase.md) | :heavy_check_mark:                                                 | Name of the Cribl product to activate the Mapping Ruleset for      |
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

Create and save a new Mapping Ruleset for the specified Cribl product.<br><br>Note: This functionality is not supported for Cribl Stream Workers on Cribl.Cloud.

### Example Usage: MappingRulesetCreateExamplesAWSEC2InstanceMapping

<!-- UsageSnippet language="go" operationID="createAdminProductsMappingsByProduct" method="post" path="/admin/products/{product}/mappings" example="MappingRulesetCreateExamplesAWSEC2InstanceMapping" -->
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

    res, err := s.Groups.Mappings.Create(ctx, components.ProductsBaseEdge, components.MappingRuleset{
        ID: "aws-ec2-mapping",
        Conf: &components.MappingRulesetConf{
            Functions: []components.Function{
                components.Function{
                    Conf: components.ConfEval{
                        Add: []components.MappingRulesetAdd{
                            components.MappingRulesetAdd{
                                Value: "'aws_prod_fleet'",
                            },
                        },
                    },
                    AdditionalProperties: map[string]any{
                        "filter": "aws?.tags?.Environment === \"Production\"",
                        "description": "Map Production EC2 instances",
                        "disabled": false,
                        "final": true,
                    },
                },
                components.Function{
                    Conf: components.ConfEval{
                        Add: []components.MappingRulesetAdd{
                            components.MappingRulesetAdd{
                                Value: "'devops_fleet'",
                            },
                        },
                    },
                    AdditionalProperties: map[string]any{
                        "filter": "aws?.tags?.Team === \"DevOps\"",
                        "description": "Map DevOps team EC2 instances",
                        "disabled": false,
                        "final": true,
                    },
                },
                components.Function{
                    Conf: components.ConfEval{
                        Add: []components.MappingRulesetAdd{
                            components.MappingRulesetAdd{
                                Value: "'default_fleet'",
                            },
                        },
                    },
                    AdditionalProperties: map[string]any{
                        "filter": "!cribl.group",
                        "description": "Default mapping",
                        "disabled": false,
                        "final": true,
                    },
                },
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedMappingRuleset != nil {
        // handle response
    }
}
```
### Example Usage: MappingRulesetCreateExamplesComplexMapping

<!-- UsageSnippet language="go" operationID="createAdminProductsMappingsByProduct" method="post" path="/admin/products/{product}/mappings" example="MappingRulesetCreateExamplesComplexMapping" -->
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

    res, err := s.Groups.Mappings.Create(ctx, components.ProductsBaseEdge, components.MappingRuleset{
        ID: "complex-mapping",
        Conf: &components.MappingRulesetConf{
            Functions: []components.Function{
                components.Function{
                    Conf: components.ConfEval{
                        Add: []components.MappingRulesetAdd{
                            components.MappingRulesetAdd{
                                Value: "'high_perf_fleet'",
                            },
                        },
                    },
                    AdditionalProperties: map[string]any{
                        "filter": "(conn_ip.startsWith(\"10.10.42.\") && cpus > 6) || env.CRIBL_HOME.match(\"DMZ\")",
                        "description": "Map high-performance nodes in specific network or DMZ",
                        "disabled": false,
                        "final": true,
                    },
                },
                components.Function{
                    Conf: components.ConfEval{
                        Add: []components.MappingRulesetAdd{
                            components.MappingRulesetAdd{
                                Value: "'database_fleet'",
                            },
                        },
                    },
                    AdditionalProperties: map[string]any{
                        "filter": "platform === \"linux\" && memory > 16000 && cribl.tags.includes(\"Database\")",
                        "description": "Map Linux nodes with high memory for database workloads",
                        "disabled": false,
                        "final": true,
                    },
                },
                components.Function{
                    Conf: components.ConfEval{
                        Add: []components.MappingRulesetAdd{
                            components.MappingRulesetAdd{
                                Value: "'default_fleet'",
                            },
                        },
                    },
                    AdditionalProperties: map[string]any{
                        "filter": "!cribl.group",
                        "description": "Default mapping",
                        "disabled": false,
                        "final": true,
                    },
                },
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedMappingRuleset != nil {
        // handle response
    }
}
```
### Example Usage: MappingRulesetCreateExamplesContainerBasedMapping

<!-- UsageSnippet language="go" operationID="createAdminProductsMappingsByProduct" method="post" path="/admin/products/{product}/mappings" example="MappingRulesetCreateExamplesContainerBasedMapping" -->
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

    res, err := s.Groups.Mappings.Create(ctx, components.ProductsBaseStream, components.MappingRuleset{
        ID: "container-mapping",
        Conf: &components.MappingRulesetConf{
            Functions: []components.Function{
                components.Function{
                    Conf: components.ConfEval{
                        Add: []components.MappingRulesetAdd{
                            components.MappingRulesetAdd{
                                Value: "'container_high_cpu_fleet'",
                            },
                        },
                    },
                    AdditionalProperties: map[string]any{
                        "filter": "hostOs?.platform === \"linux\" && hostOs?.cpu_count > 4",
                        "description": "Map containerized nodes on high-CPU Linux hosts",
                        "disabled": false,
                        "final": true,
                    },
                },
                components.Function{
                    Conf: components.ConfEval{
                        Add: []components.MappingRulesetAdd{
                            components.MappingRulesetAdd{
                                Value: "'container_fleet'",
                            },
                        },
                    },
                    AdditionalProperties: map[string]any{
                        "filter": "hostOs?.id",
                        "description": "Map all containerized Edge Nodes",
                        "disabled": false,
                        "final": true,
                    },
                },
                components.Function{
                    Conf: components.ConfEval{
                        Add: []components.MappingRulesetAdd{
                            components.MappingRulesetAdd{
                                Value: "'default_fleet'",
                            },
                        },
                    },
                    AdditionalProperties: map[string]any{
                        "filter": "!cribl.group",
                        "description": "Default mapping for non-containerized nodes",
                        "disabled": false,
                        "final": true,
                    },
                },
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedMappingRuleset != nil {
        // handle response
    }
}
```
### Example Usage: MappingRulesetCreateExamplesDefaultRuleset

<!-- UsageSnippet language="go" operationID="createAdminProductsMappingsByProduct" method="post" path="/admin/products/{product}/mappings" example="MappingRulesetCreateExamplesDefaultRuleset" -->
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

    res, err := s.Groups.Mappings.Create(ctx, components.ProductsBaseEdge, components.MappingRuleset{
        ID: "simple-default-mappings",
        Conf: &components.MappingRulesetConf{
            Functions: []components.Function{
                components.Function{
                    Conf: components.ConfEval{
                        Add: []components.MappingRulesetAdd{
                            components.MappingRulesetAdd{
                                Value: "'default_fleet'",
                            },
                        },
                    },
                    AdditionalProperties: map[string]any{
                        "filter": "!cribl.group",
                        "description": "Simple default Mapping Ruleset",
                        "disabled": false,
                        "final": true,
                    },
                },
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedMappingRuleset != nil {
        // handle response
    }
}
```
### Example Usage: MappingRulesetCreateExamplesOSBasedMapping

<!-- UsageSnippet language="go" operationID="createAdminProductsMappingsByProduct" method="post" path="/admin/products/{product}/mappings" example="MappingRulesetCreateExamplesOSBasedMapping" -->
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

    res, err := s.Groups.Mappings.Create(ctx, components.ProductsBaseStream, components.MappingRuleset{
        ID: "os-based-mapping",
        Conf: &components.MappingRulesetConf{
            Functions: []components.Function{
                components.Function{
                    Conf: components.ConfEval{
                        Add: []components.MappingRulesetAdd{
                            components.MappingRulesetAdd{
                                Value: "'linux_fleet'",
                            },
                        },
                    },
                    AdditionalProperties: map[string]any{
                        "filter": "platform === \"linux\"",
                        "description": "Map Linux Edge Nodes",
                        "disabled": false,
                        "final": true,
                    },
                },
                components.Function{
                    Conf: components.ConfEval{
                        Add: []components.MappingRulesetAdd{
                            components.MappingRulesetAdd{
                                Value: "'windows_fleet'",
                            },
                        },
                    },
                    AdditionalProperties: map[string]any{
                        "filter": "platform === \"win32\"",
                        "description": "Map Windows Edge Nodes",
                        "disabled": false,
                        "final": true,
                    },
                },
                components.Function{
                    Conf: components.ConfEval{
                        Add: []components.MappingRulesetAdd{
                            components.MappingRulesetAdd{
                                Value: "'macos_fleet'",
                            },
                        },
                    },
                    AdditionalProperties: map[string]any{
                        "filter": "platform === \"darwin\"",
                        "description": "Map macOS Edge Nodes",
                        "disabled": false,
                        "final": true,
                    },
                },
                components.Function{
                    Conf: components.ConfEval{
                        Add: []components.MappingRulesetAdd{
                            components.MappingRulesetAdd{
                                Value: "'default_fleet'",
                            },
                        },
                    },
                    AdditionalProperties: map[string]any{
                        "filter": "!cribl.group",
                        "description": "Default mapping for unmapped nodes",
                        "disabled": false,
                        "final": true,
                    },
                },
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedMappingRuleset != nil {
        // handle response
    }
}
```
### Example Usage: MappingRulesetCreateExamplesOutpostBasedMapping

<!-- UsageSnippet language="go" operationID="createAdminProductsMappingsByProduct" method="post" path="/admin/products/{product}/mappings" example="MappingRulesetCreateExamplesOutpostBasedMapping" -->
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

    res, err := s.Groups.Mappings.Create(ctx, components.ProductsBaseEdge, components.MappingRuleset{
        ID: "outpost-mapping",
        Conf: &components.MappingRulesetConf{
            Functions: []components.Function{
                components.Function{
                    Conf: components.ConfEval{
                        Add: []components.MappingRulesetAdd{
                            components.MappingRulesetAdd{
                                Value: "'outpost_a_fleet'",
                            },
                        },
                    },
                    AdditionalProperties: map[string]any{
                        "filter": "outpost.host === \"5ab6c676be6a\"",
                        "description": "Map Edge Nodes from Outpost A",
                        "disabled": false,
                        "final": true,
                    },
                },
                components.Function{
                    Conf: components.ConfEval{
                        Add: []components.MappingRulesetAdd{
                            components.MappingRulesetAdd{
                                Value: "'outpost_b_fleet'",
                            },
                        },
                    },
                    AdditionalProperties: map[string]any{
                        "filter": "outpost.guid === \"550e8400-e29b-41d4-a716-446655440000\"",
                        "description": "Map Edge Nodes from Outpost B by GUID",
                        "disabled": false,
                        "final": true,
                    },
                },
                components.Function{
                    Conf: components.ConfEval{
                        Add: []components.MappingRulesetAdd{
                            components.MappingRulesetAdd{
                                Value: "'default_fleet'",
                            },
                        },
                    },
                    AdditionalProperties: map[string]any{
                        "filter": "!cribl.group",
                        "description": "Default mapping",
                        "disabled": false,
                        "final": true,
                    },
                },
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedMappingRuleset != nil {
        // handle response
    }
}
```
### Example Usage: MappingRulesetCreateExamplesTagBasedMapping

<!-- UsageSnippet language="go" operationID="createAdminProductsMappingsByProduct" method="post" path="/admin/products/{product}/mappings" example="MappingRulesetCreateExamplesTagBasedMapping" -->
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

    res, err := s.Groups.Mappings.Create(ctx, components.ProductsBaseStream, components.MappingRuleset{
        ID: "tag-based-mapping",
        Conf: &components.MappingRulesetConf{
            Functions: []components.Function{
                components.Function{
                    Conf: components.ConfEval{
                        Add: []components.MappingRulesetAdd{
                            components.MappingRulesetAdd{
                                Value: "'laptop_fleet'",
                            },
                        },
                    },
                    AdditionalProperties: map[string]any{
                        "filter": "cribl.tags.includes(\"WinLaptop\")",
                        "description": "Map Windows Laptop Edge Nodes",
                        "disabled": false,
                        "final": true,
                    },
                },
                components.Function{
                    Conf: components.ConfEval{
                        Add: []components.MappingRulesetAdd{
                            components.MappingRulesetAdd{
                                Value: "'production_fleet'",
                            },
                        },
                    },
                    AdditionalProperties: map[string]any{
                        "filter": "cribl.tags.includes(\"Production\")",
                        "description": "Map Production Edge Nodes",
                        "disabled": false,
                        "final": true,
                    },
                },
                components.Function{
                    Conf: components.ConfEval{
                        Add: []components.MappingRulesetAdd{
                            components.MappingRulesetAdd{
                                Value: "'default_fleet'",
                            },
                        },
                    },
                    AdditionalProperties: map[string]any{
                        "filter": "!cribl.group",
                        "description": "Default mapping",
                        "disabled": false,
                        "final": true,
                    },
                },
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedMappingRuleset != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                              | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `ctx`                                                                  | [context.Context](https://pkg.go.dev/context#Context)                  | :heavy_check_mark:                                                     | The context to use for the request.                                    |
| `product`                                                              | [components.ProductsBase](../../models/components/productsbase.md)     | :heavy_check_mark:                                                     | Name of the Cribl product to create the Mapping Ruleset for            |
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

    res, err := s.Groups.Mappings.List(ctx, components.ProductsBaseEdge)
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedMappingRuleset != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                          | Type                                                               | Required                                                           | Description                                                        |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `ctx`                                                              | [context.Context](https://pkg.go.dev/context#Context)              | :heavy_check_mark:                                                 | The context to use for the request.                                |
| `product`                                                          | [components.ProductsBase](../../models/components/productsbase.md) | :heavy_check_mark:                                                 | Name of the Cribl product to list the Mapping Rulesets for         |
| `opts`                                                             | [][operations.Option](../../models/operations/option.md)           | :heavy_minus_sign:                                                 | The options for this request.                                      |

### Response

**[*operations.GetAdminProductsMappingsByProductResponse](../../models/operations/getadminproductsmappingsbyproductresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Delete

Permanently remove the specified Mapping Ruleset from the Worker Group or Edge Fleet.<br><br>Note: This functionality is not supported for Cribl Stream Workers on Cribl.Cloud.

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

    res, err := s.Groups.Mappings.Delete(ctx, components.ProductsBaseStream, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedMappingRuleset != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                          | Type                                                               | Required                                                           | Description                                                        |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `ctx`                                                              | [context.Context](https://pkg.go.dev/context#Context)              | :heavy_check_mark:                                                 | The context to use for the request.                                |
| `product`                                                          | [components.ProductsBase](../../models/components/productsbase.md) | :heavy_check_mark:                                                 | Name of the Cribl product to delete the Mapping Ruleset for        |
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

    res, err := s.Groups.Mappings.Get(ctx, components.ProductsBaseStream, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedMappingRuleset != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `product`                                                                    | [components.ProductsBase](../../models/components/productsbase.md)           | :heavy_check_mark:                                                           | Name of the Cribl product to get the Mappings for                            |
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

### Example Usage: MappingRulesetUpdateExamplesEnableDisabledRule

<!-- UsageSnippet language="go" operationID="updateAdminProductsMappingsByProductAndId" method="patch" path="/admin/products/{product}/mappings/{id}" example="MappingRulesetUpdateExamplesEnableDisabledRule" -->
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

    res, err := s.Groups.Mappings.Update(ctx, components.ProductsBaseEdge, "<id>", components.MappingRuleset{
        ID: "default",
        Conf: &components.MappingRulesetConf{
            Functions: []components.Function{
                components.Function{
                    Conf: components.ConfEval{
                        Add: []components.MappingRulesetAdd{
                            components.MappingRulesetAdd{
                                Value: "'default_fleet'",
                            },
                        },
                    },
                    AdditionalProperties: map[string]any{
                        "filter": "!cribl.group",
                        "description": "Default Mappings",
                        "disabled": false,
                        "final": true,
                    },
                },
            },
        },
        Active: criblcontrolplanesdkgo.Pointer(true),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedMappingRuleset != nil {
        // handle response
    }
}
```
### Example Usage: MappingRulesetUpdateExamplesUpdateComplexRule

<!-- UsageSnippet language="go" operationID="updateAdminProductsMappingsByProductAndId" method="patch" path="/admin/products/{product}/mappings/{id}" example="MappingRulesetUpdateExamplesUpdateComplexRule" -->
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

    res, err := s.Groups.Mappings.Update(ctx, components.ProductsBaseStream, "<id>", components.MappingRuleset{
        ID: "complex-mapping",
        Conf: &components.MappingRulesetConf{
            Functions: []components.Function{
                components.Function{
                    Conf: components.ConfEval{
                        Add: []components.MappingRulesetAdd{
                            components.MappingRulesetAdd{
                                Value: "'high_perf_fleet'",
                            },
                        },
                    },
                    AdditionalProperties: map[string]any{
                        "filter": "(conn_ip.startsWith(\"10.10.42.\") && cpus > 6) || env.CRIBL_HOME.match(\"DMZ\")",
                        "description": "Map high-performance nodes in specific network or DMZ",
                        "disabled": false,
                        "final": true,
                    },
                },
                components.Function{
                    Conf: components.ConfEval{
                        Add: []components.MappingRulesetAdd{
                            components.MappingRulesetAdd{
                                Value: "'default_fleet'",
                            },
                        },
                    },
                    AdditionalProperties: map[string]any{
                        "filter": "!cribl.group",
                        "description": "Default mapping",
                        "disabled": false,
                        "final": true,
                    },
                },
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedMappingRuleset != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                              | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `ctx`                                                                  | [context.Context](https://pkg.go.dev/context#Context)                  | :heavy_check_mark:                                                     | The context to use for the request.                                    |
| `product`                                                              | [components.ProductsBase](../../models/components/productsbase.md)     | :heavy_check_mark:                                                     | Name of the Cribl product to update the Mapping Ruleset for            |
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