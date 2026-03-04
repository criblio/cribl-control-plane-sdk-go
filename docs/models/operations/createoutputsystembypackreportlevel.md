# CreateOutputSystemByPackReportLevel

Level of ingestion status reporting. Defaults to FailuresOnly.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/operations"
)

value := operations.CreateOutputSystemByPackReportLevelFailuresOnly

// Open enum: custom values can be created with a direct type cast
custom := operations.CreateOutputSystemByPackReportLevel("custom_value")
```


## Values

| Name                                                      | Value                                                     |
| --------------------------------------------------------- | --------------------------------------------------------- |
| `CreateOutputSystemByPackReportLevelFailuresOnly`         | failuresOnly                                              |
| `CreateOutputSystemByPackReportLevelDoNotReport`          | doNotReport                                               |
| `CreateOutputSystemByPackReportLevelFailuresAndSuccesses` | failuresAndSuccesses                                      |