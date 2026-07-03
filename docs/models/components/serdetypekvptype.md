# SerdeTypeKvpType

Parser or formatter type to use

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.SerdeTypeKvpTypeCsv

// Open enum: custom values can be created with a direct type cast
custom := components.SerdeTypeKvpType("custom_value")
```


## Values

| Name                    | Value                   |
| ----------------------- | ----------------------- |
| `SerdeTypeKvpTypeCsv`   | csv                     |
| `SerdeTypeKvpTypeElff`  | elff                    |
| `SerdeTypeKvpTypeClf`   | clf                     |
| `SerdeTypeKvpTypeKvp`   | kvp                     |
| `SerdeTypeKvpTypeJSON`  | json                    |
| `SerdeTypeKvpTypeDelim` | delim                   |
| `SerdeTypeKvpTypeRegex` | regex                   |
| `SerdeTypeKvpTypeGrok`  | grok                    |
| `SerdeTypeKvpTypeAuto`  | auto                    |