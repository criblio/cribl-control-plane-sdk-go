# InputResponseInputPrometheusDiscoveryType

Target discovery mechanism. Use static to manually enter a list of targets.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.InputResponseInputPrometheusDiscoveryTypeStatic

// Open enum: custom values can be created with a direct type cast
custom := components.InputResponseInputPrometheusDiscoveryType("custom_value")
```


## Values

| Name                                              | Value                                             |
| ------------------------------------------------- | ------------------------------------------------- |
| `InputResponseInputPrometheusDiscoveryTypeStatic` | static                                            |
| `InputResponseInputPrometheusDiscoveryTypeDNS`    | dns                                               |
| `InputResponseInputPrometheusDiscoveryTypeEc2`    | ec2                                               |