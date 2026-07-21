# SerdeTypeCsvType

Parser or formatter type to use.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.SerdeTypeCsvTypeAuto

// Open enum: custom values can be created with a direct type cast
custom := components.SerdeTypeCsvType("custom_value")
```


## Values

| Name                    | Value                   |
| ----------------------- | ----------------------- |
| `SerdeTypeCsvTypeAuto`  | auto                    |
| `SerdeTypeCsvTypeCsv`   | csv                     |
| `SerdeTypeCsvTypeElff`  | elff                    |
| `SerdeTypeCsvTypeClf`   | clf                     |
| `SerdeTypeCsvTypeKvp`   | kvp                     |
| `SerdeTypeCsvTypeJSON`  | json                    |
| `SerdeTypeCsvTypeDelim` | delim                   |
| `SerdeTypeCsvTypeRegex` | regex                   |
| `SerdeTypeCsvTypeGrok`  | grok                    |