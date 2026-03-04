# CreateInputSystemByPackDiscoveryTypePrometheus

Target discovery mechanism. Use static to manually enter a list of targets.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/operations"
)

value := operations.CreateInputSystemByPackDiscoveryTypePrometheusStatic

// Open enum: custom values can be created with a direct type cast
custom := operations.CreateInputSystemByPackDiscoveryTypePrometheus("custom_value")
```


## Values

| Name                                                   | Value                                                  |
| ------------------------------------------------------ | ------------------------------------------------------ |
| `CreateInputSystemByPackDiscoveryTypePrometheusStatic` | static                                                 |
| `CreateInputSystemByPackDiscoveryTypePrometheusDNS`    | dns                                                    |
| `CreateInputSystemByPackDiscoveryTypePrometheusEc2`    | ec2                                                    |