# FormatClickHouse

Data format to use when sending data to ClickHouse. Defaults to JSON Compact.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.FormatClickHouseJSONCompactEachRowWithNames

// Open enum: custom values can be created with a direct type cast
custom := components.FormatClickHouse("custom_value")
```


## Values

| Name                                          | Value                                         |
| --------------------------------------------- | --------------------------------------------- |
| `FormatClickHouseJSONCompactEachRowWithNames` | json-compact-each-row-with-names              |
| `FormatClickHouseJSONEachRow`                 | json-each-row                                 |