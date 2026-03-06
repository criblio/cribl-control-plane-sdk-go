# CreateOutputFormatLocalSearchStorage

Data format to use when sending data. Defaults to JSON Compact.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/operations"
)

value := operations.CreateOutputFormatLocalSearchStorageJSONCompactEachRowWithNames

// Open enum: custom values can be created with a direct type cast
custom := operations.CreateOutputFormatLocalSearchStorage("custom_value")
```


## Values

| Name                                                              | Value                                                             |
| ----------------------------------------------------------------- | ----------------------------------------------------------------- |
| `CreateOutputFormatLocalSearchStorageJSONCompactEachRowWithNames` | json-compact-each-row-with-names                                  |
| `CreateOutputFormatLocalSearchStorageJSONEachRow`                 | json-each-row                                                     |