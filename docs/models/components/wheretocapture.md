# WhereToCapture

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.WhereToCaptureBeforePreProcessingPipeline

// Open enum: custom values can be created with a direct type cast
custom := components.WhereToCapture(999)
```


## Values

| Name                                         | Value                                        |
| -------------------------------------------- | -------------------------------------------- |
| `WhereToCaptureBeforePreProcessingPipeline`  | 0                                            |
| `WhereToCaptureBeforeTheRoutes`              | 1                                            |
| `WhereToCaptureBeforePostProcessingPipeline` | 2                                            |
| `WhereToCaptureBeforeTheDestination`         | 3                                            |