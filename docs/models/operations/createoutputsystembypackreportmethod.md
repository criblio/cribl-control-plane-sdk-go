# CreateOutputSystemByPackReportMethod

Target of the ingestion status reporting. Defaults to Queue.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/operations"
)

value := operations.CreateOutputSystemByPackReportMethodQueue

// Open enum: custom values can be created with a direct type cast
custom := operations.CreateOutputSystemByPackReportMethod("custom_value")
```


## Values

| Name                                                | Value                                               |
| --------------------------------------------------- | --------------------------------------------------- |
| `CreateOutputSystemByPackReportMethodQueue`         | queue                                               |
| `CreateOutputSystemByPackReportMethodTable`         | table                                               |
| `CreateOutputSystemByPackReportMethodQueueAndTable` | queueAndTable                                       |