# TriggerType

Type of the trigger condition. custom applies a kusto expression over the results, and results count applies a comparison over results count

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.TriggerTypeCustom

// Open enum: custom values can be created with a direct type cast
custom := components.TriggerType("custom_value")
```


## Values

| Name                      | Value                     |
| ------------------------- | ------------------------- |
| `TriggerTypeCustom`       | custom                    |
| `TriggerTypeResultsCount` | resultsCount              |