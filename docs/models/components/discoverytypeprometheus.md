# DiscoveryTypePrometheus

Target discovery mechanism. Use static to manually enter a list of targets.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.DiscoveryTypePrometheusStatic

// Open enum: custom values can be created with a direct type cast
custom := components.DiscoveryTypePrometheus("custom_value")
```


## Values

| Name                            | Value                           |
| ------------------------------- | ------------------------------- |
| `DiscoveryTypePrometheusStatic` | static                          |
| `DiscoveryTypePrometheusDNS`    | dns                             |
| `DiscoveryTypePrometheusEc2`    | ec2                             |