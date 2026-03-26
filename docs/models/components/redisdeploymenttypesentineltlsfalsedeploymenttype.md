# RedisDeploymentTypeSentinelTLSFalseDeploymentType

How the Redis server is configured. Defaults to Standalone

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.RedisDeploymentTypeSentinelTLSFalseDeploymentTypeStandalone

// Open enum: custom values can be created with a direct type cast
custom := components.RedisDeploymentTypeSentinelTLSFalseDeploymentType("custom_value")
```


## Values

| Name                                                          | Value                                                         |
| ------------------------------------------------------------- | ------------------------------------------------------------- |
| `RedisDeploymentTypeSentinelTLSFalseDeploymentTypeStandalone` | standalone                                                    |
| `RedisDeploymentTypeSentinelTLSFalseDeploymentTypeCluster`    | cluster                                                       |
| `RedisDeploymentTypeSentinelTLSFalseDeploymentTypeSentinel`   | sentinel                                                      |