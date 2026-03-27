# RedisDeploymentTypeClusterTLSTrueScaleReads

Which nodes read commands should be sent to

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.RedisDeploymentTypeClusterTLSTrueScaleReadsMaster

// Open enum: custom values can be created with a direct type cast
custom := components.RedisDeploymentTypeClusterTLSTrueScaleReads("custom_value")
```


## Values

| Name                                                 | Value                                                |
| ---------------------------------------------------- | ---------------------------------------------------- |
| `RedisDeploymentTypeClusterTLSTrueScaleReadsMaster`  | master                                               |
| `RedisDeploymentTypeClusterTLSTrueScaleReadsReplica` | replica                                              |
| `RedisDeploymentTypeClusterTLSTrueScaleReadsAll`     | all                                                  |