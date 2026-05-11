# FormatLocalSearchStorage

Data format to use when sending data. Defaults to JSON Compact.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.FormatLocalSearchStorageJSONCompactEachRowWithNames

// Open enum: custom values can be created with a direct type cast
custom := components.FormatLocalSearchStorage("custom_value")
```


## Values

| Name                                                  | Value                                                 |
| ----------------------------------------------------- | ----------------------------------------------------- |
| `FormatLocalSearchStorageJSONCompactEachRowWithNames` | json-compact-each-row-with-names                      |
| `FormatLocalSearchStorageJSONEachRow`                 | json-each-row                                         |