# SerializeTypeCsvType

Data output format

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.SerializeTypeCsvTypeCsv

// Open enum: custom values can be created with a direct type cast
custom := components.SerializeTypeCsvType("custom_value")
```


## Values

| Name                        | Value                       |
| --------------------------- | --------------------------- |
| `SerializeTypeCsvTypeCsv`   | csv                         |
| `SerializeTypeCsvTypeElff`  | elff                        |
| `SerializeTypeCsvTypeClf`   | clf                         |
| `SerializeTypeCsvTypeKvp`   | kvp                         |
| `SerializeTypeCsvTypeJSON`  | json                        |
| `SerializeTypeCsvTypeDelim` | delim                       |