# OutputClickHouseFormat

Data format to use when sending data to ClickHouse. Defaults to JSON Compact.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.OutputClickHouseFormatJSONCompactEachRowWithNames

// Open enum: custom values can be created with a direct type cast
custom := components.OutputClickHouseFormat("custom_value")
```


## Values

| Name                                                | Value                                               |
| --------------------------------------------------- | --------------------------------------------------- |
| `OutputClickHouseFormatJSONCompactEachRowWithNames` | json-compact-each-row-with-names                    |
| `OutputClickHouseFormatJSONEachRow`                 | json-each-row                                       |