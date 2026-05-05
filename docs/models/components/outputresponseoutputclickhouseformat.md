# OutputResponseOutputClickHouseFormat

Data format to use when sending data to ClickHouse. Defaults to JSON Compact.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.OutputResponseOutputClickHouseFormatJSONCompactEachRowWithNames

// Open enum: custom values can be created with a direct type cast
custom := components.OutputResponseOutputClickHouseFormat("custom_value")
```


## Values

| Name                                                              | Value                                                             |
| ----------------------------------------------------------------- | ----------------------------------------------------------------- |
| `OutputResponseOutputClickHouseFormatJSONCompactEachRowWithNames` | json-compact-each-row-with-names                                  |
| `OutputResponseOutputClickHouseFormatJSONEachRow`                 | json-each-row                                                     |