# RestAuthenticationBasicRestDiscoveryDiscoverTypeJSONDiscoverType

Defines how task discovery will be performed. Each entry returned by the Discover operation will result in a Collect task.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.RestAuthenticationBasicRestDiscoveryDiscoverTypeJSONDiscoverTypeHTTP

// Open enum: custom values can be created with a direct type cast
custom := components.RestAuthenticationBasicRestDiscoveryDiscoverTypeJSONDiscoverType("custom_value")
```


## Values

| Name                                                                   | Value                                                                  |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `RestAuthenticationBasicRestDiscoveryDiscoverTypeJSONDiscoverTypeHTTP` | http                                                                   |
| `RestAuthenticationBasicRestDiscoveryDiscoverTypeJSONDiscoverTypeJSON` | json                                                                   |
| `RestAuthenticationBasicRestDiscoveryDiscoverTypeJSONDiscoverTypeList` | list                                                                   |
| `RestAuthenticationBasicRestDiscoveryDiscoverTypeJSONDiscoverTypeNone` | none                                                                   |