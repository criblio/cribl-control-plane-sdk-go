# RedisDeploymentTypeStandaloneDeploymentType

How the Redis server is configured. Defaults to Standalone

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.RedisDeploymentTypeStandaloneDeploymentTypeStandalone

// Open enum: custom values can be created with a direct type cast
custom := components.RedisDeploymentTypeStandaloneDeploymentType("custom_value")
```


## Values

| Name                                                    | Value                                                   |
| ------------------------------------------------------- | ------------------------------------------------------- |
| `RedisDeploymentTypeStandaloneDeploymentTypeStandalone` | standalone                                              |
| `RedisDeploymentTypeStandaloneDeploymentTypeCluster`    | cluster                                                 |
| `RedisDeploymentTypeStandaloneDeploymentTypeSentinel`   | sentinel                                                |