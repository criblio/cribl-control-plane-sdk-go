# RedisAuthTypeManualDeploymentType

How the Redis server is configured. Defaults to Standalone

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.RedisAuthTypeManualDeploymentTypeStandalone

// Open enum: custom values can be created with a direct type cast
custom := components.RedisAuthTypeManualDeploymentType("custom_value")
```


## Values

| Name                                          | Value                                         |
| --------------------------------------------- | --------------------------------------------- |
| `RedisAuthTypeManualDeploymentTypeStandalone` | standalone                                    |
| `RedisAuthTypeManualDeploymentTypeCluster`    | cluster                                       |
| `RedisAuthTypeManualDeploymentTypeSentinel`   | sentinel                                      |