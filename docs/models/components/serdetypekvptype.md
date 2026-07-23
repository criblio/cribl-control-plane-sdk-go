# SerdeTypeKvpType

Parser or formatter type to use

## Example Usage

```go
import (
	"github.com/Cribl-Community/cribl-control-plane-sdk-go/models/components"
)

value := components.SerdeTypeKvpTypeAuto

// Open enum: custom values can be created with a direct type cast
custom := components.SerdeTypeKvpType("custom_value")
```


## Values

| Name                    | Value                   |
| ----------------------- | ----------------------- |
| `SerdeTypeKvpTypeAuto`  | auto                    |
| `SerdeTypeKvpTypeCsv`   | csv                     |
| `SerdeTypeKvpTypeElff`  | elff                    |
| `SerdeTypeKvpTypeClf`   | clf                     |
| `SerdeTypeKvpTypeKvp`   | kvp                     |
| `SerdeTypeKvpTypeJSON`  | json                    |
| `SerdeTypeKvpTypeDelim` | delim                   |
| `SerdeTypeKvpTypeRegex` | regex                   |
| `SerdeTypeKvpTypeGrok`  | grok                    |