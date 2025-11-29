# Cribl
(*System.Settings.Cribl*)

## Overview

### Available Operations

* [List](#list) - Get Cribl system settings
* [Update](#update) - Update Cribl system settings

## List

Get Cribl system settings

### Example Usage

<!-- UsageSnippet language="go" operationID="getSystemSettingsConf" method="get" path="/system/settings/conf" -->
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
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Update

Update Cribl system settings

### Example Usage

<!-- UsageSnippet language="go" operationID="updateSystemSettingsConf" method="patch" path="/system/settings/conf" -->
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
        API: components.SystemSettingsConfAPI{
            BaseURL: criblcontrolplanesdkgo.Pointer("https://both-draw.com/"),
            DisableAPICache: criblcontrolplanesdkgo.Pointer(true),
            Disabled: false,
            Headers: map[string]string{

            },
            Host: "meaty-spring.biz",
            IdleSessionTTL: criblcontrolplanesdkgo.Pointer[float64](89.27),
            ListenOnPort: criblcontrolplanesdkgo.Pointer(true),
            LoginRateLimit: criblcontrolplanesdkgo.Pointer("<value>"),
            Port: 2424.38,
            Protocol: criblcontrolplanesdkgo.Pointer("<value>"),
            Scripts: criblcontrolplanesdkgo.Pointer(true),
            SensitiveFields: []string{
                "<value 1>",
                "<value 2>",
            },
            Ssl: &components.SystemSettingsConfSsl{
                CaPath: criblcontrolplanesdkgo.Pointer("<value>"),
                CertPath: "<value>",
                Disabled: true,
                Passphrase: "<value>",
                PrivKeyPath: "<value>",
            },
            SsoRateLimit: criblcontrolplanesdkgo.Pointer("<value>"),
            WorkerRemoteAccess: criblcontrolplanesdkgo.Pointer(true),
        },
        Backups: components.CreateBackupsSettingsUnionBackupsSettings1(
            components.BackupsSettings1{
                BackupPersistence: "<value>",
                BackupsDirectory: "<value>",
            },
        ),
        CustomLogo: criblcontrolplanesdkgo.Pointer(components.CreateCustomLogoUnionCustomLogo1(
            components.CustomLogo1{
                Enabled: false,
                LogoDescription: "<value>",
                LogoImage: "<value>",
            },
        )),
        Pii: components.CreatePiiSettingsUnionPiiSettings1(
            components.PiiSettings1{
                EnablePiiDetection: false,
            },
        ),
        Proxy: components.CreateProxySettingsUnionProxySettings1(
            components.ProxySettings1{
                UseEnvVars: true,
            },
        ),
        Rollback: components.CreateRollbackSettingsUnionRollbackSettings1(
            components.RollbackSettings1{
                RollbackEnabled: false,
                RollbackRetries: criblcontrolplanesdkgo.Pointer[float64](3174.73),
                RollbackTimeout: criblcontrolplanesdkgo.Pointer[float64](1506.54),
            },
        ),
        Shutdown: components.CreateShutdownSettingsUnionShutdownSettings1(
            components.ShutdownSettings1{
                DrainTimeout: 3723.75,
            },
        ),
        Sni: components.CreateSniSettingsUnionSniSettings1(
            components.SniSettings1{
                DisableSNIRouting: false,
            },
        ),
        Sockets: &components.SystemSettingsConfSockets{
            Directory: criblcontrolplanesdkgo.Pointer("/usr/ports"),
        },
        Support: &components.SystemSettingsConfSupport{
            FeatureFlagOverrides: []components.SystemSettingsConfFeatureFlagOverride{
                components.SystemSettingsConfFeatureFlagOverride{
                    Disabled: true,
                    FlagID: "<id>",
                },
            },
        },
        System: components.SystemSettingsConfSystem{
            Intercom: false,
            Upgrade: components.SystemSettingsConfUpgradeAPI,
        },
        TLS: components.CreateTLSSettingsUnionTLSSettings1(
            components.TLSSettings1{
                DefaultCipherList: "<value>",
                DefaultEcdhCurve: "<value>",
                MaxVersion: "<value>",
                MinVersion: "<value>",
                RejectUnauthorized: true,
            },
        ),
        UpgradeGroupSettings: components.CreateUpgradeGroupSettingsUnionUpgradeGroupSettings1(
            components.UpgradeGroupSettings1{
                IsRolling: criblcontrolplanesdkgo.Pointer(false),
                Quantity: criblcontrolplanesdkgo.Pointer[float64](7915.07),
                RetryCount: criblcontrolplanesdkgo.Pointer[float64](4414.66),
                RetryDelay: criblcontrolplanesdkgo.Pointer[float64](4374.4),
            },
        ),
        UpgradeSettings: components.CreateUpgradeSettingsUnionUpgradeSettings1(
            components.UpgradeSettings1{
                AutomaticUpgradeCheckPeriod: criblcontrolplanesdkgo.Pointer("<value>"),
                DisableAutomaticUpgrade: false,
                EnableLegacyEdgeUpgrade: false,
                PackageUrls: []components.UpgradePackageUrls{
                    components.UpgradePackageUrls{
                        PackageHashURL: criblcontrolplanesdkgo.Pointer("https://thrifty-teammate.net/"),
                        PackageURL: "https://skeletal-dwell.info/",
                    },
                },
                UpgradeSource: "<value>",
            },
        ),
        Workers: components.CreateWorkersSettingsUnionWorkersSettings1(
            components.WorkersSettings1{
                Count: 2124.14,
                EnableHeapSnapshots: criblcontrolplanesdkgo.Pointer(true),
                LoadThrottlePerc: criblcontrolplanesdkgo.Pointer[float64](2538.71),
                Memory: 20.53,
                Minimum: 6157.83,
                StartupMaxConns: criblcontrolplanesdkgo.Pointer[float64](4731.29),
                StartupThrottleTimeout: criblcontrolplanesdkgo.Pointer[float64](1613.48),
                V8SingleThread: criblcontrolplanesdkgo.Pointer(true),
            },
        ),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedSystemSettings != nil {
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
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |