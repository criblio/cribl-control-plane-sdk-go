# ParquetVersionOptions

Determines which data types are supported and how they are represented

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.ParquetVersionOptionsParquet10

// Open enum: custom values can be created with a direct type cast
custom := components.ParquetVersionOptions("custom_value")
```


## Values

| Name                             | Value                            |
| -------------------------------- | -------------------------------- |
| `ParquetVersionOptionsParquet10` | PARQUET_1_0                      |
| `ParquetVersionOptionsParquet24` | PARQUET_2_4                      |
| `ParquetVersionOptionsParquet26` | PARQUET_2_6                      |