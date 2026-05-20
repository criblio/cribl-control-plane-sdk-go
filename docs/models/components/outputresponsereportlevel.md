# OutputResponseReportLevel

Level of ingestion status reporting. Defaults to FailuresOnly.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.OutputResponseReportLevelFailuresOnly

// Open enum: custom values can be created with a direct type cast
custom := components.OutputResponseReportLevel("custom_value")
```


## Values

| Name                                            | Value                                           |
| ----------------------------------------------- | ----------------------------------------------- |
| `OutputResponseReportLevelFailuresOnly`         | failuresOnly                                    |
| `OutputResponseReportLevelDoNotReport`          | doNotReport                                     |
| `OutputResponseReportLevelFailuresAndSuccesses` | failuresAndSuccesses                            |