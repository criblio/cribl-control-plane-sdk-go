# CreateOutputFormatCriblLake

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/operations"
)

value := operations.CreateOutputFormatCriblLakeJSON

// Open enum: custom values can be created with a direct type cast
custom := operations.CreateOutputFormatCriblLake("custom_value")
```


## Values

| Name                                  | Value                                 |
| ------------------------------------- | ------------------------------------- |
| `CreateOutputFormatCriblLakeJSON`     | json                                  |
| `CreateOutputFormatCriblLakeParquet`  | parquet                               |
| `CreateOutputFormatCriblLakeDdss`     | ddss                                  |
| `CreateOutputFormatCriblLakeNetskope` | netskope                              |