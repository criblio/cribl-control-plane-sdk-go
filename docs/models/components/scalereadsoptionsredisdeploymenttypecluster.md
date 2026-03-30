# ScaleReadsOptionsRedisDeploymentTypeCluster

Which nodes read commands should be sent to

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.ScaleReadsOptionsRedisDeploymentTypeClusterMaster

// Open enum: custom values can be created with a direct type cast
custom := components.ScaleReadsOptionsRedisDeploymentTypeCluster("custom_value")
```


## Values

| Name                                                 | Value                                                |
| ---------------------------------------------------- | ---------------------------------------------------- |
| `ScaleReadsOptionsRedisDeploymentTypeClusterMaster`  | master                                               |
| `ScaleReadsOptionsRedisDeploymentTypeClusterReplica` | replica                                              |
| `ScaleReadsOptionsRedisDeploymentTypeClusterAll`     | all                                                  |