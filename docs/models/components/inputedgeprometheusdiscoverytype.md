# InputEdgePrometheusDiscoveryType

Target discovery mechanism. Use static to manually enter a list of targets.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.InputEdgePrometheusDiscoveryTypeStatic

// Open enum: custom values can be created with a direct type cast
custom := components.InputEdgePrometheusDiscoveryType("custom_value")
```


## Values

| Name                                      | Value                                     |
| ----------------------------------------- | ----------------------------------------- |
| `InputEdgePrometheusDiscoveryTypeStatic`  | static                                    |
| `InputEdgePrometheusDiscoveryTypeDNS`     | dns                                       |
| `InputEdgePrometheusDiscoveryTypeEc2`     | ec2                                       |
| `InputEdgePrometheusDiscoveryTypeK8sNode` | k8s-node                                  |
| `InputEdgePrometheusDiscoveryTypeK8sPods` | k8s-pods                                  |