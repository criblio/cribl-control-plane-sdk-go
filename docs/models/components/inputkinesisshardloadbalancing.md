# InputKinesisShardLoadBalancing

The load-balancing algorithm to use for spreading out shards across Workers and Worker Processes

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.InputKinesisShardLoadBalancingConsistentHashing

// Open enum: custom values can be created with a direct type cast
custom := components.InputKinesisShardLoadBalancing("custom_value")
```


## Values

| Name                                              | Value                                             |
| ------------------------------------------------- | ------------------------------------------------- |
| `InputKinesisShardLoadBalancingConsistentHashing` | ConsistentHashing                                 |
| `InputKinesisShardLoadBalancingRoundRobin`        | RoundRobin                                        |