# ScaleReads

Which nodes read commands should be sent to

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.ScaleReadsMaster

// Open enum: custom values can be created with a direct type cast
custom := components.ScaleReads("custom_value")
```


## Values

| Name                | Value               |
| ------------------- | ------------------- |
| `ScaleReadsMaster`  | master              |
| `ScaleReadsReplica` | replica             |
| `ScaleReadsAll`     | all                 |