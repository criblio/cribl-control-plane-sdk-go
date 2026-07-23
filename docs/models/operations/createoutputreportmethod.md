# CreateOutputReportMethod

Target of the ingestion status reporting. Defaults to Queue.

## Example Usage

```go
import (
	"github.com/Cribl-Community/cribl-control-plane-sdk-go/models/operations"
)

value := operations.CreateOutputReportMethodQueue

// Open enum: custom values can be created with a direct type cast
custom := operations.CreateOutputReportMethod("custom_value")
```


## Values

| Name                                    | Value                                   |
| --------------------------------------- | --------------------------------------- |
| `CreateOutputReportMethodQueue`         | queue                                   |
| `CreateOutputReportMethodTable`         | table                                   |
| `CreateOutputReportMethodQueueAndTable` | queueAndTable                           |