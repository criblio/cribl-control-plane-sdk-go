# InputPrometheusDiscoveryType

Target discovery mechanism. Use static to manually enter a list of targets.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.InputPrometheusDiscoveryTypeStatic

// Open enum: custom values can be created with a direct type cast
custom := components.InputPrometheusDiscoveryType("custom_value")
```


## Values

| Name                                 | Value                                |
| ------------------------------------ | ------------------------------------ |
| `InputPrometheusDiscoveryTypeStatic` | static                               |
| `InputPrometheusDiscoveryTypeDNS`    | dns                                  |
| `InputPrometheusDiscoveryTypeEc2`    | ec2                                  |