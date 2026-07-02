# SerdeTypeDelimType

Parser or formatter type to use

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.SerdeTypeDelimTypeCsv

// Open enum: custom values can be created with a direct type cast
custom := components.SerdeTypeDelimType("custom_value")
```


## Values

| Name                      | Value                     |
| ------------------------- | ------------------------- |
| `SerdeTypeDelimTypeCsv`   | csv                       |
| `SerdeTypeDelimTypeElff`  | elff                      |
| `SerdeTypeDelimTypeClf`   | clf                       |
| `SerdeTypeDelimTypeKvp`   | kvp                       |
| `SerdeTypeDelimTypeJSON`  | json                      |
| `SerdeTypeDelimTypeDelim` | delim                     |
| `SerdeTypeDelimTypeRegex` | regex                     |
| `SerdeTypeDelimTypeGrok`  | grok                      |
| `SerdeTypeDelimTypeAuto`  | auto                      |