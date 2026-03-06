# CreateInputDiscoveryTypePrometheus

Target discovery mechanism. Use static to manually enter a list of targets.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/operations"
)

value := operations.CreateInputDiscoveryTypePrometheusStatic

// Open enum: custom values can be created with a direct type cast
custom := operations.CreateInputDiscoveryTypePrometheus("custom_value")
```


## Values

| Name                                       | Value                                      |
| ------------------------------------------ | ------------------------------------------ |
| `CreateInputDiscoveryTypePrometheusStatic` | static                                     |
| `CreateInputDiscoveryTypePrometheusDNS`    | dns                                        |
| `CreateInputDiscoveryTypePrometheusEc2`    | ec2                                        |