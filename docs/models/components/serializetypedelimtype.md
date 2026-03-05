# SerializeTypeDelimType

Data output format

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.SerializeTypeDelimTypeCsv

// Open enum: custom values can be created with a direct type cast
custom := components.SerializeTypeDelimType("custom_value")
```


## Values

| Name                          | Value                         |
| ----------------------------- | ----------------------------- |
| `SerializeTypeDelimTypeCsv`   | csv                           |
| `SerializeTypeDelimTypeElff`  | elff                          |
| `SerializeTypeDelimTypeClf`   | clf                           |
| `SerializeTypeDelimTypeKvp`   | kvp                           |
| `SerializeTypeDelimTypeJSON`  | json                          |
| `SerializeTypeDelimTypeDelim` | delim                         |