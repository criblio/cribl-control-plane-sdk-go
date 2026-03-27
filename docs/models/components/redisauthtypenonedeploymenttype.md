# RedisAuthTypeNoneDeploymentType

How the Redis server is configured. Defaults to Standalone

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.RedisAuthTypeNoneDeploymentTypeStandalone

// Open enum: custom values can be created with a direct type cast
custom := components.RedisAuthTypeNoneDeploymentType("custom_value")
```


## Values

| Name                                        | Value                                       |
| ------------------------------------------- | ------------------------------------------- |
| `RedisAuthTypeNoneDeploymentTypeStandalone` | standalone                                  |
| `RedisAuthTypeNoneDeploymentTypeCluster`    | cluster                                     |
| `RedisAuthTypeNoneDeploymentTypeSentinel`   | sentinel                                    |