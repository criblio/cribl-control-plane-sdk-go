# RestCollectMethodPostWithBodyRestDiscoveryDiscoverTypeListDiscoverType

Defines how task discovery will be performed. Each entry returned by the Discover operation will result in a Collect task.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.RestCollectMethodPostWithBodyRestDiscoveryDiscoverTypeListDiscoverTypeHTTP

// Open enum: custom values can be created with a direct type cast
custom := components.RestCollectMethodPostWithBodyRestDiscoveryDiscoverTypeListDiscoverType("custom_value")
```


## Values

| Name                                                                         | Value                                                                        |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `RestCollectMethodPostWithBodyRestDiscoveryDiscoverTypeListDiscoverTypeHTTP` | http                                                                         |
| `RestCollectMethodPostWithBodyRestDiscoveryDiscoverTypeListDiscoverTypeJSON` | json                                                                         |
| `RestCollectMethodPostWithBodyRestDiscoveryDiscoverTypeListDiscoverTypeList` | list                                                                         |
| `RestCollectMethodPostWithBodyRestDiscoveryDiscoverTypeListDiscoverTypeNone` | none                                                                         |