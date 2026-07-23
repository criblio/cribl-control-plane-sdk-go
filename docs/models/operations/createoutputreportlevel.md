# CreateOutputReportLevel

Level of ingestion status reporting. Defaults to FailuresOnly.

## Example Usage

```go
import (
	"github.com/Cribl-Community/cribl-control-plane-sdk-go/models/operations"
)

value := operations.CreateOutputReportLevelFailuresOnly

// Open enum: custom values can be created with a direct type cast
custom := operations.CreateOutputReportLevel("custom_value")
```


## Values

| Name                                          | Value                                         |
| --------------------------------------------- | --------------------------------------------- |
| `CreateOutputReportLevelFailuresOnly`         | failuresOnly                                  |
| `CreateOutputReportLevelDoNotReport`          | doNotReport                                   |
| `CreateOutputReportLevelFailuresAndSuccesses` | failuresAndSuccesses                          |