# ReportLevel

Level of ingestion status reporting. Defaults to FailuresOnly.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.ReportLevelFailuresOnly

// Open enum: custom values can be created with a direct type cast
custom := components.ReportLevel("custom_value")
```


## Values

| Name                              | Value                             |
| --------------------------------- | --------------------------------- |
| `ReportLevelFailuresOnly`         | failuresOnly                      |
| `ReportLevelDoNotReport`          | doNotReport                       |
| `ReportLevelFailuresAndSuccesses` | failuresAndSuccesses              |