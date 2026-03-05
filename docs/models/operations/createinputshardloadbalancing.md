# CreateInputShardLoadBalancing

The load-balancing algorithm to use for spreading out shards across Workers and Worker Processes

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/operations"
)

value := operations.CreateInputShardLoadBalancingConsistentHashing

// Open enum: custom values can be created with a direct type cast
custom := operations.CreateInputShardLoadBalancing("custom_value")
```


## Values

| Name                                             | Value                                            |
| ------------------------------------------------ | ------------------------------------------------ |
| `CreateInputShardLoadBalancingConsistentHashing` | ConsistentHashing                                |
| `CreateInputShardLoadBalancingRoundRobin`        | RoundRobin                                       |