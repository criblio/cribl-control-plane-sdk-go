# CaptureLevel

Stage at which events are captured. <br><code>0</code> == Before pre-processing Pipeline <br><code>1</code> == Before the Routes <br><code>2</code> == Before post-processing Pipeline <br><code>3</code> == Before the Destination.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.CaptureLevelBeforePreProcessingPipeline

// Open enum: custom values can be created with a direct type cast
custom := components.CaptureLevel(999)
```


## Values

| Name                                       | Value                                      |
| ------------------------------------------ | ------------------------------------------ |
| `CaptureLevelBeforePreProcessingPipeline`  | 0                                          |
| `CaptureLevelBeforeRoutes`                 | 1                                          |
| `CaptureLevelBeforePostProcessingPipeline` | 2                                          |
| `CaptureLevelBeforeDestination`            | 3                                          |