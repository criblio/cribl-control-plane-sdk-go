# CacheConnectionBackfillStatus

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.CacheConnectionBackfillStatusScheduled

// Open enum: custom values can be created with a direct type cast
custom := components.CacheConnectionBackfillStatus("custom_value")
```


## Values

| Name                                      | Value                                     |
| ----------------------------------------- | ----------------------------------------- |
| `CacheConnectionBackfillStatusScheduled`  | scheduled                                 |
| `CacheConnectionBackfillStatusPending`    | pending                                   |
| `CacheConnectionBackfillStatusStarted`    | started                                   |
| `CacheConnectionBackfillStatusFinished`   | finished                                  |
| `CacheConnectionBackfillStatusIncomplete` | incomplete                                |