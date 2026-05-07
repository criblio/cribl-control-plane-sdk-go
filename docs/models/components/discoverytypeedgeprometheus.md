# DiscoveryTypeEdgePrometheus

Target discovery mechanism. Use static to manually enter a list of targets.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.DiscoveryTypeEdgePrometheusStatic

// Open enum: custom values can be created with a direct type cast
custom := components.DiscoveryTypeEdgePrometheus("custom_value")
```


## Values

| Name                                 | Value                                |
| ------------------------------------ | ------------------------------------ |
| `DiscoveryTypeEdgePrometheusStatic`  | static                               |
| `DiscoveryTypeEdgePrometheusDNS`     | dns                                  |
| `DiscoveryTypeEdgePrometheusEc2`     | ec2                                  |
| `DiscoveryTypeEdgePrometheusK8sNode` | k8s-node                             |
| `DiscoveryTypeEdgePrometheusK8sPods` | k8s-pods                             |