# System.Settings.Cribl

## Overview

### Available Operations

* [List](#list) - Get system settings
* [Update](#update) - Update system settings

## List

Get the current Cribl system settings.

### Example Usage

<!-- UsageSnippet language="go" operationID="getSystemSettingsConf" method="get" path="/system/settings/conf" example="GetSystemSettingsConfExamplesDefault" -->
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

    res, err := s.System.Settings.Cribl.List(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedSystemSettingsConf != nil {
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

**[*operations.GetSystemSettingsConfResponse](../../models/operations/getsystemsettingsconfresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 401                | application/json   |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Update

Update the specified Cribl system settings.<br/><br/>Provide a complete representation of the system settings that you want to update in the request body. This endpoint does not support partial updates. Cribl removes any omitted fields when updating the system settings.<br/><br/>Confirm that the configuration in your request body is correct before sending the request. If the configuration is incorrect, the updated system settings might not function as expected.

### Example Usage: UpdateSystemSettingsExamplesUpdateApiSettings

<!-- UsageSnippet language="go" operationID="updateSystemSettingsConf" method="patch" path="/system/settings/conf" example="UpdateSystemSettingsExamplesUpdateApiSettings" -->
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

    res, err := s.System.Settings.Cribl.Update(ctx, components.SystemSettingsConf{
        API: components.APITypeSystemSettingsConf{
            Disabled: false,
            Host: "0.0.0.0",
            Port: 9000,
            Ssl: &components.SslTypeSystemSettingsConfAPI{
                CertPath: "/opt/cribl/local/cribl/auth/cribl.crt",
                Disabled: false,
                Passphrase: "",
                PrivKeyPath: "/opt/cribl/local/cribl/auth/cribl.key",
            },
        },
        Backups: components.CreateBackupsSettingsUnionBackupsSettings(
            components.BackupsSettings{
                BackupPersistence: "24h",
                BackupsDirectory: "$CRIBL_STATE_DIR/backups",
            },
        ),
        Pii: components.CreatePiiSettingsUnionPiiSettings(
            components.PiiSettings{
                EnablePiiDetection: false,
            },
        ),
        Proxy: components.ProxyTypeSystemSettingsConf{
            UseEnvVars: false,
        },
        Rollback: components.CreateRollbackSettingsUnionRollbackSettings(
            components.RollbackSettings{
                RollbackEnabled: true,
            },
        ),
        Shutdown: components.ShutdownTypeSystemSettingsConf{
            DrainTimeout: 10000,
        },
        Sni: components.CreateSniSettingsUnionSniSettings(
            components.SniSettings{
                DisableSNIRouting: false,
            },
        ),
        System: components.SystemTypeSystemSettingsConf{
            Intercom: true,
            Upgrade: components.UpgradeOptionsSystemSettingsConfSystemAPI,
        },
        TLS: components.CreateTLSSettingsUnionTLSSettings(
            components.TLSSettings{
                DefaultCipherList: "DEFAULT",
                DefaultEcdhCurve: "auto",
                MaxVersion: "TLSv1.3",
                MinVersion: "TLSv1.2",
                RejectUnauthorized: true,
            },
        ),
        UpgradeGroupSettings: components.UpgradeGroupSettings{
            IsRolling: criblcontrolplanesdkgo.Pointer(true),
            Quantity: criblcontrolplanesdkgo.Pointer[int64](100),
            RetryCount: criblcontrolplanesdkgo.Pointer[int64](5),
            RetryDelay: criblcontrolplanesdkgo.Pointer[int64](1000),
        },
        UpgradeSettings: components.UpgradeSettings{},
        Workers: components.WorkersTypeSystemSettingsConf{
            Count: 0,
            Memory: 0,
            Minimum: 1,
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedSystemSettingsConf != nil {
        // handle response
    }
}
```
### Example Usage: UpdateSystemSettingsResponseExamplesUpdateApiSettings

<!-- UsageSnippet language="go" operationID="updateSystemSettingsConf" method="patch" path="/system/settings/conf" example="UpdateSystemSettingsResponseExamplesUpdateApiSettings" -->
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

    res, err := s.System.Settings.Cribl.Update(ctx, components.SystemSettingsConf{
        API: components.APITypeSystemSettingsConf{
            Disabled: true,
            Host: "both-draw.com",
            Port: 379506,
        },
        Backups: components.CreateBackupsSettingsUnionEmptyObject(
            components.EmptyObject{},
        ),
        Pii: components.CreatePiiSettingsUnionEmptyObject(
            components.EmptyObject{},
        ),
        Proxy: components.ProxyTypeSystemSettingsConf{
            UseEnvVars: false,
        },
        Rollback: components.CreateRollbackSettingsUnionEmptyObject(
            components.EmptyObject{},
        ),
        Shutdown: components.ShutdownTypeSystemSettingsConf{
            DrainTimeout: 506758,
        },
        Sni: components.CreateSniSettingsUnionEmptyObject(
            components.EmptyObject{},
        ),
        System: components.SystemTypeSystemSettingsConf{
            Intercom: false,
            Upgrade: components.UpgradeOptionsSystemSettingsConfSystemAPI,
        },
        TLS: components.CreateTLSSettingsUnionEmptyObject(
            components.EmptyObject{},
        ),
        UpgradeGroupSettings: components.UpgradeGroupSettings{},
        UpgradeSettings: components.UpgradeSettings{},
        Workers: components.WorkersTypeSystemSettingsConf{
            Count: 8927,
            Memory: 142072,
            Minimum: 242438,
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedSystemSettingsConf != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [components.SystemSettingsConf](../../models/components/systemsettingsconf.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |
| `opts`                                                                         | [][operations.Option](../../models/operations/option.md)                       | :heavy_minus_sign:                                                             | The options for this request.                                                  |

### Response

**[*operations.UpdateSystemSettingsConfResponse](../../models/operations/updatesystemsettingsconfresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 401                | application/json   |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |