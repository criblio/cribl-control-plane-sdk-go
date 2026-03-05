# RunnableJobCollectionMode

Job run mode. Preview will either return up to N matching results, or will run until capture time T is reached. Discovery will gather the list of files to turn into streaming tasks, without running the data collection job. Full Run will run the collection job.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.RunnableJobCollectionModeList

// Open enum: custom values can be created with a direct type cast
custom := components.RunnableJobCollectionMode("custom_value")
```


## Values

| Name                               | Value                              |
| ---------------------------------- | ---------------------------------- |
| `RunnableJobCollectionModeList`    | list                               |
| `RunnableJobCollectionModePreview` | preview                            |
| `RunnableJobCollectionModeRun`     | run                                |