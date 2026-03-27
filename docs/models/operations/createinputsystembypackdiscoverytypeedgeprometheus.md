# CreateInputSystemByPackDiscoveryTypeEdgePrometheus

Target discovery mechanism. Use static to manually enter a list of targets.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/operations"
)

value := operations.CreateInputSystemByPackDiscoveryTypeEdgePrometheusStatic

// Open enum: custom values can be created with a direct type cast
custom := operations.CreateInputSystemByPackDiscoveryTypeEdgePrometheus("custom_value")
```


## Values

| Name                                                        | Value                                                       |
| ----------------------------------------------------------- | ----------------------------------------------------------- |
| `CreateInputSystemByPackDiscoveryTypeEdgePrometheusStatic`  | static                                                      |
| `CreateInputSystemByPackDiscoveryTypeEdgePrometheusDNS`     | dns                                                         |
| `CreateInputSystemByPackDiscoveryTypeEdgePrometheusEc2`     | ec2                                                         |
| `CreateInputSystemByPackDiscoveryTypeEdgePrometheusK8sNode` | k8s-node                                                    |
| `CreateInputSystemByPackDiscoveryTypeEdgePrometheusK8sPods` | k8s-pods                                                    |