# OutputAzureDataExplorerReportLevel

Level of ingestion status reporting. Defaults to FailuresOnly.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.OutputAzureDataExplorerReportLevelFailuresOnly

// Open enum: custom values can be created with a direct type cast
custom := components.OutputAzureDataExplorerReportLevel("custom_value")
```


## Values

| Name                                                     | Value                                                    |
| -------------------------------------------------------- | -------------------------------------------------------- |
| `OutputAzureDataExplorerReportLevelFailuresOnly`         | failuresOnly                                             |
| `OutputAzureDataExplorerReportLevelDoNotReport`          | doNotReport                                              |
| `OutputAzureDataExplorerReportLevelFailuresAndSuccesses` | failuresAndSuccesses                                     |