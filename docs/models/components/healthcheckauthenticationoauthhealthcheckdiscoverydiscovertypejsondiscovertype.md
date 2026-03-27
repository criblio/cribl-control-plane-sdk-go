# HealthCheckAuthenticationOauthHealthCheckDiscoveryDiscoverTypeJSONDiscoverType

Defines how task discovery will be performed. Use None to skip the discovery. Use HTTP Request to make a REST call to discover tasks. Use Item List to enumerate items for collect to retrieve. Use JSON Response to manually define discover tasks as a JSON array of objects. Each entry returned by the discover operation will result in a collect task.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.HealthCheckAuthenticationOauthHealthCheckDiscoveryDiscoverTypeJSONDiscoverTypeHTTP

// Open enum: custom values can be created with a direct type cast
custom := components.HealthCheckAuthenticationOauthHealthCheckDiscoveryDiscoverTypeJSONDiscoverType("custom_value")
```


## Values

| Name                                                                                 | Value                                                                                |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `HealthCheckAuthenticationOauthHealthCheckDiscoveryDiscoverTypeJSONDiscoverTypeHTTP` | http                                                                                 |
| `HealthCheckAuthenticationOauthHealthCheckDiscoveryDiscoverTypeJSONDiscoverTypeJSON` | json                                                                                 |
| `HealthCheckAuthenticationOauthHealthCheckDiscoveryDiscoverTypeJSONDiscoverTypeList` | list                                                                                 |
| `HealthCheckAuthenticationOauthHealthCheckDiscoveryDiscoverTypeJSONDiscoverTypeNone` | none                                                                                 |