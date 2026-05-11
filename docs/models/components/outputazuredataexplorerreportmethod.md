# OutputAzureDataExplorerReportMethod

Target of the ingestion status reporting. Defaults to Queue.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.OutputAzureDataExplorerReportMethodQueue

// Open enum: custom values can be created with a direct type cast
custom := components.OutputAzureDataExplorerReportMethod("custom_value")
```


## Values

| Name                                               | Value                                              |
| -------------------------------------------------- | -------------------------------------------------- |
| `OutputAzureDataExplorerReportMethodQueue`         | queue                                              |
| `OutputAzureDataExplorerReportMethodTable`         | table                                              |
| `OutputAzureDataExplorerReportMethodQueueAndTable` | queueAndTable                                      |