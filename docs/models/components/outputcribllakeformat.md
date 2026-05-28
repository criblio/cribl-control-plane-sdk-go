# OutputCriblLakeFormat

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.OutputCriblLakeFormatJSON

// Open enum: custom values can be created with a direct type cast
custom := components.OutputCriblLakeFormat("custom_value")
```


## Values

| Name                            | Value                           |
| ------------------------------- | ------------------------------- |
| `OutputCriblLakeFormatJSON`     | json                            |
| `OutputCriblLakeFormatParquet`  | parquet                         |
| `OutputCriblLakeFormatDdss`     | ddss                            |
| `OutputCriblLakeFormatNetskope` | netskope                        |