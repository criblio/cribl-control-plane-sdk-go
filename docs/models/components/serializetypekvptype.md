# SerializeTypeKvpType

Data output format

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.SerializeTypeKvpTypeCsv

// Open enum: custom values can be created with a direct type cast
custom := components.SerializeTypeKvpType("custom_value")
```


## Values

| Name                        | Value                       |
| --------------------------- | --------------------------- |
| `SerializeTypeKvpTypeCsv`   | csv                         |
| `SerializeTypeKvpTypeElff`  | elff                        |
| `SerializeTypeKvpTypeClf`   | clf                         |
| `SerializeTypeKvpTypeKvp`   | kvp                         |
| `SerializeTypeKvpTypeJSON`  | json                        |
| `SerializeTypeKvpTypeDelim` | delim                       |