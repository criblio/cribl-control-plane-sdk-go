# SerdeTypeAutoType

Parser or formatter type to use

## Example Usage

```go
import (
	"github.com/Cribl-Community/cribl-control-plane-sdk-go/models/components"
)

value := components.SerdeTypeAutoTypeAuto

// Open enum: custom values can be created with a direct type cast
custom := components.SerdeTypeAutoType("custom_value")
```


## Values

| Name                     | Value                    |
| ------------------------ | ------------------------ |
| `SerdeTypeAutoTypeAuto`  | auto                     |
| `SerdeTypeAutoTypeCsv`   | csv                      |
| `SerdeTypeAutoTypeElff`  | elff                     |
| `SerdeTypeAutoTypeClf`   | clf                      |
| `SerdeTypeAutoTypeKvp`   | kvp                      |
| `SerdeTypeAutoTypeJSON`  | json                     |
| `SerdeTypeAutoTypeDelim` | delim                    |
| `SerdeTypeAutoTypeRegex` | regex                    |
| `SerdeTypeAutoTypeGrok`  | grok                     |