# SerdeTypeRegexType

Parser or formatter type to use

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.SerdeTypeRegexTypeCsv

// Open enum: custom values can be created with a direct type cast
custom := components.SerdeTypeRegexType("custom_value")
```


## Values

| Name                      | Value                     |
| ------------------------- | ------------------------- |
| `SerdeTypeRegexTypeCsv`   | csv                       |
| `SerdeTypeRegexTypeElff`  | elff                      |
| `SerdeTypeRegexTypeClf`   | clf                       |
| `SerdeTypeRegexTypeKvp`   | kvp                       |
| `SerdeTypeRegexTypeJSON`  | json                      |
| `SerdeTypeRegexTypeDelim` | delim                     |
| `SerdeTypeRegexTypeRegex` | regex                     |
| `SerdeTypeRegexTypeGrok`  | grok                      |
| `SerdeTypeRegexTypeAuto`  | auto                      |