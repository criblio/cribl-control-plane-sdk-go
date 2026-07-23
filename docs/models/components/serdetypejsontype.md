# SerdeTypeJSONType

Parser or formatter type to use

## Example Usage

```go
import (
	"github.com/Cribl-Community/cribl-control-plane-sdk-go/models/components"
)

value := components.SerdeTypeJSONTypeAuto

// Open enum: custom values can be created with a direct type cast
custom := components.SerdeTypeJSONType("custom_value")
```


## Values

| Name                     | Value                    |
| ------------------------ | ------------------------ |
| `SerdeTypeJSONTypeAuto`  | auto                     |
| `SerdeTypeJSONTypeCsv`   | csv                      |
| `SerdeTypeJSONTypeElff`  | elff                     |
| `SerdeTypeJSONTypeClf`   | clf                      |
| `SerdeTypeJSONTypeKvp`   | kvp                      |
| `SerdeTypeJSONTypeJSON`  | json                     |
| `SerdeTypeJSONTypeDelim` | delim                    |
| `SerdeTypeJSONTypeRegex` | regex                    |
| `SerdeTypeJSONTypeGrok`  | grok                     |