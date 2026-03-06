# FormatOptions

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.FormatOptionsJSON

// Open enum: custom values can be created with a direct type cast
custom := components.FormatOptions("custom_value")
```


## Values

| Name                   | Value                  |
| ---------------------- | ---------------------- |
| `FormatOptionsJSON`    | json                   |
| `FormatOptionsParquet` | parquet                |
| `FormatOptionsDdss`    | ddss                   |