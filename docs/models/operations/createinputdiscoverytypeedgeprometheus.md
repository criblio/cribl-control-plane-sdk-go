# CreateInputDiscoveryTypeEdgePrometheus

Target discovery mechanism. Use static to manually enter a list of targets.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/operations"
)

value := operations.CreateInputDiscoveryTypeEdgePrometheusStatic

// Open enum: custom values can be created with a direct type cast
custom := operations.CreateInputDiscoveryTypeEdgePrometheus("custom_value")
```


## Values

| Name                                            | Value                                           |
| ----------------------------------------------- | ----------------------------------------------- |
| `CreateInputDiscoveryTypeEdgePrometheusStatic`  | static                                          |
| `CreateInputDiscoveryTypeEdgePrometheusDNS`     | dns                                             |
| `CreateInputDiscoveryTypeEdgePrometheusEc2`     | ec2                                             |
| `CreateInputDiscoveryTypeEdgePrometheusK8sNode` | k8s-node                                        |
| `CreateInputDiscoveryTypeEdgePrometheusK8sPods` | k8s-pods                                        |