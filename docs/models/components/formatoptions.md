# FormatOptions

Data format to use when sending data to ClickHouse. Defaults to JSON Compact.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.FormatOptionsJSONCompactEachRowWithNames

// Open enum: custom values can be created with a direct type cast
custom := components.FormatOptions("custom_value")
```


## Values

| Name                                       | Value                                      |
| ------------------------------------------ | ------------------------------------------ |
| `FormatOptionsJSONCompactEachRowWithNames` | json-compact-each-row-with-names           |
| `FormatOptionsJSONEachRow`                 | json-each-row                              |