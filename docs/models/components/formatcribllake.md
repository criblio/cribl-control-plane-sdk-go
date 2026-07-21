# FormatCriblLake

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.FormatCriblLakeJSON

// Open enum: custom values can be created with a direct type cast
custom := components.FormatCriblLake("custom_value")
```


## Values

| Name                     | Value                    |
| ------------------------ | ------------------------ |
| `FormatCriblLakeJSON`    | json                     |
| `FormatCriblLakeParquet` | parquet                  |
| `FormatCriblLakeRaw`     | raw                      |