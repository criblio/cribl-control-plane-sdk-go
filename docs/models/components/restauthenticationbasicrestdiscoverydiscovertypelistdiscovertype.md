# RestAuthenticationBasicRestDiscoveryDiscoverTypeListDiscoverType

Defines how task discovery will be performed. Each entry returned by the Discover operation will result in a Collect task.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.RestAuthenticationBasicRestDiscoveryDiscoverTypeListDiscoverTypeHTTP

// Open enum: custom values can be created with a direct type cast
custom := components.RestAuthenticationBasicRestDiscoveryDiscoverTypeListDiscoverType("custom_value")
```


## Values

| Name                                                                   | Value                                                                  |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `RestAuthenticationBasicRestDiscoveryDiscoverTypeListDiscoverTypeHTTP` | http                                                                   |
| `RestAuthenticationBasicRestDiscoveryDiscoverTypeListDiscoverTypeJSON` | json                                                                   |
| `RestAuthenticationBasicRestDiscoveryDiscoverTypeListDiscoverTypeList` | list                                                                   |
| `RestAuthenticationBasicRestDiscoveryDiscoverTypeListDiscoverTypeNone` | none                                                                   |