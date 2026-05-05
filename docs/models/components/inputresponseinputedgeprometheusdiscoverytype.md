# InputResponseInputEdgePrometheusDiscoveryType

Target discovery mechanism. Use static to manually enter a list of targets.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.InputResponseInputEdgePrometheusDiscoveryTypeStatic

// Open enum: custom values can be created with a direct type cast
custom := components.InputResponseInputEdgePrometheusDiscoveryType("custom_value")
```


## Values

| Name                                                   | Value                                                  |
| ------------------------------------------------------ | ------------------------------------------------------ |
| `InputResponseInputEdgePrometheusDiscoveryTypeStatic`  | static                                                 |
| `InputResponseInputEdgePrometheusDiscoveryTypeDNS`     | dns                                                    |
| `InputResponseInputEdgePrometheusDiscoveryTypeEc2`     | ec2                                                    |
| `InputResponseInputEdgePrometheusDiscoveryTypeK8sNode` | k8s-node                                               |
| `InputResponseInputEdgePrometheusDiscoveryTypeK8sPods` | k8s-pods                                               |