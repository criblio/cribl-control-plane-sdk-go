# CreateOutputSystemByPackFormatClickHouse

Data format to use when sending data to ClickHouse. Defaults to JSON Compact.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/operations"
)

value := operations.CreateOutputSystemByPackFormatClickHouseJSONCompactEachRowWithNames

// Open enum: custom values can be created with a direct type cast
custom := operations.CreateOutputSystemByPackFormatClickHouse("custom_value")
```


## Values

| Name                                                                  | Value                                                                 |
| --------------------------------------------------------------------- | --------------------------------------------------------------------- |
| `CreateOutputSystemByPackFormatClickHouseJSONCompactEachRowWithNames` | json-compact-each-row-with-names                                      |
| `CreateOutputSystemByPackFormatClickHouseJSONEachRow`                 | json-each-row                                                         |