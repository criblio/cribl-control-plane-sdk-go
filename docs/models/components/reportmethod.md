# ReportMethod

Target of the ingestion status reporting. Defaults to Queue.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.ReportMethodQueue

// Open enum: custom values can be created with a direct type cast
custom := components.ReportMethod("custom_value")
```


## Values

| Name                        | Value                       |
| --------------------------- | --------------------------- |
| `ReportMethodQueue`         | queue                       |
| `ReportMethodTable`         | table                       |
| `ReportMethodQueueAndTable` | queueAndTable               |