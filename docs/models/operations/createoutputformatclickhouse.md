# CreateOutputFormatClickHouse

Data format to use when sending data to ClickHouse. Defaults to JSON Compact.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/operations"
)

value := operations.CreateOutputFormatClickHouseJSONCompactEachRowWithNames

// Open enum: custom values can be created with a direct type cast
custom := operations.CreateOutputFormatClickHouse("custom_value")
```


## Values

| Name                                                      | Value                                                     |
| --------------------------------------------------------- | --------------------------------------------------------- |
| `CreateOutputFormatClickHouseJSONCompactEachRowWithNames` | json-compact-each-row-with-names                          |
| `CreateOutputFormatClickHouseJSONEachRow`                 | json-each-row                                             |