# RedisDeploymentTypeClusterTLSFalseScaleReads

Which nodes read commands should be sent to

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.RedisDeploymentTypeClusterTLSFalseScaleReadsMaster

// Open enum: custom values can be created with a direct type cast
custom := components.RedisDeploymentTypeClusterTLSFalseScaleReads("custom_value")
```


## Values

| Name                                                  | Value                                                 |
| ----------------------------------------------------- | ----------------------------------------------------- |
| `RedisDeploymentTypeClusterTLSFalseScaleReadsMaster`  | master                                                |
| `RedisDeploymentTypeClusterTLSFalseScaleReadsReplica` | replica                                               |
| `RedisDeploymentTypeClusterTLSFalseScaleReadsAll`     | all                                                   |