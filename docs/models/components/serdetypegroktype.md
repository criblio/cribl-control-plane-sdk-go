# SerdeTypeGrokType

Parser or formatter type to use.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.SerdeTypeGrokTypeAuto

// Open enum: custom values can be created with a direct type cast
custom := components.SerdeTypeGrokType("custom_value")
```


## Values

| Name                     | Value                    |
| ------------------------ | ------------------------ |
| `SerdeTypeGrokTypeAuto`  | auto                     |
| `SerdeTypeGrokTypeCsv`   | csv                      |
| `SerdeTypeGrokTypeElff`  | elff                     |
| `SerdeTypeGrokTypeClf`   | clf                      |
| `SerdeTypeGrokTypeKvp`   | kvp                      |
| `SerdeTypeGrokTypeJSON`  | json                     |
| `SerdeTypeGrokTypeDelim` | delim                    |
| `SerdeTypeGrokTypeRegex` | regex                    |
| `SerdeTypeGrokTypeGrok`  | grok                     |